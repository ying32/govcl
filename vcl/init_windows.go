package vcl

import (
	"fmt"

	. "github.com/ying32/govcl/vcl/win"
)

var (
	// StyleManager 没有实例类的，属于静态类
	StyleManager TStyleManager
)

func showError(err interface{}) {
	MessageBox(0, fmt.Sprint(err), "Error", MB_ICONERROR)
}

// 尝试加载默认Application icon
func tryLoadAppIcon() {
	hIcon := LoadIcon2(GetSelfModuleHandle(), "MAINICON")
	if hIcon != 0 {
		Application.Icon().SetHandle(hIcon)
	}
}
