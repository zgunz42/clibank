package topups

type Repository interface {
	Topup(phoneNumber string, amount float64) error
}
