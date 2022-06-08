package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
)

var list bool
var directory string
var rootCmd = &cobra.Command{
	Use:     "gls",
	Short:   "Show directory lists",
	Long:    "",
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 1 {
			display()
		} else {
			switch os.Args[1] {
			case "-l":
				pwd, _ := os.Getwd()
				listDisplay(pwd)
			case "-d":
				listDisplay(os.Args[2])
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
	rootCmd.Flags().StringVarP(&directory, "directory", "d", "", "")
}

func display() {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	for i := range fileInfoList {
		fmt.Println(fileInfoList[i].Name())
	}
}
func listDisplay(dirname string) {
	fileInfoList, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("Name\tdirectory\tMode\t\tLastModifyTime\t\t\t\tsize")
	for i := range fileInfoList {
		fmt.Printf("%s\t%v\t\t%s\t%s\t%vKB\n", fileInfoList[i].Name(), fileInfoList[i].IsDir(),
			fileInfoList[i].Mode(), fileInfoList[i].ModTime(), fileInfoList[i].Size()/1024)
	}
}
