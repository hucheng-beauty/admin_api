package marketing

import (
	"admin_api/global"
	"admin_api/internal/api/common"
	"admin_api/internal/data"
	"admin_api/internal/request"
	"admin_api/internal/response"
	send_record_service "admin_api/internal/service/store/marketing_campaign"
	"github.com/gin-gonic/gin"
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
