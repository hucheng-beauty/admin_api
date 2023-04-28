package model

import (
	"database/sql/driver"
	"encoding/json"
)

type User struct {
	BaseModel
	UserName      string `gorm:"user_name;type:varchar;comment:用户名,邮箱"`
	CompanyName   string `gorm:"company_name;type:varchar;comment:公司名称"`
	Password      string `gorm:"password;type:varchar;comment:密码"`
	ContactName   string `gorm:"contact_name;type:varchar;comment:联系人名称"`
	ContactMobile string `gorm:"contact_mobile;type:varchar;comment:联系人手机"`
	License       File   `gorm:"license;type:json;comment:营业执照"`
	Captcha       string `gorm:"captcha;type:varchar;comment:验证码"`
}

type File struct {
	Type string `json:"type"`
	URL  string `json:"url"`
}

func (j *File) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j File) Value() (driver.Value, error) {
	bs, err := json.Marshal(j)
	return bs, err
}
