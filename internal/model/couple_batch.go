package model

type CouponBatch struct {
	BM
	CouponTemplateBaseInfo `json:"-"`
	CouponBatchBaseInfo

	Comment             string `json:"comment" gorm:"type:text;comment:备注"` // 批次备注
	MarketingCampaignID string `json:"marketing_campaign_id" gorm:"type:uuid;comment:关联活动uuid"`
	StockId             string `json:"stock_id"  gorm:"type:varchar;comment:第三方批次Id"` // 批次Id
}

func (CouponBatch) TableName() string {
	return "coupon_batch"
}

type CouponBatchBaseInfo struct {
	Bin               StrSlice `gorm:"type:json;comment:银行卡json数组" json:"bin"`
	LimitPay          int64    `json:"limit_pay" gorm:"type:smallint;comment:支付方式"`
	MaxAmount         int64    `json:"max_amount" gorm:"type:integer;comment:总预算"`
	MaxAmountByDay    int64    `json:"max_amount_by_day" gorm:"type:integer;comment:单天预算发放上限"`
	MaxCoupons        int64    `json:"max_coupons" gorm:"type:integer;comment:发放总个数(上限)"`
	MaxCouponsPerUser int64    `json:"max_coupons_per_user" gorm:"type:integer;comment:单个用户可领个数"`
	TemplateID        string   `json:"template_id" gorm:"type:uuid;comment:关联券模板uuid"`
}
