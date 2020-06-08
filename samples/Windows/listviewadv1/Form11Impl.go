// +build windows

package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm11Fields struct {
}

var (
	Colors = []types.TColor{colors.ClRed, colors.ClGreen, colors.ClBlue, colors.ClLtGray}
)

func (f *TForm11) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()
	rand.Seed(time.Now().Unix())

	for i := 0; i < 20; i++ {
		f.addItem()
	}
}

func (f *TForm11) GetSubItemRect(iItem, iSubItem int32) (ret types.TRect) {
	win.ListView_GetSubItemRect(f.ListView1.Handle(), iItem, iSubItem /*win.LVIR_LABEL*/, win.LVIR_BOUNDS, &ret)
	return
}

func (f *TForm11) ColorLblClick(sender vcl.IObject) {
	if f.ColorDialog1.Execute() {
		vcl.AsLabel(sender).SetColor(f.ColorDialog1.Color())
	}
}

func (f *TForm11) addItem() {
	item := f.ListView1.Items().Add()

	// 下面的项目千万千万千万要记得释放，这里只是一个演示。
	// 比如一条记录被删除时，那必须在删除前Free这行的对象
	// 至于这行对象怎么维护起来就看你自己怎么处理了，也可以通过 item.Data设置一个指针保存。
	// 表头调整大小时也不会跟着调整，所以需要在那里处理事件，哪个事件呢？我也不知道，哈。

	item.SetCaption(fmt.Sprintf("%d", f.ListView1.Items().Count()))
	item.SubItems().Add("")
	item.SubItems().Add("")
	item.SubItems().Add("")

	r := f.GetSubItemRect(item.Index(), 1)
	r.Inflate(-2, -2)
	cbb := vcl.NewComboBox(f.ListView1)
	cbb.SetParent(f.ListView1)
	cbb.Items().Add("是")
	cbb.Items().Add("否")
	cbb.SetItemIndex(0)
	cbb.SetStyle(types.CsDropDownList)
	cbb.SetBounds(r.Left, r.Top, r.Width(), r.Height())

	r = f.GetSubItemRect(item.Index(), 2)
	r.Inflate(-2, -2)
	LColorLbl := vcl.NewLabel(f.ListView1)
	LColorLbl.SetParent(f.ListView1)
	LColorLbl.SetAutoSize(false)
	LColorLbl.SetColor(Colors[rand.Intn(len(Colors)-1)])
	LColorLbl.SetTransparent(false)
	LColorLbl.SetBounds(r.Left, r.Top, r.Width(), r.Height())
	LColorLbl.SetOnClick(f.ColorLblClick)

	r = f.GetSubItemRect(item.Index(), 3)
	r.Inflate(-2, -2)
	LSpin := vcl.NewSpinEdit(f.ListView1)
	LSpin.SetParent(f.ListView1)
	LSpin.SetValue(rand.Int31n(10))
	LSpin.SetBounds(r.Left, r.Top, r.Width(), r.Height())
}
