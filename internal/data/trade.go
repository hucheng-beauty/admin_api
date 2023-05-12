package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type tradeRepo struct {
	db *gorm.DB
}

func NewTradeRepo(db *gorm.DB) *tradeRepo {
	return &tradeRepo{db: db}
}

func (r *tradeRepo) List(query *model.Trade) ([]*model.Trade, error) {
	var res []*model.Trade
	err := r.db.Model(model.Trade{}).Where(query).Order("created_at desc").Find(&res).Error
	return res, err
}
