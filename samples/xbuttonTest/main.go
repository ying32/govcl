package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/govcl/vcl/types/colors"

	_ "github.com/ying32/govcl/pkgs/winappres"
)

type TForm1 struct {
	*vcl.TForm
	Button1 *vcl.TXButton
}

var form1 *TForm1

func main() {
	vcl.RunApp(&form1)
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {

	f.SetWidth(600)
	f.SetHeight(400)
	f.ScreenCenter()

	f.Button1 = vcl.NewXButton(f)
	f.Button1.SetParent(f)
	f.Button1.SetDrawMode(types.DimCenter)

	f.Button1.Picture().LoadFromFile("favicon.png")

	f.Button1.SetBackColor(colors.ClAzure)
	f.Button1.SetNormalFontColor(colors.ClBlue)

	f.Button1.SetHoverColor(colors.ClLinen)
	f.Button1.SetHoverFontColor(colors.ClGreen)

	f.Button1.SetDownColor(colors.ClSilver)
	f.Button1.SetDownFontColor(colors.ClFuchsia)

	f.Button1.SetBorderWidth(1)
	f.Button1.SetBorderColor(colors.ClBrown)

	f.Button1.SetCaption("文字")
	//f.Button1.SetShowCaption(false)

	f.Button1.SetBounds(10, 10, 80, 40)
}

func (f *TForm1) OnButton1Click(object vcl.IObject) {
	vcl.ShowMessage("Click")
}
