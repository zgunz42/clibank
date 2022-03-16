package topups

type ITopupService interface {
	Topup(phoneNumber string, amount float64) error
}
