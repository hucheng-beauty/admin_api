package marketing

import (
	"admin_api/global"
	"admin_api/internal/api/common"
	"admin_api/internal/data"
	"admin_api/internal/event/publish"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/internal/service/fetch/wechat"
	consumer_service "admin_api/internal/service/store/consumer"
	send_record_service "admin_api/internal/service/store/marketing_campaign"
	"errors"
	"github.com/gin-gonic/gin"
	"go.uber.org/zap"
	"net/http"
)

/*
	Action:CreateSendRecord
	1.获取手机号、手机号去重、活动Id
	2.获取活动名称、活动剩余数量
	3.获取单个活动下的所有券批次 Id
	4.通过手机号获取用户的 openid
	5.发送记录入库
	6.获取商家 appid、密钥、私钥等相关信息(TODO)
	7.微信发券
	8.使用记录入库
	9.微信获取券详情
	10.券详情入库
	11.更新发送记录
	12.通知营销活动活动发送总量
*/

type SendRecord struct{}

// DescribeSendRecord 发送记录列表
func (SendRecord) DescribeSendRecord(c *gin.Context) {
	srs := send_record_service.NewSendRecordService(data.NewSendRecordRepo(global.DB))
	resp, total, err := srs.DescribeSendRecord(&request.DescribeSendRecord{Pagination: common.OffsetAndLimitHandle(c)})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	c.JSON(http.StatusOK, response.SuccessWithPagination(resp, response.NewPagination(total)))
	return
}

// CreateSendRecord 创建定向发送记录
func (SendRecord) CreateSendRecord(c *gin.Context) {
	var req *request.CreateSendRecord
	err := c.ShouldBind(&req)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	//切片去重
	var (
		Unique = func(slice []string) []string {
			m := make(map[string]bool)
			var list []string
			for _, v := range slice {
				if _, ok := m[v]; !ok {
					m[v] = true
					list = append(list, v)
				}
			}
			return list
		}
		consumers []*response.DescribeConsumer
		batch     []*response.DescribeCouponBatch
		mcd       *response.GetMarketingCampaignDetail
	)

	// 获取手机号、手机号去重、活动Id
	req.CampaignId = c.Param("campaign_id")
	if req.CampaignId == "" {
		c.JSON(http.StatusOK, response.Error(1, errors.New("活动Id不能为空").Error()))
		return
	}
	//账户id去重
	req.AccountIds = Unique(req.AccountIds)
	if len(req.AccountIds) <= 0 {
		c.JSON(http.StatusOK, response.Error(1, errors.New("有效账号Id数量不能为空").Error()))
		return
	}

	// 获取活动名称、活动剩余数量
	mcs := send_record_service.NewMarketingCampaignService(data.NewMarketingCampaignRepo(global.DB))
	mcd, err = mcs.GetMarketingCampaignDetail(&request.GetMarketingCampaignDetail{Id: req.CampaignId})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	req.CampaignName = mcd.CampaignName

	// 获取单个活动下的所有券批次 Id
	cbs := send_record_service.NewCouponBatchService(data.NewCouponBatchRepo(global.DB))
	batch, err = cbs.DescribeCouponBatch(&request.DescribeCouponBatch{MarketingCampaignId: req.CampaignId})
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	if len(batch) <= 0 {
		c.JSON(http.StatusOK, response.Error(1, errors.New("该活动下未查到券").Error()))
		return
	}
	req.TotalCount = len(req.AccountIds) * len(batch)

	// 获取用户openid
	cs := consumer_service.NewConsumerService(data.NewConsumerRepo(global.DB))
	consumers, _, err = cs.DescribeConsumer(&request.DescribeConsumer{AccountIds: req.AccountIds}, nil)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}
	if len(consumers) <= 0 {
		c.JSON(http.StatusOK, response.Error(1, errors.New("未查到用户").Error()))
		return
	}

	// 发送记录入库
	srs := send_record_service.NewSendRecordService(data.NewSendRecordRepo(global.DB))
	resp, err := srs.CreateSendRecord(req)
	if err != nil {
		c.JSON(http.StatusOK, response.Error(1, err.Error()))
		return
	}

	go func(ccs []*response.DescribeConsumer) {
		zap.S().Info("开始发券")
		zap.S().Infof("campaign_name:%s,campaign_id:%s,len(stock_id):%d,len(consumers):%d",
			req.CampaignName, req.CampaignId, len(batch), len(consumers))

		//发送账户数量
		count := len(req.AccountIds)
		totalSuccessCount := 0
		var stockSendRecordInfo []*request.StockSendRecord
		//卷批次
		for _, b := range batch {
			successCount := 0
			for _, cc := range ccs {
				// 微信发券   卷批次id openid
				sendResp, sendErr := wechat.NewDefaultInfo().Send(b.StockId, cc.OpenId)
				if sendErr != nil {
					zap.S().Errorf("微信发券失败,err:%v,stock_id:%s,account_id:%s,openid:%s\n",
						sendErr.Error(), b.StockId, cc.AccountId, cc.OpenId)
					continue
				}

				zap.S().Infof("微信获取券,coupon_id:%s\n", sendResp.CouponId)

				// 使用记录入库
				crs := consumer_service.NewCouponRecordService(data.NewCouponRecordRepo(global.DB))
				//创建发送记录
				_, createErr := crs.CreateCouponRecord(&request.CreateCouponRecord{
					ConsumerId:   cc.ConsumerId,
					CouponId:     sendResp.CouponId,
					StockId:      b.StockId,
					StockName:    b.StockName,
					CampaignId:   req.CampaignId,
					CampaignName: req.CampaignName,
					AccountId:    cc.AccountId,
					BelongTo:     b.BelongTo,
				})
				if createErr != nil {
					zap.S().Errorf("使用记录入库失败,err:%v,stock_id:%s,account_id:%s,coupon_id:%s\n",
						createErr, b.StockId, cc.AccountId, sendResp.CouponId)
					continue
				}
				successCount++

				// 微信获取券详情
				CouponDetail, getErr := wechat.NewDefaultInfo().GetCouponDetail(sendResp.CouponId, cc.OpenId)
				if getErr != nil {
					zap.S().Errorf("微信获取券详情失败,err:%v,stock_id:%s,account_id:%s,openid:%s\n",
						getErr, b.StockId, cc.AccountId, cc.OpenId)
					continue
				}

				couponService := consumer_service.NewCouponService(data.NewCouponRepo(global.DB))
				_, createCouponErr := couponService.CreateCoupon(&request.CreateCoupon{
					CouponId:     CouponDetail.CouponId,
					CouponName:   CouponDetail.CouponName,
					CouponType:   CouponDetail.CouponType,
					CouponAmount: CouponDetail.NormalCouponInformation.CouponAmount,
					CreateTime:   CouponDetail.CreateTime,
					Description:  CouponDetail.Description,
					Status:       CouponDetail.Status,
				})
				if createCouponErr != nil {
					zap.S().Errorf("券详情入库失败,err:%v,stock_id:%s,account_id:%s,coupon_id:%s\n",
						createCouponErr, b.StockId, cc.AccountId, sendResp.CouponId)
					continue
				}
			}

			stockSendRecordInfo = append(stockSendRecordInfo, &request.StockSendRecord{
				StockId:      b.StockId,
				Count:        count,
				SuccessCount: successCount,
			})
			totalSuccessCount += successCount

			zap.S().Errorf("campaign_id:%s,stock_id:%s,count:%d,success_count:%d\n",
				req.CampaignId, b.StockId, count, successCount)
		}

		// 更新发送记录
		updateSendRecordReq := &request.UpdateSendRecord{
			Id:                  resp.Id,
			CampaignId:          req.CampaignId,
			SurplusCount:        int(mcd.CouponSurplusNumber) - totalSuccessCount,
			TotalCount:          req.TotalCount,
			TotalSuccessCount:   totalSuccessCount,
			StockSendRecordInfo: stockSendRecordInfo,
		}
		_, updateErr := srs.UpdateSendRecord(updateSendRecordReq)
		if updateErr != nil {
			zap.S().Errorf("更新发送记录失败,err:%v,campaign_id:%s,total_count:%d,total_success_count:%d\n",
				updateErr, req.CampaignId, req.TotalCount, totalSuccessCount)
		}

		// 通知营销活动活动发送总量
		go publish.WithMarketingCampaignSendRecord(c, updateSendRecordReq)

		zap.S().Info("发券完成")
	}(consumers)

	c.JSON(http.StatusOK, response.NewResponse(resp))
	return
}
