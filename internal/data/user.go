package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
)

type UserRepo struct {
	db *gorm.DB
}

func NewUserRepo(db *gorm.DB) *UserRepo {
	return &UserRepo{db: db}
}

func (u *UserRepo) IsExist(id string, username string) bool {
	var (
		total int64
		uu    *model.User
	)
	//使用数据库模型
	tx := u.db.Model(&uu)
	if id != "" {
		tx = tx.Where("id = ?", id)
	}
	if username != "" {
		tx = tx.Where("user_name = ?", username)
	}
	//计算select count(*)后的条数==0不存在
	tx = tx.Count(&total)
	if tx.Error != nil || tx.RowsAffected == 0 {
		return false
	}
	if total == 0 {
		return false
	}
	return true
}

func (u *UserRepo) Save(user *model.User) (*model.User, error) {
	r := u.db.Create(&user)
	return user, r.Error
}
