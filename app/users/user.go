package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email string `gorm:"type:varchar(100);unique"`
	Pin   string `gorm:"type:varchar(100)"`
}
