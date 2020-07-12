package model

import "github.com/jinzhu/gorm"

type User struct {
	gorm.Model
	Name string  `gorm:"type:varchar(20);not null"`    // 目前没限制，可重复
	Telephone string  `gorm:"type:varchar(11);not null"`
	Password string  `gorm:"size:255; not null"`
}