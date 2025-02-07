package model

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email    string `gorm:"nuiqueIndex;type:varchar(128) not null"`
	Password string `gorm:"type:varchar(64) not null"`
}
// 设置生成的表名
func (User) TableName() string {
	return "user"
}