package users

import "errors"

type IUserService interface {
	GetUser(phone string) (*User, error)
	GetUsers() ([]*User, error)
	CreateUser(data CreateUserDto) error
	UpdateUser(phone string, pin string, data UpdateUserDto) error
}

type UserService struct {
	userRepository IUserRepository
}

func (u *UserService) Init(userRepository IUserRepository) {
	u.userRepository = userRepository
}

func (u *UserService) CreateUser(data CreateUserDto) (User, error) {
	if data.Pin != data.ConfirmPin {
		return User{}, errors.New("pin and confirm pin must be same")
	}
	user, err := u.userRepository.Create(data)

	return user, err
}

func (u *UserService) GetUser(phone string) (*User, error) {
	return u.userRepository.FindByPhone(phone)
}

func (u *UserService) UpdateUser(phone string, pin string, data UpdateUserDto) (User, error) {
	// update user data
	return u.userRepository.Update(phone, data)
}
