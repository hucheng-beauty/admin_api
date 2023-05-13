package response

import "admin_api/internal/model"

type CreateUser struct {
	Id string `json:"id"`
}

type GetUserByUserName struct {
	Id            string `json:"id"`
	UserName      string `json:"username" `
	Password      string `json:"password"`
	CompanyName   string `json:"account_name" `
	ContactName   string `json:"contact_name" `
	ContactMobile string `json:"contact_mobile" `
	License       File   `json:"license"`
	Industry      string `json:"industry"`
	Subject       string `json:"subject"`
	Captcha       string `json:"captcha"`
	CreatedAt     string `json:"created_at" `
}

type GetUserById struct {
	Id            string `json:"id"`
	UserName      string `json:"username" `
	Password      string `json:"password"`
	CompanyName   string `json:"account_name" `
	ContactName   string `json:"contact_name" `
	ContactMobile string `json:"contact_mobile" `
	License       File   `json:"license"`
	Industry      string `json:"industry"`
	Subject       string `json:"subject"`
	Captcha       string `json:"captcha"`
	CreatedAt     string `json:"created_at" `
}

type GetUserDetail struct {
	Id            string `json:"id"`
	UserName      string `json:"username" `
	CompanyName   string `json:"company_name" `
	ContactName   string `json:"contact_name" `
	ContactMobile string `json:"contact_mobile" `
	License       File   `json:"license"`
	Industry      string `json:"industry"`
	Subject       string `json:"subject"`
	CreatedAt     string `json:"created_at" `
}

type GetUserAmount struct {
	UserId        string `json:"user_id"`
	Amount        int    `json:"amount"`
	UsableBalance int    `json:"usable_balance"`
	FrozenAmount  int    `json:"frozen_amount"`
}

type GetAmountId struct {
	Userid string `json:"Amount_id"`
}

type DescribeUserTrade struct {
	TradeId     string `json:"trade_id"`
	TradeType   string `json:"trade_type"`
	TradeAmount int    `json:"trade_amount"`
	CreatedAt   string `json:"created_at"`
	Remark      string `json:"remark"`
}

func (d *DescribeUserTrade) ModelToResp(trade *model.Trade) *DescribeUserTrade {
	c := d
	if c == nil {
		c = new(DescribeUserTrade)
	}
	c.TradeId = trade.TradeId
	c.TradeType = trade.TradeType
	c.TradeAmount = trade.TradeAmount
	c.Remark = trade.Remark
	c.CreatedAt = trade.CreatedAt.Local().Format("2006-01-02 15:04:05")
	return c
}

type DescribeSendRecord struct {
	Id                string `json:"id"`
	CampaignId        string `json:"campaign_id"`
	CampaignName      string `json:"campaign_name"`
	SurplusCount      int    `json:"surplus_count"`
	TotalCount        int    `json:"count"` // the front end is difficult to communicate
	TotalSuccessCount int    `json:"success_count"`
	TotalFailCount    int    `json:"fail_count"`
	CreatedAt         string `json:"created_at"`
}

type DescribeConsumer struct {
	ConsumerId string `json:"consumer_id"`
	AccountId  string `json:"account_id"`
	OpenId     string `json:"open_id"`
}

type DescribeCouponBatch struct {
	MarketingCampaignID string `json:"marketing_campaign_id"`
	StockId             string `json:"stock_id"`
	StockName           string `json:"stock_name"`
	BelongTo            string `json:"belong_to"`
}

type GetMarketingCampaignDetail struct {
	CampaignId          string `json:"campaign_id"`
	CreatedAt           string `json:"created_at"`
	UpdatedAt           string `json:"updated_at"`
	AvailableBeginTime  string `json:"available_begin_time"`
	AvailableEndTime    string `json:"available_end_time"`
	CampaignName        string `json:"campaign_name"`
	CouponBatchNumber   int64  `json:"coupon_batch_number"`
	CouponNumber        int64  `json:"coupon_number"`
	CouponSurplusNumber int64  `json:"coupon_surplus_number"`
	FreezeAmount        int64  `json:"freeze_amount"`
	MaxAmount           int64  `json:"max_amount"`
	CampaignNu          string `json:"campaign_nu"`
	State               string `json:"state"`
}

type CreateSendRecord struct {
	Id string `json:"id"`
}

type CreateCouponRecord struct {
	Id string `json:"id"`
}

type CreateCoupon struct {
	Id       string `json:"id"`
	CouponId string `json:"coupon_id"`
}

type UpdateSendRecord struct {
	Id                string `json:"id"`
	CampaignId        string `json:"campaign_id"`
	CampaignName      string `json:"campaign_name"`
	TotalCount        int    `json:"total_count"`
	TotalSuccessCount int    `json:"total_success_count"`
	CreatedAt         string `json:"created_at"`
}
