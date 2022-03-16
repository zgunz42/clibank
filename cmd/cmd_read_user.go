package cmd

import "context"

type CmdReadUser struct {
}

func (c CmdReadUser) Execute(ctx context.Context) error {
	return nil
}
