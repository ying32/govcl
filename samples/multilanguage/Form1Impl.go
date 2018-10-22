// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/multilang"
	"github.com/ying32/govcl/vcl/types"
)

var (
	testMessage = "这是一个测试消息！"
)

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	multilang.InitComponentLang(f)

	f.MainMenu1.SetAutoHotkeys(types.MaManual)
	item := vcl.NewMenuItem(f)
	item.SetCaption("Languages")
	for key, val := range multilang.LocalLangs {
		subitem := vcl.NewMenuItem(f)
		subitem.SetCaption(fmt.Sprintf("%d - %s", val.Language.Id, val.Language.Description))
		subitem.SetOnClick(f.OnLanguageMenuItemClick)
		subitem.SetTag(key)
		item.Add(subitem)
	}
	f.MainMenu1.Items().Add(item)
}

func (f *TForm1) OnLanguageMenuItemClick(sender vcl.IObject) {
	id := vcl.MenuItemFromObj(sender).Tag()
	if lang, ok := multilang.LocalLangs[id]; ok {
		fmt.Println(lang)
		multilang.ChangeLang(lang.Language.Name)
	}
}

func (f *TForm1) OnButton1Click(sender vcl.IObject) {
	vcl.ShowMessage(testMessage)
}

// 初始就注册
func init() {
	multilang.RegsiterVarString("testMessage", &testMessage)
	//multilang.RegsiterVar(&testMessage)
}
