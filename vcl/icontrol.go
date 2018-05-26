package vcl

import (
	. "gitee.com/ying32/govcl/vcl/types"
)

type IControl interface {
	IComponent
	// ---------------------方法

	Perform(uint32, uintptr, int) int
	SendToBack()
	Show()
	Hide()
	Refresh()
	Repaint()
	Invalidate()

	// ---------------------属性

	Align() TAlign
	SetAlign(TAlign)

	Visible() bool
	SetVisible(bool)

	Enabled() bool
	SetEnabled(bool)

	Left() int32
	SetLeft(int32)

	Top() int32
	SetTop(int32)

	Width() int32
	SetWidth(int32)

	Height() int32
	SetHeight(int32)

	Hint() string
	SetHint(string)

	ShowHint() bool
	SetShowHint(bool)

	Parent() *TControl
	SetParent(IControl)

	SetBounds(int32, int32, int32, int32)
	// ---------------------事件

	//SetOnClick(TNotifyEvent)
}
