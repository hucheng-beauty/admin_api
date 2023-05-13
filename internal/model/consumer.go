package model

type Consumer struct {
	BaseModel
	AccountId string `gorm:"account_id;type:varchar;comment:账户Id;手机、邮箱"`
	OpenId    string `gorm:"open_id;type:varchar;comment:用户openid"`
}
