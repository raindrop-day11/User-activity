package cmd

import (
	"User-activity/model"
	"errors"

	"github.com/spf13/cobra"
)

func RunUserActivity(cmd *cobra.Command, args []string) error {
	if len(args) != 1 {
		return errors.New("wrong number of parameters")
	}

	username := args[0]
	return model.FetchTheRecentActivity(username)

}
