// 在这里写你的事件

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"log"
	"math"
	"math/rand"
	"runtime"
	"unsafe"

	"github.com/ying32/govcl/vcl/types"

	"github.com/golang/freetype/truetype"
	"golang.org/x/image/font/gofont/goregular"

	"github.com/fogleman/gg"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	drawFunds map[string]func()
}

// 使用gg 2d图形库来绘制，最终显示到govcl控件上。
// gg部分的代码来自gg的例子
// https://github.com/fogleman/gg

func ifThen(val bool, atrue, afalse float64) float64 {
	if val {
		return atrue
	}
	return afalse
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.drawFunds = map[string]func(){
		"ggDrawLines":          f.ggDrawLines,
		"ggDrawEllipse":        f.ggDrawEllipse,
		"ggDrawGradientLinear": f.ggDrawGradientLinear,
		"ggDrawGradientRadial": f.ggDrawGradientRadial,
		"ggDrawBeziers":        f.ggDrawBeziers,
		"ggDrawCircle":         f.ggDrawCircle,
		"ggDrawCrisp":          f.ggDrawCrisp,
		"ggDrawCubic":          f.ggDrawCubic,
		"ggDrawFont":           f.ggDrawFont,
		"ggDrawGradientText":   f.ggDrawGradientText,
		"ggDrawInvertMask":     f.ggDrawInvertMask,
		"ggDrawLineWidth":      f.ggDrawLineWidth,
		"ggDrawLorem":          f.ggDrawLorem,
		"ggDrawRotatedText":    f.ggDrawRotatedText,
		"ggDrawRotatedImage":   f.ggDrawRotatedImage,
		"ggDrawWarpText":       f.ggDrawWarpText,
		"ggDrawStar":           f.ggDrawStar,
		"ggDrawStars":          f.ggDrawStars,
	}

}

// 32bit bmp，丢失透明度
func imageToBitmap(img image.Image) *vcl.TBitmap {
	srcData, ok := img.(*image.RGBA)
	if !ok {
		return nil
	}
	srcP := img.Bounds().Size()
	bmp := vcl.NewBitmap()
	bmp.SetPixelFormat(types.Pf32bit)
	bmp.SetSize(int32(srcP.X), int32(srcP.Y))

	if runtime.GOOS == "windows" || runtime.GOOS == "linux" {
		for h := srcP.Y - 1; h >= 0; h-- {
			ptr := bmp.ScanLine(int32(h))
			for w := 0; w < srcP.X*4; w++ {
				index := h*(srcP.X*4) + w
				*(*byte)(unsafe.Pointer(ptr + uintptr(w))) = srcData.Pix[index]
			}
		}
	} else {
		// 填充，左下角为起点
		for h := srcP.Y - 1; h >= 0; h-- {
			ptr := bmp.ScanLine(int32(h))
			for w := 0; w < int(srcP.X); w++ {
				index := (h*int(srcP.X) + w) * 4
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4))) = srcData.Pix[index+3]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+1))) = srcData.Pix[index+2]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+2))) = srcData.Pix[index+1]
				*(*byte)(unsafe.Pointer(ptr + uintptr(w*4+3))) = srcData.Pix[index]
			}
		}
	}
	return bmp
}

func ggImageToBitmap(dc *gg.Context) *vcl.TBitmap {
	return imageToBitmap(dc.Image())
}

func ggImageToPng(dc *gg.Context) *vcl.TPngImage {
	bs := bytes.NewBuffer([]byte{})
	if err := dc.EncodePNG(bs); err != nil {
		return nil
	}
	mem := vcl.NewMemoryStreamFromBytes(bs.Bytes())
	defer mem.Free()
	mem.SetPosition(0)
	pngObj := vcl.NewPngImage()
	pngObj.LoadFromStream(mem)
	return pngObj
}

func (f *TForm1) ggDrawImage(dc *gg.Context, isPng bool) {
	//if isPng {
	png := ggImageToPng(dc)
	if png != nil {
		defer png.Free()
		f.Canvas().Draw(0, 0, png)
	}
	//} else {
	//	bmp := ggImageToBitmap(dc)
	//	if bmp != nil {
	//		defer bmp.Free()
	//		f.Canvas().Draw(0, 0, bmp)
	//	}
	//}
}

func getFontFullPathName(name string) string {
	switch runtime.GOOS {
	case "windows":
		return "C:\\Windows\\Fonts\\" + name
	case "linux":
		return " /user/share/fonts/" + name
	case "darwin":
		return "/Library/Fonts/" + name
	}
	return name
}

// Lines
func (f *TForm1) ggDrawLines() {

	const W, H = 560, 560

	dc := gg.NewContext(W, H)
	dc.SetRGB(0, 0, 0)
	dc.Clear()
	for i := 0; i < 1000; i++ {
		x1 := rand.Float64() * W
		y1 := rand.Float64() * H
		x2 := rand.Float64() * W
		y2 := rand.Float64() * H
		r := rand.Float64()
		g := rand.Float64()
		b := rand.Float64()
		a := rand.Float64()*0.5 + 0.5
		w := rand.Float64()*4 + 1
		dc.SetRGBA(r, g, b, a)
		dc.SetLineWidth(w)
		dc.DrawLine(x1, y1, x2, y2)
		dc.Stroke()
	}

	f.ggDrawImage(dc, false)
}

// Gradient Linear
func (f *TForm1) ggDrawGradientLinear() {

	dc := gg.NewContext(400, 400)

	grad := gg.NewLinearGradient(20, 320, 400, 20)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	grad.AddColorStop(0.5, color.RGBA{255, 0, 0, 255})

	dc.SetColor(color.White)
	dc.DrawRectangle(20, 20, 400-20, 300)
	dc.Stroke()

	dc.SetStrokeStyle(grad)
	dc.SetLineWidth(4)
	dc.MoveTo(10, 10)
	dc.LineTo(410, 10)
	dc.LineTo(410, 100)
	dc.LineTo(10, 100)
	dc.ClosePath()
	dc.Stroke()

	dc.SetFillStyle(grad)
	dc.MoveTo(10, 120)
	dc.LineTo(410, 120)
	dc.LineTo(410, 300)
	dc.LineTo(10, 300)
	dc.ClosePath()
	dc.Fill()

	f.ggDrawImage(dc, true)
}

// Gradient Radial
func (f *TForm1) ggDrawGradientRadial() {

	dc := gg.NewContext(400, 200)

	grad := gg.NewRadialGradient(100, 100, 10, 100, 120, 80)
	grad.AddColorStop(0, color.RGBA{0, 255, 0, 255})
	grad.AddColorStop(1, color.RGBA{0, 0, 255, 255})

	dc.SetFillStyle(grad)
	dc.DrawRectangle(0, 0, 200, 200)
	dc.Fill()

	dc.SetColor(color.White)
	dc.DrawCircle(100, 100, 10)
	dc.Stroke()
	dc.DrawCircle(100, 120, 80)
	dc.Stroke()

	f.ggDrawImage(dc, true)
}

//beziers

func random() float64 {
	return rand.Float64()*2 - 1
}

func point() (x, y float64) {
	return random(), random()
}

func drawCurve(dc *gg.Context) {
	dc.SetRGBA(0, 0, 0, 0.1)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetLineWidth(12)
	dc.Stroke()
}

func drawPoints(dc *gg.Context) {
	dc.SetRGBA(1, 0, 0, 0.5)
	dc.SetLineWidth(2)
	dc.Stroke()
}

func randomQuadratic(dc *gg.Context) {
	x0, y0 := point()
	x1, y1 := point()
	x2, y2 := point()
	dc.MoveTo(x0, y0)
	dc.QuadraticTo(x1, y1, x2, y2)
	drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	drawPoints(dc)
}

func randomCubic(dc *gg.Context) {
	x0, y0 := point()
	x1, y1 := point()
	x2, y2 := point()
	x3, y3 := point()
	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	drawCurve(dc)
	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	drawPoints(dc)
}

func (f *TForm1) ggDrawBeziers() {

	const (
		S = 70 //256
		W = 8
		H = 8
	)
	dc := gg.NewContext(S*W, S*H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	for j := 0; j < H; j++ {
		for i := 0; i < W; i++ {
			x := float64(i)*S + S/2
			y := float64(j)*S + S/2
			dc.Push()
			dc.Translate(x, y)
			dc.Scale(S/2, S/2)
			if j%2 == 0 {
				randomCubic(dc)
			} else {
				randomQuadratic(dc)
			}
			dc.Pop()
		}
	}

	f.ggDrawImage(dc, false)
}

// Circle
func (f *TForm1) ggDrawCircle() {

	dc := gg.NewContext(400, 400)
	dc.DrawCircle(201, 201, 200)
	dc.SetRGB(1, 0, 0)
	dc.Fill()

	f.ggDrawImage(dc, true)
}

// Crisp
func (f *TForm1) ggDrawCrisp() {

	const W = 750
	const H = 500
	const Minor = 10
	const Major = 100

	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// minor grid
	for x := Minor; x < W; x += Minor {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Minor; y < H; y += Minor {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.25)
	dc.Stroke()

	// major grid
	for x := Major; x < W; x += Major {
		fx := float64(x) + 0.5
		dc.DrawLine(fx, 0, fx, H)
	}
	for y := Major; y < H; y += Major {
		fy := float64(y) + 0.5
		dc.DrawLine(0, fy, W, fy)
	}
	dc.SetLineWidth(1)
	dc.SetRGBA(0, 0, 0, 0.5)
	dc.Stroke()

	f.ggDrawImage(dc, false)
}

// ggDrawCubic
func (f *TForm1) ggDrawCubic() {
	const S = 500
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.Translate(S/2, S/2)
	dc.Scale(40, 40)

	var x0, y0, x1, y1, x2, y2, x3, y3 float64
	x0, y0 = -10, 0
	x1, y1 = -8, -8
	x2, y2 = 8, 8
	x3, y3 = 10, 0

	dc.MoveTo(x0, y0)
	dc.CubicTo(x1, y1, x2, y2, x3, y3)
	dc.SetRGBA(0, 0, 0, 0.2)
	dc.SetLineWidth(8)
	dc.FillPreserve()
	dc.SetRGB(0, 0, 0)
	dc.SetDash(16, 24)
	dc.Stroke()

	dc.MoveTo(x0, y0)
	dc.LineTo(x1, y1)
	dc.LineTo(x2, y2)
	dc.LineTo(x3, y3)
	dc.SetRGBA(1, 0, 0, 0.4)
	dc.SetLineWidth(2)
	dc.SetDash(4, 8, 1, 8)
	dc.Stroke()

	f.ggDrawImage(dc, false)
}

// ggDrawEllipse
func (f *TForm1) ggDrawEllipse() {
	const S = 500
	dc := gg.NewContext(S, S)
	dc.SetRGBA(0, 0, 0, 0.1)
	for i := 0; i < 360; i += 15 {
		dc.Push()
		dc.RotateAbout(gg.Radians(float64(i)), S/2, S/2)
		dc.DrawEllipse(S/2, S/2, S*7/16, S/8)
		dc.Fill()
		dc.Pop()
	}
	if im, err := gg.LoadImage("gopher.png"); err == nil {
		dc.DrawImageAnchored(im, S/2, S/2, 0.5, 0.5)
	}
	f.ggDrawImage(dc, true)
}

// ggDrawEllipse
func (f *TForm1) ggDrawFont() {
	// 这个字体不支持中文
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		log.Println(err)
		return
	}

	face := truetype.NewFace(font, &truetype.Options{Size: 48})

	dc := gg.NewContext(750, 500)
	dc.SetFontFace(face)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	dc.DrawStringAnchored("Hello, World!", 200, 200, 0.5, 0.5)

	f.ggDrawImage(dc, false)
}

// Gradient text
func (f *TForm1) ggDrawGradientText() {
	const (
		W = 750
		H = 500
	)
	dc := gg.NewContext(W, H)

	// draw text
	dc.SetRGB(0, 0, 0)

	dc.LoadFontFace(getFontFullPathName("impact.ttf"), 128)
	dc.DrawStringAnchored("Gradient Text", W/2, H/2, 0.5, 0.5)

	// get the context as an alpha mask
	mask := dc.AsMask()

	// clear the context
	dc.SetRGB(1, 1, 1)
	dc.Clear()

	// set a gradient
	g := gg.NewLinearGradient(0, 0, W, H)
	g.AddColorStop(0, color.RGBA{255, 0, 0, 255})
	g.AddColorStop(1, color.RGBA{0, 0, 255, 255})
	dc.SetFillStyle(g)

	// using the mask, fill the context with the gradient
	dc.SetMask(mask)
	dc.DrawRectangle(0, 0, W, H)
	dc.Fill()

	f.ggDrawImage(dc, false)
}

// ggDrawInvertMask
func (f *TForm1) ggDrawInvertMask() {
	const (
		w = 750
		h = 500
	)
	dc := gg.NewContext(w, h)
	dc.DrawCircle(200, 200, 200)
	dc.Clip()
	dc.InvertMask()
	dc.DrawRectangle(0, 0, w, h)
	dc.SetRGB(0, 0, 0)
	dc.Fill()

	f.ggDrawImage(dc, true)
}

// ggDrawLineWidth
func (f *TForm1) ggDrawLineWidth() {

	dc := gg.NewContext(750, 500)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	w := 0.1
	for i := 100; i <= 900; i += 20 {
		x := float64(i)
		dc.DrawLine(x+50, 0, x-50, 400)
		dc.SetLineWidth(w)
		dc.Stroke()
		w += 0.1
	}

	f.ggDrawImage(dc, false)
}

// ggDrawLorem
func (f *TForm1) ggDrawLorem() {
	var lines = []string{
		"Lorem ipsum dolor sit amet, consectetur adipiscing elit, sed do eiusmod",
		"tempor incididunt ut labore et dolore magna aliqua. Ut enim ad minim veniam,",
		"quis nostrud exercitation ullamco laboris nisi ut aliquip ex ea commodo",
		"consequat. Duis aute irure dolor in reprehenderit in voluptate velit esse",
		"cillum dolore eu fugiat nulla pariatur. Excepteur sint occaecat cupidatat",
		"non proident, sunt in culpa qui officia deserunt mollit anim id est laborum.",
	}
	const W = 750
	const H = 500
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	// dc.LoadFontFace("/Library/Fonts/Arial.ttf", 18)
	const h = 24
	for i, line := range lines {
		y := H/2 - h*len(lines)/2 + i*h
		dc.DrawStringAnchored(line, 400, float64(y), 0.5, 0.5)
	}
	f.ggDrawImage(dc, false)
}

// ggDrawRotatedText
func (f *TForm1) ggDrawRotatedText() {
	const S = 400
	dc := gg.NewContext(S, S)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.SetRGB(0, 0, 0)
	font, err := truetype.Parse(goregular.TTF)
	if err != nil {
		panic("")
	}
	face := truetype.NewFace(font, &truetype.Options{
		Size: 40,
	})
	dc.SetFontFace(face)
	text := "Hello, world!"
	w, h := dc.MeasureString(text)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 180, w, h)
	dc.Stroke()
	dc.DrawStringAnchored(text, 100, 180, 0.0, 0.0)

	f.ggDrawImage(dc, false)
}

// ggDrawRotatedImage
func (f *TForm1) ggDrawRotatedImage() {
	const W = 400
	const H = 500
	im, err := gg.LoadPNG("gopher.png")
	if err != nil {
		panic(err)
	}
	iw, ih := im.Bounds().Dx(), im.Bounds().Dy()
	dc := gg.NewContext(W, H)
	// draw outline
	dc.SetHexColor("#ff0000")
	dc.SetLineWidth(1)
	dc.DrawRectangle(0, 0, float64(W), float64(H))
	dc.Stroke()
	// draw full image
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.DrawRectangle(100, 210, float64(iw), float64(ih))
	dc.Stroke()
	dc.DrawImage(im, 100, 210)
	// draw image with current matrix applied
	dc.SetHexColor("#0000ff")
	dc.SetLineWidth(2)
	dc.Rotate(gg.Radians(10))
	dc.DrawRectangle(100, 0, float64(iw), float64(ih)/2+20.0)
	dc.StrokePreserve()
	dc.Clip()
	dc.DrawImageAnchored(im, 100, 0, 0.0, 0.0)

	f.ggDrawImage(dc, true)
}

// ggDrawWarpText
func (f *TForm1) ggDrawWarpText() {
	const TEXT = "Call me Ishmael. Some years ago—never mind how long precisely—having little or no money in my purse, and nothing particular to interest me on shore, I thought I would sail about a little and see the watery part of the world. It is a way I have of driving off the spleen and regulating the circulation. Whenever I find myself growing grim about the mouth; whenever it is a damp, drizzly November in my soul; whenever I find myself involuntarily pausing before coffin warehouses, and bringing up the rear of every funeral I meet; and especially whenever my hypos get such an upper hand of me, that it requires a strong moral principle to prevent me from deliberately stepping into the street, and methodically knocking people's hats off—then, I account it high time to get to sea as soon as I can. This is my substitute for pistol and ball. With a philosophical flourish Cato throws himself upon his sword; I quietly take to the ship. There is nothing surprising in this. If they but knew it, almost all men in their degree, some time or other, cherish very nearly the same feelings towards the ocean with me."

	const W = 700
	const H = 500
	const P = 16
	dc := gg.NewContext(W, H)
	dc.SetRGB(1, 1, 1)
	dc.Clear()
	dc.DrawLine(W/2, 0, W/2, H)
	dc.DrawLine(0, H/2, W, H/2)
	dc.DrawRectangle(P, P, W-P-P, H-P-P)
	dc.SetRGBA(0, 0, 1, 0.25)
	dc.SetLineWidth(3)
	dc.Stroke()
	dc.SetRGB(0, 0, 0)
	// impact.ttf Arial Bold.ttf
	// 没粗体的，就两种吧
	if err := dc.LoadFontFace(getFontFullPathName("impact.ttf"), 22); err != nil {
		log.Println(err)
		return
	}
	dc.DrawStringWrapped("UPPER LEFT", P, P, 0, 0, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("UPPER RIGHT", W-P, P, 1, 0, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("BOTTOM LEFT", P, H-P, 0, 1, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("BOTTOM RIGHT", W-P, H-P, 1, 1, 0, 1.5, gg.AlignRight)
	dc.DrawStringWrapped("UPPER MIDDLE", W/2, P, 0.5, 0, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LOWER MIDDLE", W/2, H-P, 0.5, 1, 0, 1.5, gg.AlignCenter)
	dc.DrawStringWrapped("LEFT MIDDLE", P, H/2, 0, 0.5, 0, 1.5, gg.AlignLeft)
	dc.DrawStringWrapped("RIGHT MIDDLE", W-P, H/2, 1, 0.5, 0, 1.5, gg.AlignRight)
	if err := dc.LoadFontFace(getFontFullPathName("Arial.ttf"), 12); err != nil {
		log.Println(err)
		return
	}
	dc.DrawStringWrapped(TEXT, W/2-P, H/2-P, 1, 1, W/3, 1.75, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2-P, 0, 1, W/3, 2, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2-P, H/2+P, 1, 0, W/3, 2.25, gg.AlignLeft)
	dc.DrawStringWrapped(TEXT, W/2+P, H/2+P, 0, 0, W/3, 2.5, gg.AlignLeft)

	f.ggDrawImage(dc, false)
}

// ggDrawStar

type Point struct {
	X, Y float64
}

func (f *TForm1) ggDrawStar() {

	Polygon := func(n int, x, y, r float64) []Point {
		result := make([]Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = Point{x + r*math.Cos(a), y + r*math.Sin(a)}
		}
		return result
	}

	n := 5
	points := Polygon(n, 300, 300, 200)
	dc := gg.NewContext(600, 600)
	dc.SetHexColor("fff")
	dc.Clear()
	for i := 0; i < n+1; i++ {
		index := (i * 2) % n
		p := points[index]
		dc.LineTo(p.X, p.Y)
	}
	dc.SetRGBA(0, 0.5, 0, 1)
	dc.SetFillRule(gg.FillRuleEvenOdd)
	dc.FillPreserve()
	dc.SetRGBA(0, 1, 0, 0.5)
	dc.SetLineWidth(16)
	dc.Stroke()

	f.ggDrawImage(dc, false)
}

// ggDrawStars
func (f *TForm1) ggDrawStars() {

	Polygon := func(n int) []Point {
		result := make([]Point, n)
		for i := 0; i < n; i++ {
			a := float64(i)*2*math.Pi/float64(n) - math.Pi/2
			result[i] = Point{math.Cos(a), math.Sin(a)}
		}
		return result
	}

	const W = 600
	const H = 120
	const S = 100
	dc := gg.NewContext(W, H)
	dc.SetHexColor("#FFFFFF")
	dc.Clear()
	n := 5
	points := Polygon(n)
	for x := S / 2; x < W; x += S {
		dc.Push()
		s := rand.Float64()*S/4 + S/4
		dc.Translate(float64(x), H/2)
		dc.Rotate(rand.Float64() * 2 * math.Pi)
		dc.Scale(s, s)
		for i := 0; i < n+1; i++ {
			index := (i * 2) % n
			p := points[index]
			dc.LineTo(p.X, p.Y)
		}
		dc.SetLineWidth(10)
		dc.SetHexColor("#FFCC00")
		dc.StrokePreserve()
		dc.SetHexColor("#FFE43A")
		dc.Fill()
		dc.Pop()
	}

	f.ggDrawImage(dc, false)
}

func (f *TForm1) OnFormPaint(sender vcl.IObject) {
	if f.ListBox1.ItemIndex() == -1 {
		return
	}
	fn, ok := f.drawFunds[f.ListBox1.Items().Strings(f.ListBox1.ItemIndex())]
	if !ok {
		fmt.Println("没有找到指定的滤镜。")
		return
	}
	fn()
}

func (f *TForm1) OnListBox1Click(sender vcl.IObject) {
	f.Repaint()
}
