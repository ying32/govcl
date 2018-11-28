package main

import (
	"fmt"
	"runtime"

	"github.com/ying32/govcl/vcl"

	"github.com/ying32/govcl/vcl/exts/tools"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {
	tools.RunWithMacOSApp()
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)

	trayicon := vcl.NewTrayIcon(mainForm)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetCaption("ShowBalloon")
	btn.SetLeft(20)
	btn.SetTop(30)
	btn.SetWidth(100)
	btn.SetOnClick(func(vcl.IObject) {
		trayicon.SetBalloonTitle("test")
		trayicon.SetBalloonTimeout(1000)
		trayicon.SetBalloonHint("我是提示正文啦")
		trayicon.ShowBalloonHint()
	})

	pm := vcl.NewPopupMenu(mainForm)

	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	pm.Items().Add(item)
	trayicon.SetPopupMenu(pm)
	// lcl库得手指定
	if rtl.LcLLoaded() {
		if runtime.GOOS != "windows" {
			icon := vcl.NewIcon()
			icon.LoadFromFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "/2.ico")
			trayicon.SetIcon(icon)
			icon.Free()
		} else {
			trayicon.SetIcon(vcl.Application.Icon())
		}
	}
	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)

	if runtime.GOOS != "darwin" {
		trayicon.SetOnDblClick(func(vcl.IObject) {
			// macOS似乎不支持双击
			trayicon.SetBalloonTitle("test")
			trayicon.SetBalloonTimeout(10000)
			trayicon.SetBalloonHint("我是提示正文啦")
			trayicon.ShowBalloonHint()
			fmt.Println("TrayIcon DClick.")
		})
	}
	vcl.Application.Run()
}
