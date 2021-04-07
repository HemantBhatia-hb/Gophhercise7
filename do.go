package cmd

import (
	"fmt"
	"strconv"

	"github.com/gophercise7/db"
	"github.com/spf13/cobra"
)

var doCmd = &cobra.Command{
	Use:   "do",                       // name of the command
	Short: "Marks a task as complete", // short description about command.

	Run: func(cmd *cobra.Command, args []string) {
		var ids []int
		for _, arg := range args {
			id, err := strconv.Atoi(arg) // bcoz of...do 1 2 3 ....these no. are string..
			//we have to convert in integer.

			if err != nil {
				fmt.Println("Failed to parse the arguments:", arg)
			} else {
				ids = append(ids, id)
			}

		}
		tasks, err := db.AllTasks()
		if err != nil {
			fmt.Println("Something went Wrong", err)
			return
		}

		for _, id := range ids {
			if id <= 0 || id > len(tasks) {
				fmt.Println("Invalid task number:", id)
				continue
			}
			tasks := tasks[id-1]
			err := db.DeleteTask(tasks.Key)
			if err != nil {
				fmt.Println("Failed to mark \"%d\" as complete. Error:%s\n", id, err)
			} else {
				fmt.Printf("Marked \"%d\" as complete.\n", id)
			}
		}
	},
}

func init() {
	RootCmd.AddCommand(doCmd)
}
