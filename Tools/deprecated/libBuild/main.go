package main

import (
	"fmt"
	"log"
	"runtime"
)

func main() {
	fmt.Println("请选择要编译的目标库:\r\n1、libvcl\r\n2、liblcl\r\n3、all")
	var libNumber, libArchNumber int
	_, err := fmt.Scan(&libNumber)
	if err != nil || (libNumber <= 0 || libNumber > 3) {
		fmt.Println("输入错误。")
	}
	switch runtime.GOOS {
	case "windows":
		fmt.Println("请选择arch模式: \r\n1、x86\r\n2、x64\r\n3、all")
		_, err := fmt.Scan(&libArchNumber)
		if err != nil || (libArchNumber <= 0 || libArchNumber > 3) {
			fmt.Println("输入错误。")
		}
	case "linux", "darwin":
		// linux与darwin上总是liblcl和x64的。
		libNumber = 2
		libArchNumber = 2
	default:
		log.Fatal("不支持：", runtime.GOOS)
	}
	fmt.Println(libNumber, libArchNumber)
}
