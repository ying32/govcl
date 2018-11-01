package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TForm1 struct {
	*vcl.TForm
}

var form1 *TForm1

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&form1)
	vcl.Application.Run()
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	form1.SetCaption("Message Test")
	form1.SetPosition(types.PoScreenCenter)
	form1.EnabledMaximize(false)
	form1.SetWidth(500)
	form1.SetHeight(400)
	form1.SetOnWndProc(f.OnFormWindProc)
}

const (
	WM_MOUSEMOVE     = 0x0200
	WM_LBUTTONDOWN   = 0x0201
	WM_LBUTTONUP     = 0x0202
	WM_LBUTTONDBLCLK = 0x0203
	WM_RBUTTONDOWN   = 0x0204
	WM_RBUTTONUP     = 0x0205
	WM_RBUTTONDBLCLK = 0x0206
)

func (f *TForm1) OnFormWindProc(msg *types.TMessage, handled *bool) {
	switch msg.Msg {
	case WM_MOUSEMOVE:
	case WM_LBUTTONDOWN:
		fmt.Println("左键接下")
	case WM_LBUTTONUP:
		fmt.Println("左键抬起")
	case WM_LBUTTONDBLCLK:
		fmt.Println("左键双击")
	case WM_RBUTTONDOWN:
		fmt.Println("右键接下")
	case WM_RBUTTONUP:
		fmt.Println("右键抬起")
	case WM_RBUTTONDBLCLK:
		fmt.Println("右键双击")
	}
}
