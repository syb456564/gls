package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"os"
)

var list bool
var all bool
var sortOrder string
var unit string
var rootCmd = &cobra.Command{
	Use:     "gls",
	Short:   "Show directory lists",
	Long:    "",
	Args:    cobra.MaximumNArgs(2),
	Version: "3.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 1 {
			display()
		} else {
			switch os.Args[1] {
			case "-l":
				pwd, _ := os.Getwd()
				listDisplay(pwd)
			case "-a":
				allDisplay()
			case "-s":
				sortDisplay(sortOrder)
			case "-u":
				unitDisplay(unit)
			default:
				listDisplay(os.Args[1])
			}
		}
	},
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "")
	rootCmd.Flags().StringVarP(&sortOrder, "sort", "s", "", "")
	rootCmd.Flags().StringVarP(&unit, "unit", "u", "", "")
}
