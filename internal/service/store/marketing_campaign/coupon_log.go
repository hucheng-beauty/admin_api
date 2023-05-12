package marketing_campaign

import "admin_api/internal/model"

func (s *MarCampaignService) CreateCouponLogWithMarCampaign(campaign *model.MarketingCampaign) error {
	cbs, err := s.CouponBatchByMarId(campaign.Id)
	if err != nil {
		return err
	}
	return s.CreateCouponLogWithModel(campaign.Id, cbs, campaign.State)
}

// ///
func (s *MarCampaignService) CreateCouponLogWithModel(marCampaignId string, couponBatches []*model.CouponBatch, state model.State) error {
	couponBatchIds := make([]string, 0, len(couponBatches))
	for _, item := range couponBatches {
		couponBatchIds = append(couponBatchIds, item.Id)
	}
	return s.CreateCouponLog(marCampaignId, state, couponBatchIds...)
}

// 创建卷批次日志
func (s *MarCampaignService) CreateCouponLog(marCampaignId string, state model.State, couponBatchIds ...string) error {
	var cts = make([]*model.CouponLog, 0, len(couponBatchIds))
	for _, item := range couponBatchIds {
		var m = &model.CouponLog{}
		m.BM = NewBM()
		m.MarketingCampaignId = marCampaignId
		m.CouponBatchId = item
		m.State = state
		cts = append(cts, m)
	}

	return s.clr.Create(cts)
}
