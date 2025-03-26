package main

import (
	"User-activity/cmd"
	"fmt"

	"github.com/fatih/color"
	"github.com/spf13/cobra"
)

func main() {
	var rootCmd = &cobra.Command{
		Use:  "root",
		RunE: cmd.RunUserActivity,
	}

	rootCmd.SetHelpFunc(func(cmd *cobra.Command, args []string) {
		yellow := color.New(color.FgYellow).SprintFunc()
		fmt.Println(yellow("usage example:"))
		fmt.Println(yellow(".\\github-activity <username>"))

		fmt.Print("\n")

		green := color.New(color.FgGreen).SprintFunc()
		fmt.Println(green("it will output like:"))
		fmt.Println(green("Output:"))
		fmt.Println(green("- Pushed 3 commits to kamranahmedse/developer-roadmap"))
		fmt.Println(green("- Opened a new issue in kamranahmedse/developer-roadmap"))
		fmt.Println(green("- Starred kamranahmedse/developer-roadmap"))
		fmt.Println(green("- ..."))
	})

	err := rootCmd.Execute()
	if err != nil {
		fmt.Println(err.Error())
	}
}
