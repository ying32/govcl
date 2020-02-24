
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package vcl

import . "github.com/ying32/govcl/vcl/types"

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
	Realign()
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

	Parent() *TWinControl
	SetParent(IWinControl)

	SetBounds(int32, int32, int32, int32)

	Caption() string
	SetCaption(string)

	Cursor() TCursor
	SetCursor(TCursor)

	Action() *TAction
	SetAction(IComponent)

	Anchors() TAnchors
	SetAnchors(TAnchors)

	BoundsRect() TRect
	SetBoundsRect(TRect)

	ClientHeight() int32
	SetClientHeight(int32)

	ClientWidth() int32
	SetClientWidth(int32)

	ClientOrigin() TPoint
	ClientRect() TRect

	SetTextBuf(string)

	AlignWithMargins() bool
	SetAlignWithMargins(bool)

	Margins() *TMargins
	SetMargins(*TMargins)

	PopupMenu() *TPopupMenu
	SetPopupMenu(IComponent)

	Font() *TFont
	SetFont(*TFont)

	CustomHint() *TCustomHint
	SetCustomHint(IComponent)

	ParentCustomHint() bool
	SetParentCustomHint(bool)

	StyleElements() TStyleElements
	SetStyleElements(TStyleElements)

	Floating() bool
	SetFloating(bool)

	Color() TColor
	SetColor(TColor)

	ControlStyle() TControlStyle
	SetControlStyle(TControlStyle)

	ControlState() TControlState
	SetControlState(TControlState)

	// ---------------------事件

	//SetOnClick(TNotifyEvent)
}
