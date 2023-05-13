package marketing_campaign

import (
	"admin_api/internal/enum"
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"errors"
	"go.uber.org/zap"
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
	ListByMarketingCampaignId(string) ([]*model.CouponBatch, error)
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

type MarketingCampaignRepo interface {
	GetById(string) (*model.MarketingCampaign, error)
}

type MarketingCampaign struct {
	repo MarketingCampaignRepo
}

func NewMarketingCampaignService(repo MarketingCampaignRepo) *MarketingCampaign {
	return &MarketingCampaign{repo: repo}
}

// 获取活动信息
func (mc *MarketingCampaign) GetMarketingCampaignDetail(req *request.GetMarketingCampaignDetail) (*response.GetMarketingCampaignDetail, error) {
	//获取活动id
	detail, err := mc.repo.GetById(req.Id)
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("获取营销活动详情失败")
	}
	if detail == nil {
		return nil, errors.New("营销活动不存在")
	}
	return &response.GetMarketingCampaignDetail{
		CampaignId:          detail.Id,
		CreatedAt:           detail.CreatedAt.Local().Format("2006-01-02 15:04:05"),
		UpdatedAt:           detail.UpdatedAt.Local().Format("2006-01-02 15:04:05"),
		AvailableBeginTime:  detail.AvailableBeginTime.Local().Format("2006-01-02 15:04:05"),
		AvailableEndTime:    detail.AvailableEndTime.Local().Format("2006-01-02 15:04:05"),
		CampaignName:        detail.CampaignName,
		CouponBatchNumber:   detail.CouponBatchNumber,
		CouponNumber:        detail.CouponNumber,
		CouponSurplusNumber: detail.CouponSurplusNumber,
		FreezeAmount:        detail.FreezeAmount,
		MaxAmount:           detail.MaxAmount,
		CampaignNu:          detail.CampaignNu,
		State:               enum.StatusMapResponse[int(detail.State)],
	}, nil
}

type CouponBatch struct {
	repo CouponBatchRepo
}

func NewCouponBatchService(repo CouponBatchRepo) *CouponBatch {
	return &CouponBatch{repo: repo}
}

func (cb *CouponBatch) DescribeCouponBatch(req *request.DescribeCouponBatch) ([]*response.DescribeCouponBatch, error) {
	//查询卷批次
	batch, err := cb.repo.ListByMarketingCampaignId(req.MarketingCampaignId)
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("获取券批次列表失败")
	}

	var resp []*response.DescribeCouponBatch
	for _, b := range batch {
		resp = append(resp, &response.DescribeCouponBatch{
			MarketingCampaignID: b.MarketingCampaignID,
			StockId:             b.StockId,
			StockName:           b.StockName,
			BelongTo:            enum.BelongToMapResponse[int(b.Platform)],
		})
	}
	return resp, nil
}
