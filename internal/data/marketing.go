package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type marketingCampaignRepo struct {
	db *gorm.DB
}

func NewMarketingCampaignRepo(db *gorm.DB) *marketingCampaignRepo {
	return &marketingCampaignRepo{db: db}
}

// 保存营销活动
func (d *marketingCampaignRepo) Create(campagin *model.MarketingCampaign) error {
	return d.db.Save(campagin).Error
}

func (d *marketingCampaignRepo) FindById(id string) (*model.MarketingCampaign, error) {
	var m model.MarketingCampaign
	r := d.db.Model(model.MarketingCampaign{}).Where("id= ?", id).First(&m)
	if r.RowsAffected == 0 {
		return nil, nil
	}
	return &m, nil
}

func (d *marketingCampaignRepo) UpdateStateById(id string, state model.State) error {
	return d.db.Model(model.MarketingCampaign{}).Where("id= ?", id).Update("state", state).Error
}

func (d *marketingCampaignRepo) UpdateSurplusNumberById(id string, successCount int) error {
	r := d.db.Model(model.MarketingCampaign{}).Where("id = ?", id).Update("coupon_surplus_number", gorm.Expr("coupon_surplus_number - ?", successCount))

	return r.Error

}
