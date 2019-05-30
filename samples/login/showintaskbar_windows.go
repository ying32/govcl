package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/win"
)

func SetShowInTaskBar(f vcl.IWinControl) {
	win.SetWindowLong(f.Handle(), win.GWL_EXSTYLE, uintptr(win.GetWindowLong(f.Handle(), win.GWL_EXSTYLE)|win.WS_EX_APPWINDOW))
}
