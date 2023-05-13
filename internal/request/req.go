package request

import "time"

// 创建用户
type CreateUser struct {
	Username      string `json:"username" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	CompanyName   string `json:"company_name" binding:"required"`
	ContactName   string `json:"contact_name" binding:"required"`
	ContactMobile string `json:"contact_mobile" binding:"required"`
	License       File   `json:"license" binding:"required"`
}

type PasswordLogin struct {
	UserName string `json:"username" binding:"required,email"`
	Password string `json:"password" binding:"required"`
}

type GetUserByUserName struct {
	UserName string `json:"username"`
}

type CheckPassword struct {
	Password          string `json:"password"`
	EncryptedPassword string `json:"encrypted_password"`
}

type GetUserDetail struct {
	Id       string `json:"id"`
	UserName string `json:"username"`
}

type GetUserById struct {
	Id string `json:"id"`
}

type UpdatePassword struct {
	Password string `json:"password" binding:"required"`
	Id       string `json:"id"`
}

// 具体查询页码和判断条件
type DescribeUserTrade struct {
	Pagination
	FilterString
	UserId string
}

// FilterString example: `name eq 'hello' and age eq 18`
type FilterString struct {
	Filter string `json:"filter"`
}

type DescribeSendRecord struct {
	Pagination
}

type CreateSendRecord struct {
	AccountIds          []string          `json:"account_ids" binding:"required"`
	CampaignId          string            `json:"campaign_id"`
	CampaignName        string            `json:"campaign_name"`
	StockSendRecordInfo []StockSendRecord `json:"stock_send_record_info"`
	TotalCount          int               `json:"total_count"`
}

type StockSendRecord struct {
	Count        int    `json:"count"`
	SuccessCount int    `json:"success_count"`
	StockId      string `json:"stock_id"`
}

type GetMarketingCampaignDetail struct {
	Id string `json:"id"`
}

type DescribeCouponBatch struct {
	MarketingCampaignId string `json:"marketing_campaign_id"`
}

type DescribeConsumer struct {
	AccountIds []string `json:"account_id"`
}

type CreateCouponRecord struct {
	ConsumerId   string `json:"consumer_id"`
	CouponId     string `json:"coupon_id"`
	StockId      string `json:"stock_id"`
	StockName    string `json:"stock_name"`
	CampaignId   string `json:"campaign_id"`
	CampaignName string `json:"campaign_name"`
	AccountId    string `json:"account_id"`
	BelongTo     string `json:"belong_to"`
}

type CreateCoupon struct {
	CouponId     string    `json:"coupon_id"`
	CouponName   string    `json:"coupon_name"`
	CouponType   string    `json:"coupon_type"`
	CouponAmount int       `json:"coupon_amount"`
	CreateTime   time.Time `json:"create_time"`
	Description  string    `json:"description"`
	Status       string    `json:"status"`
}

type UpdateSendRecord struct {
	Id                  string             `json:"id"`
	CampaignId          string             `json:"campaign_id"`
	SurplusCount        int                `json:"surplus_count"`
	TotalCount          int                `json:"total_count"`
	TotalSuccessCount   int                `json:"total_success_count"`
	StockSendRecordInfo []*StockSendRecord `json:"stock_send_record_info"`
}
