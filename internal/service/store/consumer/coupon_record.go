package consumer

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/pkg/uuid"
	"errors"
	"go.uber.org/zap"
	"time"
)

type CouponRecordRepo interface {
	Save(*model.CouponRecord) (*model.CouponRecord, error)
}

type CouponRecord struct {
	repo CouponRecordRepo
}

func NewCouponRecordService(repo CouponRecordRepo) *CouponRecord {
	return &CouponRecord{repo: repo}
}

func (cr *CouponRecord) CreateCouponRecord(req *request.CreateCouponRecord) (*response.CreateCouponRecord, error) {
	record, err := cr.repo.Save(&model.CouponRecord{
		Id:           uuid.New(),
		ConsumerId:   req.ConsumerId,
		AccountId:    req.AccountId,
		CouponId:     req.CouponId,
		CampaignId:   req.CampaignId,
		BelongTo:     req.BelongTo,
		CampaignName: req.CampaignName,
		StockId:      req.StockId,
		StockName:    req.StockName,
		CreatedAt:    time.Now().Local(),
	})
	if err != nil {
		zap.S().Errorf("%+#v\n", err)
		return nil, errors.New("新建领取记录失败")
	}
	return &response.CreateCouponRecord{Id: record.Id}, nil
}
