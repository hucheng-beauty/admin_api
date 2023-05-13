package model

import "time"

type CouponRecord struct {
	Id           string    `gorm:"primarykey;type:uuid;comment:主键Id"`
	ConsumerId   string    `gorm:"consumer_id;type:uuid;comment:用户Id"`
	AccountId    string    `gorm:"account_id;type:varchar;comment:账户Id"`
	CouponId     string    `gorm:"coupon_id;type:varchar;comment:券Id"`
	CampaignId   string    `gorm:"campaign_id;type:uuid;comment:活动Id"`
	BelongTo     string    `gorm:"belong_to;comment:归属平台,Alipay-支付宝,WeChat-微信"`
	CampaignName string    `gorm:"campaign_name;comment:活动名称"`
	StockId      string    `gorm:"stock_id;type:varchar;comment:券批次Id"`
	StockName    string    `gorm:"stock_name;type:varchar;comment:券批次名称"`
	CreatedAt    time.Time `gorm:"column:created_at;sort:desc;index:idx_created_at;comment:创建时间"`
}

type Coupon struct {
	BaseModel
	CouponId     string    `gorm:"coupon_id;comment:券Id"`
	CouponName   string    `gorm:"coupon_name;type:varchar;comment:券名称"`
	CouponType   string    `gorm:"coupon_type;comment:券类型;NORMAL-满减券,CUT_TO-减至券"`
	CouponAmount int       `gorm:"coupon_amount;comment:券面额;单位分"`
	CreateTime   time.Time `gorm:"create_time;comment:领取时间"`
	UsedTime     time.Time `gorm:"used_time;comment:使用时间"`
	Description  string    `gorm:"description;type:varchar;comment:使用说明"`
	Status       string    `gorm:"status;comment:券状态SENDED-已领取,USED-已使用,EXPIRED-已失效"`
}
