package data

import "gorm.io/gorm"

type accountRepo struct {
	db *gorm.DB
}

func NewAccountRepo(db *gorm.DB) *accountRepo {
	return &accountRepo{db: db}
}
