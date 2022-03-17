package users

import (
	"errors"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/wallets"
	"github.com/go-sql-driver/mysql"
	"gorm.io/gorm"
)

type IUserRepository interface {
	Init(db *gorm.DB)
	FindByID(id int) (*User, error)
	FindByPhone(phone string) (*User, error)
	Create(user CreateUserDto) (User, error)
	CheckPin(phone string, pin string) (bool, error)
	Update(phone string, data UpdateUserDto) (User, error)
}

type UserRepository struct {
	db *gorm.DB
}

func (u *UserRepository) Init(db *gorm.DB) {
	u.db = db
}

func (u *UserRepository) FindByID(id int) (*User, error) {
	var user User
	err := u.db.Where("id = ?", id).First(&user).Error
	return &user, err
}

func (u *UserRepository) FindByPhone(phone string) (*User, error) {
	var user *User
	err := u.db.Joins("Account").Where(&User{PhoneNumber: phone}).Find(&user).Error
	return user, err
}

func (u *UserRepository) Create(data CreateUserDto) (User, error) {

	user := &User{
		Email:       data.Email,
		Pin:         data.Pin,
		PhoneNumber: data.Phone,
		Account: &Account{
			Name: "",
			Wallet: wallets.Wallet{
				Balance:  0,
				Status:   "active",
				Currency: "IDR",
			},
		},
	}

	err := u.db.Create(user).Error
	if err != nil {
		var mysqlError *mysql.MySQLError
		if errors.As(err, &mysqlError) && mysqlError.Number == 1062 {
			return *user, errors.New("email already exist")
		}
		return *user, err
	}
	return *user, err
}

func (r *UserRepository) CheckPin(phone string, pin string) (bool, error) {
	var user User
	err := r.db.Where("phone = ? AND pin = ?", phone, pin).First(&user).Error
	if err != nil {
		return false, err
	}
	return true, nil
}

func (r *UserRepository) Update(phone string, data UpdateUserDto) (User, error) {
	user := &User{}
	err := r.db.Joins("Account").Where("phone_number = ?", phone).First(&user).Error
	if err != nil {
		return User{}, err
	}

	if data.Name != "" {
		// update account name
		user.Account.Name = data.Name
		err = r.db.Save(&user.Account).Error
		if err != nil {
			return User{}, err
		}
	}
	if data.Pin != "" && len(data.Pin) == 6 {
		user.Pin = data.Pin
		err = r.db.Save(user).Error
	}

	return *user, err
}
