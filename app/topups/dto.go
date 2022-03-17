package topups

type TopupWalletDTO struct {
	PhoneNumber string  `json:"phone_number"`
	Method      string  `json:"method"`
	Amount      float64 `json:"amount"`
}

type AddTopupOptionDTO struct {
	Name  string `json:"name"`
	Code  string `json:"code"`
	AccNo string `json:"acc_no"`
}
