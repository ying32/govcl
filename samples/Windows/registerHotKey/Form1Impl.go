// 在这里写你的事件

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/keys"
	"github.com/ying32/govcl/vcl/types/messages"
	"github.com/ying32/govcl/vcl/win"
)

//::private::
type TForm1Fields struct {
	hotKeyId types.ATOM
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.SetCaption("Press Ctrl+F1")
	f.ScreenCenter()
	f.hotKeyId = win.GlobalAddAtom("HotKeyId") - 0xC000
	//rtl.ShortCutToText()
	shift := types.NewSet(types.SsCtrl)
	// 注册Ctrl+F1
	// rtl.ShiftStateToWord(shift) 这个只是更容易理解，也可以使用 MOD_CONTROL | MOD_ALT 方法
	if !win.RegisterHotKey(f.Handle(), int32(f.hotKeyId), rtl.ShiftStateToWord(shift), keys.VkF1) {
		vcl.ShowMessage("注册热键失败。")
	}

}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {
	if f.hotKeyId > 0 {
		win.UnregisterHotKey(f.Handle(), int32(f.hotKeyId))
		win.GlobalDeleteAtom(f.hotKeyId)
	}
}

func (f *TForm1) OnFormWndProc(msg *types.TMessage) {

	f.InheritedWndProc(msg)
	/*
	  TWMHotKey = record
	    Msg: Cardinal;
	    MsgFiller: TDWordFiller;
	    HotKey: WPARAM;
	    Unused: LPARAM;
	    Result: LRESULT;
	  end;
	*/
	if msg.Msg == messages.WM_HOTKEY {
		if msg.WParam == types.WPARAM(f.hotKeyId) {
			vcl.ShowMessage("按下了Ctrl+F1")
		}
	}
}
