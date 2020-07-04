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

func showError(err interface{}) {
	MessageBox(0, fmt.Sprint(err), "Error", MB_ICONERROR)
}

// 尝试加载默认Application icon
//
// Try to load the default Application icon.
func tryLoadAppIcon() {
	// 将来会删除此代码，因为liblcl中已经做了处理了
	if Application.Icon().Handle() == 0 {
		hIcon := LoadIcon2(GetSelfModuleHandle(), "MAINICON")
		if hIcon != 0 {
			Application.Icon().SetHandle(hIcon)
		}
	}
}
