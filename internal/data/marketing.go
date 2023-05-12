package data

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/utils"
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

func (d *marketingCampaignRepo) FilterWithPage(mr *model.MarketingCampaign, query *request.Query) ([]*model.MarketingCampaign, int, error) {
	var count int64
	filter := utils.Filter(d.db, query.Filter)
	err := d.db.Model(model.MarketingCampaign{}).Where(mr).Where(filter).Count(&count).Error
	if err != nil || count == 0 {
		return nil, 0, err
	}
	var res []*model.MarketingCampaign
	err = d.db.Model(model.MarketingCampaign{}).Where(mr).Where(filter).Offset(query.Offset).Limit(query.Limit).Order("created_at desc").Find(&res).Error

	return res, int(count), err
}
