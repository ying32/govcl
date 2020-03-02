package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetShowHint(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetPosition(types.PoScreenCenter)

	form2 := vcl.Application.CreateForm()
	form2.SetCaption("Form2")
	stb := vcl.NewStatusBar(form2)
	stb.SetParent(form2)
	stb.SetSimpleText("状态条")

	// 窗口全局show提示
	mainForm.SetShowHint(true)

	button := vcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("按钮1")

	// 长提示，后面使用  | 分割
	button.SetHint("这是一个提示|后面可以是长提示")
	button.SetBounds(10, 20, 100, 30)
	button.SetOnClick(func(sender vcl.IObject) {
		form2.Show()
	})

	button = vcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("按钮2")
	button.SetHint("提示。好好好好好。。。。。。。。。。。。。")
	button.SetBounds(10, 60, 100, 30)

	statusbar := vcl.NewStatusBar(mainForm)
	statusbar.SetParent(mainForm)
	statusbar.SetName("statusbar")

	// 右下角出现可调整窗口三角形，默认显示
	//statusbar.SetSizeGrip(false)

	// 当开启后，菜单的或者按钮的提示会出现在panel上
	statusbar.SetAutoHint(true)

	// 一般配合AutoHint使用
	statusbar.SetSimplePanel(true)

	// 原本是不需这下面这样的，在dll中创建的貌似有些问题，所以手动来操作了
	vcl.Application.SetOnHint(func(sender vcl.IObject) {
		if statusbar.IsValid() {
			if statusbar.SimplePanel() {
				statusbar.SetSimpleText(vcl.Application.Hint())
			} else {
				if statusbar.Panels().Count() > 0 {
					statusbar.Panels().Items(0).SetText(vcl.Application.Hint())
				}
			}
		}
	})

	pnl := statusbar.Panels().Add()
	pnl.SetText("pnl1")
	pnl.SetWidth(100)

	pnl = statusbar.Panels().Add()
	pnl.SetText("pnl1")
	pnl.SetAlignment(types.TaCenter)
	pnl.SetWidth(100)

	pnl = statusbar.Panels().Add()
	pnl.SetText("第二格")
	pnl.SetAlignment(types.TaRightJustify)
	pnl.SetWidth(100)

	var i int32
	for i = 0; i < statusbar.Panels().Count(); i++ {
		fmt.Println("text:", statusbar.Panels().Items(i).Text())
	}

	vcl.Application.Run()
}
