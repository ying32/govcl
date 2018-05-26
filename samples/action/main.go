package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	icon := vcl.NewIcon()
	defer icon.Free()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetIcon(icon)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)
	// 全局设置提示
	mainForm.SetShowHint(true)

	imgList := vcl.NewImageList(mainForm)
	imgList.AddIcon(icon)

	actList := vcl.NewActionList(mainForm)
	actList.SetImages(imgList)

	// 顶部工具条
	tlbar := vcl.NewToolBar(mainForm)
	tlbar.SetParent(mainForm)
	tlbar.SetImages(imgList)

	tlbtn := vcl.NewToolButton(mainForm)
	tlbtn.SetParent(tlbar)

	// 底部状态条
	stbar := vcl.NewStatusBar(mainForm)
	stbar.SetParent(mainForm)
	stbar.SetAutoHint(true)
	stbar.SetSimplePanel(true)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetLeft(80)
	btn.SetTop(tlbar.Top() + tlbar.Height() + 10)
	btn.SetImages(imgList)

	chk := vcl.NewCheckBox(mainForm)
	chk.SetParent(mainForm)
	chk.SetCaption("action状态演示")
	chk.SetLeft(btn.Left())
	chk.SetTop(btn.Top() + btn.Height() + 10)
	chk.SetChecked(true)

	// action
	act := vcl.NewAction(mainForm)
	act.SetCaption("action")
	act.SetImageIndex(0)
	act.SetHint("这是一个提示|长提示了")

	act.SetOnExecute(func(vcl.IObject) {
		vcl.ShowMessage("点击了action")
	})
	act.SetOnUpdate(func(sender vcl.IObject) {
		vcl.ActionFromObj(sender).SetEnabled(chk.Checked())
	})

	mainMenu := vcl.NewMainMenu(mainForm)
	mainMenu.SetImages(imgList)

	menu := vcl.NewMenuItem(mainForm)
	menu.SetCaption("菜单")
	mainMenu.Items().Add(menu)
	subMenu := vcl.NewMenuItem(mainForm)
	subMenu.SetAction(act)
	menu.Add(subMenu)

	btn.SetAction(act)
	tlbtn.SetAction(act)

	vcl.Application.Run()
}
