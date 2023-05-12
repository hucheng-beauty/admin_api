package marketing_campaign

import "admin_api/internal/model"

func (s *MarCampaignService) CreateMarketingCampaignLogWithModel(m *model.MarketingCampaign) error {
	var log = &model.MarketingCampaignLog{}
	log.BM = NewBM()
	log.MarketingCampaignId = m.Id
	log.State = m.State
	//保存记录
	return s.mclr.Create(log)
}

// 活动状态日志
func (s *MarCampaignService) MarketingCampaignLogsByMarCampaignId(marCampaignId string) ([]*model.MarketingCampaignLog, error) {

	return s.mclr.List(&model.MarketingCampaignLog{MarketingCampaignId: marCampaignId})
}
