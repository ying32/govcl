//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

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
type TDrawItemEvent func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)

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

//TStartDragEvent = procedure(Sender: TObject; var DragObject: TDragObject) of object;
//type TStartDragEvent func(sender IObject, dragObject **TDragObject)

//TEndDragEvent = procedure(Sender, Target: TObject; X, Y: Integer) of object;
type TEndDragEvent func(sender, target IObject, x, y int32)

//TDockDropEvent = procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer) of object;
type TDockDropEvent func(sender IObject, source *TDragDockObject, x, y int32)

//TDockOverEvent = procedure(Sender: TObject; Source: TDragDockObject; X, Y: Integer; State: TDragState; var Accept: Boolean) of object;
type TDockOverEvent func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)

//TUnDockEvent = procedure(Sender: TObject; Client: TControl; NewTarget: TWinControl; var Allow: Boolean) of object;
type TUnDockEvent func(sender IObject, client *TControl, newTarget *TControl, allow *bool)

//TStartDockEvent = procedure(Sender: TObject; var DragObject: TDragDockObject) of object;
type TStartDockEvent func(sender IObject, dragObject **TDragDockObject)

//TGetSiteInfoEvent = procedure(Sender: TObject; DockClient: TControl; var InfluenceRect: TRect; MousePos: TPoint; var CanDock: Boolean) of object;
type TGetSiteInfoEvent func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)

//TMouseWheelUpDownEvent = procedure(Sender: TObject; Shift: TShiftState; MousePos: TPoint; var Handled: Boolean) of object;
type TMouseWheelUpDownEvent func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)

// ---- grid

//TGridOperationEvent = procedure (Sender: TObject; IsColumn:Boolean;  sIndex, tIndex: Integer) of object;
type TGridOperationEvent func(sender IObject, isColumn bool, sIndex, tIndex int32)

//TDrawCellEvent = procedure (Sender: TObject; ACol, ARow: Longint; Rect: TRect; State: TGridDrawState) of object;
type TDrawCellEvent func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)

//TFixedCellClickEvent = procedure (Sender: TObject; ACol, ARow: Longint) of object;
type TFixedCellClickEvent func(sender IObject, aCol, aRow int32)

// TGetEditEvent = procedure (Sender: TObject; ACol, ARow: Longint; var Value: string) of object;
type TGetEditEvent func(sender IObject, aCol, aRow int32, value *string)

// TSelectCellEvent = procedure (Sender: TObject; ACol, ARow: Longint; var CanSelect: Boolean) of object;
type TSelectCellEvent func(sender IObject, aCol, aRow int32, canSelect *bool)

// TSetEditEvent = procedure (Sender: TObject; ACol, ARow: Longint; const Value: string) of object;
type TSetEditEvent func(sender IObject, aCol, aRow int32, value string)

// ---- headercontrol

// TDrawSectionEvent = procedure(HeaderControl: THeaderControl; Section: THeaderSection; const Rect: TRect; Pressed: Boolean) of object;
type TDrawSectionEvent func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)

// TSectionNotifyEvent = procedure(HeaderControl: THeaderControl; Section: THeaderSection) of object;
type TSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)

//TSectionTrackEvent = procedure(HeaderControl: THeaderControl; Section: THeaderSection; Width: Integer; State: TSectionTrackState) of object;
type TSectionTrackEvent func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)

// TSectionDragEvent = procedure (Sender: TObject; FromSection, ToSection: THeaderSection; var AllowDrag: Boolean) of object;
type TSectionDragEvent func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)

// TCustomSectionNotifyEvent = procedure(HeaderControl: TCustomHeaderControl; Section: THeaderSection) of object;
type TCustomSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)

//
// TGestureEvent = procedure(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean) of object;
//type TGestureEvent func(sender IObject, eventInfo TGestureEventInfo, handled *bool)

// TMouseActivateEvent = procedure(Sender: TObject; Button: TMouseButton; Shift: TShiftState; X, Y: Integer; HitTest: Integer; var MouseActivate: TMouseActivate) of object;
type TMouseActivateEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)

// TLBGetDataEvent = procedure(Control: TWinControl; Index: Integer; var Data: string) of object;
type TLBGetDataEvent func(control *TWinControl, index int32, data *string)

// TLBGetDataObjectEvent = procedure(Control: TWinControl; Index: Integer; var DataObject: TObject) of object;
type TLBGetDataObjectEvent func(control *TWinControl, index int32, dataObject IObject)

// TLBFindDataEvent = function(Control: TWinControl; FindString: string): Integer of object;
type TLBFindDataEvent func(control *TWinControl, findString string) int32

// TMeasureItemEvent = procedure(Control: TWinControl; Index: Integer; var Height: Integer) of object;
type TMeasureItemEvent func(control *TWinControl, index int32, height *int32)

// TLVChangingEvent = procedure(Sender: TObject; Item: TListItem; Change: TItemChange; var AllowChange: Boolean) of object;
type TLVChangingEvent func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)

// TLVDataEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVDataEvent func(sender IObject, item *TListItem)

// TLVDataFindEvent = procedure(Sender: TObject; Find: TItemFind; const FindString: string;
//  const FindPosition: TPoint; FindData: TCustomData; StartIndex: Integer; Direction: TSearchDirection; Wrap: Boolean; var Index: Integer) of object;
type TLVDataFindEvent func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32, direction TSearchDirection, warp bool, index *int32)

//TLVDeletedEvent = procedure(Sender: TObject; Item: TListItem) of object;
type TLVDeletedEvent func(sender IObject, item *TListItem)

//TLVEditingEvent = procedure(Sender: TObject; Item: TListItem; var AllowEdit: Boolean) of object;
type TLVEditingEvent func(sender IObject, item *TListItem, allowEdit *bool)

//TLVEditedEvent = procedure(Sender: TObject; Item: TListItem; var S: string) of object;
type TLVEditedEvent func(sender IObject, item *TListItem, s *string)

// TMenuMeasureItemEvent = procedure (Sender: TObject; ACanvas: TCanvas; var Width, Height: Integer) of object;
type TMenuMeasureItemEvent func(sender IObject, aCanvas *TCanvas, width, height *int32)

// TTabChangingEvent = procedure(Sender: TObject; var AllowChange: Boolean) of object;
type TTabChangingEvent func(sender IObject, allowChange *bool)

// TTVChangingEvent = procedure(Sender: TObject; Node: TTreeNode; var AllowChange: Boolean) of object;
type TTVChangingEvent func(sender IObject, node *TTreeNode, allowChange *bool)

// TTVCollapsingEvent = procedure(Sender: TObject; Node: TTreeNode; var AllowCollapse: Boolean) of object;
type TTVCollapsingEvent func(sender IObject, node *TTreeNode, allowCollapse *bool)

// TTVEditedEvent = procedure(Sender: TObject; Node: TTreeNode; var S: string) of object;
type TTVEditedEvent func(sender IObject, node *TTreeNode, s *string)

// TTVEditingEvent = procedure(Sender: TObject; Node: TTreeNode; var AllowEdit: Boolean) of object;
type TTVEditingEvent func(sender IObject, node *TTreeNode, allowEdit *bool)

// TTVExpandingEvent = procedure(Sender: TObject; Node: TTreeNode; var AllowExpansion: Boolean) of object;
type TTVExpandingEvent func(sender IObject, node *TTreeNode, allowExpansion *bool)

// TTVHintEvent = procedure(Sender: TObject; const Node: TTreeNode; var Hint: String) of object;
type TTVHintEvent func(sender IObject, node *TTreeNode, hint *string)

// TUDChangingEvent = procedure (Sender: TObject; var AllowChange: Boolean) of object;
type TUDChangingEvent func(sender IObject, allowChange *bool)

// TCreatingListErrorEvent = procedure (Sender: TObject; WinErrorCode: Cardinal; const ErrorDescription: string; var Handled: Boolean) of object;
type TCreatingListErrorEvent func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)

//--

//TLVCustomDrawEvent = procedure(Sender: TCustomListView; const ARect: TRect; var DefaultDraw: Boolean) of object;
type TLVCustomDrawEvent func(sender *TListView, aRect TRect, defaultDraw *bool)

//TLVCustomDrawItemEvent = procedure(Sender: TCustomListView; Item: TListItem; State: TCustomDrawState; var DefaultDraw: Boolean) of object;
type TLVCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, defaultDraw *bool)

//TLVCustomDrawSubItemEvent = procedure(Sender: TCustomListView; Item: TListItem; SubItem: Integer; State: TCustomDrawState; var DefaultDraw: Boolean) of object;
type TLVCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, defaultDraw *bool)

//TLVDrawItemEvent = procedure(Sender: TCustomListView; Item: TListItem; Rect: TRect; State: TOwnerDrawState) of object;
type TLVDrawItemEvent func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)

//TLVDataHintEvent = procedure(Sender: TObject; StartIndex, EndIndex: Integer) of object;
type TLVDataHintEvent func(sender IObject, startIndex, endIndex int32)

//TTVCustomDrawEvent = procedure(Sender: TCustomTreeView; const ARect: TRect;var DefaultDraw: Boolean) of object;
type TTVCustomDrawEvent func(sender *TTreeView, aRect TRect, defaultDraw *bool)

//TTVCustomDrawItemEvent = procedure(Sender: TCustomTreeView; Node: TTreeNode; State: TCustomDrawState; var DefaultDraw: Boolean) of object;
type TTVCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, defaultDraw *bool)

// 消息过程

//TWndProcEvent = procedure(Sender: TObject; var AMsg: TMessage) of object;
type TWndProcEvent func(msg *TMessage)

// TWebTitleChangeEvent = procedure(Sender: TObject; const Text: string) of object;
type TWebTitleChangeEvent func(sender IObject, text string)

// TWebJSExternalEvent = procedure(Sender: TObject; const func: string; const args: WideString; var retval: WideString) of object;
type TWebJSExternalEvent func(sender IObject, funcName, args string, retVal *string)

//TTaskDlgClickEvent = procedure(Sender: TObject; ModalResult: TModalResult; var CanClose: Boolean) of object;
type TTaskDlgClickEvent func(sender IObject, modalResult TModalResult, canClose *bool)

//TTaskDlgTimerEvent = procedure(Sender: TObject; TickCount: Cardinal; var Reset: Boolean) of object;
type TTaskDlgTimerEvent func(sender IObject, tickCount uint32, reset *bool)

//TAlignPositionEvent = procedure(Sender: TWinControl; Control: TControl; var NewLeft, NewTop, NewWidth, NewHeight: Integer; var AlignRect: TRect; AlignInfo: TAlignInfo) of object;
type TAlignPositionEvent func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)

//TCheckGroupClicked = procedure(Sender: TObject; Index: integer) of object;
type TCheckGroupClicked func(sender IObject, index int32)

// --------------

//TOnSelectEvent = procedure(Sender: TObject; aCol, aRow: Integer) of object;
type TOnSelectEvent func(sender IObject, aCol, aRow int32)

//TToggledCheckboxEvent = procedure(sender: TObject; aCol, aRow: Integer; aState: TCheckboxState) of object;
type TToggledCheckboxEvent func(sender IObject, aCol, aRow int32, aState TCheckBoxState)

//TOnCompareCells = procedure (Sender: TObject; ACol, ARow, BCol,BRow: Integer;var Result: integer) of object;
type TOnCompareCells func(sender IObject, ACol, ARow, BCol, BRow int32, result *int32)

//TGetCellHintEvent = procedure (Sender: TObject; ACol, ARow: Integer; var HintText: String) of object;
type TGetCellHintEvent func(sender IObject, ACol, ARow int32, hintText *string)

//TGetCheckboxStateEvent = procedure (Sender: TObject; ACol, ARow: Integer; var Value: TCheckboxState) of object;
type TGetCheckboxStateEvent func(sender IObject, ACol, ARow int32, value *TCheckBoxState)

//TSetCheckboxStateEvent = procedure (Sender: TObject; ACol, ARow: Integer; const Value: TCheckboxState) of object;
type TSetCheckboxStateEvent func(sender IObject, ACol, ARow int32, Value TCheckBoxState)

//THdrEvent = procedure(Sender: TObject; IsColumn: Boolean; Index: Integer) of object;
type THdrEvent func(sender IObject, isColumn bool, index int32)

//THeaderSizingEvent = procedure(sender: TObject; const IsColumn: boolean; const aIndex, aSize: Integer) of object;
type THeaderSizingEvent func(sender IObject, isColumn bool, aIndex, aSize int32)

//TSelectEditorEvent = procedure(Sender: TObject; aCol, aRow: Integer; var Editor: TWinControl) of object;
type TSelectEditorEvent func(sender IObject, aCol, aRow int32, editor **TWinControl)

//TUserCheckBoxBitmapEvent = procedure(Sender: TObject; const aCol, aRow: Integer; const CheckedState: TCheckboxState; var ABitmap: TBitmap) of object;
type TUserCheckBoxBitmapEvent func(sender IObject, aCol, aRow int32, CheckedState TCheckBoxState, aBitmap **TBitmap)

//TValidateEntryEvent = procedure(sender: TObject; aCol, aRow: Integer; const OldValue: string; var NewValue: String) of object;
type TValidateEntryEvent func(sender IObject, aCol, aRow int32, oldValue string, newValue *string)

// TOnPrepareCanvasEvent = procedure(sender: TObject; aCol, aRow: Integer; aState: TGridDrawState) of object;
type TOnPrepareCanvasEvent = func(sender IObject, aCol, aRow int32, aState TGridDrawState)
