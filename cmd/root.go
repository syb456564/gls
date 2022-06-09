package cmd

import (
	"fmt"
	"github.com/spf13/cobra"
	"io/ioutil"
	"log"
	"os"
	"strconv"
)

var list bool
var all bool
var sort int
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
			case "-s":
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
	rootCmd.Flags().IntVarP(&sort, "sort", "s", 0, "")
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
	fmt.Println("LastModifyTime\t\tDirectory\tMode\t\tSize\tName")
	for i := range fileInfoList {
		fmt.Printf("%v\t", fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"))
		fmt.Printf("%v\t\t", fileInfoList[i].IsDir())
		fmt.Printf("%v\t", fileInfoList[i].Mode())
		fmt.Printf("%vKB\t", fileInfoList[i].Size()/1024)
		fmt.Printf("%v\t\n", fileInfoList[i].Name())
	}
	var str = make([]string, 5)
	for i := range fileInfoList {
		str[0] = fileInfoList[i].Name()
		str[1] = strconv.FormatBool(fileInfoList[i].IsDir())
		str[2] = strconv.Itoa(int(fileInfoList[i].Mode()))
		str[3] = fileInfoList[i].ModTime().Format("2006-01-02 15:04:05")
		str[4] = strconv.Itoa(int(fileInfoList[i].Size()))
	}
}
