// +build darwin

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/rtl"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
	isDown    bool
	downPoint types.TPoint
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	nsWin := f.PlatformWindow()

	fmt.Println("nsWindow:", nsWin)
	//	// 不显示标题栏
	nsWin.SetTitleVisibility(types.NSWindowTitleHidden)
	// 标题栏上面透明
	nsWin.SetTitleBarAppearsTransparent(true)
	//	// 不显示icon
	nsWin.SetRepresentedURL(0)
	//	// 视图的类型，应该是系统按钮什么的都算在主体视图中了
	nsWin.SetStyleMask(nsWin.StyleMask() | types.NSWindowStyleMaskFullSizeContentView)

	f.Panel1.SetOnAlignPosition(f.onCenterAlign)
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {

}

func (f *TForm1) onCenterAlign(sender *vcl.TWinControl, control *vcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
	*newLeft = (alignRect.Width() - *newWidth) / 2
	*newTop = (alignRect.Height() - *newHeight) / 2
}

/*
Lwin.setTitleVisibility(NSWindowTitleHidden);
Lwin.SetTitlebarAppearsTransparent(True);
Lwin.setStyleMask(Lwin.styleMask or NSWindowStyleMaskFullSizeContentView);
Lwin.setRepresentedURL(nil);
*/

func (f *TForm1) OnPanel1Click(sender vcl.IObject) {

}

func (f *TForm1) OnPanel1MouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft && !rtl.IsZoomed(f.Handle()) {
		f.isDown = true
		f.downPoint = types.Point(x, y)
	}
}

func (f *TForm1) OnPanel1MouseMove(sender vcl.IObject, shift types.TShiftState, x, y int32) {
	if f.isDown {
		f.SetLeft(f.Left() + (x - f.downPoint.X))
		f.SetTop(f.Top() + (y - f.downPoint.Y))
	}
}

func (f *TForm1) OnPanel1MouseUp(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	f.isDown = false
}
