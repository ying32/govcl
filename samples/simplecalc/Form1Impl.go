// 由res2go自动生成。
// 在这里写你的事件。

package main

import (
	"fmt"
	"strings"

	"github.com/Knetic/govaluate"
	"github.com/ying32/govcl/vcl"
)

//::private::
type TForm1Fields struct {
}

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	//f.Edit1.Clear()
	//f.Label1.SetCaption("")
	// 本程序布局解析：
	/*
							   TForm
							     |- TPanel <Align: alClient; BorderSpacing： left, top, bottom, right=8>
							        |- TPanel <
				                             	Align: alClient;
					                            ChildSizing:
					                              ControlsPerLine=5                           // 行数
					                              EnlargeHorizontal=crsHomogenousChildResize  // 子控件横向宽度
					                              EnlargeVertical=crsHomogenousChildResize    // 子控件纵向高度
			                                      Layout=cclTopToBottomThenLeftToRight        // 从上到下再左到右排序
					                        >
							             |- TSpeedButton < Flat=true >  // 20个按钮
							               ....
		                            |- TEdit <Align: alTop>
							        |- TLabel < Align: alTop>

	*/
}

func (f *TForm1) OnSpeedButton1Click(sender vcl.IObject) {
	// 所有按钮共用一个事件
	btn := vcl.AsSpeedButton(sender)
	text := btn.Caption()

	switch text {
	case "C":
		f.Edit1.Clear()
		f.Label1.SetCaption("")
	case "=":
		s1 := f.Edit1.Text()
		s := strings.Replace(strings.Replace(s1, "÷", "/", -1), "x", "*", -1)

		expression, err := govaluate.NewEvaluableExpression(s)
		if err != nil {
			f.Edit1.SetText(err.Error())
			return
		}
		result, err := expression.Evaluate(nil)
		if err != nil {
			f.Edit1.SetText(err.Error())
			return
		}
		f.Label1.SetCaption(s1)
		f.Edit1.SetText(fmt.Sprint(result))
	default:
		f.Edit1.SetText(f.Edit1.Text() + text)
	}
}
