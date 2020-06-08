// +build darwin

package main

/*

   #cgo CFLAGS: -x objective-c
   #cgo LDFLAGS: -framework Cocoa

   #include <Cocoa/Cocoa.h>

  static NSWindow* toNSWindow(void* ptr) {
      return ((__bridge NSWindow*)ptr); //
  }

  static void resizeAction(void *ptr) {
     if(ptr == nil) return;
     // MRC方式
     // NSAutoreleasePool *pool = [[NSAutoreleasePool alloc] init];
     // [pool release];
     // ARC方式。 以防万一，我还是这么加个，也不知道这个有效还是MRC方式有效果。。。
     @autoreleasepool {
		 //NSLog(@"resize\n");
		 NSWindow *win = toNSWindow(ptr);
		 // 修改系统按钮位置
		 NSView *themeFrameView = win.contentView.superview;
		 NSView *titleBarView = themeFrameView.subviews.count > 1 ? themeFrameView.subviews[1] : nil;
		 if(titleBarView == nil) return;
		 //NSLog(@"titleBarView.superviews: %@\n", titleBarView.superview);

		 // 原始的方式操作。。。
		 //NSRect r = win.frame;
		 //titleBarView.autoresizesSubviews = YES;
		 //[titleBarView setFrame: CGRectMake(10, r.size.height - 40, 60, 30)];

		 // VFL 方式约束
		 NSDictionary *views = NSDictionaryOfVariableBindings(titleBarView);
		 titleBarView.translatesAutoresizingMaskIntoConstraints = NO;
		 // 横向，距离父视图左边8位置，titleView视图宽为60
		 [themeFrameView addConstraints: [NSLayoutConstraint constraintsWithVisualFormat: @"H:|-8-[titleBarView(60)]" options:0 metrics:nil views: views]];
		 // 纵向，距离父视图顶部11位置，titleView视图高为30
		 [themeFrameView addConstraints: [NSLayoutConstraint constraintsWithVisualFormat: @"V:|-11-[titleBarView(30)]" options:0 metrics:nil views: views]];

     }
  }

  static void setNoTitlebarWindow(void* ptr) {
     if(ptr == nil) return;
     @autoreleasepool{
		 NSWindow *win = toNSWindow(ptr);
		 win.movableByWindowBackground = YES;
		 win.titleVisibility = NSWindowTitleHidden;
		 win.titlebarAppearsTransparent = YES;
		 win.representedURL = nil;
		 win.styleMask = win.styleMask | NSWindowStyleMaskFullSizeContentView;
     }
  }

*/
import "C"

import (
	"fmt"
	"unsafe"

	"github.com/ying32/govcl/vcl/rtl"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
	isDown    bool
	downPoint types.TPoint
	nsWinPtr  unsafe.Pointer
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {

	f.ScreenCenter()
	nsWin := f.PlatformWindow()
	f.nsWinPtr = unsafe.Pointer(nsWin)

	fmt.Println("nsWindow:", nsWin)

	// 用原生objective-c代码操作
	C.setNoTitlebarWindow(f.nsWinPtr)
	// 约束位置
	C.resizeAction(f.nsWinPtr)

	//-----------------
	// 不显示标题栏，用lcl里面提供函数
	//nsWin.SetTitleVisibility(types.NSWindowTitleHidden)
	// 标题栏上面透明
	//nsWin.SetTitleBarAppearsTransparent(true)
	//	// 不显示icon
	//nsWin.SetRepresentedURL(0)
	//	// 视图的类型，应该是系统按钮什么的都算在主体视图中了
	//nsWin.SetStyleMask(nsWin.StyleMask() | types.NSWindowStyleMaskFullSizeContentView)

	// 先不显示
	//f.Panel1.Hide()
	//f.Panel2.Hide()
	//f.Panel3.Hide()
	f.Panel1.SetOnAlignPosition(f.onCenterAlign)

}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {

}

func (f *TForm1) OnFormShow(sender vcl.IObject) {

}

func (f *TForm1) OnFormResize(sender vcl.IObject) {
	// 应该是监视最大化和还原，
	C.resizeAction(f.nsWinPtr)
}

func (f *TForm1) onCenterAlign(sender *vcl.TWinControl, control *vcl.TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *types.TRect, alignInfo types.TAlignInfo) {
	*newLeft = (alignRect.Width() - *newWidth) / 2
	*newTop = (alignRect.Height() - *newHeight) / 2
}

// lcl好像是有个属性可以判断的，但不知道有没有效果，暂暂时也没提供接口，所以就这样凑合着用吧
func (f *TForm1) isFullScreen() bool {
	return f.Width() == vcl.Screen.Width() && f.Height() == vcl.Screen.Height()
}

func (f *TForm1) OnPanel1DblClick(sender vcl.IObject) {
	// 全屏了
	if f.isFullScreen() {
		return
	}
	// 有bug。。。。
	// 换个方式吧
	if f.WindowState() != types.WsNormal {
		f.SetWindowState(types.WsNormal)
	} else {
		f.SetWindowState(types.WsMaximized)
	}
}

func (f *TForm1) OnPanel1MouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if f.isFullScreen() {
		return
	}
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
