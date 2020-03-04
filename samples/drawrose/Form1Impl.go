// 在这里写你的事件

package main

import (
	"fmt"
	"math"
	"math/rand"
	"time"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

// https://www.oschina.net/code/snippet_212659_48907

const (
	RAND_MAX  = 0x7FFF
	urosesize = 500
	uh        = -250
)

type TDOT struct {
	x float64
	y float64
	z float64
	r float64 // 红色
	g float64 // 绿色
	// b(蓝色) 通过 r 计算
}

func ifThen(val bool, atrue, afalse float64) float64 {
	if val {
		return atrue
	}
	return afalse
}

func calc(a, b, c float64, d *TDOT) bool {
	var j, n, o, w, z, _A, _B float64

	if c > 60 { // 花柄
		d.x = math.Sin(a*7)*(13+5/(0.2+math.Pow(b*4, 4))) - math.Sin(b)*50
		d.y = b*urosesize + 50
		d.z = 625 + math.Cos(a*7)*(13+5/(0.2+math.Pow(b*4, 4))) + b*400
		d.r = a*1 - b/2
		d.g = a
		return true
	}
	_A = a*2 - 1
	_B = b*2 - 1

	if _A*_A+_B*_B < 1 {
		if c > 37 { // 叶
			j = float64(int(math.Trunc(c)) & 1)
			n = ifThen(j != 0, 6, 4)
			o = 0.5/(a+0.01) + math.Cos(b*125)*3 - a*300
			w = b * uh

			d.x = o*math.Cos(n) + w*math.Sin(n) + j*610 - 390
			d.y = o*math.Sin(n) - w*math.Cos(n) + 550 - j*350
			d.z = 1180 + math.Cos(_B+_A)*99 - j*300
			d.r = 0.4 - a*0.1 + math.Pow(1-_B*_B, -uh*6)*0.15 - a*b*0.4 + math.Cos(a+b)/5 + math.Pow(math.Cos((o*(a+1)+ifThen(_B > 0, w, -w))/25), 30)*0.1*(1-_B*_B)
			d.g = o/1000 + 0.7 - o*w*0.000003
			return true
		}

		if c > 32 { // 花萼

			c = c*1.16 - 0.15
			o = a*45 - 20
			w = b * b * uh
			z = o*math.Sin(c) + w*math.Cos(c) + 620

			d.x = o*math.Cos(c) - w*math.Sin(c)
			d.y = 28 + math.Cos(_B*0.5)*99 - b*b*b*60 - z/2 - uh
			d.z = z
			d.r = (b*b*0.3 + math.Pow(1-(_A*_A), 7)*0.15 + 0.3) * b
			d.g = b * 0.7
			return true
		}

		// 花
		o = _A * (2 - b) * (80 - c*2)
		w = 99 - math.Cos(_A)*120 - math.Cos(b)*(-uh-c*4.9) + math.Cos(math.Pow(1-b, 7))*50 + c*2
		z = o*math.Sin(c) + w*math.Cos(c) + 700

		d.x = o*math.Cos(c) - w*math.Sin(c)
		d.y = _B*99 - math.Cos(math.Pow(b, 7))*50 - c/3 - z/1.35 + 450
		d.z = z
		d.r = (1-b/1.2)*0.9 + a*0.1
		d.g = math.Pow((1-b), 20)/4 + 0.05
		return true
	}

	return false
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()

}

func (f *TForm1) drawRose() {
	var zBuffer []int
	var i, j, x, y, z, zBufferIndex int
	var dot TDOT
	var r, g, b int

	zBuffer = make([]int, urosesize*urosesize)
	canvas := f.Canvas()

	canvas.Brush().SetColor(colors.ClWhite)
	canvas.FillRect(types.TRect{0, 0, f.Width(), f.Height()})
	rand.Seed(time.Now().UnixNano())
	for j = 0; j <= 1999; j++ {
		for i = 0; i <= 9999; i++ {
			if calc(float64(rand.Intn(RAND_MAX))/RAND_MAX, float64(rand.Intn(RAND_MAX))/RAND_MAX, float64(rand.Intn(RAND_MAX)%46)/0.74, &dot) {
				z = int(math.Trunc(dot.z + 0.5))
				x = int(math.Trunc(dot.x*float64(urosesize)/float64(z) - uh + 0.5))
				y = int(math.Trunc(dot.y*float64(urosesize)/float64(z) - uh + 0.5))
				if y >= urosesize {
					continue
				}
				zBufferIndex = y*urosesize + x
				if (!(zBuffer[zBufferIndex] != 0)) || (zBuffer[zBufferIndex] > z) {
					zBuffer[zBufferIndex] = z

					// 画点
					r = ^int(math.Trunc(dot.r * uh))
					if r < 0 {
						r = 0
					}
					if r > 255 {
						r = 255
					}
					g = ^int(math.Trunc(dot.g * uh))
					if g < 0 {
						g = 0
					}
					if g > 255 {
						g = 255
					}
					b = ^int(math.Trunc(dot.r * dot.r * -80))
					if b < 0 {
						b = 0
					}
					if b > 255 {
						b = 255
					}
					var color types.TColor
					canvas.SetPixels(int32(x+50), int32(y-20), color.RGB(byte(r), byte(g), byte(b)))
				}
			}
			vcl.Application.ProcessMessages()
		}
		time.Sleep(time.Millisecond)
	}
	fmt.Println("done.")
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	f.drawRose()
}

func (f *TForm1) OnFormPaint(sender vcl.IObject) {

}
