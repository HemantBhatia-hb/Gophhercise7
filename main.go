package main

import (
	"fmt"
	"os"

	"path/filepath"

	"github.com/gophercise7/cmd"
	"github.com/gophercise7/db"
	homedir "github.com/mitchellh/go-homedir"
)

func main() {
	home, _ := homedir.Dir() // this is used for detecting user's home directory.
	dbPath := filepath.Join(home, "task.db")
	must(db.Init(dbPath))
	must(cmd.RootCmd.Execute())
}
func must(err error) {
	if err != nil {
		fmt.Println(err.Error())
		os.Exit(1)
	}
}
