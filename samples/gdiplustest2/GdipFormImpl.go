// 在这里写你的事件。
// Write your event here.

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/tryor/gdiplus"
	"github.com/tryor/winapi"
)

func (f *TGdipForm) OnGdipFormCreate(sender vcl.IObject) {
	fmt.Println("OnCreate")

}

// 改自Delphi IGDIPlus例程

func GetStringBoundingBoxF(g *gdiplus.Graphics, str string, font *gdiplus.Font, origin *gdiplus.PointF) *gdiplus.RectF {

	var nFont *gdiplus.Font
	if font.IsAvailable() {
		nFont = font
	}
	rect := &gdiplus.RectF{origin.X, origin.Y, 0.0, 0.0}

	r, s := g.MeasureString4(str, nFont, rect)
	if s == gdiplus.Ok {
		return r
	}
	return nil
}

func GPInflateRectF(rect *gdiplus.RectF, cx, cy gdiplus.REAL) (result *gdiplus.RectF) {
	result = new(gdiplus.RectF)
	result.X = rect.X - cx
	result.Y = rect.Y - cy
	result.Width = rect.Width + cx*2
	result.Height = rect.Height + cy*2
	return
}

func AddRectangleF(path *gdiplus.GraphicsPath, rect *gdiplus.RectF) {
	path.AddRectangle(rect)
}

func AddRoundRectangleF(path *gdiplus.GraphicsPath, rect *gdiplus.RectF, cornerSize *gdiplus.SizeF) {
	if rect.Width == 0 || rect.Height == 0 {
		return
	}
	cornerSizeX := cornerSize
	if cornerSizeX.Width < 0 {
		cornerSizeX.Width = 0
	}
	if cornerSizeX.Height < 0 {
		cornerSizeX.Height = 0
	}
	if cornerSizeX.Width == 0 || cornerSizeX.Height == 0 {
		AddRectangleF(path, rect)
	}
	cornerSizeX.Width = cornerSizeX.Width * 2
	cornerSizeX.Height = cornerSizeX.Height * 2
	rectRight := rect.X + rect.Width
	if cornerSizeX.Width > rect.Width {
		cornerSizeX.Width = rect.Width
	}
	if cornerSizeX.Height > rect.Height {
		cornerSizeX.Height = rect.Height
	}
	path.StartFigure()
	path.AddArc(rect.X, rect.Y, cornerSizeX.Width, cornerSizeX.Height, 180, 90)
	path.AddArc(rectRight-cornerSizeX.Width, rect.Y, cornerSizeX.Width, cornerSizeX.Height, 270, 90)
	path.AddArc(rectRight-cornerSizeX.Width, rect.Y+rect.Height-cornerSizeX.Height, cornerSizeX.Width, cornerSizeX.Height, 0, 90)
	path.AddArc(rect.X, rect.Y+rect.Height-cornerSizeX.Height, cornerSizeX.Width, cornerSizeX.Height, 90, 90)
	path.CloseFigure()
}

//func DrawRoundRectangleF(g *gdiplus.Graphics, path *gdiplus.GraphicsPath, pen *gdiplus.Pen) {
//
//}
//
//func FillRoundRectangleF(g *gdiplus.Graphics, path *gdiplus.GraphicsPath, brush gdiplus.IBrush) {
//
//}

const (
	aclGreen  = 0xFF008000
	aclCyan   = 0xFF00FFFF
	aclYellow = 0xFFFFFF00
	aclRed    = 0xFFFF0000
	aclBlue   = 0xFF0000FF
)

func (f *TGdipForm) UpdateLayer() {
	defer func() {
		if err := recover(); err != nil {
			fmt.Println("err:", err)
		}
	}()

	//aWidth := f.Width()
	//aHeight := f.Height()

	// 这里写gdi+处理

	g, err := gdiplus.FromHDC(winapi.HDC(f.Canvas().Handle()))
	if err != nil {
		return
	}
	defer g.Release()

	g.SetSmoothingMode(gdiplus.SmoothingModeAntiAlias)
	//g.SetInterpolationMode(gdiplus.InterpolationModeHighQualityBicubic)
	g.SetTextRenderingHint(gdiplus.TextRenderingHintAntiAlias)

	leftTopCorner := gdiplus.NewPointF(20.0, 20.0)
	font, err := gdiplus.NewFont2("Microsoft Sans Serif", 40, gdiplus.FontStyleBold, gdiplus.UnitPixel, nil)
	defer font.Release()
	if err != nil {
		return
	}

	rect := GetStringBoundingBoxF(g, "Welcome to GDI+", font, leftTopCorner)
	rect1 := GPInflateRectF(rect, 10, 10)

	// 画笔
	brush, _ := gdiplus.NewLinearGradientBrushFromRect(GPInflateRectF(rect1, 2, 2), gdiplus.NewColor(aclRed), gdiplus.NewColor(aclBlue), gdiplus.LinearGradientModeVertical)
	defer brush.Release()
	pen, err := gdiplus.NewPen2(brush, 4)
	defer pen.Release()

	path, _ := gdiplus.NewGraphicsPath()
	defer path.Release()

	AddRoundRectangleF(path, rect1, gdiplus.NewSizeF(20, 20))

	pathbrush, _ := gdiplus.NewPathGradientBrush2(path)
	defer pathbrush.Release()
	pathbrush.SetInterpolationColors([]gdiplus.Color{gdiplus.NewColor(aclGreen), gdiplus.NewColor(aclCyan), gdiplus.NewColor(aclYellow)}, []gdiplus.REAL{0, 0.3, 1})
	pathbrush.SetCenterPoint(gdiplus.NewPointF(250, 50))
	pathbrush.SetFocusScales(0.87, 0.2)

	g.FillPath(pathbrush, path)
	g.DrawPath(pen, path)

	// Draw a text with semitransparent shadow.
	sbrush, _ := gdiplus.NewSolidBrush(gdiplus.NewColor3(50, 0, 0, 0))
	defer sbrush.Release()
	r2 := gdiplus.NewRectF(23, 23, 0, 0)
	g.DrawString("Welcome to GDI+", font, r2, nil, sbrush)
	r2 = gdiplus.NewRectF(leftTopCorner.Y, leftTopCorner.Y, 0, 0)
	sbrush2, _ := gdiplus.NewLinearGradientBrushFromRect(rect, gdiplus.NewColor(aclRed), gdiplus.NewColor(aclBlue), gdiplus.LinearGradientModeForwardDiagonal)
	defer sbrush2.Release()
	g.DrawString("Welcome to GDI+", font, r2, nil, sbrush2)

	// Draw a closed curve.
	pen2, _ := gdiplus.NewPen(gdiplus.NewColor(aclRed), 3)
	defer pen2.Release()
	sbrush3, _ := gdiplus.NewSolidBrush(gdiplus.NewColor(aclBlue))
	defer sbrush3.Release()

	ps := []gdiplus.PointF{gdiplus.PointF{60, 160}, gdiplus.PointF{150, 180},
		gdiplus.PointF{200, 140}, gdiplus.PointF{180, 220}, gdiplus.PointF{120, 200}, gdiplus.PointF{80, 260}}
	g.FillClosedCurve(sbrush3, ps)
	g.DrawClosedCurve(pen2, ps)

	// Draw a semitransparent star.
	path2, _ := gdiplus.NewGraphicsPath()
	defer path2.Release()
	path2.AddLines([]gdiplus.PointF{gdiplus.PointF{75, 0}, gdiplus.PointF{100, 50}, gdiplus.PointF{150, 50},
		gdiplus.PointF{112, 75}, gdiplus.PointF{150, 150}, gdiplus.PointF{75, 100}, gdiplus.PointF{0, 150},
		gdiplus.PointF{37, 75}, gdiplus.PointF{0, 50}, gdiplus.PointF{50, 50}})

	g.TranslateTransform(420, 30)
	sbrush4, _ := gdiplus.NewPathGradientBrush2(path2)
	defer sbrush4.Release()
	sbrush4.SetCenterColor(gdiplus.NewColor3(200, 255, 0, 0))
	sbrush4.SetSurroundColors([]gdiplus.Color{gdiplus.NewColor3(80, 0, 0, 0),
		gdiplus.NewColor3(80, 0, 255, 0), gdiplus.NewColor3(80, 0, 0, 255),
		gdiplus.NewColor3(80, 255, 255, 255), gdiplus.NewColor3(80, 0, 0, 0),
		gdiplus.NewColor3(80, 0, 255, 0), gdiplus.NewColor3(80, 0, 0, 255),
		gdiplus.NewColor3(80, 255, 255, 255), gdiplus.NewColor3(80, 0, 0, 0),
		gdiplus.NewColor3(80, 0, 255, 0)})
	g.FillPath(sbrush4, path2)

	// Draw rotated ellipse.
	g.ResetTransform()
	g.TranslateTransform(300, 160)
	g.RotateTransform(30)
	pen4, _ := gdiplus.NewPen(gdiplus.NewColor(aclRed), 3)
	defer pen4.Release()
	sbrush5, _ := gdiplus.NewLinearGradientBrush(gdiplus.NewPointF(0, 0), gdiplus.NewPointF(20, 20), gdiplus.NewColor(aclYellow), gdiplus.NewColor(aclGreen))
	defer sbrush5.Release()
	sbrush5.SetWrapMode(gdiplus.WrapModeTileFlipX)
	g.FillEllipse(sbrush5, 0, 0, 200, 80)
	g.DrawEllipse(pen4, 0, 0, 200, 80)

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

func (f *TGdipForm) OnGdipFormPaint(vcl.IObject) {
	fmt.Println("OnPaint")
	f.UpdateLayer()
}
