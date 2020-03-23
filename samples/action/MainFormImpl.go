package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainForm struct {
	*vcl.TForm
	ImgList  *vcl.TImageList
	ActList  *vcl.TActionList
	Tlbar    *vcl.TToolBar
	Tlbtn    *vcl.TToolButton
	Stbar    *vcl.TStatusBar
	Btn      *vcl.TButton
	Chk      *vcl.TCheckBox
	Act      *vcl.TAction
	MainMenu *vcl.TMainMenu
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

	f.SetCaption("Hello")
	f.SetPosition(types.PoScreenCenter)
	f.EnabledMaximize(false)
	f.SetWidth(300)
	f.SetHeight(200)
	// 全局设置提示
	f.SetShowHint(true)

	// 动态创建
	f.initComponents()
}

func (f *TMainForm) OnActExecute(sender vcl.IObject) {
	vcl.ShowMessage("点击了action")
}

func (f *TMainForm) OnActUpdate(sender vcl.IObject) {
	vcl.AsAction(sender).SetEnabled(f.Chk.Checked())
}

func (f *TMainForm) initComponents() {
	f.ImgList = vcl.NewImageList(f)
	f.ImgList.AddIcon(vcl.Application.Icon())

	f.ActList = vcl.NewActionList(f)
	f.ActList.SetImages(f.ImgList)

	// 顶部工具条
	f.Tlbar = vcl.NewToolBar(f)
	f.Tlbar.SetParent(f)
	f.Tlbar.SetImages(f.ImgList)

	f.Tlbtn = vcl.NewToolButton(f)
	f.Tlbtn.SetParent(f.Tlbar)

	// 底部状态条
	f.Stbar = vcl.NewStatusBar(f)
	f.Stbar.SetParent(f)
	f.Stbar.SetAutoHint(true)
	f.Stbar.SetSimplePanel(true)

	f.Btn = vcl.NewButton(f)
	f.Btn.SetParent(f)
	f.Btn.SetLeft(80)
	f.Btn.SetTop(f.Tlbar.Top() + f.Tlbar.Height() + 10)
	f.Btn.SetImages(f.ImgList)

	f.Chk = vcl.NewCheckBox(f)
	f.Chk.SetParent(f)
	f.Chk.SetCaption("action状态演示")
	f.Chk.SetLeft(f.Btn.Left())
	f.Chk.SetTop(f.Btn.Top() + f.Btn.Height() + 10)
	f.Chk.SetChecked(true)

	// action
	f.Act = vcl.NewAction(f)
	f.Act.SetCaption("action")
	f.Act.SetImageIndex(0)
	f.Act.SetHint("这是一个提示|长提示了")

	// mainMenu
	f.MainMenu = vcl.NewMainMenu(f)
	f.MainMenu.SetImages(f.ImgList)

	menu := vcl.NewMenuItem(f)
	menu.SetCaption("菜单")
	f.MainMenu.Items().Add(menu)
	subMenu := vcl.NewMenuItem(f)
	subMenu.SetAction(f.Act)
	menu.Add(subMenu)

	f.Btn.SetAction(f.Act)
	f.Tlbtn.SetAction(f.Act)
}
