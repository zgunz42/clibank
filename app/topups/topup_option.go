package topups

import "gorm.io/gorm"

type TopupOption struct {
	gorm.Model
	Name   string `gorm:"type:varchar(100)"`
	Code   string `gorm:"type:varchar(100);uniqueIndex"`
	AccNo  string `gorm:"type:varchar(100);unique"`
	Status string `gorm:"type:enum('maintance','inactive','active');default:'active'"`
}
