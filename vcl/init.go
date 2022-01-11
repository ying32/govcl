//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"fmt"
	"os"
	"runtime"

	. "github.com/ying32/govcl/vcl/api"
)

const (
	// CN: 要求最小liblcl二进制版本
	// EN: Requires a minimum liblcl binary version.
	requireMinBinaryVersion = 0x02020000
)

var (
	// 几个实例类，不需要Create即可访问，同时也不需要手动Free

	Application *TApplication // 应用程序管理
	Screen      *TScreen      // 屏幕
	Mouse       *TMouse       // 鼠标
	Clipboard   *TClipboard   // 剪切板
	Printer     *TPrinter     // 打印机
)

func toVersionString(ver uint32) string {
	if byte(ver) == 0 {
		return fmt.Sprintf("%d.%d.%d", byte(ver>>24), byte(ver>>16), byte(ver>>8))
	}
	return fmt.Sprintf("%d.%d.%d.%d", byte(ver>>24), byte(ver>>16), byte(ver>>8), byte(ver))
}

func init() {
	if !DEBUG {
		defer func() {
			if err := recover(); err != nil {
				showError(err)
				os.Exit(1)
			}
		}()
	}
	libVersion := DLibVersion()
	fmt.Println("Library Version:", toVersionString(libVersion))
	if libVersion < requireMinBinaryVersion {
		panic(fmt.Sprintf("Require liblcl binary version >=%s. Please go to \"https://github.com/ying32/govcl\" to download the latest binary.", toVersionString(requireMinBinaryVersion)))
	}
	// 这个似乎得默认加上，锁定主线程，防止中间被改变
	runtime.LockOSThread()
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(eventCallback)
	// 消息回调
	SetMessageCallback(messageCallback)
	// 线程同步回调
	SetThreadSyncCallback(threadSyncCallback)
	// 调求回调CreateParams方法
	SetRequestCallCreateParamsCallback(requestCallCreateParamsCallback)

	// 导入几个实例类
	Application = AsApplication(Application_Instance())
	Screen = AsScreen(Screen_Instance())
	Mouse = AsMouse(Mouse_Instance())
	Clipboard = AsClipboard(Clipboard_Instance())
	Printer = AsPrinter(Printer_Instance())

	// 尝试加载ICON，仅Windows下有效，尝试加载名为MAINICON的图标
	tryLoadAppIcon()
}
