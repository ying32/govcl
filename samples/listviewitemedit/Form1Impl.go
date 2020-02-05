// 在这里写你的事件

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl"
)

// --------------------------------------------

//::private::
type TForm1Fields struct {
	subItemHit TLVHitTestInfo
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {

	for i := 0; i < 10; i++ {
		f.ComboBox1.Items().Add(fmt.Sprintf("Item %d", i))
	}
	f.ListView1.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		item := f.ListView1.Items().Add()
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("Edit-%d", i))
		item.SubItems().Add(fmt.Sprintf("Combobox-%d", i))
	}
	f.ListView1.Items().EndUpdate()
}

func (f *TForm1) OnComboBox1Exit(sender vcl.IObject) {
	f.ComboBox1.Hide()
	if f.subItemHit.iItem != -1 {
		if f.subItemHit.iSubItem == 2 {
			f.ListView1.Items().Item(f.subItemHit.iItem).SubItems().SetStrings(f.subItemHit.iSubItem-1, f.ComboBox1.Text())
		}
	}
}

func (f *TForm1) OnEdit1Exit(sender vcl.IObject) {
	f.Edit1.Hide()
	if f.subItemHit.iItem != -1 {
		if f.subItemHit.iSubItem == 1 {
			f.ListView1.Items().Item(f.subItemHit.iItem).SubItems().SetStrings(f.subItemHit.iSubItem-1, f.Edit1.Text())
		}
	}
}

func (f *TForm1) OnListView1DblClick(sender vcl.IObject) {
	p := f.ListView1.ScreenToClient(vcl.Mouse.CursorPos())
	f.subItemHit.pt.X = p.X
	f.subItemHit.pt.Y = p.Y
	ListView_SubItemHitTest(f.ListView1.Handle(), &f.subItemHit)
	if f.subItemHit.iItem != -1 {
		fmt.Println(fmt.Sprintf("f.subItemHit.iItem=%d, f.subItemHit.iSubItem=%d", f.subItemHit.iItem, f.subItemHit.iSubItem))

		var r types.TRect
		ListView_GetItemRect(f.ListView1.Handle(), f.subItemHit.iItem, &r, 0)

		colWidht := ListView_GetColumnWidth(f.ListView1.Handle(), f.subItemHit.iSubItem)
		//var itemPoint types.TPoint
		//ListView_GetItemPosition(f.ListView1.Handle(), f.subItemHit.iItem, &itemPoint)

		fmt.Println("subItemHit.iSubItem =", f.subItemHit.iSubItem)
		//fmt.Println("colWidht=", colWidht, ", pos=", itemPoint)

		var left, i int32
		// 差2个像素???????????????????????????????????
		left += 2
		r.Top += 2
		for i = 0; i < f.subItemHit.iSubItem; i++ {
			left += ListView_GetColumnWidth(f.ListView1.Handle(), i)
		}
		fmt.Println("left:", left)

		switch f.subItemHit.iSubItem {
		case 0: // 不处理
		case 1:
			// edit
			f.Edit1.SetText(f.ListView1.Items().Item(f.subItemHit.iItem).SubItems().Strings(f.subItemHit.iSubItem - 1))
			f.Edit1.SetBounds(left, f.ListView1.Top()+r.Top, colWidht, r.Bottom-r.Top)
			f.Edit1.Show()
			f.Edit1.SetFocus()
		case 2:
			//combobox
			f.ComboBox1.SetText(f.ListView1.Items().Item(f.subItemHit.iItem).SubItems().Strings(f.subItemHit.iSubItem - 1))
			f.ComboBox1.SetBounds(left, f.ListView1.Top()+r.Top, colWidht, r.Bottom-r.Top)
			f.ComboBox1.Show()
			f.ComboBox1.SetFocus()
		}
	}
}
