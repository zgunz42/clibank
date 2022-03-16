package topups

type TopupRequest struct {
	PhoneNumber string  `json:"phone_number"`
	Amount      float64 `json:"amount"`
}
