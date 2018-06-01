package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	pm := vcl.NewPopupMenu(mainForm)
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	pm.Items().Add(item)

	trayicon := vcl.NewTrayIcon(mainForm)
	//trayicon.SetIcon(icon) 不设置则使用Application.Icon现有的
	// lcl库得手指定
	if rtl.LcLLoaded() {
		trayicon.SetIcon(vcl.Application.Icon())
	}
	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)
	trayicon.SetOnDblClick(func(vcl.IObject) {
		trayicon.SetBalloonTitle("test")
		trayicon.SetBalloonTimeout(10000)
		trayicon.SetBalloonHint("我是提示正文啦")
		trayicon.ShowBalloonHint()
		fmt.Println("TrayIcon Click.")
	})
	trayicon.SetPopupMenu(pm)
	// 其它事件请看源代码中以 SetOn 开头的
	//	trayicon.SetOnMouseDown(func(Sender vcl.IObject, Button, Shift, X, Y int32) {
	//		if Button == types.MbRight {
	//			vcl.ShowMessage("fff")
	//		}
	//	})

	vcl.Application.Run()
}
