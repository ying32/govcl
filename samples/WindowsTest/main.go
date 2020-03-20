// +build windows

package main

import (
	"fmt"
	"os"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/win"
)

func main() {

	// 利用互斥来演示exe单一运行，当然不止这一种方法了
	// GetLastError 无法获取错误，待研究原因所在。
	/*
		Mutex := win.CreateMutex(nil, true, "SingleRunTest")
		defer win.ReleaseMutex(Mutex)
		fmt.Println("Mutex:", Mutex)
		fmt.Println("win.GetLastError:", win.GetLastError(), syscall.GetLastError())
		if win.GetLastError() == win.ERROR_ALREADY_EXISTS {
			win.MessageBox(0, "我已经在运行中啦！", "运行提示", win.MB_OK+win.MB_ICONINFORMATION)
			os.Exit(1)
		}

		Mutex3 := win.CreateMutex(nil, true, "SingleRunTest")

		fmt.Println("win.GetLastError:", win.GetLastError())
		defer win.ReleaseMutex(Mutex3)
	*/

	// Exe运行在Administrator权限下检测
	fmt.Println("IsAdministrator:", win.IsAdministrator())

	// 注：只有当exe为Win32并在64位系统上运行时才返回true, 否则都会返回false
	fmt.Println("IsWow64:", win.IsWow64())

	// 在资源管理器中定位指定文件名
	// win.OpenInExplorer("F:\\Golang\\bin\\libvclx64.dll")
	var s string
	fmt.Scan(&s)

	// 以管理员权限运行一个程序
	if cmdFileName, ok := os.LookupEnv("ComSpec"); ok {
		fmt.Println(win.RunAsAdministrator(cmdFileName, "/c ping 192.168.1.1", ""))
	}

}
