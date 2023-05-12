package data

import (
	"admin_api/internal/model"
	"github.com/pkg/errors"
	"gorm.io/gorm"
)

type sendRecordRepo struct {
	db *gorm.DB
}

func NewSendRecordRepo(db *gorm.DB) *sendRecordRepo {
	return &sendRecordRepo{db: db}
}

// 查询发送记录
func (sr *sendRecordRepo) ListWithPagination(p *model.Pagination) ([]*model.SendRecord, int64, error) {
	var total int64
	tx := sr.db.Model(model.SendRecord{})
	tx = tx.Count(&total)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, -1, errors.Wrap(tx.Error, "original error")
	}

	if p != nil {
		if p.Limit >= 0 {
			tx = tx.Limit(p.Limit)
		}
		if p.Offset >= 0 {
			tx = tx.Offset(p.Offset)
		}
	}

	var records []*model.SendRecord
	tx = tx.Order("created_at desc").Find(&records)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, 0, nil
		}
		return nil, -1, errors.Wrap(tx.Error, "original error")
	}
	return records, total, nil
}
