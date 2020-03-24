// +build windows

package main

import (
	"fmt"
	"syscall"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

var (
	oldWndPrc uintptr
)

type TMainForm struct {
	*vcl.TForm
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(object vcl.IObject) {
	f.SetCaption("Windows Messages")
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	newWndProc := syscall.NewCallback(WndProc)
	oldWndPrc = win.SetWindowLongPtr(mainForm.Handle(), win.GWL_WNDPROC, newWndProc)
	fmt.Println("newWndProc:", newWndProc)
	fmt.Println("oldWndPro:", oldWndPrc)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("按钮1")
	btn.SetLeft(50)
	btn.SetTop(50)
}

func (f *TMainForm) OnFormDestroy(object vcl.IObject) {
	fmt.Println("FormOnDestroy")
	// 完成后要恢复的
	win.SetWindowLongPtr(mainForm.Handle(), win.GWL_WNDPROC, oldWndPrc)
}

func main() {
	vcl.RunApp(&mainForm)
}

func WndProc(hWnd uintptr, message uint32, wParam, lParam uintptr) uintptr {
	switch message {
	case win.WM_SYSCOMMAND:
		switch wParam {
		case win.SC_MAXIMIZE:
			fmt.Println("最大化")
		case win.SC_MINIMIZE:
			fmt.Println("最小化")
		case win.SC_RESTORE:
			fmt.Println("恢复")
		case win.SC_CLOSE:
			fmt.Println("关闭")
		}
	}
	return win.CallWindowProc(oldWndPrc, types.HWND(hWnd), message, wParam, lParam)
}
