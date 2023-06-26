package data

import (
	"admin_api/internal/model"
	"fmt"
	"gorm.io/gorm"
)

type CacheRepo struct {
	db *gorm.DB
}

func NewCacheRepo(db *gorm.DB) *CacheRepo {
	return &CacheRepo{db: db}
}

func (c CacheRepo) Create(name, val string) (ret bool) {
	var cache *model.Cache
	//使用数据库模型
	fmt.Println("11111111", name, val)
	tx := c.db.Model(&cache)

	value := &model.Cache{
		Name:  name,
		Value: val,
	}

	tx = c.db.Create(value)
	if tx.Error != nil {
		fmt.Println("Mysql Create value failed:", tx.Error)
		return false
	}
	fmt.Println("Mysql Create value successfully")
	return true
}

func (c CacheRepo) Find(name string) (*model.Cache, error) {
	var cache *model.Cache
	tx := c.db.Model(&cache)
	fmt.Println(0000000000000, name)
	tx.Where("name = ?", name)
	tx = tx.First(&cache)
	if tx.Error != nil {
		if tx.Error == gorm.ErrRecordNotFound {
			// 如果查询结果不存在，则返回空值和 nil 错误
			return nil, fmt.Errorf("数据库中没有该条记录")
		}
		return nil, tx.Error
	}

	return cache, nil
}
