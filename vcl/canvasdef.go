package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// BrushCopy 画刷复制
func (c *TCanvas) BrushCopy(dest TRect, bitmap IObject, source TRect, color TColor) {
	Canvas_BrushCopy(c.instance, dest, CheckPtr(bitmap), source, color)
}

// CopyRect 复制矩形
func (c *TCanvas) CopyRect(dest TRect, canvas IObject, source TRect) {
	Canvas_CopyRect(c.instance, dest, CheckPtr(canvas), source)
}

// Draw 绘制graphic对象
func (c *TCanvas) Draw(x, y int32, graphic IObject) {
	Canvas_Draw1(c.instance, x, y, CheckPtr(graphic))
}

// Draw2 绘制graphic对象，opacity透明度
func (c *TCanvas) Draw2(x, y int32, graphic IObject, opacity uint8) {
	Canvas_Draw2(c.instance, x, y, CheckPtr(graphic), opacity)
}

// DrawFocusRect 画焦点矩形
func (c *TCanvas) DrawFocusRect(aRect TRect) {
	Canvas_DrawFocusRect(c.instance, aRect)
}

// FillRect 填充矩形
func (c *TCanvas) FillRect(aRect TRect) {
	Canvas_FillRect(c.instance, aRect)
}

// FrameRect 绘制边框
func (c *TCanvas) FrameRect(aRect TRect) {
	Canvas_FrameRect(c.instance, aRect)
}

// StretchDraw 拉伸绘制
func (c *TCanvas) StretchDraw(aRect TRect, graphic IObject) {
	Canvas_StretchDraw(c.instance, aRect, CheckPtr(graphic))
}

// TextRect 在矩形内绘制文字
func (c *TCanvas) TextRect(aRect TRect, x, y int32, text string) {
	Canvas_TextRect1(c.instance, aRect, x, y, text)
}

// TextRect2 在矩形内绘制文字
func (c *TCanvas) TextRect2(aRect *TRect, text *string, textFormat TTextFormat) {
	Canvas_TextRect2(c.instance, aRect, text, textFormat)
}

// TextRect3 在矩形内绘制文字
func (c *TCanvas) TextRect3(aRect *TRect, text string, textFormat TTextFormat) {
	Canvas_TextRect3(c.instance, aRect, text, textFormat)
}
