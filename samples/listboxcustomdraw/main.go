package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(400)

	var itemHeight int32 = 30
	listbox := vcl.NewListBox(mainForm)
	listbox.SetParent(mainForm)
	listbox.SetStyle(types.LbOwnerDrawFixed)
	listbox.SetAlign(types.AlClient)
	listbox.Items().Add("第一项")
	listbox.Items().Add("第二项")
	listbox.Items().Add("第三项")
	listbox.Items().Add("第四项")
	listbox.Items().Add("第五项")
	listbox.Items().Add("第六项")
	listbox.Items().Add("第七项")
	listbox.Items().Add("第八项")
	listbox.SetItemHeight(itemHeight)
	listbox.SetOnDrawItem(func(control vcl.IControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
		canvas := listbox.Canvas()
		s := listbox.Items().Strings(index)
		fw := canvas.TextWidth(s)
		fh := canvas.TextHeight(s)
		canvas.Font().SetColor(types.ClBlack)
		canvas.Brush().SetColor(types.ClBtnFace)
		canvas.FillRect(aRect)
		canvas.Brush().SetColor(0x00FFF7F7)
		canvas.Pen().SetColor(types.ClSkyblue)
		canvas.Rectangle(aRect.Left+1, aRect.Top+1, aRect.Right-1, aRect.Bottom-1)
		canvas.Rectangle(aRect.Left, aRect.Top, aRect.Right, aRect.Bottom)
		if rtl.InSets(uint32(state), types.OdSelected) {
			canvas.Brush().SetColor(0x00FFB2B5)
			canvas.Rectangle(aRect.Left+1, aRect.Top+1, aRect.Right-1, aRect.Bottom-1)
			canvas.Font().SetColor(types.ClBlue)
			if rtl.InSets(uint32(state), types.OdFocused) {
				canvas.DrawFocusRect(aRect)
			}
		}
		canvas.TextOut(aRect.Left+(aRect.Right-fw)/2, aRect.Top+(itemHeight-fh)/2, s)
	})

	vcl.Application.Run()
}
