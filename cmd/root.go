/*
Copyright © 2023 NAME HERE <EMAIL ADDRESS>
*/
package cmd

import (
	"fmt"
	"os"

	"github.com/spf13/cobra"
)

var rootCmd = &cobra.Command{
	Use:   "golang-cli-app",
	Short: "brief",
	Long: `A longer description 
	for the first command`,

	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("go-cli-app")
	},
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}
