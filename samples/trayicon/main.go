package main

import (
	"fmt"
	"runtime"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	//"github.com/ying32/govcl/vcl/rtl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)

	mMenu := vcl.NewMainMenu(mainForm)
	mmItem := vcl.NewMenuItem(mainForm)
	mmItem.SetCaption("File")
	mMenu.Items().Add(mmItem)

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
	item.SetCaption("显示(&S)")
	item.SetOnClick(func(vcl.IObject) {

		mainForm.Show()
		// Windows上为了最前面显示，有时候要调用SetForegroundWindow
		//rtl.SetForegroundWindow(mainForm.Handle())
		//vcl.Application.Restore()
		//vcl.Application.BringToFront()
	})
	pm.Items().Add(item)

	item = vcl.NewMenuItem(mainForm)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(vcl.IObject) {
		// 主窗口关闭
		mainForm.Close()
		// 或者使用
		//		vcl.Application.Terminate()
	})
	pm.Items().Add(item)
	trayicon.SetPopupMenu(pm)
	// lcl库得手指定，在windows下，如果实例资源中存在一个名为 MAINICON 的图标资源，则会自动加载，下面只是应对于linux与macOS下

	if runtime.GOOS != "windows" {
		// 这是使用文件加载方法，不考虑外部文件的话，可以用新的内存方法加载
		//icon := vcl.NewIcon()
		//icon.LoadFromFile(rtl.ExtractFilePath(vcl.Application.ExeName()) + "/2.ico")
		//trayicon.SetIcon(icon)
		//icon.Free()
		// 将图标应用到Application中的Icon中，到时候随时可以使用
		// 但也可不使用
		//loadMainIconFromStream(vcl.Application.Icon())
		loadMainIconFromStream(trayicon.Icon())

	} else {
		//trayicon.SetIcon(vcl.Application.Icon())
	}

	trayicon.SetHint(mainForm.Caption())
	trayicon.SetVisible(true)

	// 捕捉最小化
	vcl.Application.SetOnMinimize(func(sender vcl.IObject) {
		mainForm.Hide() // 主窗口最隐藏掉
	})

	// 这里写啥好呢，macOS下似乎这些事件跟PopupMenu有冲突
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

	// 托盘图片可闪烁 1 秒闪一次
	//tmr1 := vcl.NewTimer(mainForm)
	//tmr1.SetOnTimer(func(sender vcl.IObject) {
	//	if trayicon.Icon().Empty() {
	//		trayicon.SetIcon(vcl.Application.Icon())
	//	} else {
	//		trayicon.SetIcon(nil)
	//	}
	//})
	//tmr1.SetEnabled(true)

	// 加载其他格式的ico方式一
	//png := vcl.NewPngImage()
	//png.LoadFromFile("bow.png")
	//trayicon.Icon().Assign(png)
	//png.Free()
	// 方式二可以通过imagelist操作
	//imglist := vcl.NewImageList(mainForm)
	//png := vcl.NewPngImage()
	//png.LoadFromFile("bow.png")
	//imglist.Add(png, nil)
	//png.Free()
	//imglist.GetIcon(0, trayicon.Icon())

	vcl.Application.Run()
}

// 主要是用于linux跟macOS下，因为不能像Windows一样直接内置到资源中
func loadMainIconFromStream(outIcon *vcl.TIcon) {
	if outIcon.IsValid() {
		//mem := vcl.NewMemoryStreamFromBytes(mainIconBytes)
		//defer mem.Free() // 不要在阻塞的时候使用defer不然会一直到阻塞结束才释放，这里使用是因为这个函数结束了就释放了
		//mem.SetPosition(0)
		//outIcon.LoadFromStream(mem)
		outIcon.LoadFromBytes(mainIconBytes)
	}
}
