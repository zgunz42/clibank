package wallets

import "gorm.io/gorm"

type IWalletRepository interface {
	GetWalletByPhoneNumber(phoneNumber string) (Wallet, error)
	GetWalletById(walletId uint) (Wallet, error)
	UpdateWallet(wallet *Wallet) error
}

type WalletRepository struct {
	db *gorm.DB
}

func (rep *WalletRepository) Init(db *gorm.DB) {
	rep.db = db
}

func (repo *WalletRepository) GetWalletByPhoneNumber(phoneNumber string) (Wallet, error) {
	var wallet Wallet
	err := repo.db.Model(&Wallet{}).Where("id = (SELECT wallet_id FROM accounts WHERE phone_number = ?)", phoneNumber).First(&wallet).Error
	return wallet, err
}

func (repo *WalletRepository) UpdateWallet(wallet *Wallet) error {
	if wallet.Balance > 50000 {
		wallet.Status = "active"
	} else {
		wallet.Status = "inactive"
	}
	err := repo.db.Save(wallet).Error
	return err
}

func (repo *WalletRepository) GetWalletById(walletId uint) (Wallet, error) {
	var wallet Wallet
	err := repo.db.Model(&Wallet{}).Where("id = ?", walletId).First(&wallet).Error
	return wallet, err
}
