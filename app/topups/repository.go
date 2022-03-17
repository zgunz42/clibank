package topups

import (
	"gorm.io/gorm"
)

type ITopupRepository interface {
	Topup(topupWallet *TopupWallet) error
	GetTopupOptions() ([]TopupOption, error)
	GetTopupOption(method string) (TopupOption, error)
	CreateTopupOption(topupOption *TopupOption) error
}

type TopupRepository struct {
	db *gorm.DB
}

func (rep *TopupRepository) Init(db *gorm.DB) {
	rep.db = db
}

func (repo *TopupRepository) Topup(topupWallet *TopupWallet) error {
	return repo.db.Create(topupWallet).Error
}

func (repo *TopupRepository) GetTopupOptions() ([]TopupOption, error) {
	topupOption := []TopupOption{}
	err := repo.db.Find(&topupOption).Error
	return topupOption, err
}

func (repo *TopupRepository) GetTopupOption(method string) (TopupOption, error) {
	topupOption := TopupOption{}
	err := repo.db.Where("code = ?", method).First(&topupOption).Error
	return topupOption, err
}

func (repo *TopupRepository) CreateTopupOption(topupOption *TopupOption) error {
	return repo.db.Create(topupOption).Error
}
