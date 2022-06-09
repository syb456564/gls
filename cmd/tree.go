package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var length int
var treeCmd = &cobra.Command{
	Use:   "tree",
	Short: "Tree display directory",
	Args:  cobra.MaximumNArgs(2),
	Run: func(cmd *cobra.Command, args []string) {
		pwd, _ := os.Getwd()
		treeDisplay(pwd, length)
	},
}

func init() {
	treeCmd.Flags().IntVarP(&length, "length", "L", 3, "")
	rootCmd.AddCommand(treeCmd)
}
