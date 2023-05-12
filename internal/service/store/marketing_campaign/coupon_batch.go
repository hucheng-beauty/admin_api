package marketing_campaign

import "admin_api/internal/model"

func (s *MarCampaignService) CouponBatchByMarId(marId string) ([]*model.CouponBatch, error) {
	return s.cbr.List(&model.CouponBatch{MarketingCampaignID: marId})
}
