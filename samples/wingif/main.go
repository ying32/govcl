package main

import (
	"runtime"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("GIF动画")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	img := vcl.NewImage(mainForm)
	img.SetParent(mainForm)
	img.SetBounds(20, 20, 60, 60)
	img.Picture().LoadFromFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "test.gif")
	// 动画只支持Windows下的libvcl
	if runtime.GOOS == "windows" && !rtl.LcLLoaded() {
		vcl.GIFImageFromObj(img.Picture().Graphic()).SetAnimate(true)

	}

	vcl.Application.Run()
}
