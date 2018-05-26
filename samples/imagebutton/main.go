package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/exts/tools"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	// mac下记得发布时去掉
	tools.RunWithMacOSApp()

	path := rtl.ExtractFilePath(vcl.Application.ExeName())

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetDoubleBuffered(true) // 最好开启，以免闪烁
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(400)

	mainForm.SetShowHint(true)

	btnClose := vcl.NewImageButton(mainForm)
	btnClose.SetParent(mainForm)

	btnClose.SetImageCount(4)
	btnClose.SetAutoSize(true)
	btnClose.Picture().LoadFromFile(path + "btn_close.png")
	btnClose.SetLeft(mainForm.ClientWidth() - btnClose.Width() - 3)
	btnClose.SetHint("关闭")

	btnClose.SetOnClick(func(vcl.IObject) {
		vcl.ShowMessage("close")
	})

	btnMax := vcl.NewImageButton(mainForm)
	btnMax.SetParent(mainForm)
	btnMax.SetImageCount(4)
	btnMax.SetAutoSize(true)
	btnMax.Picture().LoadFromFile(path + "btn_max.png")
	btnMax.SetLeft(btnClose.Left() - btnMax.Width())
	btnMax.SetHint("最大化")

	btnMin := vcl.NewImageButton(mainForm)
	btnMin.SetParent(mainForm)
	btnMin.SetImageCount(4)
	btnMin.SetAutoSize(true)
	btnMin.Picture().LoadFromFile(path + "btn_min.png")
	btnMin.SetLeft(btnMax.Left() - btnMin.Width())
	btnMin.SetHint("最小化")

	btnSkin := vcl.NewImageButton(mainForm)
	btnSkin.SetParent(mainForm)
	btnSkin.SetImageCount(3)
	btnSkin.SetAutoSize(true)
	btnSkin.Picture().LoadFromFile(path + "btn_skin.png")
	btnSkin.SetLeft(btnMin.Left() - btnSkin.Width())
	btnSkin.SetHint("皮肤")

	btnScan := vcl.NewImageButton(mainForm)
	btnScan.SetParent(mainForm)
	btnScan.SetImageCount(3)
	btnScan.SetAutoSize(true)
	btnScan.Picture().LoadFromFile(path + "btn_scan.png")
	btnScan.SetTop((mainForm.ClientHeight() - btnScan.Height()) / 2)
	btnScan.SetLeft((mainForm.ClientWidth() - btnScan.Width()) / 2)
	btnScan.SetHint("全盘扫描")

	vcl.Application.Run()
}
