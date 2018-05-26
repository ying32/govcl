package main

import (
	"gitee.com/ying32/govcl/vcl"

	"fmt"

	"gitee.com/ying32/govcl/vcl/types"
)

// 事件公用。啥叫事件公用呢，比如btn, btn2两个控件，处理的大部分是相同的，只是根据不同按钮处理进行不同的选择

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("事件公用演示")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	btn := vcl.NewButton(mainForm)
	btn.SetParent(mainForm)
	btn.SetName("btn1")
	btn.SetCaption("按钮1")
	btn.SetTag(1)
	btn.SetLeft(10)
	btn.SetTop(50)
	btn.SetOnClick(buttonOnClick)

	btn2 := vcl.NewButton(mainForm)
	btn2.SetParent(mainForm)
	btn2.SetName("btn2")
	btn2.SetCaption("按钮2")
	btn2.SetTag(2)
	btn2.SetLeft(10)
	btn2.SetTop(90)
	btn2.SetOnClick(buttonOnClick)

	pnl1 := vcl.NewPanel(mainForm)
	pnl1.SetParent(mainForm)
	pnl1.SetBounds(150, 20, 60, 30)
	pnl1.SetParentBackground(false)
	pnl1.SetColor(types.ClRed)
	pnl1.SetOnMouseEnter(pnlOnMouseEnter)
	pnl1.SetOnMouseLeave(pnlOnMouseLeave)

	pnl2 := vcl.NewPanel(mainForm)
	pnl2.SetParent(mainForm)
	pnl2.SetBounds(150, pnl1.Top()+pnl1.Height()+5, 60, 30)
	pnl2.SetParentBackground(false)
	pnl2.SetColor(types.ClGreen)
	pnl2.SetOnMouseEnter(pnlOnMouseEnter)
	pnl2.SetOnMouseLeave(pnlOnMouseLeave)

	pnl3 := vcl.NewPanel(mainForm)
	pnl3.SetParent(mainForm)
	pnl3.SetBounds(150, pnl2.Top()+pnl2.Height()+5, 60, 30)
	pnl3.SetParentBackground(false)
	pnl3.SetColor(types.ClBlue)
	pnl3.SetOnMouseEnter(pnlOnMouseEnter)
	pnl3.SetOnMouseLeave(pnlOnMouseLeave)

	vcl.Application.Run()
}

// 两个按钮使用同一个事件回调,Delphi里称为方法(method)
func buttonOnClick(sender vcl.IObject) {
	// 这里就可以根据sender去做选择了
	btn := vcl.ButtonFromObj(sender)
	switch btn.Tag() {
	case 1:
		fmt.Println("按钮1的")
	case 2:
		fmt.Println("按钮2的")
	}
	vcl.ShowMessage("消息，Caption：" + btn.Caption() + ", Name：" + btn.Name())
}

func pnlOnMouseEnter(sender vcl.IObject) {
	fmt.Println("鼠标进入")
	pnl := vcl.PanelFromObj(sender)
	pnl.SetLeft(pnl.Left() + 10)
}

func pnlOnMouseLeave(sender vcl.IObject) {
	fmt.Println("鼠标离开")
	pnl := vcl.PanelFromObj(sender)
	pnl.SetLeft(pnl.Left() - 10)
}
