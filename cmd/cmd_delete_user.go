package cmd

import "context"

type CmdDeleteUser struct {
}

func (c CmdDeleteUser) Execute(ctx context.Context) error {
	return nil
}
