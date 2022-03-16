package wallets

type WalletRepository interface {
	Create(wallet *Wallet) error
	Update(wallet *Wallet) error
	Delete(wallet *Wallet) error
	FindByID(id int) (*Wallet, error)
	FindByUserID(userID int) ([]*Wallet, error)
}
