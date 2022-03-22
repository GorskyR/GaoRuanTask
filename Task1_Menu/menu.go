package main

import "fmt"

func main() {
	var str string
	for {
		fmt.Print("请输入命令：")
		fmt.Scan(&str)
		if str == "quit" {
			break
		}
		switch str {
		case "help":
			fmt.Println("帮助信息")
		case "list":
			fmt.Println("列表信息")
		default:
			fmt.Println("无法识别的指令")
		}
	}
}
