package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

func init() {
	rootCmd.AddCommand(add)
}

var add = &cobra.Command{
	Use:   "add",
	Short: "Add a new task to your TODO list",
	Long:  "some stuff",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(args)
	},
}
