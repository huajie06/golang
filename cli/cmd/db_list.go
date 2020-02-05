package cmd

import (
	"cli/db"
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(db_list)
}

var db_list = &cobra.Command{
	Use:   "db_list",
	Short: "list all your TODO list from Database",
	Run: func(cmd *cobra.Command, args []string) {
		ret := db.RetrieveAll()
		fmt.Println(ret)
	},
}
