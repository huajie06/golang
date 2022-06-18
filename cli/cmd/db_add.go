package cmd

import (
	"learn/test/cli/db"
	"log"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	err := db.Init()
	if err != nil {
		log.Println(err)
	}
	rootCmd.AddCommand(db_add)
}

var db_add = &cobra.Command{
	Use:   "db_add",
	Short: "Add a new task to your TODO list to database",
	Long:  "some stuff",
	Run: func(cmd *cobra.Command, args []string) {

		err := db.WriteToDB(strings.Join(args, " "))

		if err != nil {
			log.Println(err)
		}
	},
}
