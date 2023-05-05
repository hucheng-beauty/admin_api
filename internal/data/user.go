package data

import (
	"admin_api/internal/model"
	"gorm.io/gorm"
	"time"
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

func (u *UserRepo) Detail(user *model.User) (*model.User, error) {
	var uu *model.User
	tx := u.db.Model(&uu)

	if user.BaseModel.Id != "" {
		tx = tx.Where("id = ?", user.BaseModel.Id)
	}
	if user.UserName != "" {
		tx = tx.Where("user_name = ?", user.UserName)
	}

	tx = tx.First(&uu)
	if tx.Error != nil {
		return nil, tx.Error
	}
	return uu, nil
}

func (u *UserRepo) UpdatePassword(user *model.User) (*model.User, error) {
	uu := &model.User{
		BaseModel: model.BaseModel{
			UpdatedAt: time.Now().Local(),
		},
		Password: user.Password,
	}
	//更新数据库中密码
	r := u.db.Where("id = ?", user.Id).Updates(&uu)
	if r.Error != nil {
		return nil, r.Error
	}
	return uu, nil
}
