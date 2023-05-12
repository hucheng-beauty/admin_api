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

// 券批次日志Map {"template_id": template}
func (s *MarCampaignService) CouponLogsByMarCampaignId2CouponBatchIdMap(marCampaignId string) (map[string][]*model.CouponLog, error) {

	var mp = map[string][]*model.CouponLog{}

	res, err := s.CouponLogsByMarCampaignId(marCampaignId)
	if err != nil {
		return nil, err
	}
	for _, item := range res {
		mp[item.CouponBatchId] = append(mp[item.CouponBatchId], item)
	}

	return mp, nil
}

func (s *MarCampaignService) CouponLogsByMarCampaignId(marCampaignId string) ([]*model.CouponLog, error) {
	return s.clr.List(&model.CouponLog{MarketingCampaignId: marCampaignId})
}
