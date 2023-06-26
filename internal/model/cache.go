package model

type Cache struct {
	ID    int    `gorm:"user_id;primaryKey;autoIncrement;type:varchar;comment:记录号"`
	Name  string `gorm:"name;type:varchar;comment:key"`
	Value string `gorm:"value;type:varchar;comment:缓存数据"`
}
