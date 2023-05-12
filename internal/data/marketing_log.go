package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type marketingCampaignLogRepo struct {
	db *gorm.DB
}

func NewMarketingCampaignLogRepo(db *gorm.DB) *marketingCampaignLogRepo {
	return &marketingCampaignLogRepo{db: db}
}

func (d *marketingCampaignLogRepo) Create(log *model.MarketingCampaignLog) error {
	return d.db.Save(log).Error
}

func (d *marketingCampaignLogRepo) List(log *model.MarketingCampaignLog) ([]*model.MarketingCampaignLog, error) {
	var res []*model.MarketingCampaignLog
	err := d.db.Model(model.MarketingCampaignLog{}).Where(log).Order("created_at desc").Find(&res).Error
	return res, err
}
