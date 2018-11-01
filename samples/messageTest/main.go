package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/message"
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

func (f *TForm1) OnFormWindProc(msg *types.TMessage, handled *bool) {
	switch msg.Msg {
	case message.WM_MOUSEMOVE:

	case message.WM_LBUTTONDOWN:
		fmt.Println("左键接下")

	case message.WM_LBUTTONUP:
		fmt.Println("左键抬起")

	case message.WM_LBUTTONDBLCLK:
		fmt.Println("左键双击")

	case message.WM_RBUTTONDOWN:
		fmt.Println("右键接下")

	case message.WM_RBUTTONUP:
		fmt.Println("右键抬起")

	case message.WM_RBUTTONDBLCLK:
		fmt.Println("右键双击")

	}
}
