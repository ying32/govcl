package main

import (
	"runtime"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
)

type TMainForm struct {
	*vcl.TForm
	img *vcl.TImage
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(object vcl.IObject) {
	f.SetCaption("GIF动画")
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()

	f.img = vcl.NewImage(f)
	f.img.SetParent(f)
	f.img.SetBounds(20, 20, 60, 60)
	f.img.Picture().LoadFromFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "test.gif")
	// 动画只支持Windows下的libvcl
	if runtime.GOOS == "windows" && !rtl.LcLLoaded() {
		vcl.AsGIFImage(f.img.Picture().Graphic()).SetAnimate(true)
	}
}

func main() {
	vcl.RunApp(&mainForm)
}
