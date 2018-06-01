package main

import (
	"fmt"

	. "github.com/ying32/govcl/vcl/rtl/version"
)

func main() {

	fmt.Println("Major: ", OSVersion.Major)
	fmt.Println("Minor: ", OSVersion.Minor)
	fmt.Println("Name: ", OSVersion.Name)
	fmt.Println("String: ", OSVersion.ToString())
	switch OSVersion.Platform {
	case PfWindows:
		if OSVersion.CheckMajor(10) {
			fmt.Println("当前为Windows10")
		} else {
			fmt.Println("当前不是Windows10")
		}
	case PfLinux:
		fmt.Println("LIBC Version:", OSVersion.LibCVersionMajor, ", ", OSVersion.LibCVersionMinor)
	case PfMacOS:
		//
	}

}
