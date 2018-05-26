// +build windows

package main

import (
	"gitee.com/ying32/govcl/vcl"

	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.SetIconResId(3)
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
	vcl.GIFImageFromObj(img.Picture().Graphic()).SetAnimate(true)

	vcl.Application.Run()
}
