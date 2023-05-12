package model

type CouponLog struct {
	BM

	MarketingCampaignId string `json:"marketing_campaign_id" gorm:"type:uuid;comment:关联活动uuid"`
	CouponBatchId       string `json:"coupon_batch_id" gorm:"type:uuid;comment:批次id"`

	State State `json:"state" gorm:"type:smallint;comment:状态"`
}

func (CouponLog) TableName() string {
	return "coupon_log"
}
