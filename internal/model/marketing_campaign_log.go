package model

type MarketingCampaignLog struct {
	BM
	MarketingCampaignId string `json:"marketing_campaign_id" gorm:"type:uuid;comment:关联活动uuid"`
	State               State  `json:"state" gorm:"type:smallint;comment:状态"`
}

func (m MarketingCampaignLog) TableName() string {
	return "marketing_campaign_log"
}
