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
	FindByPhone(phone string) ([]*User, error)
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

func (u *UserRepository) FindByPhone(phone string) ([]*User, error) {
	var users []*User
	err := u.db.Where("phone = ?", phone).Find(&users).Error
	return users, err
}

func (u *UserRepository) Create(data CreateUserDto) (User, error) {

	user := &User{
		Email: data.Email,
		Pin:   data.Pin,
	}

	err := u.db.Transaction(func(tx *gorm.DB) error {
		err := tx.Create(user).Error
		if err != nil {
			var mysqlError *mysql.MySQLError
			if errors.As(err, &mysqlError) && mysqlError.Number == 1062 {
				return errors.New("email already exist")
			}
			return err
		}

		wallet := &wallets.Wallet{
			Balance:  0,
			Status:   "active",
			Currency: "IDR",
		}

		if err := tx.Create(wallet).Error; err != nil {
			return err
		}

		account := &Account{
			PhoneNumber: data.Phone,
			Wallet:      *wallet,
			User:        user,
		}
		if err := tx.Create(account).Error; err != nil {
			return err
		}

		return nil
	})
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
	account := &Account{}
	err := r.db.Joins("User").Where("phone_number = ?", phone).First(&account).Error
	if err != nil {
		return User{}, err
	}

	if data.Name != "" {
		// update account name
		account.Name = data.Name
		err = r.db.Save(&account).Error
		if err != nil {
			return User{}, err
		}
	}
	if data.Pin != "" && len(data.Pin) == 6 {
		user := account.User
		user.Pin = data.Pin
		err = r.db.Save(user).Error
	}

	return *account.User, err
}
