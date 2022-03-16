package cmd

import "context"

type CmdGetUser struct {
}

func (c CmdGetUser) Execute(ctx context.Context) error {
	return nil
}
