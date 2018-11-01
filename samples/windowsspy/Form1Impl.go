// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

var (
	lastRect types.TRect
	lasthWnd types.HWND
)

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.SetLeft(vcl.Screen.Width() - f.Width() - 5)
	f.SetTop(vcl.Screen.WorkAreaHeight() - f.Height() - 5)
	vcl.Screen.SetCursors(1, f.Img3.Picture().Icon().Handle())
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	vcl.Screen.SetCursor(types.CrDefault)
	f.clearDesktopRect()
}

func (f *TForm1) OnImg1MouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if rtl.InSets(shift, types.SsLeft) {
		lasthWnd = 0 //win.WindowFromPoint(vcl.Mouse.CursorPos())
		vcl.Screen.SetCursor(types.TCursor(1))
		f.Img2.SetVisible(true)
		f.Img1.SetVisible(false)
	}
}

func (f *TForm1) OnImg1MouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	vcl.Screen.SetCursor(types.CrDefault)
	f.Img1.SetVisible(true)
	f.Img2.SetVisible(false)
	f.clearDesktopRect()
	lasthWnd = 0
}

func (f *TForm1) clearDesktopRect() {
	if lastRect.IsEmpty() {
		return
	}
	f.drawRect(lastRect)
	lastRect.Empty()
}

func (f *TForm1) drawRect(rc types.TRect) {
	dc := win.GetDC(0)
	defer win.ReleaseDC(0, dc)
	newPenDC := win.CreatePen(0, 3, 0xFF)
	defer win.DeleteObject(newPenDC)
	oldRop := win.SetROP2(dc, win.R2_NOTXORPEN)
	oldPen := win.SelectObject(dc, newPenDC)
	win.Rectangle(dc, rc.Left, rc.Top, rc.Right, rc.Bottom)
	win.SelectObject(dc, oldPen)
	win.SetROP2(dc, oldRop)
}

func (f *TForm1) drawDesktopRect(rc types.TRect) {
	f.clearDesktopRect()
	f.drawRect(rc)
	lastRect = rc
}

func (f *TForm1) OnImg1MouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	if rtl.InSets(shift, types.SsLeft) {
		pt, _ := win.GetCursorPos2()
		hWnd := win.WindowFromPoint(pt)
		if hWnd != lasthWnd {
			lasthWnd = hWnd
			if hWnd != f.Handle() && !win.IsChild(f.Handle(), hWnd) {

				caption, _ := win.GetWindowText(hWnd)
				className, _ := win.GetClassName(hWnd)
				threadId, processId := win.GetWindowThreadProcessId2(hWnd)
				wndStyle := win.GetWindowLong(hWnd, win.GWL_STYLE)
				wRect, _ := win.GetWindowRect2(hWnd)

				f.Edit_Caption.SetText(caption)
				f.Edit_Handle.SetText(fmt.Sprintf("%d", hWnd))
				f.Edit_ClassName.SetText(className)
				f.Edit_Style.SetText(fmt.Sprintf("%d", wndStyle))
				f.Edit_ThreadId.SetText(fmt.Sprintf("%d", threadId))
				f.Edit_Pid.SetText(fmt.Sprintf("%d", processId))
				f.Edit_Width.SetText(fmt.Sprintf("%dx%d", wRect.Width(), wRect.Height()))
				f.Edit_Top.SetText(fmt.Sprintf("%dx%d", wRect.Left, wRect.Top))

				f.drawDesktopRect(wRect)
			}
		}
	}
}

func (f *TForm1) OnChkonTopClick(sender vcl.IObject) {

	if f.ChkonTop.Checked() {
		f.SetFormStyle(types.FsStayOnTop)
	} else {
		f.SetFormStyle(types.FsNormal)
	}
	// 好吧，不生效了。。。
	//if f.ChkonTop.Checked() {
	//	fmt.Println(win.SetWindowPos(f.Handle(), win.HWND_TOPMOST, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE))
	//} else {
	//	win.SetWindowPos(f.Handle(), win.HWND_TOP, 0, 0, 0, 0, win.SWP_NOMOVE|win.SWP_NOSIZE)
	//}
}

func (f *TForm1) OnChkSwitchHexClick(sender vcl.IObject) {

}
