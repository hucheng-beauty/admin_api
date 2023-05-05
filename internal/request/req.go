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
