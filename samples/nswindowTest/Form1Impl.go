// +build darwin

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()

}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {

}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	nsWin := f.Window()
	// 不显示标题栏
	nsWin.SetTitleVisibility(types.NSWindowTitleHidden)
	// 标题栏上面透明
	nsWin.SetTitleBarAppearsTransparent(true)
	// 不显示icon
	nsWin.SetRepresentedURL(0)
	// 视图的类型，应该是系统按钮什么的都算在主体视图中了
	nsWin.SetStyleMask(nsWin.StyleMask() | types.NSWindowStyleMaskFullSizeContentView)
}

/*
Lwin.setTitleVisibility(NSWindowTitleHidden);
Lwin.SetTitlebarAppearsTransparent(True);
Lwin.setStyleMask(Lwin.styleMask or NSWindowStyleMaskFullSizeContentView);
Lwin.setRepresentedURL(nil);
*/
