//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import . "github.com/ying32/govcl/vcl/api"

// 从资源中创建Form，不使用Application.CreateForm
func CreateResForm(owner IComponent, fields ...interface{}) {
	resObjectBuild(1, owner, 0, fields...)
}

// 窗口消息过程
func (f *TForm) SetOnWndProc(fn TWndProcEvent) {
	Form_SetOnWndProc(f.instance, fn)
}

func (f *TForm) setGoPtr(ptr uintptr) {
	Form_SetGoPtr(f.instance, ptr)
}

// 这个方法主要是用于当不使用资源窗口创建时用，这个方法要用于设置了Width, Height或者ClientWidth、ClientHeight之后
func (f *TForm) ScaleSelf() {
	if Application.Scaled() {
		f.SetClientWidth(int32(float64(f.ClientWidth()) * (float64(Screen.PixelsPerInch()) / 96.0)))
		f.SetClientHeight(int32(float64(f.ClientHeight()) * (float64(Screen.PixelsPerInch()) / 96.0)))
	}
}
