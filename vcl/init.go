package vcl

import (
	"runtime"

	. "github.com/ying32/govcl/vcl/api"
)

var (
	// 四个实例类，不需要Create和Free即可访问
	Application *TApplication
	Screen      *TScreen
	Mouse       *TMouse
	Clipboard   *TClipboard
)

func init() {
	defer func() {
		if err := recover(); err != nil {
			showError(err)
		}
	}()
	// 这个似乎得默认加上，锁定主线程，防止中间被改变
	runtime.LockOSThread()
	// 设置事件的回调函数，因go中callback数量有限，只好折中处理
	SetEventCallback(callbackStdcall)

	// 导入四个实例类
	Application = ApplicationFromInst(Application_Instance())
	Screen = ScreenFromInst(Screen_Instance())
	Mouse = MouseFromInst(Mouse_Instance())
	Clipboard = ClipboardFromInst(Clipboard_Instance())

	// 尝试加载ICON，仅Windows下有限，尝试加载名为MAINICON的
	tryLoadAppIcon()
}
