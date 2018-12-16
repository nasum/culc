package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

var RootCmd = &cobra.Command{
	Use:   "culc",
	Short: "commandline culculator",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("root command")
	},
}
