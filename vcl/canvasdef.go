package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

func (c *TCanvas) BrushCopy(dest TRect, bitmap IObject, source TRect, color TColor) {
    Canvas_BrushCopy(c.instance, dest, CheckPtr(bitmap), source, color)
}
 
func (c *TCanvas) CopyRect(dest TRect, canvas IObject, source TRect) {
    Canvas_CopyRect(c.instance, dest, CheckPtr(canvas), source)
}

func (c *TCanvas) Draw(x, y int32, graphic IObject) {
    Canvas_Draw1(c.instance, x, y, CheckPtr(graphic))
}

func (c *TCanvas) Draw2 (x, y int32, graphic IObject, opacity uint8) {
    Canvas_Draw2(c.instance, x, y, CheckPtr(graphic), opacity)
}

func (c *TCanvas) DrawFocusRect(aRect TRect) {
    Canvas_DrawFocusRect(c.instance, aRect)
}

func (c *TCanvas) FillRect(aRect TRect) {
    Canvas_FillRect(c.instance, aRect)
}

func (c *TCanvas) FrameRect (aRect TRect) {
    Canvas_FrameRect(c.instance, aRect)
}

func (c *TCanvas) StretchDraw(aRect TRect, graphic IObject) {
    Canvas_StretchDraw(c.instance, aRect, CheckPtr(graphic))
}

func (c *TCanvas) TextRect(aRect TRect, x, y int32, text string) {
    Canvas_TextRect1(c.instance, aRect, x, y, text)
}

func (c *TCanvas) TextRect2(aRect *TRect, text *string, textFormat TTextFormat) {
    Canvas_TextRect2(c.instance, aRect, text, textFormat)
}

