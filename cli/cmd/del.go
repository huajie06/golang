package cmd

import (
	"errors"
	"fmt"
	"log"
	"strconv"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(del)
}

var del = &cobra.Command{
	Use:   "del",
	Short: "Delete an existing task from your TODO list",
	Long:  "some stuff",
	Run: func(cmd *cobra.Command, args []string) {
		if len(args) >= 2 {
			fmt.Println("Please enter only 1")
			return
		}
		deln, err := strconv.Atoi(args[0])
		if err != nil {
			strconvErr := errors.New("ID is not valid")
			log.Println(strconvErr)
			return
		}
		err = delNthLine(fname, deln)

		if err != nil {
			log.Println(err)
		}
	},
}
