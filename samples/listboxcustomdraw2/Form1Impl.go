// 在这里写你的事件

package main

import (
	"fmt"
	"math"
	"strings"

	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

//::private::
type TForm1Fields struct {
}

// 这个例了是改的原来网上一个Delphi的TListBox自绘的，
// 原地址： https://blog.csdn.net/nhconch/article/details/205127

func (f *TForm1) OnFormCreate(sender vcl.IObject) {
	f.ScreenCenter()

	// 这里为显示，手动设置，一般可以设计器时设置
	f.LstLeft.SetStyle(types.LbOwnerDrawFixed)
	f.LstLeft.SetCtl3D(false)
	f.LstLeft.SetItemHeight(90)
	f.LstLeft.Items().Add("第一项")
	f.LstLeft.Items().Add("第二项")
	f.LstLeft.Items().Add("第三项")
	f.LstLeft.Items().Add("第四项")

	f.LstRight.SetStyle(types.LbOwnerDrawFixed)
	f.LstRight.SetCtl3D(false)
	f.LstRight.SetItemHeight(53)
	for i := int32(0); i < 20; i++ {
		f.LstRight.Items().Add(fmt.Sprintf("ListBox Item of %d\nSecond of %d\nThird of %d", i, i, i))
	}

}

func (f *TForm1) OnFormShow(object vcl.IObject) {
	f.LstLeft.SetItemIndex(0)
	f.LstLeft.Repaint()

	f.LstRight.SetItemIndex(0)
	f.LstRight.Repaint()
}

func (f *TForm1) OnFormDestroy(sender vcl.IObject) {

}

func (f *TForm1) OnLstLeftDrawItem(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
	r := types.TRect{}
	canvas := f.LstLeft.Canvas()
	brush := canvas.Brush()
	pen := canvas.Pen()
	font := canvas.Font()

	//设置填充的背景颜色并填充背景
	brush.SetColor(colors.ClWhite)
	canvas.FillRect(aRect)

	//绘制圆角矩形
	if state.In(types.OdSelected) {
		pen.SetColor(0xFFB2B5) //选中项的圆角矩形颜色
	} else {
		pen.SetColor(colors.ClSilver) //未选中项的圆角矩形颜色
	}
	brush.SetStyle(types.BsClear)

	r = types.TRect{Left: aRect.Left + 3, Top: aRect.Top + 3, Right: aRect.Right - 3, Bottom: aRect.Bottom - 3}
	canvas.RoundRect(r.Left, r.Top, r.Right, r.Bottom, 10, 10)

	//画出图标
	if state.In(types.OdSelected) {
		pic := f.Image1.Picture()
		canvas.Draw(r.Left+(r.Right-r.Left-pic.Width())/2, r.Top+2, pic.Graphic()) //选中项的图像
	} else {
		pic := f.Image2.Picture()
		canvas.Draw(r.Left+(r.Right-r.Left-pic.Width())/2, r.Top+2, pic.Graphic()) //未选中项的图像
	}

	//填充文字区背景
	r.Top = r.Bottom - int32(math.Abs(float64(font.Height()))) - 4
	brush.SetStyle(types.BsSolid)
	if state.In(types.OdSelected) {
		brush.SetColor(0xFFB2B5) //选中项的背景颜色
	} else {
		brush.SetColor(colors.ClSilver) //未选中项的背景颜色
	}
	canvas.FillRect(r)

	//输出文字，仅支持单行
	font.SetColor(colors.ClBlack)

	//计算文字顶点位置， 水平，垂直居中，文字超出绘制圆点
	r.Top += 2
	canvas.TextRect3(&r, f.LstLeft.Items().S(index), types.NewSet(types.TfVerticalCenter, types.TfCenter, types.TfWordEllipsis))

	//画焦点虚框，当系统再绘制时，变成XOR运算，从而达到擦除焦点虚框的目的
	if state.In(types.OdFocused) {
		canvas.DrawFocusRect(aRect)
	}

}

func (f *TForm1) OnLstRightDrawItem(control vcl.IWinControl, index int32, aRect types.TRect, state types.TOwnerDrawState) {
	canvas := f.LstRight.Canvas()
	brush := canvas.Brush()
	pen := canvas.Pen()
	font := canvas.Font()

	//文字颜色
	font.SetColor(colors.ClBlack)

	//设置背景颜色并填充背景
	brush.SetColor(colors.ClWhite)
	canvas.FillRect(aRect)

	//设置圆角矩形颜色并画出圆角矩形
	brush.SetColor(types.TColor(0x00FFF7F7))
	pen.SetColor(types.TColor(0x00131315))
	canvas.RoundRect(aRect.Left+3, aRect.Top+3, aRect.Right-2, aRect.Bottom-2, 8, 8)

	//以不同的宽度和高度再画一次，实现立体效果
	canvas.RoundRect(aRect.Left+3, aRect.Top+3, aRect.Right-3, aRect.Bottom-3, 5, 5)

	//如果是当前选中项
	if state.In(types.OdSelected) {
		//以不同的背景色画出选中项的圆角矩形
		brush.SetColor(types.TColor(0x00FFB2B5))
		canvas.RoundRect(aRect.Left+3, aRect.Top+3, aRect.Right-3, aRect.Bottom-3, 5, 5)

		//选中项的文字颜色
		font.SetColor(colors.ClBlue)

		//如果当前项拥有焦点，画焦点虚框，当系统再绘制时变成XOR运算从而达到擦除焦点虚框的目的
		if state.In(types.OdFocused) {
			canvas.DrawFocusRect(aRect)
		}
	}

	idx := 0
	if index%2 != 0 {
		idx = 1
	}
	//画出图标
	f.ImageList1.Draw(canvas, aRect.Left+7, aRect.Top+(f.LstRight.ItemHeight()-f.ImageList1.Height())/2, int32(idx), true)

	//分别绘出三行文字
	sArr := strings.Split(f.LstRight.Items().S(index), "\n")
	if len(sArr) == 3 {
		canvas.TextOut(aRect.Left+32+10, aRect.Top+4, sArr[0])
		canvas.TextOut(aRect.Left+32+10, aRect.Top+18, sArr[1])
		canvas.TextOut(aRect.Left+32+10, aRect.Top+32, sArr[2])
	}
}
