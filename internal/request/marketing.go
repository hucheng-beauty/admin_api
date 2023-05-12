package request

import (
	"admin_api/internal/model"
	"fmt"
	"time"
)

type CreateCouponBatchReq struct {
	model.CouponBatchBaseInfo
	Comment string `json:"comment"`
}

type CreateMarketingCampaignRequest struct {
	model.MarketingCampaignBaseInfo
	CouponBatches []CreateCouponBatchReq `json:"coupon_batches"`
}

func (c CreateMarketingCampaignRequest) Validate() error {
	if c.AvailableBeginTime.Unix() <= time.Now().Unix() {
		return fmt.Errorf("活动开始时间不能小于当前时间")
	}
	if c.AvailableBeginTime.Unix() > c.AvailableEndTime.Unix() {
		return fmt.Errorf("活动开始时间不能大于结束时间")
	}
	return nil

}

// 使用营销活动和卷批次
func (c CreateMarketingCampaignRequest) ToModel() (*model.MarketingCampaign, []*model.CouponBatch) {
	var (
		mr  = new(model.MarketingCampaign)
		cbs []*model.CouponBatch
	)
	//营销活动基础活动信息
	mr.MarketingCampaignBaseInfo = c.MarketingCampaignBaseInfo

	//可能存在很多卷批次
	for _, item := range c.CouponBatches {
		var cb = new(model.CouponBatch)
		cb.CouponBatchBaseInfo = item.CouponBatchBaseInfo
		cb.Comment = item.Comment
		cbs = append(cbs, cb)
	}

	return mr, cbs
}
