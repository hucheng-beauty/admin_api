package data

import (
	"admin_api/internal/model"
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
