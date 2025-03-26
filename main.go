package main

import (
	"User-activity/cmd"
	"fmt"

	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:   "root",
		Short: "list one user's action",
		RunE:  cmd.RunUserActivity,
	}

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
