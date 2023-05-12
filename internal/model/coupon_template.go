package model

import (
	"gorm.io/gorm"
	"time"
)

type BM struct {
	Id        string         `gorm:"primarykey;type:uuid;comment:主键id,uuid" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp with time zone;comment:创建时间,带时区" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp with time zone;comment:修改时间,带时区" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
}

type CouponTemplateBaseInfo struct {
	Platform int8 `json:"platform" gorm:"type:smallint;comment:所属平台枚举(1-支付宝、2-微信)"` // 所属平台，枚举(1-支付宝、2-微信)

	MerchantName string `json:"merchant_name" gorm:"type:text;comment:服务商名称"` // 服务商名称
	StockName    string `json:"stock_name" gorm:"type:text;comment:批次名称"`     // 批次名称

	CouponAmount       int64 `json:"coupon_amount" gorm:"type:integer;comment:面额"`             // 面额,单位分，单位分,整数;1元<面额<=500元
	TransactionMinimum int64 `json:"transaction_minimum" gorm:"type:integer;comment:使用限制(门槛)"` // 使用限制(门槛)，单位分,必须大于优惠金额
	Type               int64 `json:"type" gorm:"type:integer;comment:类型"`                      // 枚举(1-满减)
}

// gorm 获取表名通用设置，卷模板
type CouponTemplate struct {
	BM
	CouponTemplateBaseInfo
}

func (CouponTemplate) TableName() string {
	return "coupon_template"
}
