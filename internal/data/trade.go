package data

import "gorm.io/gorm"

type tradeRepo struct {
	db *gorm.DB
}

func NewTradeRepo(db *gorm.DB) *tradeRepo {
	return &tradeRepo{db: db}
}
