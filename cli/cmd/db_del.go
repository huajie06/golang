package cmd

import (
	"errors"
	"fmt"
	"learn/test/cli/db"
	"log"
	"strconv"
	"strings"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(db_del)
}

var db_del = &cobra.Command{
	Use:   "db_del",
	Short: "delete your TODO from Database",
	Run: func(cmd *cobra.Command, args []string) {
		k, err := strconv.Atoi(strings.Join(args[:1], ""))
		if err != nil {
			fmt.Println(errors.New("invalid id"))
		}

		// v:= strings.Join(args[1:], "")
		err = db.DelFromDB(k)
		if err != nil {
			log.Println(err)
		}
	},
}
