package cmd

import (
	"fmt"
	"log"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(list)
}

var list = &cobra.Command{
	Use:   "list",
	Short: "list all your TODO list",
	Run: func(cmd *cobra.Command, args []string) {
		r, err := getAllTask(fname)
		if err != nil {
			log.Println(err)
		}

		fmt.Println("You have the following tasks")
		fmt.Println("----------------------------")
		for i := 1; i <= len(r); i++ {
			fmt.Printf("Task #%v: %v\n", i, r[i])
		}
	},
}
