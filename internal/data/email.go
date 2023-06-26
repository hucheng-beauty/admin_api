package data

import (
	"admin_api/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type EmailRepo struct {
	db *gorm.DB
}

func NewEmailRepo(db *gorm.DB) *EmailRepo {
	return &EmailRepo{db: db}
}

func (e EmailRepo) Create(emailID string) (ret bool) {
	var email *model.Email
	//使用数据库模型
	//fmt.Println("11111111", emailID)
	tx := e.db.Model(&email)

	value := &model.Email{
		Email: emailID,
	}

	tx = e.db.Create(value)
	if tx.Error != nil {
		fmt.Println("Create email send serial ID failed:", tx.Error)
		return false
	}
	fmt.Println("Create email send serial ID successfully")
	return true
}

func (e EmailRepo) A(emailID string) {
	fmt.Println(emailID)
}
