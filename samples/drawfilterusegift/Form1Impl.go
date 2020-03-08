// 在这里写你的事件

package main

import (
	"bytes"
	"fmt"
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/disintegration/gift"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	srcImage image.Image
}

// 使用gift来对图片做滤镜处理
//
// https://github.com/disintegration/gift

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()

}

// 24bit bmp，丢失透明度
// 这个在linux和macOS下还得修改，windows是正常的
//func imageToBitmap24(img image.Image) *vcl.TBitmap {
//
//	srcData, ok := img.(*image.RGBA) // RGBA
//	if !ok {
//
//		return nil
//	}
//	bmp := vcl.NewBitmap()
//	bmp.SetPixelFormat(types.Pf24bit)
//	srcP := img.Bounds().Size()
//	bmp.SetSize(int32(srcP.X), int32(srcP.Y))
//	//bmp.SetHandleType(types.BmDIB)
//	// 透明，看情况了
//	//bmp.SetTransparent(true)
//	//bmp.SetTransparentColor(colors.ClBlack)
//
//	setValue := func(ptr uintptr, x int, v uint8) {
//		*((*uint8)(unsafe.Pointer(ptr + uintptr(x)))) = v
//	}
//
//	for y := 0; y < srcP.Y; y++ {
//		DestPtr := bmp.ScanLine(int32(y))
//		pos := srcData.PixOffset(0, y)
//		for x := 0; x < srcP.X; x++ {
//			// Delphi BGR  Go: RGBA
//			setValue(DestPtr, x*3+0, srcData.Pix[pos+x*4+2])
//			setValue(DestPtr, x*3+1, srcData.Pix[pos+x*4+1])
//			setValue(DestPtr, x*3+2, srcData.Pix[pos+x*4+0])
//		}
//	}
//	return bmp
//}

func loadImage(filename string) image.Image {
	f, err := os.Open(filename)
	if err != nil {
		log.Fatalf("os.Open failed: %v", err)
	}
	defer f.Close()
	img, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("image.Decode failed: %v", err)
	}
	return img
}

var (
	filters = map[string]gift.Filter{
		"resize":              gift.Resize(100, 0, gift.LanczosResampling),
		"crop_to_size":        gift.CropToSize(100, 100, gift.LeftAnchor),
		"rotate_180":          gift.Rotate180(),
		"rotate_30":           gift.Rotate(30, color.Transparent, gift.CubicInterpolation),
		"brightness_increase": gift.Brightness(30),
		"brightness_decrease": gift.Brightness(-30),
		"contrast_increase":   gift.Contrast(30),
		"contrast_decrease":   gift.Contrast(-30),
		"saturation_increase": gift.Saturation(50),
		"saturation_decrease": gift.Saturation(-50),
		"gamma_1.5":           gift.Gamma(1.5),
		"gamma_0.5":           gift.Gamma(0.5),
		"gaussian_blur":       gift.GaussianBlur(1),
		"unsharp_mask":        gift.UnsharpMask(1, 1, 0),
		"sigmoid":             gift.Sigmoid(0.5, 7),
		"pixelate":            gift.Pixelate(5),
		"colorize":            gift.Colorize(240, 50, 100),
		"grayscale":           gift.Grayscale(),
		"sepia":               gift.Sepia(100),
		"invert":              gift.Invert(),
		"mean":                gift.Mean(5, true),
		"median":              gift.Median(5, true),
		"minimum":             gift.Minimum(5, true),
		"maximum":             gift.Maximum(5, true),
		"hue_rotate":          gift.Hue(45),
		"color_balance":       gift.ColorBalance(10, -10, -10),
		"color_func": gift.ColorFunc(
			func(r0, g0, b0, a0 float32) (r, g, b, a float32) {
				r = 1 - r0   // invert the red channel
				g = g0 + 0.1 // shift the green channel by 0.1
				b = 0        // set the blue channel to 0
				a = a0       // preserve the alpha channel
				return r, g, b, a
			},
		),
		"convolution_emboss": gift.Convolution(
			[]float32{
				-1, -1, 0,
				-1, 1, 1,
				0, 1, 1,
			},
			false, false, false, 0.0,
		),
	}
)

func (f *TForm1) OnFormPaint(sender vcl.IObject) {
	if f.ListBox1.ItemIndex() == -1 {
		return
	}

	if f.srcImage == nil {
		f.srcImage = loadImage("applogo.png")
	}
	if f.srcImage == nil {
		fmt.Println("没有加载applog.png。")
		return
	}

	fi, ok := filters[f.ListBox1.Items().Strings(f.ListBox1.ItemIndex())]
	if !ok {
		fmt.Println("没有找到指定的滤镜。")
		return
	}

	g := gift.New(fi)
	// 有损的
	//dst := image.NewNRGBA(g.Bounds(src.Bounds()))
	// 无损的
	dst := image.NewRGBA(g.Bounds(f.srcImage.Bounds()))
	g.Draw(dst, f.srcImage)

	//if f.ChkUsePng.Checked() {
	bs := bytes.NewBuffer([]byte{})
	if err := png.Encode(bs, dst); err != nil {
		return
	}
	mem := vcl.NewMemoryStreamFromBytes(bs.Bytes())
	defer mem.Free()
	mem.SetPosition(0)
	pngObj := vcl.NewPngImage()
	defer pngObj.Free()
	pngObj.LoadFromStream(mem)
	f.Canvas().Draw(0, 0, pngObj)
	//} else {
	//bmp := imageToBitmap24(dst)
	//if bmp != nil {
	//	defer bmp.Free()
	//	f.Canvas().Draw(0, 0, bmp)
	//}
	//}
}

func (f *TForm1) OnListBox1Click(sender vcl.IObject) {
	f.Repaint()
}
