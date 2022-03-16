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

	fmt.Print("Enter Phone Number: ")
	fmt.Scanln(&user.Phone)
	fmt.Print("Enter your email: ")
	fmt.Scanln(&user.Email)
	fmt.Print("Enter your pin: ")
	fmt.Scanln(&user.Pin)
	fmt.Print("Enter your pin again: ")
	fmt.Scanln(&user.ConfirmPin)

	var userService users.UserService = service.(users.UserService)
	userDb, err := userService.CreateUser(user)
	if err != nil {
		fmt.Printf("Error: %v\n", err)
		return err
	}

	fmt.Printf("User %v has been created", userDb.ID)
	return nil
}
