package response

import (
	"admin_api/internal/model"
	"time"
)

type Id struct {
	Id string `json:"id"`
}

type MarketingCampaignListResponse struct {
	TotalCount int `json:"total_count"`

	Data []*model.MarketingCampaign `json:"data"`
}

type MarketingCampaignResponse struct {
	*model.MarketingCampaign
	CouponBatches []*CouponBatchResponse `json:"coupon_batches"`
	Logs          []Log                  `json:"logs"`
	Trades        []*DescribeUserTrade   `json:"trades"`
}

type Log struct {
	State     model.State `json:"state"`
	CreatedAt time.Time   `json:"created_at"`
}

type CouponBatchResponse struct {
	*model.CouponBatch
	Template *model.CouponTemplate `json:"template"`
	Logs     []Log                 `json:"logs"`
}

func (m *MarketingCampaignResponse) Model2Resp(mc *model.MarketingCampaign, cbs []*model.CouponBatch, tMap map[string]*model.CouponTemplate, logMap map[string][]*model.CouponLog, clogs []*model.MarketingCampaignLog) {
	m.MarketingCampaign = mc
	for _, item := range cbs {
		c := &CouponBatchResponse{
			CouponBatch: item,
			Template:    tMap[item.TemplateID],
		}
		c.Logs2Resp(logMap[item.Id])
		m.CouponBatches = append(m.CouponBatches, c)
	}
	for _, item := range clogs {
		m.Logs = append(m.Logs, Log{
			State:     item.State,
			CreatedAt: item.CreatedAt,
		})
	}

}

func (c *CouponBatchResponse) Logs2Resp(logs []*model.CouponLog) {
	for _, item := range logs {
		c.Logs = append(c.Logs, Log{
			State:     item.State,
			CreatedAt: item.CreatedAt,
		})
	}
}

func (m *MarketingCampaignResponse) SetTradeModel2Resp(mc []*model.Trade) {
	for _, item := range mc {
		m.Trades = append(m.Trades, new(DescribeUserTrade).ModelToResp(item))
	}
}

type CouponLogsResponse struct {
	Data       []*CouponLogRsp `json:"data"`
	TotalCount int             `json:"total_count"`
}

type CouponLogRsp struct {
	AvailableBeginTime time.Time `json:"available_begin_time"`
	AvailableEndTime   time.Time `json:"available_end_time"`
	MerchantName       string    `json:"merchant_name"`
	StockName          string    `json:"stock_name"`
	CouponAmount       int       `json:"coupon_amount"`
	TransactionMinimum int       `json:"transaction_minimum"`
	Type               int       `json:"type"`
	Platform           int       `json:"platform"`
	LimitPay           int       `json:"limit_pay"`
	MaxAmount          int       `json:"max_amount"`
	MaxCoupons         int       `json:"max_coupons"`
	MaxCouponsPerUser  int       `json:"max_coupons_per_user"`
	State              int       `json:"state"`
	CreatedAt          time.Time `json:"created_at"`
}
