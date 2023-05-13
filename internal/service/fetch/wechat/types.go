package wechat

import "time"

type SendResp struct {
	CouponId string `json:"coupon_id"`
}

type SendReq struct {
	StockId           string `json:"stock_id"`
	OutRequestNo      string `json:"out_request_no"`
	AppId             string `json:"appid"`
	StockCreatorMchid string `json:"stock_creator_mchid"`
}

type GetCouponDetailResp struct {
	StockCreatorMchid string `json:"stock_creator_mchid"`
	StockId           string `json:"stock_id"`
	CouponId          string `json:"coupon_id"`
	CutToMessage      struct {
		SinglePriceMax int `json:"single_price_max"`
		CutToPrice     int `json:"cut_to_price"`
	} `json:"cut_to_message"`
	CouponName              string    `json:"coupon_name"`
	Status                  string    `json:"status"`
	Description             string    `json:"description"`
	CreateTime              time.Time `json:"create_time"`
	CouponType              string    `json:"coupon_type"`
	NoCash                  bool      `json:"no_cash"`
	AvailableBeginTime      time.Time `json:"available_begin_time"`
	AvailableEndTime        time.Time `json:"available_end_time"`
	Singleitem              bool      `json:"singleitem"`
	NormalCouponInformation struct {
		CouponAmount       int `json:"coupon_amount"`
		TransactionMinimum int `json:"transaction_minimum"`
	} `json:"normal_coupon_information"`
}
