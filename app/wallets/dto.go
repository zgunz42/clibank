package wallets

type AddWalletRequest struct {
	UserID    int
	Currency  string
	Balance   float64
	CreatedAt string
	UpdatedAt string
}

type WalleFilter struct {
	PhoneNumber string
}

type UpdateWalletDto struct {
	ID       uint
	Balance  float64
	Currency string
}
