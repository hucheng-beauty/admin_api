package request

// 创建用户
type CreateUser struct {
	Username      string `json:"username" binding:"required,email"`
	Password      string `json:"password" binding:"required"`
	CompanyName   string `json:"company_name" binding:"required"`
	ContactName   string `json:"contact_name" binding:"required"`
	ContactMobile string `json:"contact_mobile" binding:"required"`
	License       File   `json:"license" binding:"required"`
}
