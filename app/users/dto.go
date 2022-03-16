package users

type CreateUserDto struct {
	Email      string `json:"email"`
	Phone      string `json:"phone"`
	Pin        string `json:"pin"`
	ConfirmPin string `json:"confirm_pin"`
}
