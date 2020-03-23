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

	. "github.com/ying32/govcl/vcl/win"
)

var (
	// CN: StyleManager 没有实例类的，属于静态类。
	// EN: StyleManager does not have an instance class, it is a static class.
	StyleManager TStyleManager
)

func showError(err interface{}) {
	MessageBox(0, fmt.Sprint(err), "Error", MB_ICONERROR)
}

// CN: 尝试加载默认Application icon
// EN: Try to load the default Application icon.
func tryLoadAppIcon() {
	hIcon := LoadIcon2(GetSelfModuleHandle(), "MAINICON")
	if hIcon != 0 {
		Application.Icon().SetHandle(hIcon)
	}
}
