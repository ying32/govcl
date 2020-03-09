package main

import (
	"fmt"
	"runtime"

	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

type TMainForm struct {
	*vcl.TForm
	mainMenu *vcl.TMainMenu
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("Menu example")
	f.ScreenCenter()

	// TMainMenu
	f.mainMenu = vcl.NewMainMenu(f)
	// 不自动生成热键
	if !rtl.LcLLoaded() {
		f.mainMenu.SetAutoHotkeys(types.MaManual)
	}

	// macOS下专有的
	if runtime.GOOS == "darwin" {
		fmt.Println("darwin")
		// https://wiki.lazarus.freepascal.org/Mac_Preferences_and_About_Menu
		// 动态添加的，静态好像是通过设计器将顶级的菜单标题设置为应用程序名，但动态的就是另一种方式
		appMenu := vcl.NewMenuItem(f)
		// 动态添加的，设置一个Unicode Apple logo char
		appMenu.SetCaption(string([]byte{0xEF, 0xA3, 0xBF}))
		subItem := vcl.NewMenuItem(f)
		// ----
		subItem.SetCaption("关于")
		subItem.SetOnClick(func(sender vcl.IObject) {
			vcl.ShowMessage("About")
		})
		appMenu.Add(subItem)
		// --
		subItem = vcl.NewMenuItem(f)
		subItem.SetCaption("-")
		appMenu.Add(subItem)

		// ---
		subItem = vcl.NewMenuItem(f)
		subItem.SetCaption("首选项...")
		subItem.SetShortCutFromString("Meta+,")
		subItem.SetOnClick(func(sender vcl.IObject) {
			vcl.ShowMessage("Preferences")
		})
		appMenu.Add(subItem)
		// 添加
		f.mainMenu.Items().Insert(0, appMenu)
	}

	// 一级菜单
	item := vcl.NewMenuItem(f)
	item.SetCaption("文件(&F)")

	subMenu := vcl.NewMenuItem(f)
	subMenu.SetCaption("新建(&N)")
	subMenu.SetShortCutFromString("Ctrl+N")
	subMenu.SetOnClick(func(vcl.IObject) {
		fmt.Println("单击了新建")
	})
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("打开(&O)")
	subMenu.SetShortCutFromString("Ctrl+O")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("保存(&S)")
	subMenu.SetShortCutFromString("Ctrl+S")
	item.Add(subMenu)

	// 分割线
	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("-")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("历史记录...")
	item.Add(subMenu)

	m := vcl.NewMenuItem(f)
	m.SetCaption("第一个历史记录")
	subMenu.Add(m)

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("-")
	item.Add(subMenu)

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("退出(&Q)")
	subMenu.SetShortCutFromString("Ctrl+Q")
	subMenu.SetOnClick(func(vcl.IObject) {
		f.Close()
	})
	item.Add(subMenu)

	f.mainMenu.Items().Add(item)

	item = vcl.NewMenuItem(f)
	item.SetCaption("关于(&A)")

	subMenu = vcl.NewMenuItem(f)
	subMenu.SetCaption("帮助(&H)")
	item.Add(subMenu)
	f.mainMenu.Items().Add(item)

	// TPopupMenu
	pm := vcl.NewPopupMenu(f)
	item = vcl.NewMenuItem(f)
	item.SetCaption("退出(&E)")
	item.SetOnClick(func(vcl.IObject) {
		f.Close()
	})
	pm.Items().Add(item)

	// 将窗口设置一个弹出菜单，右键单击就可显示
	f.SetPopupMenu(pm)
}

func (f *TMainForm) OnFormPaint(sender vcl.IObject) {
	///r := types.TRect{0, 0, f.Width(), f.Height()}
	//f.Canvas().TextRect(r, 0, 0, "右键弹出菜单")
	f.Canvas().Brush().SetStyle(types.BsClear)
	f.Canvas().Font().SetColor(colors.ClGreen)
	f.Canvas().TextOut(10, 80, "右键弹出菜单")
}
