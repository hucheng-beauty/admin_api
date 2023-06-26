package data

import (
	"admin_api/internal/model"
	"admin_api/internal/request"
	"admin_api/internal/response"
	"admin_api/utils"
	"gorm.io/gorm"
)

type couponLogRepo struct {
	db *gorm.DB
}

func NewCouponLogRepo(db *gorm.DB) *couponLogRepo {
	return &couponLogRepo{db: db}
}

// 创建卷批次日志
func (d *couponLogRepo) Create(ct []*model.CouponLog) error {
	return d.db.Model(model.CouponLog{}).Create(ct).Error
}

func (d *couponLogRepo) List(c *model.CouponLog) ([]*model.CouponLog, error) {
	var cbs []*model.CouponLog
	r := d.db.Model(c).Where(c).Order("created_at desc").Find(&cbs)
	return cbs, r.Error
}

func (d *couponLogRepo) ListWitPage(mrID string, query *request.Query) ([]*response.CouponLogRsp, int, error) {
	//select available_begin_time,
	//	available_end_time,
	//	merchant_name,
	//	stock_name,
	//	coupon_amount,
	//	transaction_minimum,
	//	type,
	//	platform,
	//		limit_pay,
	//		coupon_batch.max_amount,
	//		coupon_batch.max_coupons,
	//		max_coupons_per_user,
	//		coupon_log.state,
	//		coupon_log.created_at
	//from coupon_batch join  coupon_log on coupon_batch.id = coupon_log.coupon_batch_id join marketing_campaign mc on coupon_batch.marketing_campaign_id = mc.id;

	filter := utils.Filter(d.db, query.Filter)
	db := d.db.Model(model.CouponLog{}).Select(
		"available_begin_time",
		"available_end_time",
		"merchant_name",
		"stock_name",
		"coupon_amount",
		"transaction_minimum",
		"type",
		"platform",
		"limit_pay",
		"cb.max_amount",
		"cb.max_coupons",
		"max_coupons_per_user",
		"coupon_log.state",
		"coupon_log.created_at",
	).Joins("join  coupon_batch as cb on cb.id = coupon_log.coupon_batch_id").
		Joins("join marketing_campaign mc on cb.marketing_campaign_id = mc.id").
		Where("cb.marketing_campaign_id = ?", mrID).Where(filter)

	var count int64
	err := db.Count(&count).Error
	if err != nil || count == 0 {
		return nil, 0, nil
	}

	var cbs []*response.CouponLogRsp
	err = db.Offset(query.Offset).Limit(query.Limit).Order("created_at desc").Find(&cbs).Error
	return cbs, int(count), err
}
