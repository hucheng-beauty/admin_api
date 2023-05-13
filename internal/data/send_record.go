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

func (sr *sendRecordRepo) Save(record *model.SendRecord) (*model.SendRecord, error) {
	r := sr.db.Create(&record)
	if r.Error != nil {
		return nil, errors.Wrap(r.Error, "original error")
	}
	return record, nil
}

func (sr *sendRecordRepo) Update(record *model.SendRecord) (*model.SendRecord, error) {
	tx := sr.db.Model(model.SendRecord{})
	if record.Id != "" {
		tx = tx.Where("id = ?", record.Id)
	}
	if record.CampaignId != "" {
		tx = tx.Where("campaign_id = ?", record.CampaignId)
	}

	tx = tx.Updates(&record)
	if tx.Error != nil {
		return nil, errors.Wrap(tx.Error, "original error")
	}
	return record, nil
}

func (sr *sendRecordRepo) Detail(id string) (*model.SendRecord, error) {
	var record *model.SendRecord
	tx := sr.db.Model(model.SendRecord{}).First(&record, "id = ?", id)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(tx.Error, "original error")
	}
	return record, nil
}
