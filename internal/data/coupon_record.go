package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type couponRecordRepo struct {
	db *gorm.DB
}

func NewCouponRecordRepo(db *gorm.DB) *couponRecordRepo {
	return &couponRecordRepo{db: db}
}

func (c *couponRecordRepo) Save(record *model.CouponRecord) (*model.CouponRecord, error) {
	r := c.db.Create(&record)
	if r.Error != nil {
		return nil, r.Error
	}
	return record, nil
}
