package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Drap And Drop")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(500)
	mainForm.SetHeight(400)

	// tv 演示自身内的拖放
	tv := vcl.NewTreeView(mainForm)
	tv.SetParent(mainForm)
	tv.SetAlign(types.AlLeft)
	tv.SetWidth(250)
	tv.SetDragMode(types.DmAutomatic)

	tv.SetOnDragOver(func(sender, source vcl.IObject, x, y int32, state types.TDragState, accept *bool) {
		*accept = false
		if source.IsValid() {
			node := vcl.TreeViewFromObj(source).GetNodeAt(x, y)
			selnode := vcl.TreeViewFromObj(source).Selected()
			if node.IsValid() && selnode.IsValid() {
				*accept = selnode.Parent().Instance() != node.Parent().Instance()
			}
		}
	})

	tv.SetOnEndDrag(func(sender, target vcl.IObject, x, y int32) {
		if target.IsValid() {
			node := tv.GetNodeAt(x, y)
			selnode := tv.Selected()
			if node.IsValid() && selnode.IsValid() {
				selnode.MoveTo(node, types.NaInsert) // NaAdd
			}
		}
	})

	tv.Items().BeginUpdate()
	for i := 0; i < 10; i++ {
		node := tv.Items().AddChild(nil, fmt.Sprintf("第%d个项目", i))
		for j := 0; j < 5; j++ {
			tv.Items().AddChild(node, fmt.Sprintf("第几%d个子项目，父索引%d", j, i))
		}
	}
	tv.Items().EndUpdate()

	// 演示从tv拖到tv2
	tv2 := vcl.NewTreeView(mainForm)
	tv2.SetParent(mainForm)
	tv2.SetAlign(types.AlRight)
	tv2.SetWidth(250)
	tv2.SetDragMode(types.DmAutomatic)

	//tv2.SetOnDragOver(func(sender, source vcl.IObject, x, y int32, state types.TDragState, accept *bool) {
	//	//*accept = false
	//	//node := tv2.GetNodeAt(x, y)
	//	//selnode := tv2.Selected()
	//	//if node.IsValid() && selnode.IsValid() {
	//	//	*accept = selnode.Parent().Instance() != node.Parent().Instance()
	//	//}
	//})
	//
	//tv2.SetOnEndDrag(func(sender, target vcl.IObject, x, y int32) {
	//	if target.IsValid() {
	//		fmt.Println("有目标，目标是否为自己：", target.Instance() == tv2.Instance())
	//		//node := tv.GetNodeAt(x, y)
	//		//selnode := tv.Selected()
	//		//if node.IsValid() && selnode.IsValid() {
	//		//	selnode.MoveTo(node, types.NaInsert) // NaAdd
	//		//}
	//	}
	//})

	tv2.Items().BeginUpdate()
	for i := 0; i < 10; i++ {
		node := tv2.Items().AddChild(nil, fmt.Sprintf("第%d个项目", i))
		for j := 0; j < 5; j++ {
			tv2.Items().AddChild(node, fmt.Sprintf("第几%d个子项目", j))
		}
	}
	tv2.Items().EndUpdate()

	vcl.Application.Run()
}
