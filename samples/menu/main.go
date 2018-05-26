package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"

	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)

	// TMainMenu
	mainMenu := vcl.NewMainMenu(mainForm)
	// 不自动生成热键
	mainMenu.SetAutoHotkeys(types.MaManual)
	// 一级菜单
	item := vcl.NewMenuItem(mainForm)
	item.SetCaption("文件(&F)")

	subMenu := vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("新建(&N)")
	subMenu.SetShortCutFromString("Ctrl+N")
	subMenu.SetOnClick(func(vcl.IObject) {
		fmt.Println("单击了新建")
	})
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("打开(&O)")
	subMenu.SetShortCutFromString("Ctrl+O")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("保存(&S)")
	subMenu.SetShortCutFromString("Ctrl+S")
	item.Add(subMenu)

	// 分割线
	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("-")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("历史记录...")
	item.Add(subMenu)

	m := vcl.NewMenuItem(mainForm)
	m.SetCaption("第一个历史记录")
	subMenu.Add(m)

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("-")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("退出(&Q)")
	subMenu.SetShortCutFromString("Ctrl+Q")
	subMenu.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	item.Add(subMenu)

	mainMenu.Items().Add(item)

	item = vcl.NewMenuItem(mainForm)
	item.SetCaption("关于(&A)")

	subMenu = vcl.NewMenuItem(mainForm)
	subMenu.SetCaption("帮助(&H)")
	item.Add(subMenu)
	mainMenu.Items().Add(item)

	// TPopupMenu
	pm := vcl.NewPopupMenu(mainForm)
	item = vcl.NewMenuItem(mainForm)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(vcl.IObject) {
		mainForm.Close()
	})
	pm.Items().Add(item)

	// 将窗口设置一个弹出菜单，右键单击就可显示
	mainForm.SetPopupMenu(pm)

	vcl.Application.Run()
}
