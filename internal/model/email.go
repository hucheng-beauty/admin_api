package model

type Email struct {
	ID    int    `gorm:"user_id;primaryKey;autoIncrement;type:varchar;comment:记录号"`
	Email string `gorm:"email;type:varchar;comment:邮箱"`
}
