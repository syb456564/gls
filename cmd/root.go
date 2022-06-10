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
	Args:    cobra.MaximumNArgs(2),
	Version: "3.0",
	Run: func(cmd *cobra.Command, args []string) {
		if len(os.Args) == 1 {
			display()
		} else {
			switch os.Args[1] {
			case "-l": //显示当前目录下的文件清单及相关信息
				pwd, _ := os.Getwd()
				listDisplay(pwd)
			case "-a": //显示隐藏文件
				allDisplay()
			case "-s": //对文件排序输出
				sortDisplay(sortOrder)
			case "-u": //指定文件大小的单位
				unitDisplay(unit)
			default: //显示指定目录下的相关信息
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
	rootCmd.Flags().BoolVarP(&list, "list", "l", false, "列表形式展示当前目录下文件清单及文件类型、大小等信息")
	rootCmd.Flags().BoolVarP(&all, "all", "a", false, "显示隐藏文件信息")
	rootCmd.Flags().StringVarP(&sortOrder, "sort", "s", "", "根据指定参数值对展示列表进行排序，参数列表如下：\n"+
		"TimeUp：按修改时间升序输出，TimeDown：按修改时间降序输出\n"+
		"SizeUp：按文件大小升序输出，SizeDown：按修改时间降序输出\n"+
		"NameUp：按文件名升序输出，  NameDown：按文件名降序输出")
	rootCmd.Flags().StringVarP(&unit, "unit", "u", "", "根据指定参数使用不同单位显示文件大小信息，默认单位为KB，参数列表如下：\n"+
		"B：单位为Byte，K：单位为KB，M：单位为MB，G：单位为GB")
}
