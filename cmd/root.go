package cmd

import (
	"fmt"
	"github.com/liushuochen/gotable"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var list bool
var all bool
var rootCmd = &cobra.Command{
	Use:     "gls",
	Short:   "Show directory lists",
	Long:    "",
	Args:    cobra.MaximumNArgs(1),
	Version: "1.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 1 {
			display()
		} else {
			switch os.Args[1] {
			case "-l":
				pwd, _ := os.Getwd()
				listDisplay(pwd)
			case "-a":
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
	table, err := gotable.Create("Name", "directory", "Mode", "LastModifyTime", "Size")
	if err != nil {
		fmt.Println("Create table failed: ", err.Error())
		return
	}
	var str = make([]string, 5)
	for i := range fileInfoList {
		str[0] = fileInfoList[i].Name()
		str[1] = strconv.FormatBool(fileInfoList[i].IsDir())
		str[2] = strconv.Itoa(int(fileInfoList[i].Mode()))
		str[3] = fileInfoList[i].ModTime().Format("2006-01-02 15:04:05")
		str[4] = strconv.Itoa(int(fileInfoList[i].Size()))
		table.AddRow(str)
	}
	fmt.Println(table)
}
