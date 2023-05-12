package model

type TradeType int

const (
	Recharge = 1 //充值
	Freeze   = 2 //冻结
	Release  = 3 //释放
	Settle   = 4 //结算
)

type Trade struct {
	BM
	TradeId     string `json:"trade_id" gorm:"trade_id;comment:交易Id"`
	TradeType   string `json:"trade_type" gorm:"trade_type;comment:交易类型;Recharge-充值,Freeze-冻结,Release-释放,Settle-结算"`
	TradeAmount int    `json:"trade_amount" gorm:"trade_amount;comment:交易金额,单位分"`
	Remark      string `json:"remark" gorm:"remark;comment:备注"`
	CampaignId  string `json:"campaign_id" gorm:"campaign_id;type:varchar(36);comment:活动Id"`
	UserId      string `json:"user_id" gorm:"user_id;type:varchar(36);comment:用户Id"`
}

func (Trade) TableName() string {
	return "trade"
}
