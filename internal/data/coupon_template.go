package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type couponTemplateRepo struct {
	db *gorm.DB
}

func NewCouponTemplateRepo(db *gorm.DB) *couponTemplateRepo {
	return &couponTemplateRepo{db: db}
}

// 查询卷模板
func (d *couponTemplateRepo) FindByIds(ids []string) ([]*model.CouponTemplate, error) {
	if len(ids) == 0 {
		return nil, nil
	}

	var res []*model.CouponTemplate
	err := d.db.Model(model.CouponTemplate{}).Where("id IN ?", ids).Find(&res).Error

	return res, err
}

func (d *couponTemplateRepo) Create(ct *model.CouponTemplate) error {
	return d.db.Save(ct).Error
}
