package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type couponRepo struct {
	db *gorm.DB
}

func NewCouponRepo(db *gorm.DB) *couponRepo {
	return &couponRepo{db: db}
}
func (c *couponRepo) Save(record *model.Coupon) (*model.Coupon, error) {
	r := c.db.Create(&record)
	if r.Error != nil {
		return nil, r.Error
	}
	return record, nil
}
