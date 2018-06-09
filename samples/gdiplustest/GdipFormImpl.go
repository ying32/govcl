// 在这里写你的事件。
// Write your event here.

package main

import (
	"fmt"

	"unsafe"

	"github.com/tryor/gdiplus"
	"github.com/tryor/winapi"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/win"
)

func (f *TGdipForm) OnFormCreate(sender vcl.IObject) {
	fmt.Println("OnCreate")
	style := win.GetWindowLongPtr(f.Handle(), win.GWL_EXSTYLE) | win.WS_EX_LAYERED | win.WS_EX_TOOLWINDOW
	win.SetWindowLongPtr(f.Handle(), win.GWL_EXSTYLE, uintptr(style))
	f.UpdateLayer()
}

func (f *TGdipForm) UpdateLayer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()
	hWd := f.Handle()
	formDC := win.GetDC(hWd)
	if formDC == 0 {
		return
	}
	defer win.ReleaseDC(hWd, formDC)
	aWidth := f.Width()
	aHeight := f.Height()

	memDC := win.CreateCompatibleDC(formDC)
	if memDC == 0 {
		return
	}
	defer win.DeleteDC(memDC)

	bmih := win.TBitmapInfo{}
	bmih.BmiHeader.BiSize = uint32(unsafe.Sizeof(bmih.BmiHeader))
	bmih.BmiHeader.BiWidth = uint32(aWidth)
	bmih.BmiHeader.BiHeight = uint32(aHeight)
	bmih.BmiHeader.BiPlanes = 1
	bmih.BmiHeader.BiBitCount = 32
	bmih.BmiHeader.BiCompression = win.BI_RGB
	bmih.BmiHeader.BiSizeImage = 0
	bmih.BmiHeader.BiXPelsPerMeter = 0
	bmih.BmiHeader.BiYPelsPerMeter = 0
	bmih.BmiHeader.BiClrUsed = 0
	bmih.BmiHeader.BiClrImportant = 0

	var pBits uintptr
	hBitMap := win.CreateDIBSection(0, &bmih, 0, &pBits, 0, 0)
	if hBitMap == 0 {
		return
	}
	defer win.DeleteObject(types.HGDIOBJ(hBitMap))
	win.SelectObject(memDC, types.HGDIOBJ(hBitMap))

	// 这里写gdi+处理

	g, err := gdiplus.FromHDC(winapi.HDC(memDC))
	if err != nil {
		return
	}
	defer g.Release()

	g.SetSmoothingMode(gdiplus.SmoothingModeAntiAlias)
	g.SetInterpolationMode(gdiplus.InterpolationModeHighQualityBicubic)
	g.SetTextRenderingHint(gdiplus.TextRenderingHintAntiAlias)

	// 自己画一个半透明的作为背景
	brush, err := gdiplus.NewSolidBrush(gdiplus.NewColor3(0xB2, 0, 0, 0))
	defer brush.Release()
	if err == nil {
		g.FillRectangle(brush, 0, 0, gdiplus.REAL(aWidth), gdiplus.REAL(aHeight))
	}

	// 使用带有透明通道的图片作为背景
	//r := f.ClientRect()
	//bmp, err := gdiplus.NewBitmap("./testbk.png")
	//defer bmp.Release()
	//if err != nil {
	//	return
	//}
	//fmt.Println(bmp.GetWidth(), ", ", bmp.GetHeight())
	//destR := gdiplus.Rect{0, 0, winapi.INT(r.Width()), winapi.INT(r.Height())}
	//g.DrawImageI7(bmp, &destR, 0, 0, winapi.INT(bmp.GetWidth()-1), winapi.INT(bmp.GetHeight()-1), gdiplus.UnitPixel, nil, nil, 0)

	// 画文字
	//path, _ := gdiplus.NewGraphicsPath()
	//defer path.Release()

	family, _ := gdiplus.NewFontFamily("黑体", nil)
	defer family.Release()

	strFormat, _ := gdiplus.NewStringFormat()
	defer strFormat.Release()
	strFormat.SetFormatFlags(winapi.INT(gdiplus.StringFormatFlagsNoWrap))
	strFormat.SetAlignment(gdiplus.StringAlignmentCenter)

	// 红色
	f.DrawText("文字String", 0, g, family, strFormat, gdiplus.NewColor3(0xFF, 0xB3, 0x00, 0x00), gdiplus.NewColor3(0xFF, 0xFE, 0x95, 0x00))
	// 蓝色
	f.DrawText("文字String", 200, g, family, strFormat, gdiplus.NewColor3(0xFF, 0x00, 0x43, 0x93), gdiplus.NewColor3(0xFF, 0x02, 0xB4, 0xEA))
	// 绿色
	f.DrawText("文字String", 400, g, family, strFormat, gdiplus.NewColor3(0xFF, 0x0B, 0x63, 0x00), gdiplus.NewColor3(0xFF, 0x8A, 0xF6, 0x22))
	// 更新分层窗口
	winsize := types.TSize{aWidth, aHeight}
	scrPoint := types.TPoint{}
	blend := win.TBlendFunction{}
	blend.BlendOp = win.AC_SRC_OVER
	blend.BlendFlags = 0
	blend.AlphaFormat = win.AC_SRC_ALPHA
	blend.SourceConstantAlpha = 255

	ret := win.UpdateLayeredWindow(hWd, formDC, nil, &winsize, memDC, &scrPoint, 0, &blend, win.ULW_ALPHA)
	fmt.Println("更新结果：", ret)
}

func (f *TGdipForm) DrawText(s string, top int, g *gdiplus.Graphics, family *gdiplus.FontFamily, strFormat *gdiplus.StringFormat, color1, color2 gdiplus.Color) {
	rF := gdiplus.RectF{0, gdiplus.REAL(top), gdiplus.REAL(f.Width()), 0}
	var fontHeight gdiplus.REAL = 100

	path, _ := gdiplus.NewGraphicsPath()
	defer path.Release()

	path.AddString(s, family, gdiplus.FontStyleBold, fontHeight, &rF, strFormat)

	// 画笔
	pen, err := gdiplus.NewPen(gdiplus.NewColor3(155, 215, 215, 215), 3)
	defer pen.Release()
	if err != nil {
		return
	}
	pen.SetColor(gdiplus.NewColor3(65, 1, 3, 3))
	pen.SetLineJoin(gdiplus.LineJoinRound)
	g.DrawPath(pen, path)

	brush2, err := gdiplus.NewLinearGradientBrush(
		gdiplus.NewPointF(0, 0),
		gdiplus.NewPointF(0, fontHeight), color1, color2)
	defer brush2.Release()
	if err != nil {
		return
	}

	// 多绘制几次
	for i := 0; i < 4; i++ {
		pen.SetWidth(gdiplus.REAL(i))
		g.DrawPath(pen, path)
	}
	g.FillPath(brush2, path)
}

func (f *TGdipForm) OnFormMouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if button == types.MbLeft {
		win.ReleaseCapture()
		f.Perform(win.WM_SYSCOMMAND, win.SC_MOVE+1, 0)
	}
}
