package cmd

import (
	"context"
	"fmt"

	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/app/users"
	"github.com/ALTA-BE7-I-Kadek-Adi-Gunawan/clibank/platform"
)

type CmdAddUser struct {
}

func (cmd CmdAddUser) Execute(ctx context.Context) error {
	service := ctx.Value(platform.UserServiceKey)
	user := users.CreateUserDto{}

	fmt.Println("Enter Phone Number: ")
	fmt.Scan(&user.Phone)
	fmt.Println("Enter your email: ")
	fmt.Scan(&user.Email)
	fmt.Println("Enter your pin: ")
	fmt.Scan(&user.Pin)
	fmt.Println("Enter your pin again: ")
	fmt.Scan(&user.ConfirmPin)

	var userService users.UserService = service.(users.UserService)
	userDb, err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("User %v has been created", userDb.ID)
	return nil
}
