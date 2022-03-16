package wallets

type AddWalletRequest struct {
	UserID    int
	Currency  string
	Balance   float64
	CreatedAt string
	UpdatedAt string
}
