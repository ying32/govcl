// 在这里写你的事件

package main

import (
	"fmt"
	"image"
	"runtime"

	"github.com/ying32/govcl/vcl/bitmap"

	"github.com/disintegration/gift"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
	srcImage image.Image
	buffBmp  *vcl.TBitmap
}

// 使用gift来对图片做滤镜处理
// gift部分的代码来自gift的例子
// https://github.com/disintegration/gift

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	f.buffBmp = vcl.NewBitmap()
	// 先不使用吧
	f.ChkUsePng.SetChecked(false)
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	if f.buffBmp != nil && f.buffBmp.IsValid() {
		f.buffBmp.Free()
	}
}

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

	// linux跟macOS下有些效果不太好，所以用png，具体原因还待分析。
	if runtime.GOOS != "windows" {
		pngObj, err := bitmap.ToPngImage(dst)
		if err != nil {
			return
		}
		if pngObj != nil {
			defer pngObj.Free()
			f.Canvas().Draw(0, 0, pngObj)
		}
	} else {
		if bitmap.ToBitmap2(dst, f.buffBmp) == nil {
			f.Canvas().Draw(0, 0, f.buffBmp)
		}
	}

}

func (f *TForm1) OnListBox1Click(sender vcl.IObject) {
	f.Repaint()
}
