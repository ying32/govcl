// 在这里写你的事件

package main

import (
	"math"
	"time"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.drawClock()
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {

}

func (f *TForm1) OnPaintBox1Paint(sender vcl.IObject) {
	f.drawClock()
}

func (f *TForm1) OnTimer1Timer(sender vcl.IObject) {
	f.PaintBox1.Repaint()
}

func (f *TForm1) drawClock() {
	canvas := f.PaintBox1.Canvas()
	pen := canvas.Pen()
	brush := canvas.Brush()

	brush.SetColor(colors.ClSkyblue)
	brush.SetStyle(types.BsClear)

	// 圆线宽和颜色
	pen.SetColor(colors.ClSkyblue)
	pen.SetWidth(2)

	r := f.PaintBox1.ClientRect()
	cp := r.CenterPoint()

	// 画圆
	canvas.Ellipse(1, 1, r.Width()-1, r.Height()-1)

	angle := 0.0

	// 画刻度
	for i := 0; i < 60; i++ {
		var len int
		//刻度，每五分钟一次
		if i%5 == 0 {
			len = 12
			pen.SetWidth(2)
		} else {
			len = 6
			pen.SetWidth(1)
		}
		pen.SetColor(colors.ClTeal)

		sline := cp.X - pen.Width()

		canvas.MoveTo(int32(math.Trunc(float64(cp.X)+float64(sline)*math.Cos(angle))),
			int32(math.Trunc(float64(cp.Y)+float64(sline)*math.Sin(angle))))

		sline -= int32(len)

		canvas.LineTo(int32(math.Trunc(float64(cp.X)+float64(sline)*math.Cos(angle))),
			int32(math.Trunc(float64(cp.Y)+float64(sline)*math.Sin(angle))))

		// 角度
		angle += math.Pi / 30.0
	}

	// 画12、3、6、9
	canvas.Font().SetSize(13)
	// 12
	canvas.TextOut(r.Width()/2-canvas.TextWidth("12")/2, 20, "12")
	// 3
	canvas.TextOut(r.Width()-32, cp.Y-canvas.TextHeight("3")/2, "3")
	// 6
	canvas.TextOut(r.Width()/2-canvas.TextWidth("6")/2, r.Height()-32-canvas.TextHeight("6")/2, "6")
	// 9
	canvas.TextOut(22, cp.Y-canvas.TextHeight("9")/2, "9")

	// 时分秒
	h, m, s := time.Now().Clock()

	hh := 2.0 * math.Pi * (float64(h) + float64(m)/60.0) / 12.0
	mm := float64(m) / 60.0 * 2.0 * math.Pi
	ss := float64(s) / 60.0 * 2.0 * math.Pi

	pl := float64(r.Width() / 2)

	// 开始画指针
	// 时针
	pen.SetColor(colors.ClRed)
	pen.SetWidth(2)
	canvas.MoveTo(cp.X, cp.Y)
	canvas.LineTo(int32(math.Trunc(float64(cp.X)+math.Sin(hh)*(pl/3))),
		int32(math.Trunc(float64(cp.Y)-math.Cos(hh)*(pl/3))))

	// 分针
	pen.SetColor(colors.ClFuchsia)
	pen.SetWidth(2)
	canvas.MoveTo(cp.X, cp.Y)
	canvas.LineTo(int32(math.Trunc(float64(cp.X)+math.Sin(mm)*(pl/2))),
		int32(math.Trunc(float64(cp.Y)-math.Cos(mm)*(pl/2))))

	// 秒针
	pen.SetColor(colors.ClGreen)
	pen.SetWidth(1)
	canvas.MoveTo(cp.X, cp.Y)
	canvas.LineTo(int32(math.Trunc(float64(cp.X)+math.Sin(ss)*(pl-20))),
		int32(math.Trunc(float64(cp.Y)-math.Cos(ss)*(pl-20))))
}
