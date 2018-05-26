package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
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
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	jpgimg := vcl.NewJPEGImage()
	defer jpgimg.Free()
	jpgimg.LoadFromFile("..//../imgs/1.jpg")

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(400)
	mainForm.SetHeight(600)
	mainForm.SetDoubleBuffered(true)

	mainForm.SetOnPaint(func(vcl.IObject) {

		canvas := mainForm.Canvas()

		canvas.MoveTo(10, 10)
		canvas.LineTo(50, 10)
		s := "这是一段文字"
		canvas.Font().SetColor(types.ClRed) // red
		canvas.Font().SetSize(20)
		style := canvas.Font().Style()
		canvas.Font().SetStyle(types.TFontStyles(rtl.Include(uint32(style), types.FsBold, types.FsItalic)))
		canvas.TextOut(100, 30, s)

		r := types.TRect{0, 0, 80, 80}

		// 计算文字
		//fmt.Println("TfSingleLine: ", types.TfSingleLine)
		s = "由于现有第三方的Go UI库不是太宠大就是用的不习惯，或者组件太少。"
		canvas.TextRect2(&r, &s, types.TTextFormat(rtl.Include(0,
			types.TfCenter, types.TfVerticalCenter, types.TfSingleLine)))
		//fmt.Println("r: ", r, ", s: ", s)

		s = "测试输出"
		r = types.TRect{0, 0, 80, 80}
		// brush
		canvas.Brush().SetColor(types.ClGreen)
		canvas.FillRect(r)

		// font
		canvas.Font().SetStyle(0)
		canvas.Font().SetSize(9)
		canvas.Font().SetColor(types.ClBlue)

		// pen
		canvas.Pen().SetWidth(2)
		canvas.Pen().SetColor(types.ClFuchsia)
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		textFmt := rtl.Include(0, types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
		//fmt.Println("format: ", textFmt)
		//		canvas.TextRect(r, 0, 0, s)
		canvas.TextRect2(&r, &s, types.TTextFormat(textFmt))

		canvas.Draw(0, 80, jpgimg)
		//canvas.Draw2(0, 200, jpgimg, 10)

	})

	paintbox := vcl.NewPaintBox(mainForm)
	paintbox.SetParent(mainForm)
	paintbox.SetAlign(types.AlBottom)
	paintbox.SetHeight(mainForm.Height() - 280)
	paintbox.SetOnPaint(func(vcl.IObject) {
		canvas := paintbox.Canvas()
		canvas.Pen().SetColor(types.ClRed)
		r := paintbox.ClientRect()
		canvas.Rectangle(r.Left, r.Top, r.Right, r.Bottom)

		canvas.Font().SetColor(types.ClSkyblue)
		rect := paintbox.ClientRect()
		s := "在这可以用鼠标绘制"
		textFmt := types.TTextFormat(rtl.Include(0, types.TfCenter, types.TfSingleLine, types.TfVerticalCenter))
		canvas.TextRect2(&rect, &s, textFmt)

		canvas.Pen().SetColor(types.ClGreen)
		for _, p := range points {
			if p.Down {
				canvas.MoveTo(p.X, p.Y)
			} else {
				canvas.LineTo(p.X, p.Y)
			}
		}
	})

	paintbox.SetOnMouseDown(func(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
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
