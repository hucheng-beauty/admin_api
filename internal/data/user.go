package data

import (
	"admin_api/internal/model"
	"github.com/pkg/errors"
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

type Wallet struct {
	db *gorm.DB
}

func NewWalletService(db *gorm.DB) *Wallet {
	return &Wallet{db: db}
}

func (w *Wallet) Detail(wallet *model.Wallet) (*model.Wallet, error) {
	var ww *model.Wallet

	tx := w.db.Model(model.Wallet{}).Where("user_id = ?", wallet.UserId).First(&ww)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			return nil, nil
		}
		return nil, errors.Wrap(tx.Error, "original error")
	}
	return ww, nil
}

func (w *Wallet) Save(wallet *model.Wallet) (*model.Wallet, error) {
	r := w.db.Create(&wallet)
	return wallet, r.Error
}

func (w *Wallet) CheckId(id string) (bool, error) {
	var a int64
	w.db.Model(model.Wallet{}).Where("user_id=?", id).Count(&a)
	if a == 0 {
		return false, nil
	}
	return true, nil
}

func (w *Wallet) Update(id string) (*model.Wallet, error) {
	var ww *model.Wallet
	w.db.Model(model.Wallet{}).Where("user_id=?", id).Update("amount", gorm.Expr("amount + ?", 1000000))
	w.db.Model(model.Wallet{}).Where("user_id=?", id).First(&ww)
	return ww, nil
}
