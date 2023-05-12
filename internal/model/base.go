package model

import (
	"database/sql/driver"
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type IntSlice []int

func (j *IntSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j IntSlice) Value() (driver.Value, error) {
	if j == nil {
		return json.Marshal(IntSlice{})
	}
	return json.Marshal(j)
}

type StrSlice []string

func (j *StrSlice) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j StrSlice) Value() (driver.Value, error) {
	if j == nil {
		return json.Marshal(StrSlice{})
	}
	return json.Marshal(j)
}

type Object map[string]interface{}

func (j *Object) Scan(value interface{}) error {
	return json.Unmarshal(value.([]byte), j)
}

func (j Object) Value() (driver.Value, error) {
	if j == nil {
		return json.Marshal(Object{})
	}
	return json.Marshal(j)
}

type BaseModel struct {
	Id        string         `gorm:"primarykey;comment:主键id;not null" json:"id"`
	CreatedAt time.Time      `gorm:"type:timestamp with time zone;comment:创建时间,带时区;not null" json:"created_at"`
	UpdatedAt time.Time      `gorm:"type:timestamp with time zone;comment:修改时间,带时区;not null" json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"-"`
	IsDelete  bool           `gorm:"is_delete"`
}

type Pagination struct {
	Offset int
	Limit  int
}
