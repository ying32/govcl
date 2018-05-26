package vcl

import (
	. "gitee.com/ying32/govcl/vcl/win"
)

var (
	// StyleManager 没有实例类的，属于静态类
	StyleManager TStyleManager
)

func showError(err interface{}) {
	MessageBox(0, err.(error).Error(), "Error", MB_ICONERROR)
}
