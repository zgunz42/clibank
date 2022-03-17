package wallets

type IWalletService interface {
	GetWalletByPhoneNumber(phoneNumber string) (Wallet, error)
	UpdateWallet(dto UpdateWalletDto) error
}

type WalletService struct {
	walletRepo IWalletRepository
}

func (s *WalletService) Init(walletRepo IWalletRepository) {
	s.walletRepo = walletRepo
}

func (s *WalletService) GetWalletByPhoneNumber(phoneNumber string) (Wallet, error) {
	return s.walletRepo.GetWalletByPhoneNumber(phoneNumber)
}

func (s *WalletService) UpdateWallet(dto UpdateWalletDto) error {
	wallet, err := s.walletRepo.GetWalletById(dto.ID)
	if err != nil {
		return err
	}
	wallet.Balance = dto.Balance
	wallet.Currency = dto.Currency
	return s.walletRepo.UpdateWallet(&wallet)
}
