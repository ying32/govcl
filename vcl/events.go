package vcl

import (
	. "github.com/ying32/govcl/vcl/types"
)

// TNotifyEvent
type TNotifyEvent func(sender IObject)

// TUDClickEvent TUpDown  TUDBtnType
type TUDClickEvent func(sender IObject, button TUDBtnType)

// TCloseEvent Form
type TCloseEvent func(sender IObject, action *TCloseAction) // Action *uintptr

// TCloseQueryEvent Form
type TCloseQueryEvent func(sender IObject, canClose *bool) //CanClose *uintptr

// TMenuChangeEvent Menu
type TMenuChangeEvent func(sender IObject, source *TMenuItem, rebuild bool)

// TSysLinkEvent LinkLabel
type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType

// TExceptionEvent TApplication
type TExceptionEvent func(sender IObject, e *Exception)

// TKeyEvent = procedure(Sender: TObject; var Key: Word; Shift: TShiftState)
type TKeyEvent func(sender IObject, key *Char, shift TShiftState)

// TKeyPressEvent = procedure(Sender: TObject; var Key: Char) of object;
type TKeyPressEvent func(sender IObject, key *Char)

// TMouseEvent = procedure(Sender: TObject; Button: TMouseButton;
//     Shift: TShiftState; X, Y: Integer) of object;
type TMouseEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)

// TMouseMoveEvent = procedure(Sender: TObject; Shift: TShiftState;
//    X, Y: Integer) of object;
type TMouseMoveEvent func(sender IObject, shift TShiftState, x, y int32)

// TMouseWheelEvent = procedure(Sender: TObject; Shift: TShiftState;
//    WheelDelta: Integer; MousePos: TPoint; var Handled: Boolean) of object;
type TMouseWheelEvent func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)

//  TDrawItemEvent = procedure(Control: TWinControl; Index: Integer;
//    Rect: TRect; State: TOwnerDrawState) of object;
type TDrawItemEvent func(control IControl, index int32, aRect TRect, state TOwnerDrawState)

//  TMenuDrawItemEvent = procedure (Sender: TObject; ACanvas: TCanvas;
//    ARect: TRect; Selected: Boolean) of object;
type TMenuDrawItemEvent func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)

// ---------------TListView
// TLVColumnClickEvent = procedure(Sender: TObject; Column: TListColumn) of object;
type TLVColumnClickEvent func(sender IObject, column *TListColumn)

// TLVColumnRClickEvent = procedure(Sender: TObject; Column: TListColumn; Point: TPoint) of object;
type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)

// TLVSelectItemEvent = procedure(Sender: TObject; Item: TListItem;  Selected: Boolean) of object;
type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)

// TLVCheckedItemEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVCheckedItemEvent func(sender IObject, item *TListItem)

// TLVCompareEvent = procedure(Sender: TObject; Item1, Item2: TListItem;
// 	Data: Integer; var Compare: Integer) of object;
type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)

// TLVChangeEvent TListView TTItemChange
type TLVChangeEvent func(sender IObject, item *TListItem, change TItemChange)

// TLVNotifyEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVNotifyEvent func(sender IObject, item *TListItem)

//TLVAdvancedCustomDrawEvent = procedure(Sender: TCustomListView; const ARect: TRect;
//Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TLVAdvancedCustomDrawItemEvent = procedure(Sender: TCustomListView; Item: TListItem;
//State: TCustomDrawState; Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)

//TLVAdvancedCustomDrawSubItemEvent = procedure(Sender: TCustomListView; Item: TListItem;
//SubItem: Integer; State: TCustomDrawState; Stage: TCustomDrawStage;
//var DefaultDraw: Boolean) of object;
type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)

// ----------------TTreeView
// TTVCompareEvent = procedure(Sender: TObject; Node1, Node2: TTreeNode;
// 	Data: Integer; var Compare: Integer) of object;
type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)

// TTVExpandedEvent = procedure(Sender: TObject; Node: TTreeNode) of object;
type TTVExpandedEvent func(sender IObject, node *TTreeNode)

// TTVChangedEvent TTreeView
type TTVChangedEvent func(sender IObject, node *TTreeNode)

//TTVAdvancedCustomDrawEvent = procedure(Sender: TCustomTreeView; const ARect: TRect;
//  Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TTVAdvancedCustomDrawItemEvent = procedure(Sender: TCustomTreeView; Node: TTreeNode;
//  State: TCustomDrawState; Stage: TCustomDrawStage; var PaintImages,
//DefaultDraw: Boolean) of object;
type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)

// ----------------TPageControl
// TTabGetImageEvent = procedure(Sender: TObject; TabIndex: Integer; var ImageIndex: Integer) of object;
type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)

// -------------------------TToolBar
//TTBAdvancedCustomDrawEvent = procedure(Sender: TToolBar; const ARect: TRect;
//  Stage: TCustomDrawStage; var DefaultDraw: Boolean) of object;
type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

//TTBAdvancedCustomDrawBtnEvent = procedure(Sender: TToolBar; Button: TToolButton;
//  State: TCustomDrawState; Stage: TCustomDrawStage;
//  var Flags: TTBCustomDrawFlags; var DefaultDraw: Boolean) of object;
type TTBAdvancedCustomDrawBtnEvent func(sender *TToolBar, button *TToolButton, state TCustomDrawState, stage TCustomDrawStage, flags *TTBCustomDrawFlags, defaultDraw *bool)

// TThreadProc
type TThreadProc func()

// TDropFilesEvent
// 注意，当在Windows上使用时如果使用了UAC，则无法收到消息
// 需要使用未公开的winapi   ChangeWindowMessageFilter 或 ChangeWindowMessageFilterEx 根据系统版本不同使用其中的，然后添加
// ChangeWindowMessageFilterEx(pnl_Drag.Handle, WM_DROPFILES, MSGFLT_ALLOW, LChangeFilterStruct);消息
type TDropFilesEvent func(sender IObject, aFileNames []string)

// 约束调整大小事件
type TConstrainedResizeEvent func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)

// THelpEvent = function(Command: Word; Data: THelpEventData; var CallHelp: Boolean): Boolean of object;
type THelpEvent func(command uint16, data THelpEventData, callhelp, result *bool)

// TShortCutEvent = procedure (var Msg: TWMKey; var Handled: Boolean) of object;
type TShortCutEvent func(msg *TWMKey, handled *bool)

// TContextPopupEvent = procedure(Sender: TObject; MousePos: TPoint; var Handled: Boolean) of object;
type TContextPopupEvent func(sender IObject, mousePos TPoint, handled *bool)

//TDragOverEvent = procedure(Sender, Source: TObject; X, Y: Integer; State: TDragState; var Accept: Boolean) of object;
type TDragOverEvent func(sender, source IObject, x, y int32, state TDragState, accept *bool)

//TDragDropEvent = procedure(Sender, Source: TObject; X, Y: Integer) of object;
type TDragDropEvent func(sender, source IObject, x, y int32)

//TStartDragEvent = procedure(Sender: TObject;var DragObject: TDragObject) of object;
type TStartDragEvent func(sender IObject, dragObject *TDragObject)

//TEndDragEvent = procedure(Sender, Target: TObject; X, Y: Integer) of object;
type TEndDragEvent func(sender, target IObject, x, y int32)

//TDockDropEvent = procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer) of object;
type TDockDropEvent func(sender IObject, source *TDragDockObject, x, y int32)

//TDockOverEvent = procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer; State: TDragState; var Accept: Boolean) of object;
type TDockOverEvent func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)

//TUnDockEvent = procedure(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean) of object;
type TUnDockEvent func(sender IObject, client *TControl, newTarget *TControl, allow *bool)

//TStartDockEvent = procedure(Sender: TObject;var DragObject: TDragDockObject) of object;
type TStartDockEvent func(sender IObject, dragObject *TDragDockObject)

//TGetSiteInfoEvent = procedure(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean) of object;
type TGetSiteInfoEvent func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)

//TMouseWheelUpDownEvent = procedure(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean) of object;
type TMouseWheelUpDownEvent func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)

// TMessageEvent = procedure (var Msg: TMsg; var Handled: Boolean) of object;
type TMessageEvent func(msg *TMsg, handled *bool)
