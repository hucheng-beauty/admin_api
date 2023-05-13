package data

import (
	"admin_api/internal/model"
	errors2 "github.com/pkg/errors"
	"gorm.io/gorm"
)

type couponBatchRepo struct {
	db *gorm.DB
}

func NewCouponBatchRepo(db *gorm.DB) *couponBatchRepo {
	return &couponBatchRepo{db: db}
}

// 创建卷批次
func (c *couponBatchRepo) InsertMany(cbs []*model.CouponBatch) error {
	if len(cbs) == 0 {
		return nil
	}
	return c.db.Create(cbs).Error
}

func (c *couponBatchRepo) List(batch *model.CouponBatch) ([]*model.CouponBatch, error) {
	var cbs []*model.CouponBatch
	r := c.db.Model(batch).Where(batch).Find(&cbs)
	return cbs, r.Error
}

// 查询卷批次
func (c *couponBatchRepo) ListByMarketingCampaignId(marketingCampaignId string) ([]*model.CouponBatch, error) {
	var cbs []*model.CouponBatch
	r := c.db.Where("marketing_campaign_id = ?", marketingCampaignId).Find(&cbs)
	if r.Error != nil {
		return nil, errors2.Wrap(r.Error, "original error")
	}
	return cbs, nil
}
