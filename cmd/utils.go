package cmd

import (
	"fmt"
	"io/fs"
	"io/ioutil"
	"log"
	"os"
	"strings"
)

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
	myUnitPrint(fileInfoList, "K")
}

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

func sortDisplay(sortOrder string) { //1代表按修改时间排序,2代表按文件大小排序,3代表按文件名排序,order为true代表升序,false代表降序
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

func treePrint(l int, fileName string) {
	if l == 0 {
		fmt.Printf("|—— %v\n", fileName)
	} else if l == 1 {
		fmt.Printf("|  |—— %v\n", fileName)
	} else if l == 2 {
		fmt.Printf("|     |—— %v\n", fileName)
	} else if l == 3 {
		fmt.Printf("|        |—— %v\n", fileName)
	} else {
		fmt.Printf("|        |—— %v\n", fileName)
	}
}
