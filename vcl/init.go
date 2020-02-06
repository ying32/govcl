package vcl

import (
	"fmt"
	"os"
	"runtime"

	. "github.com/ying32/govcl/vcl/api"
)

const (
	// 要求最小liblcl或者libvcl二进制版本为1.2.6
	requireMinBinaryVersion = 0x01020600
)

var (
	// 几个实例类，不需要Create即可访问，同时也不需要手动Free
	Application *TApplication
	Screen      *TScreen
	Mouse       *TMouse
	Clipboard   *TClipboard
	Printer     *TPrinter
)

func toVersionString(ver uint32) string {
	return fmt.Sprintf("%d.%d.%d.%d", byte(ver>>24), byte(ver>>16), byte(ver>>8), byte(ver))
}

func init() {
	defer func() {
		if err := recover(); err != nil {
			showError(err)
			os.Exit(1)
		}
	}()
	libVersion := DLibVersion()
	fmt.Println("Library Version:", toVersionString(libVersion))
	if libVersion < requireMinBinaryVersion {
		panic("要求libvcl或liblcl二进制版本>=1.2.6。\r\n(Require libvcl or liblcl binary version >=1.2.6.)")
	}
	// 这个似乎得默认加上，锁定主线程，防止中间被改变
	runtime.LockOSThread()
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(eventCallback)
	// 消息回调
	SetMessageCallback(messageCallback)
	// 线程同步回调
	SetThreadSyncCallback(threadSyncCallback)

	// 导入几个实例类
	Application = ApplicationFromInst(Application_Instance())
	Screen = ScreenFromInst(Screen_Instance())
	Mouse = MouseFromInst(Mouse_Instance())
	Clipboard = ClipboardFromInst(Clipboard_Instance())
	Printer = PrinterFromInst(Printer_Instance())

	// 尝试加载ICON，仅Windows下有效，尝试加载名为MAINICON的图标
	tryLoadAppIcon()
}
