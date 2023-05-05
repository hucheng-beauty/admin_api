package response

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
