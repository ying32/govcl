package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"

	_ "github.com/ying32/govcl/pkgs/winappres"
)

type TMainForm struct {
	*vcl.TForm
	ButtonPanel *vcl.TButtonPanel
}

var (
	mainForm *TMainForm
)

func main() {
	vcl.DEBUG = true
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Hello")
	f.EnabledMaximize(false)
	f.SetWidth(600)
	f.SetHeight(400)
	f.ScreenCenter()

	f.ButtonPanel = vcl.NewButtonPanel(f)
	f.ButtonPanel.SetParent(f)
	// 默认按钮
	f.ButtonPanel.SetDefaultButton(types.PbCancel)

	// 隐藏横线
	//f.ButtonPanel.SetShowBevel(false)
	// 设置要显示的按钮，默认是全部显示，这里移除掉关闭按钮
	f.ButtonPanel.SetShowButtons(f.ButtonPanel.ShowButtons().Exclude(types.PbClose))

	// 设置按钮图标的显示，默认全部显示，这里移除掉ok按钮的
	f.ButtonPanel.SetShowGlyphs(f.ButtonPanel.ShowGlyphs().Exclude(types.PbOK))

	// 按钮顺序
	f.ButtonPanel.SetButtonOrder(types.BoCloseCancelOK)

	f.ButtonPanel.OKButton().SetOnClick(f.onOKButtonClick)
	//f.ButtonPanel.CancelButton().SetOnClick()
	//f.ButtonPanel.CloseButton().SetOnClick()
	//f.ButtonPanel.HelpButton().SetOnClick()
}

func (f *TMainForm) onOKButtonClick(sender vcl.IObject) {
	vcl.ShowMessage("单击ok")
}
