package users

import "gorm.io/gorm"

type User struct {
	gorm.Model
	Email       string  `gorm:"type:varchar(100);unique;not null"`
	Pin         string  `gorm:"type:varchar(100);not null"`
	PhoneNumber string  `gorm:"type:varchar(12);uniqueIndex;not null" json:"phone_number"`
	Account     Account `gorm:"foreignKey:PhoneNumber;references:PhoneNumber;constraint:OnUpdate:CASCADE,OnDelete:CASCADE;"`
}
