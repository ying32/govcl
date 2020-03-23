//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import . "github.com/ying32/govcl/vcl/types"

type IWinControl interface {
	IControl
	Handle() HWND
	SetHandle(HWND)

	TabStop() bool
	SetTabStop(bool)

	TabOrder() TTabOrder
	SetTabOrder(TTabOrder)

	Showing() bool

	ParentWindow() HWND
	SetParentWindow(HWND)

	ParentDoubleBuffered() bool
	SetParentDoubleBuffered(bool)

	ControlCount() int32
	Controls(index int32) *TControl

	Brush() *TBrush
	//Padding() *TPadding

	SetFocus()
	Update()

	CanFocus() bool
	ContainsControl(IControl) bool
	DisableAlign()
	DockClientCount() int32
	DockClients(int32) *TControl

	DoubleBuffered() bool
	SetDoubleBuffered(bool)

	EnableAlign()

	FindChildControl(string) *TControl
	FlipChildren(bool)

	Focused() bool

	InsertControl(IControl)
	RemoveControl(IControl)

	ScaleBy(M int32, D int32)
	ScrollBy(DeltaX int32, DeltaY int32)

	MouseInClient() bool
	AlignDisabled() bool
	//Invalidate()

	UpdateControlState()

	HandleAllocated() bool
	PaintTo(DC HDC, X int32, Y int32)

	ClientToScreen(Point TPoint) TPoint
	ClientToParent(Point TPoint, AParent IWinControl) TPoint

	UseDockManager() bool
	SetUseDockManager(value bool)
}
