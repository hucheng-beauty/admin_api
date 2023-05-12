package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type couponLogRepo struct {
	db *gorm.DB
}

func NewCouponLogRepo(db *gorm.DB) *couponLogRepo {
	return &couponLogRepo{db: db}
}

// 创建卷批次日志
func (d *couponLogRepo) Create(ct []*model.CouponLog) error {
	return d.db.Model(model.CouponLog{}).Create(ct).Error
}
