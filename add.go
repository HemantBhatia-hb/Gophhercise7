package cmd

import (
	"fmt"
	"strings"

	"github.com/gophercise7/db"
	"github.com/spf13/cobra"
)

var addCmd = &cobra.Command{
	Use:   "add",                           // name of the command
	Short: "Add A  task to your task list", // short description about command.

	Run: func(cmd *cobra.Command, args []string) {
		task := strings.Join(args, " ")
		_, err := db.CreateTask(task)
		if err != nil {
			fmt.Println("Something went wrong!", err)
			return
		}
		fmt.Print("Added \"%s\" to your task list.\n", task)
		fmt.Println()
	},
}

func init() {
	RootCmd.AddCommand(addCmd)
}
