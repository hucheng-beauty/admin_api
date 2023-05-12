package marketing

import (
	"admin_api/global"
	"admin_api/internal/data"
	"admin_api/internal/event/publish"
	"admin_api/internal/request"
	"admin_api/internal/response"
	marService "admin_api/internal/service/store/marketing_campaign"
	"admin_api/middlewares"
	"github.com/gin-gonic/gin"
	"time"
)

type MarketingCampaignApi struct {
}

func NewMarCampaignService() *marService.MarCampaignService {
	return marService.NewMarCampaignService(
		data.NewMarketingCampaignRepo(global.DB),
		data.NewCouponBatchRepo(global.DB),
		data.NewCouponTemplateRepo(global.DB),
		data.NewCouponLogRepo(global.DB),
		data.NewMarketingCampaignLogRepo(global.DB),
	)
}

// 创建活动和券批次
func (c MarketingCampaignApi) Create(ctx *gin.Context, in *request.CreateMarketingCampaignRequest, out *response.Id) error {
	//获取userid
	userId := middlewares.GetUserId(ctx)
	//时间判断
	if err := in.Validate(); err != nil {
		return err
	}

	//添加时间
	by, bm, bd := in.AvailableBeginTime.Date()
	ey, em, ed := in.AvailableEndTime.Date()
	// fix url https://www.tapd.cn/62185617/bugtrace/bugs/view/1162185617001000141
	in.AvailableBeginTime = time.Date(by, bm, bd, 0, 0, 0, 0, time.Local)
	in.AvailableEndTime = time.Date(ey, em, ed, 23, 59, 59, 0, time.Local)
	//创建营销活动和卷批次
	res, _, err := NewMarCampaignService().CreateMarCampaignAndCouponBatch(in, userId)
	if err != nil {
		return err
	}
	// 创建券日志
	_, err = NewMarCampaignService().SetCampaignState(res.Id, res.State)
	if err != nil {
		return err
	}
	//////
	go publish.WithMarketingCampaignCreate(ctx, res)
	out.Id = res.Id
	return nil
}
