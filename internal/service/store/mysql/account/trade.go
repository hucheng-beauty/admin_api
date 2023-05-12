package account

import "admin_api/internal/model"

type TradeRepo interface {
	List(query *model.Trade) ([]*model.Trade, error)
}

type TradeService struct {
	repo TradeRepo
}

func NewTradeService(repo TradeRepo) *TradeService {
	return &TradeService{repo: repo}
}

func (t TradeService) TradesByCampaignId(campaignId string) ([]*model.Trade, error) {
	return t.repo.List(&model.Trade{CampaignId: campaignId})
}
