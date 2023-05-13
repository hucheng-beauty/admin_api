package data

import (
	"admin_api/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type consumerRepo struct {
	db *gorm.DB
}

func NewConsumerRepo(db *gorm.DB) *consumerRepo {
	return &consumerRepo{db: db}
}

func (c *consumerRepo) ListByAccountIds(accountIds []string) ([]*model.Consumer, int64, error) {
	var total int64
	var consumers []*model.Consumer
	tx := c.db.Model(&consumers)
	if len(accountIds) > 0 {
		tx = tx.Where("account_id IN ?", accountIds)
	}
	tx = tx.Count(&total)
	if total == 0 {
		return []*model.Consumer{}, 0, nil
	}
	tx = tx.Find(&consumers)
	if tx.Error != nil {
		return nil, -1, errors.Wrap(tx.Error, "original error")
	}
	return consumers, total, nil
}
