package model

import (
	"database/sql/driver"
	"encoding/json"
)

type SendRecord struct {
	BaseModel
	CampaignId        string             `gorm:"campaign_id;type:uuid;comment:活动Id"`
	CampaignName      string             `gorm:"campaign_name;comment:活动名称"`
	SurplusCount      int                `gorm:"surplus_count;comment:每次发送之后活动剩余数量"`
	AccountIds        StrSlice           `gorm:"account_ids;comment:每次提交手机号的数量"`
	TotalCount        int                `gorm:"total_count;comment:单个活动下发送总数量"`
	TotalSuccessCount int                `gorm:"total_success_count;comment:单个活动下发送总成功数量"`
	StockSendInfo     StockSendRecordSli `gorm:"stock_send_info;type:json;comment:单个活动下所有券的发送情况"`
}

func (SendRecord) TableName() string {
	return "send_record"
}

type StockSendRecordSli []StockSendRecord

func (j *StockSendRecordSli) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j StockSendRecordSli) Value() (driver.Value, error) {
	bs, err := json.Marshal(j)
	return bs, err
}

type StockSendRecord struct {
	StockId      string `json:"stock_id"`
	Count        int    `json:"count"`
	SuccessCount int    `json:"success_count"`
}
