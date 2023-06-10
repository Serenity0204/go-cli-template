package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var testCmd = &cobra.Command{
	Use:     "test",
	Aliases: []string{"test"},
	Short:   "test a command",
	Args:    cobra.ExactArgs(1),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("A command for testing, the stuff you typed: ", args[0])
	},
}

func init() {
	rootCmd.AddCommand(testCmd)
}
