package model

import "time"

type MarketingCampaignBaseInfo struct {
	AvailableBeginTime time.Time `json:"available_begin_time" gorm:"type:timestamp with time zone;comment:开始时间,带时区"`
	AvailableEndTime   time.Time `json:"available_end_time" gorm:"type:timestamp with time zone;comment:结束时间,带时区"`
	CampaignName       string    `json:"campaign_name" gorm:"type:varchar(255);comment:活动名称"`
	CouponBatchNumber  int64     `json:"coupon_batch_number"  gorm:"type:integer;comment:活动券批次数量"`

	CouponNumber        int64 `json:"coupon_number"  gorm:"type:integer;comment:活动券数量"`
	CouponSurplusNumber int64 `json:"coupon_surplus_number,omitempty" gorm:"type:integer;comment:活动剩余数量"`
	FreezeAmount        int64 `json:"freeze_amount,omitempty" gorm:"type:integer;comment:冻结金额冻"`
	MaxAmount           int64 `json:"max_amount" gorm:"type:integer;comment:总预算"`
}

type State int

const (
	Stop        State = -1 // 暂停
	Submit      State = 1  // 已提交
	NoEffective State = 3  // 未生效(已激活)
	Effective   State = 4  // 已生效
	Expire      State = 5  // 已失效
	Fail        State = 6  // -失败
)

type MarketingCampaign struct {
	BM
	MarketingCampaignBaseInfo

	CampaignNu string `json:"campaign_nu" gorm:"type:varchar(16);comment:活动编号"` // 活动编号
	State      State  `json:"state" gorm:"type:smallint;comment:状态"`            // 状态-1-暂停 1-已提交,2-已创建,3-未生效,4-已生效,5-已失效,6-失败
	IsEnd      int    `json:"-" gorm:"type:smallint;comment:是否结束;default:0"`    // 1-> 是
	UserId     string `json:"-" gorm:"type:varchar(36);comment:账户id"`
}

func (m MarketingCampaign) TableName() string {
	return "marketing_campaign"
}

func (m MarketingCampaign) CanAction() bool {
	if m.IsEnd == 1 {
		return false
	}
	switch m.State {
	case NoEffective, Effective:
		return true
	}
	return false
}
