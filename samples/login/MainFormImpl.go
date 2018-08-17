// 在这里写你的事件

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

var (
	// 是否登录
	isLogin bool
)

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {

}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *TMainForm) OnFormCloseQuery(sender vcl.IObject, canClose *bool) {
	if isLogin {
		*canClose = vcl.MessageDlg("是否退出?", types.MtConfirmation, types.MbYes, types.MbNo) == types.MrYes
	}
}
