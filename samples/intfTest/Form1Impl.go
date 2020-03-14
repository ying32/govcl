// 在这里写你的事件

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
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
	//

	dc := rtl.GetDC(0)         // windows桌面
	defer rtl.ReleaseDC(0, dc) // 一定要释放
	// canvas
	canvas := vcl.NewCanvas()
	defer canvas.Free()

	canvas.SetHandle(dc)

	w := vcl.Screen.Width()
	h := vcl.Screen.Height()

	bmp := vcl.NewBitmap()
	defer bmp.Free()
	bmp.SetSize(w, h)
	bmp.SetPixelFormat(types.Pf32bit)

	bmp.Canvas().CopyRect(types.Rect(0, 0, w, h), canvas, types.Rect(0, 0, w, h))

	f.Image1.Picture().Assign(bmp)
}
