package marketing_campaign

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
)

type MarCampaignRepo interface {
	Create(*model.MarketingCampaign) error
	FindById(id string) (*model.MarketingCampaign, error)
	FilterWithPage(mr *model.MarketingCampaign, query *request.Query) ([]*model.MarketingCampaign, int, error)
	UpdateStateById(id string, state model.State) error
	UpdateSurplusNumberById(id string, successCount int) error
}

type CouponBatchRepo interface {
	InsertMany([]*model.CouponBatch) error
	List(batch *model.CouponBatch) ([]*model.CouponBatch, error)
}

type CouponTemplateRepo interface {
	Create(ct *model.CouponTemplate) error
	FindByIds(ids []string) ([]*model.CouponTemplate, error)
}

type CouponLogRepo interface {
	Create([]*model.CouponLog) error
	List(*model.CouponLog) ([]*model.CouponLog, error)
}

type MarCampaignLogRepo interface {
	Create(log *model.MarketingCampaignLog) error
	List(log *model.MarketingCampaignLog) ([]*model.MarketingCampaignLog, error)
}

type MarCampaignService struct {
	mcr  MarCampaignRepo
	cbr  CouponBatchRepo
	ctr  CouponTemplateRepo
	clr  CouponLogRepo
	mclr MarCampaignLogRepo
}

func NewMarCampaignService(r1 MarCampaignRepo, r2 CouponBatchRepo, r3 CouponTemplateRepo, r4 CouponLogRepo, r5 MarCampaignLogRepo) *MarCampaignService {
	return &MarCampaignService{
		mcr:  r1,
		cbr:  r2,
		ctr:  r3,
		clr:  r4,
		mclr: r5,
	}
}
