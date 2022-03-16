package cmd

import "context"

type CmdHistoryTransaction struct {
}

func (c CmdHistoryTransaction) Execute(ctx context.Context) error {
	return nil
}
