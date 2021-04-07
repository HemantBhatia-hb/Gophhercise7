package cmd

import (
	"fmt"
	"os"

	"github.com/gophercise7/db"
	"github.com/spf13/cobra"
)

var ListCmd = &cobra.Command{
	Use:   "list",                   // name of the command.
	Short: "list all of your tasks", // short description about command.

	Run: func(cmd *cobra.Command, args []string) {
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went wrong:", err)
			os.Exit(1)
		}
		if len(tasks) == 0 {
			fmt.Println("You have no task to complete!")
			return
		}
		fmt.Println("You have the following tasks:")
		for i, task := range tasks {
			fmt.Printf("%d. %s,Key=%d\n", i+1, task.Value, task.Key)
		}
	},
}

func init() {
	RootCmd.AddCommand(ListCmd)
}
