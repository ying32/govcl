//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// 画刷复制
func (c *TCanvas) BrushCopy(dest TRect, bitmap IObject, source TRect, color TColor) {
	Canvas_BrushCopy(c.instance, dest, CheckPtr(bitmap), source, color)
}

// 复制矩形
func (c *TCanvas) CopyRect(dest TRect, canvas IObject, source TRect) {
	Canvas_CopyRect(c.instance, dest, CheckPtr(canvas), source)
}

// 绘制graphic对象
func (c *TCanvas) Draw(x, y int32, graphic IGraphic) {
	Canvas_Draw1(c.instance, x, y, CheckPtr(graphic))
}

// 绘制graphic对象，opacity透明度
func (c *TCanvas) Draw2(x, y int32, graphic IGraphic, opacity uint8) {
	Canvas_Draw2(c.instance, x, y, CheckPtr(graphic), opacity)
}

// 画焦点矩形
func (c *TCanvas) DrawFocusRect(aRect TRect) {
	Canvas_DrawFocusRect(c.instance, aRect)
}

// 填充矩形
func (c *TCanvas) FillRect(aRect TRect) {
	Canvas_FillRect(c.instance, aRect)
}

// 绘制边框
func (c *TCanvas) FrameRect(aRect TRect) {
	Canvas_FrameRect(c.instance, aRect)
}

// 在矩形内绘制文字
func (c *TCanvas) TextRect(aRect TRect, x, y int32, text string) {
	Canvas_TextRect1(c.instance, aRect, x, y, text)
}

// 在矩形内绘制文字
func (c *TCanvas) TextRect2(aRect *TRect, text string, textFormat TTextFormat) {
	Canvas_TextRect2(c.instance, aRect, text, textFormat)
}

// 填充多边形
func (c *TCanvas) Polygon(points []TPoint) {
	Canvas_Polygon(c.instance, points)
}

// 画多边形，不填充
func (c *TCanvas) Polyline(points []TPoint) {
	Canvas_Polyline(c.instance, points)
}

// 多边形贝塞尔曲线
func (c *TCanvas) PolyBezier(points []TPoint) {
	Canvas_PolyBezier(c.instance, points)
}
