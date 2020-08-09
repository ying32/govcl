package main

import (
	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

type TPoint struct {
	X, Y int32
	Down bool
}

var (
	isMouseDown bool
	points      = make([]TPoint, 0)
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	jpgFileName := "./imgs/1.jpg"
	canLoad := rtl.FileExists(jpgFileName)
	var jpgimg *vcl.TJPEGImage
	if canLoad {
		jpgimg = vcl.NewJPEGImage()
		jpgimg.LoadFromFile(jpgFileName)
	}

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(600)
	mainForm.SetDoubleBuffered(true)
	mainForm.SetOnDestroy(func(sender vcl.IObject) {
		if jpgimg != nil {
			jpgimg.Free()
		}
	})

	mainForm.SetOnPaint(func(vcl.IObject) {

		canvas := mainForm.Canvas()

		canvas.MoveTo(10, 10)
		canvas.LineTo(50, 10)
		s := "这是一段文字"
		canvas.Font().SetColor(colors.ClRed) // red
		canvas.Font().SetSize(20)
		style := canvas.Font().Style()
		canvas.Brush().SetStyle(types.BsClear)
		canvas.Font().SetStyle(style.Include(types.FsBold, types.FsItalic))
		canvas.TextOut(100, 30, s)

		r := types.TRect{0, 0, 80, 80}

		// 计算文字
		//fmt.Println("TfSingleLine: ", types.TfSingleLine)
		s = "由于现有第三方的Go UI库不是太庞大就是用的不习惯，或者组件太少。"
		canvas.TextRect2(&r, s, types.NewSet(types.TfCenter, types.TfVerticalCenter, types.TfSingleLine))
		//fmt.Println("r: ", r, ", s: ", s)

		s = "测试输出"
		r = types.TRect{0, 0, 80, 80}
		// brush
		canvas.Brush().SetColor(colors.ClGreen)
		canvas.FillRect(r)

		// font
		canvas.Font().SetStyle(0)
		canvas.Font().SetSize(9)
		canvas.Font().SetColor(colors.ClBlue)

		// pen
		canvas.Pen().SetWidth(2)
		canvas.Pen().SetColor(colors.ClFuchsia)
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		//fmt.Println("format: ", textFmt)
		//		canvas.TextRect(r, 0, 0, s)
		canvas.TextRect2(&r, s, textFmt)

		canvas.Draw(0, 80, jpgimg)
		//canvas.Draw2(0, 200, jpgimg, 10)

		// 画多边形

		canvas.Brush().SetColor(colors.ClYellow)
		canvas.Polygon([]types.TPoint{{15, 40}, {43, 123}, {81, 42}, {45, 11}})

		canvas.Polyline([]types.TPoint{{15 + 100, 40}, {43 + 100, 123}, {81 + 100, 42}, {45 + 100, 11}})
	})

	paintbox := vcl.NewPaintBox(mainForm)
	paintbox.SetParent(mainForm)
	paintbox.SetAlign(types.AlBottom)
	paintbox.SetHeight(mainForm.Height() - 280)
	paintbox.SetOnPaint(func(vcl.IObject) {
		canvas := paintbox.Canvas()
		canvas.Pen().SetColor(colors.ClRed)
		r := paintbox.ClientRect()
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		canvas.Font().SetColor(colors.ClSkyblue)
		rect := paintbox.ClientRect()
		s := "在这可以用鼠标绘制"
		textFmt := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		canvas.TextRect2(&rect, s, textFmt)

		canvas.Pen().SetColor(colors.ClGreen)
		for _, p := range points {
			if p.Down {
				canvas.MoveTo(p.X, p.Y)
			} else {
				canvas.LineTo(p.X, p.Y)
			}
		}
	})

	paintbox.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse down")
		if button == types.MbLeft {
			points = append(points, TPoint{X: x, Y: y, Down: true})
			isMouseDown = true
		}
	})

	paintbox.SetOnMouseMove(func(sender vcl.IObject, shift types.TShiftState, x, y int32) {
		if isMouseDown {
			points = append(points, TPoint{X: x, Y: y, Down: false})
			paintbox.Repaint()
		}
	})

	paintbox.SetOnMouseUp(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
		fmt.Println("mouse SetOnMouseUp")
		if button == types.MbLeft {
			isMouseDown = false
		}
	})

	btnClear := vcl.NewButton(mainForm)
	btnClear.SetParent(mainForm)
	btnClear.SetCaption("清除绘制")
	btnClear.SetLeft(mainForm.Width() - btnClear.Width() - 20)
	btnClear.SetTop(10)
	btnClear.SetOnClick(func(vcl.IObject) {
		points = make([]TPoint, 0)
		paintbox.Repaint()
	})

	vcl.Application.Run()
}
