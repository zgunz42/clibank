package wallets

import (
	"gorm.io/gorm"
)

type Wallet struct {
	gorm.Model
	Currency string  `gorm:"type:enum('IDR','USD','EUR');not null"`
	Balance  float64 `gorm:"type:decimal(10,2);not null"`
	Status   string  `gorm:"type:enum('active','inactive');not null"`
}
