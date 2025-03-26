package cmd

import (
	"errors"

	"github.com/spf13/cobra"
)

func RunUserActivity(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("wrong number of parameters")
	}

	if args[0] == "--help" || args[0] == "-h" {
		return SupportHelp()
	} else {
		return FindUserActivity()
	}
}

func SupportHelp() error {
	return nil
}

func FindUserActivity() error {
	return nil
}
