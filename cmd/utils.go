package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

//列出文件信息，同时根据文件类型使用不同颜色区别显示，对文件大小使用指定单位
func myUnitPrint(fileInfoList []fs.FileInfo, unit string) {
	var rate float64
	if unit == "B" {
		rate = 1
	} else if unit == "K" {
		rate = 1024
	} else if unit == "M" {
		rate = 1024 * 1024
	} else if unit == "G" {
		rate = 1024 * 1024 * 1024
	} else {
		fmt.Println("单位未知...")
		rate = 1
	}
	fmt.Println("LastModifyTime\t\tDirectory\tMode\t\tSize\t\tName")
	for i := range fileInfoList {
		if fileInfoList[i].Mode()&fs.ModeDir == fs.ModeDir {
			fmt.Printf("\033[1;34;40m%v\t%v\t\t%v\t%.2f %v\t\t%v\033[0m\n", fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				fileInfoList[i].IsDir(), fileInfoList[i].Mode(), float64(fileInfoList[i].Size())/rate, unit, fileInfoList[i].Name())
		} else if strings.Contains(fileInfoList[i].Name(), ".exe") {
			fmt.Printf("\033[1;33;40m%v\t%v\t\t%v\t%.2f %v\t\t%v\033[0m\n", fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				fileInfoList[i].IsDir(), fileInfoList[i].Mode(), float64(fileInfoList[i].Size())/rate, unit, fileInfoList[i].Name())
		} else {
			fmt.Printf("%v\t%v\t\t%v\t%.2f %v\t\t%v\n", fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				fileInfoList[i].IsDir(), fileInfoList[i].Mode(), float64(fileInfoList[i].Size())/rate, unit, fileInfoList[i].Name())
		}
	}
}

//列出当前目录下的文件
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

//列表形式展示当前目录下文件清单及文件类型、大小等信息
func listDisplay(dirname string) {
	fileInfoList, err := ioutil.ReadDir(dirname)
	if err != nil {
		log.Fatal(err)
	}
	myUnitPrint(fileInfoList, "K")
}

//显示隐藏文件信息
func allDisplay() {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	count := 0
	for i := range fileInfoList {
		str := fileInfoList[i].Name()
		if str[0] == '.' {
			fmt.Println("LastModifyTime\t\tDirectory\tMode\t\tSize\tName")
			fmt.Printf("%v\t%v\t\t%v\t%vKB\t%v\t\n", fileInfoList[i].ModTime().Format("2006-01-02 15:04:05"),
				fileInfoList[i].IsDir(), fileInfoList[i].Mode(), fileInfoList[i].Size()/1024, fileInfoList[i].Name())
			count++
		}
	}
	if count == 0 {
		fmt.Println("没有隐藏文件")
	}
}

//根据指定参数值对展示列表进行排序
func sortDisplay(sortOrder string) {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	switch sortOrder {
	case "TimeUp":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].ModTime().After(fileInfoList[j+1].ModTime()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	case "TimeDown":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].ModTime().Before(fileInfoList[j+1].ModTime()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	case "SizeUp":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].Size() > (fileInfoList[j+1].Size()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	case "SizeDown":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].Size() < (fileInfoList[j+1].Size()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	case "NameUp":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].Name() > (fileInfoList[j+1].Name()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	case "NameDown":
		for i := range fileInfoList {
			for j := 0; j < len(fileInfoList)-1-i; j++ {
				if fileInfoList[j].Name() < (fileInfoList[j+1].Name()) {
					fileInfoList[j], fileInfoList[j+1] = fileInfoList[j+1], fileInfoList[j]
				}
			}
		}
		myUnitPrint(fileInfoList, "K")
	}
}

//根据指定参数使用不同单位显示文件大小信息
func unitDisplay(unit string) {
	pwd, _ := os.Getwd()
	fileInfoList, err := ioutil.ReadDir(pwd)
	if err != nil {
		log.Fatal(err)
	}
	switch unit {
	case "B":
		myUnitPrint(fileInfoList, unit)
	case "K":
		myUnitPrint(fileInfoList, unit)
	case "M":
		myUnitPrint(fileInfoList, unit)
	case "G":
		myUnitPrint(fileInfoList, unit)
	}
}

//以树形展示目录结构
func treeDisplay(dir string, l int) {
	files, err := ioutil.ReadDir(dir)
	if err != nil {
		log.Fatal(err)
	}
	currentLen := length - l
	for i := range files {
		if files[i].IsDir() {
			if currentLen <= length {
				treePrint(currentLen, files[i].Name())
				currentLen++
				currentDir := dir + "\\" + files[i].Name()
				tempLen := l - 1
				treeDisplay(currentDir, tempLen)
				currentLen--
			}
		} else {
			if currentLen <= length {
				treePrint(currentLen, files[i].Name())
			}
		}
	}
}

//以树形展示目录结构
func treePrint(l int, fileName string) {
	var treeFormat string
	for i := 0; i < (l-1)*3+2; i++ {
		treeFormat += " "
	}
	if l == 0 {
		fmt.Printf("|—— %v\n", fileName)
	} else {
		fmt.Printf("|%s|—— %v\n", treeFormat, fileName)
	}
}
