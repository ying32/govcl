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

type TNotifyEvent func(sender IObject)

type TUDClickEvent func(sender IObject, button TUDBtnType)

type TCloseEvent func(sender IObject, action *TCloseAction)

type TCloseQueryEvent func(sender IObject, canClose *bool)

type TMenuChangeEvent func(sender IObject, source *TMenuItem, rebuild bool)

type TSysLinkEvent func(sender IObject, link string, linkType TSysLinkType)

type TExceptionEvent func(sender IObject, e *Exception)

type TKeyEvent func(sender IObject, key *Char, shift TShiftState)

type TKeyPressEvent func(sender IObject, key *Char)

type TMouseEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)

type TMouseMoveEvent func(sender IObject, shift TShiftState, x, y int32)

type TMouseWheelEvent func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)

type TDrawItemEvent func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)

type TLVColumnClickEvent func(sender IObject, column *TListColumn)

type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)

type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)

type TLVCheckedItemEvent func(sender IObject, item *TListItem)

type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)

type TLVChangeEvent func(sender IObject, item *TListItem, change TItemChange)

type TLVNotifyEvent func(sender IObject, item *TListItem)

type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)

type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)

type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)

type TTVExpandedEvent func(sender IObject, node *TTreeNode)

type TTVChangedEvent func(sender IObject, node *TTreeNode)

type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)

type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)

type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)

type TThreadProc func()

// TDropFilesEvent
// 注意，当在Windows上使用时如果使用了UAC，则无法收到消息
// 需要使用未公开的winapi   ChangeWindowMessageFilter 或 ChangeWindowMessageFilterEx 根据系统版本不同使用其中的，然后添加
// ChangeWindowMessageFilterEx(pnl_Drag.Handle, WM_DROPFILES, MSGFLT_ALLOW, LChangeFilterStruct);消息
type TDropFilesEvent func(sender IObject, aFileNames []string)

type TConstrainedResizeEvent func(sender IObject, minWidth, minHeight, maxWidth, maxHeight *int32)

type THelpEvent func(command uint16, data THelpEventData, callhelp, result *bool)

type TShortCutEvent func(msg *TWMKey, handled *bool)

type TContextPopupEvent func(sender IObject, mousePos TPoint, handled *bool)

type TDragOverEvent func(sender, source IObject, x, y int32, state TDragState, accept *bool)

type TDragDropEvent func(sender, source IObject, x, y int32)

//TStartDragEvent = procedure(Sender: TObject; var DragObject: TDragObject) of object;
//type TStartDragEvent func(sender IObject, dragObject **TDragObject)

type TEndDragEvent func(sender, target IObject, x, y int32)

type TDockDropEvent func(sender IObject, source *TDragDockObject, x, y int32)

type TDockOverEvent func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)

type TUnDockEvent func(sender IObject, client *TControl, newTarget *TControl, allow *bool)

type TStartDockEvent func(sender IObject, dragObject **TDragDockObject)

type TGetSiteInfoEvent func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)

type TMouseWheelUpDownEvent func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)

type TGridOperationEvent func(sender IObject, isColumn bool, sIndex, tIndex int32)

type TDrawCellEvent func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)

type TFixedCellClickEvent func(sender IObject, aCol, aRow int32)

type TGetEditEvent func(sender IObject, aCol, aRow int32, value *string)

type TSelectCellEvent func(sender IObject, aCol, aRow int32, canSelect *bool)

type TSetEditEvent func(sender IObject, aCol, aRow int32, value string)

type TDrawSectionEvent func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)

type TSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)

type TSectionTrackEvent func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)

type TSectionDragEvent func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)

type TCustomSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)

//
// TGestureEvent = procedure(Sender: TObject; const EventInfo: TGestureEventInfo; var Handled: Boolean) of object;
//type TGestureEvent func(sender IObject, eventInfo TGestureEventInfo, handled *bool)

type TMouseActivateEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)

type TLBGetDataEvent func(control *TWinControl, index int32, data *string)

type TLBGetDataObjectEvent func(control *TWinControl, index int32, dataObject IObject)

type TLBFindDataEvent func(control *TWinControl, findString string) int32

type TMeasureItemEvent func(control *TWinControl, index int32, height *int32)

type TLVChangingEvent func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)

type TLVDataEvent func(sender IObject, item *TListItem)

type TLVDataFindEvent func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32, direction TSearchDirection, warp bool, index *int32)

type TLVDeletedEvent func(sender IObject, item *TListItem)

type TLVEditingEvent func(sender IObject, item *TListItem, allowEdit *bool)

type TLVEditedEvent func(sender IObject, item *TListItem, s *string)

type TMenuMeasureItemEvent func(sender IObject, aCanvas *TCanvas, width, height *int32)

type TTabChangingEvent func(sender IObject, allowChange *bool)

type TTVChangingEvent func(sender IObject, node *TTreeNode, allowChange *bool)

type TTVCollapsingEvent func(sender IObject, node *TTreeNode, allowCollapse *bool)

type TTVEditedEvent func(sender IObject, node *TTreeNode, s *string)

type TTVEditingEvent func(sender IObject, node *TTreeNode, allowEdit *bool)

type TTVExpandingEvent func(sender IObject, node *TTreeNode, allowExpansion *bool)

type TTVHintEvent func(sender IObject, node *TTreeNode, hint *string)

type TUDChangingEvent func(sender IObject, allowChange *bool)

type TCreatingListErrorEvent func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)

type TLVCustomDrawEvent func(sender *TListView, aRect TRect, defaultDraw *bool)

type TLVCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, defaultDraw *bool)

type TLVCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, defaultDraw *bool)

type TLVDrawItemEvent func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)

type TLVDataHintEvent func(sender IObject, startIndex, endIndex int32)

type TTVCustomDrawEvent func(sender *TTreeView, aRect TRect, defaultDraw *bool)

type TTVCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, defaultDraw *bool)

type TWndProcEvent func(msg *TMessage)

type TWebTitleChangeEvent func(sender IObject, text string)

type TWebJSExternalEvent func(sender IObject, funcName, args string, retVal *string)

type TTaskDlgClickEvent func(sender IObject, modalResult TModalResult, canClose *bool)

type TTaskDlgTimerEvent func(sender IObject, tickCount uint32, reset *bool)

type TAlignPositionEvent func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)

type TCheckGroupClicked func(sender IObject, index int32)

type TOnSelectEvent func(sender IObject, aCol, aRow int32)

type TToggledCheckboxEvent func(sender IObject, aCol, aRow int32, aState TCheckBoxState)

type TOnCompareCells func(sender IObject, ACol, ARow, BCol, BRow int32, result *int32)

type TGetCellHintEvent func(sender IObject, ACol, ARow int32, hintText *string)

type TGetCheckboxStateEvent func(sender IObject, ACol, ARow int32, value *TCheckBoxState)

type TSetCheckboxStateEvent func(sender IObject, ACol, ARow int32, Value TCheckBoxState)

type THdrEvent func(sender IObject, isColumn bool, index int32)

type THeaderSizingEvent func(sender IObject, isColumn bool, aIndex, aSize int32)

type TSelectEditorEvent func(sender IObject, aCol, aRow int32, editor **TWinControl)

type TUserCheckBoxBitmapEvent func(sender IObject, aCol, aRow int32, CheckedState TCheckBoxState, aBitmap **TBitmap)

type TValidateEntryEvent func(sender IObject, aCol, aRow int32, oldValue string, newValue *string)

type TOnPrepareCanvasEvent = func(sender IObject, aCol, aRow int32, aState TGridDrawState)

type TAcceptFileNameEvent = func(sender IObject, value *string)

type TCheckItemChange = func(sender IObject, index int32)

type TUTF8KeyPressEvent = func(sender IObject, utf8key *TUTF8Char)

type TMenuDrawItemEvent = func(sender IObject, aCanvas *TCanvas, aRect TRect, aState TOwnerDrawState)
