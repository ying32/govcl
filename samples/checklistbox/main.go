package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("TCheckListBox测试")
	mainForm.ScreenCenter()

	chkListBox := vcl.NewCheckListBox(mainForm)
	chkListBox.SetParent(mainForm)
	chkListBox.SetAlign(types.AlClient)
	chkListBox.SetOnClickCheck(func(sender vcl.IObject) {
		fmt.Println("check单击。")
	})

	for i := 1; i < 100; i++ {
		chkListBox.Items().Add(fmt.Sprintf("第%d个项目", i))
	}
	button := vcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetAlign(types.AlBottom)
	button.SetCaption("项目启用")
	button.SetOnClick(func(sender vcl.IObject) {
		//fmt.Println("选中数：", chkListBox.Checked())
		chkListBox.SetItemEnabled(0, !chkListBox.ItemEnabled(0))
	})

	// 获取/设置项目启用
	//chkListBox.SetItemEnabled()
	//chkListBox.ItemEnabled()

	// 获取/设置项目状态
	//chkListBox.State()
	//chkListBox.SetState()

	// 获取/设置项目选中
	//chkListBox.Checked()
	//chkListBox.SetChecked()

	button = vcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("全选")
	button.SetAlign(types.AlBottom)
	button.SetOnClick(func(sender vcl.IObject) {
		chkListBox.CheckAll(types.CbChecked, true, true)
	})

	button = vcl.NewButton(mainForm)
	button.SetParent(mainForm)
	button.SetCaption("取消全选")
	button.SetAlign(types.AlBottom)
	button.SetOnClick(func(sender vcl.IObject) {
		chkListBox.CheckAll(types.CbUnchecked, true, true)
	})

	vcl.Application.Run()
}
