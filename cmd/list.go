package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
)

var listCmd = &cobra.Command{
	Use:   "gls",
	Short: "Show directory lists",
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println("listCmd called")
	},
}

func init() {
	rootCmd.AddCommand(listCmd)
}
