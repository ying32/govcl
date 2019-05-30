// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
)

func (f *TLoginForm) OnFormCreate(sender vcl.IObject) {
	//f.SetParentWindow(0)
	SetShowInTaskBar(f)
}

func (f *TLoginForm) OnFormClose(sender vcl.IObject, action *types.TCloseAction) {
	if !isLogin {
		vcl.Application.Terminate()
	}
}

func (f *TLoginForm) OnButtonLoginClick(sender vcl.IObject) {
	usr := f.EditUserName.Text()
	if usr == "" {
		fmt.Println("输入用户名吧")
		f.EditUserName.SetFocus()
		return
	}
	pwd := f.EditPassword.Text()
	if pwd == "" {
		fmt.Println("输入密码吧。")
		f.EditPassword.SetFocus()
		return
	}
	if usr == "admin" && pwd == "admin" {
		isLogin = true
		f.Close()
		//MainForm.Show()
		vcl.Application.MainForm().Show()
	}
}

func (f *TLoginForm) OnFormKeyPress(sender vcl.IObject, key *types.Char) {
	fmt.Println("key:", *key)
	if *key == keys.VkReturn {
		f.ButtonLogin.Click()
	}
}
