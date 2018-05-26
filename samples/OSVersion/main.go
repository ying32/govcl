package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl/rtl"
)

func main() {

	fmt.Println("Major: ", rtl.OSVersion.Major)
	fmt.Println("Minor: ", rtl.OSVersion.Minor)
	fmt.Println("Name: ", rtl.OSVersion.Name)
	fmt.Println("String: ", rtl.OSVersion.ToString())
	switch rtl.OSVersion.Platform {
	case rtl.PfWindows:
		if rtl.OSVersion.CheckMajor(10) {
			fmt.Println("当前为Windows10")
		} else {
			fmt.Println("当前不是Windows10")
		}
	case rtl.PfLinux:
		fmt.Println("LIBC Version:", rtl.OSVersion.LibCVersionMajor, ", ", rtl.OSVersion.LibCVersionMinor)
	case rtl.PfMacOS:
		//
	}

}
