// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/i18n"
	"github.com/ying32/govcl/vcl/types"
)

var (
	testMessage  = "这是一个测试消息！"
	testMessage2 = "你确定么？"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	i18n.InitComponentLang(f)
	f.ScreenCenter()

	item := vcl.NewMenuItem(f)
	item.SetCaption("Languages")
	for key, val := range i18n.LocalLangs {
		subitem := vcl.NewMenuItem(f)
		subitem.SetGroupIndex(1)
		subitem.SetRadioItem(true)
		subitem.SetCaption(fmt.Sprintf("%d - %s", val.Language.Id, val.Language.Description))
		subitem.SetOnClick(f.OnLanguageMenuItemClick)
		subitem.SetTag(key)
		if i18n.CurrentLang == val.Language.Name {
			subitem.SetChecked(true)
		}
		item.Add(subitem)
	}
	f.MainMenu1.Items().Add(item)
}

func (f *TForm1) OnLanguageMenuItemClick(sender vcl.IObject) {
	id := vcl.AsMenuItem(sender).Tag()
	if lang, ok := i18n.LocalLangs[id]; ok {
		fmt.Println(lang)
		i18n.ChangeLang(lang.Language.Name)
		i18n.WriteSetLang(lang.Language.Name)
		vcl.AsMenuItem(sender).SetChecked(true)
	}
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	vcl.ShowMessage(testMessage)
}

func (f *TForm1) OnButton2Click(sender vcl.IObject) {
	if vcl.MessageDlg(testMessage2, types.MtConfirmation, types.MbYes, types.MbNo) == types.IdYes {

	}
}

func (f *TForm1) OnButton3Click(sender vcl.IObject) {
	vcl.ShowMessage(i18n.IdRes("testMessage3"))
}

// 初始就注册
func init() {
	i18n.RegsiterVarString("testMessage", &testMessage)
	i18n.RegsiterVarString("testMessage2", &testMessage2)
	//multilang.RegsiterVar(&testMessage)
}
