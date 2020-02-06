// +build windows

// 这个功能暂时只能在Windows下使用，至于linux跟macOS的没研究过。。

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(600)
	mainForm.SetHeight(500)

	// -----------TreeView 不同Node弹出不同菜单，两个右键例程不同使用

	tvpm1 := vcl.NewPopupMenu(mainForm)
	mItem := vcl.NewMenuItem(mainForm)
	mItem.SetCaption("选中")
	tvpm1.Items().Add(mItem)

	tv := vcl.NewTreeView(mainForm)
	tv.SetParent(mainForm)
	tv.SetAlign(types.AlClient)
	tv.SetPopupMenu(tvpm1)

	mItem.SetOnClick(func(sender vcl.IObject) {
		node := tv.Selected()
		if node != nil && node.IsValid() {
			NodeChecked(node, true)
		}
	})

	// 设置显示TreeView的CheckBox
	win.SetWindowLong(tv.Handle(), win.GWL_STYLE, uintptr(win.GetWindowLong(tv.Handle(), win.GWL_STYLE))|win.TVS_CHECKBOXES)

	//	tv.Items().Clear()
	// 第一个节点
	node := tv.Items().AddChild(nil, "首个")

	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))

	}
	tv.Items().EndUpdate()
	// 展开
	node.Expand(true)

	// 第二个节点
	node = tv.Items().AddChild(nil, "第二个节点")
	// 批量添加最好使用BeginUpdate与EndUpdate组合
	tv.Items().BeginUpdate()
	for i := 0; i < 30; i++ {
		tv.Items().AddChild(node, fmt.Sprintf("Node%d", i))
	}
	tv.Items().EndUpdate()

	tv.SetOnClick(func(vcl.IObject) {
		node := tv.Selected()
		if node != nil && node.IsValid() {
			fmt.Println("选中：", IsNodeChecked(node))
		}
	})

	vcl.Application.Run()
}

// 判断某个节点的checkbox是否选中
func IsNodeChecked(node *vcl.TTreeNode) bool {
	item := win.TTVItem{}
	item.Mask = win.TVIF_STATE
	item.HItem = node.ItemId()
	win.TreeView_GetItem(node.TreeView().Handle(), &item)
	return item.State&win.TVIS_CHECKED == win.TVIS_CHECKED
}

// 节点选中
func NodeChecked(node *vcl.TTreeNode, checked bool) {
	item := win.TTVItem{}

	item.HItem = node.ItemId()
	item.Mask = win.TVIF_STATE
	item.StateMask = win.TVIS_STATEIMAGEMASK
	if checked {
		item.State = win.TVIS_CHECKED
	} else {
		item.State = win.TVIS_CHECKED >> 1
	}
	win.TreeView_SetItem(node.TreeView().Handle(), item)
}
