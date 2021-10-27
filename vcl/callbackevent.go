//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// 外部回调事件
// 参数一：函数地址
// 参数二：获取参数值的函数
// 返回值：如果为true则终止事件传递，如果为false则继续向后转发事件。
type TExtEventCallback func(fn interface{}, getVal func(idx int) uintptr) bool

// 外部扩展的事件回调，先不管重复注册的问题
var extEventCallback []TExtEventCallback

// 注册外部扩展回调事件
//
// Registering external extension callback events.
func RegisterExtEventCallback(callback TExtEventCallback) {
	extEventCallback = append(extEventCallback, callback)
}

// getParam 从指定索引和地址获取事件中的参数
// 不再使用Delphi导出的了，直接在这处理
func getParamOf(index int, ptr uintptr) uintptr {
	return *(*uintptr)(unsafe.Pointer(ptr + uintptr(index)*unsafe.Sizeof(ptr)))
}

// 回调过程
func eventCallbackProc(f uintptr, args uintptr, _ int) uintptr {
	v, ok := EventCallbackOf(f)
	if ok {

		// 获取值
		getVal := func(i int) uintptr {
			return getParamOf(i, args)
		}

		// 指针
		getPtr := func(i int) unsafe.Pointer {
			return unsafe.Pointer(getVal(i))
		}

		// 指针的值
		getPtrVal := func(i int) uintptr {
			return *(*uintptr)(getPtr(i))
		}

		setPtrVal := func(i int, val uintptr) {
			*(*uintptr)(getPtr(i)) = val
		}

		getBoolPtr := func(i int) *bool {
			return (*bool)(getPtr(i))
		}

		getI32Ptr := func(i int) *int32 {
			return (*int32)(getPtr(i))
		}

		getRectPtr := func(i int) *TRect {
			return (*TRect)(getPtr(i))
		}

		getPointPtr := func(i int) *TPoint {
			return (*TPoint)(getPtr(i))
		}

		// 调用外部注册的事件回调过程
		for n := 0; n < len(extEventCallback); n++ {
			// 外部返回True则不继续下去了
			if extEventCallback[n](v, getVal) {
				return 0
			}
		}

		switch v.(type) {
		// func(sender IObject)
		case TNotifyEvent:
			v.(TNotifyEvent)(
				AsForm(getVal(0)))

		// func(sender IObject, button TUDBtnType)
		case TUDClickEvent:
			v.(TUDClickEvent)(
				AsObject(getVal(0)),
				TUDBtnType(getVal(1)))

		// func(sender IObject, item *TListItem, change int32)
		case TLVChangeEvent:
			v.(TLVChangeEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				TItemChange(getVal(2)))

		// func(sender IObject, action *TCloseAction) // Action *uintptr
		case TCloseEvent:
			v.(TCloseEvent)(
				AsObject(getVal(0)),
				(*TCloseAction)(getPtr(1)))

		// func(sender IObject, canClose *bool) //CanClose *uintptr
		case TCloseQueryEvent:
			v.(TCloseQueryEvent)(
				AsObject(getVal(0)),
				getBoolPtr(1))

		// func(sender IObject, source *TMenuItem, rebuild bool)
		case TMenuChangeEvent:
			v.(TMenuChangeEvent)(
				AsObject(getVal(0)),
				AsMenuItem(getVal(1)),
				DBoolToGoBool(getVal(2)))

		// func(sender IObject, node *TreeNode)
		case TTVChangedEvent:
			v.(TTVChangedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)))

		// func(sender IObject, link string, linkType TSysLinkType) // TSysLinkType
		case TSysLinkEvent:
			v.(TSysLinkEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)),
				TSysLinkType(getVal(2)))

		// func(sender, e IObject)
		case TExceptionEvent:
			v.(TExceptionEvent)(
				AsObject(getVal(0)),
				AsException(getVal(1)))

		// func(sender IObject, key *Char, shift TShiftState)
		case TKeyEvent:

			v.(TKeyEvent)(
				AsObject(getVal(0)),
				(*Char)(getPtr(1)),
				TShiftState(getVal(2)))

		// func(sender IObject, key *Char)
		case TKeyPressEvent:

			v.(TKeyPressEvent)(
				AsObject(getVal(0)),
				(*Char)(getPtr(1)))

		// func(sender IObject, button TMouseButton, shift TShiftState, x, y int32)
		case TMouseEvent:
			v.(TMouseEvent)(
				AsObject(getVal(0)),
				TMouseButton(getVal(1)),
				TShiftState(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)))

		// func(sender IObject, shift TShiftState, x, y int32)
		case TMouseMoveEvent:
			v.(TMouseMoveEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

		// func(sender IObject, shift TShiftState, wheelDelta, x, y int32, handled *bool)
		case TMouseWheelEvent:
			v.(TMouseWheelEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				getBoolPtr(5))

			// func(control IWinControl, index int32, aRect TRect, state TOwnerDrawState)
		case TDrawItemEvent:
			v.(TDrawItemEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				*getRectPtr(2),
				TOwnerDrawState(getVal(3)))

			// func(sender IObject, aCanvas *TCanvas, aRect TRect, selected bool)
		case TMenuDrawItemEvent:
			v.(TMenuDrawItemEvent)(
				AsObject(getVal(0)),
				AsCanvas(getVal(1)),
				*getRectPtr(2),
				DBoolToGoBool(getVal(3)))

			// type TLVNotifyEvent func(sender IObject, item *TListItem)
		case TLVNotifyEvent:
			v.(TLVNotifyEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			// type TLVColumnClickEvent func(sender IObject, column *TListColumn)
		case TLVColumnClickEvent:
			v.(TLVColumnClickEvent)(
				AsObject(getVal(0)),
				AsListColumn(getVal(1)))

			// type TLVColumnRClickEvent func(sender IObject, column *TListColumn, point TPoint)
		case TLVColumnRClickEvent:
			v.(TLVColumnRClickEvent)(
				AsObject(getVal(0)),
				AsListColumn(getVal(1)),
				TPoint{X: int32(getVal(2)), Y: int32(getVal(3))})

			// type TLVSelectItemEvent func(sender IObject, item *TListItem, selected bool)
		case TLVSelectItemEvent:
			v.(TLVSelectItemEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				DBoolToGoBool(getVal(2)))

			// type TLVCheckedItemEvent func(sender IObject, item *TListItem)
		case TLVCheckedItemEvent:
			v.(TLVCheckedItemEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			// type TTabGetImageEvent func(sender IObject, tabIndex int32, imageIndex *int32)
		case TTabGetImageEvent:
			v.(TTabGetImageEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				getI32Ptr(2))

			// type TTVExpandedEvent func(sender IObject, node *TTreeNode)
		case TTVExpandedEvent:
			v.(TTVExpandedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)))

		//type TLVCompareEvent func(sender IObject, item1, item2 *TListItem, data int32, compare *int32)
		case TLVCompareEvent:
			v.(TLVCompareEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				AsListItem(getVal(2)),
				int32(getVal(3)),
				getI32Ptr(4))

		//type TTVCompareEvent func(sender IObject, node1, node2 *TTreeNode, data int32, compare *int32)
		case TTVCompareEvent:
			v.(TTVCompareEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				AsTreeNode(getVal(2)),
				int32(getVal(3)),
				getI32Ptr(4))

		//type TTVAdvancedCustomDrawEvent func(sender *TTreeView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTVAdvancedCustomDrawEvent:

			v.(TTVAdvancedCustomDrawEvent)(
				AsTreeView(getVal(0)),
				*getRectPtr(1),
				TCustomDrawStage(getVal(2)),
				getBoolPtr(3))

		//type TTVAdvancedCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawState, stage TCustomDrawStage, paintImages, defaultDraw *bool)
		case TTVAdvancedCustomDrawItemEvent:
			v.(TTVAdvancedCustomDrawItemEvent)(
				AsTreeView(getVal(0)),
				AsTreeNode(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				getBoolPtr(4),
				getBoolPtr(5))

			//---------------------------------------

		//type TLVAdvancedCustomDrawEvent func(sender *TListView, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawEvent:
			v.(TLVAdvancedCustomDrawEvent)(
				AsListView(getVal(0)),
				*getRectPtr(1),
				TCustomDrawStage(getVal(2)),
				getBoolPtr(3))

		//type TLVAdvancedCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawState, Stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawItemEvent:
			v.(TLVAdvancedCustomDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				TCustomDrawState(getVal(2)),
				TCustomDrawStage(getVal(3)),
				getBoolPtr(4))

		//type TLVAdvancedCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawState, stage TCustomDrawStage, defaultDraw *bool)
		case TLVAdvancedCustomDrawSubItemEvent:
			v.(TLVAdvancedCustomDrawSubItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				int32(getVal(2)),
				TCustomDrawState(getVal(3)),
				TCustomDrawStage(getVal(4)),
				getBoolPtr(5))

		//-----------------------------
		//type TTBAdvancedCustomDrawEvent func(sender *TToolBar, aRect TRect, stage TCustomDrawStage, defaultDraw *bool)
		case TTBAdvancedCustomDrawEvent:
			v.(TTBAdvancedCustomDrawEvent)(
				AsToolBar(getVal(0)),
				*getRectPtr(1),
				TCustomDrawStage(getVal(2)),
				getBoolPtr(3))

		// TDropFilesEvent
		case TDropFilesEvent:
			nLen := int(getVal(2))
			tempArr := make([]string, nLen)
			p := getVal(1)
			for i := 0; i < nLen; i++ {
				tempArr[i] = DGetStringArrOf(p, i)
			}
			v.(TDropFilesEvent)(
				AsObject(getVal(0)),
				tempArr)

			// TConstrainedResizeEvent
		case TConstrainedResizeEvent:
			v.(TConstrainedResizeEvent)(
				AsObject(getVal(0)),
				getI32Ptr(1),
				getI32Ptr(2),
				getI32Ptr(3),
				getI32Ptr(4))

			// func(command uint16, data THelpEventData, callhelp *bool) bool
		case THelpEvent:
			v.(THelpEvent)(
				uint16(getVal(0)),
				THelpEventData(getVal(1)),
				getBoolPtr(2),
				getBoolPtr(3))

			// func(msg *TWMKey, handled *bool)
		case TShortCutEvent:
			v.(TShortCutEvent)(
				(*TWMKey)(getPtr(0)),
				getBoolPtr(1))

			// func(sender IObject, mousePos TPoint, handled *bool)
		case TContextPopupEvent:
			v.(TContextPopupEvent)(
				AsObject(getVal(0)),
				*getPointPtr(1),
				getBoolPtr(2))

			// func(sender, source IObject, x, y int32, state TDragState, accept *bool)
		case TDragOverEvent:
			v.(TDragOverEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				getBoolPtr(5))

			//func(sender, source IObject, x, y int32)
		case TDragDropEvent:
			v.(TDragDropEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, dragObject *TDragObject)
		//case TStartDragEvent:
		//	obj := AsDragObject(getVal(1))
		//	v.(TStartDragEvent)(
		//		AsObject(getVal(0)),
		//		obj)
		//	if obj != nil {
		//		*(*uintptr)(unsafe.Pointer(getVal(1))) = obj.instance
		//	}

		//func(sender, target IObject, x, y int32)
		case TEndDragEvent:
			v.(TEndDragEvent)(
				AsObject(getVal(0)),
				AsObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			// func(sender IObject, source *TDragDockObject, x, y int32)
		case TDockDropEvent:
			v.(TDockDropEvent)(
				AsObject(getVal(0)),
				AsDragDockObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//func(sender IObject, source *TDragDockObject, x, y int32, state TDragState, accept *bool)
		case TDockOverEvent:
			v.(TDockOverEvent)(
				AsObject(getVal(0)),
				AsDragDockObject(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				TDragState(getVal(4)),
				getBoolPtr(5))

			//func(sender IObject, client *TControl, newTarget *TControl, allow *bool)
		case TUnDockEvent:
			v.(TUnDockEvent)(
				AsObject(getVal(0)),
				AsControl(getVal(1)),
				AsControl(getVal(2)),
				getBoolPtr(3))

			//func(sender IObject, dragObject **TDragDockObject)
		case TStartDockEvent:
			obj := AsDragDockObject(getPtrVal(1))
			v.(TStartDockEvent)(
				AsObject(getVal(0)),
				&obj)
			if obj != nil {
				setPtrVal(1, obj.instance)
			}

			//func(sender IObject, dockClient *TControl, influenceRect *TRect, mousePos TPoint, canDock *bool)
		case TGetSiteInfoEvent:
			v.(TGetSiteInfoEvent)(
				AsObject(getVal(0)),
				AsControl(getVal(1)),
				getRectPtr(2),
				*getPointPtr(3),
				getBoolPtr(4))

			//func(sender IObject, shift TShiftState, mousePos TPoint, handled *bool)
		case TMouseWheelUpDownEvent:
			v.(TMouseWheelUpDownEvent)(
				AsObject(getVal(0)),
				TShiftState(getVal(1)),
				*getPointPtr(2),
				getBoolPtr(3))

		// ---- grid
		//type TGridOperationEvent func(sender IObject, isColumn bool, sIndex, tIndex int32)
		case TGridOperationEvent:
			v.(TGridOperationEvent)(
				AsObject(getVal(0)),
				DBoolToGoBool(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

			//type TDrawCellEvent func(sender IObject, aCol, aRow int32, aRect TRect, state TGridDrawState)
		case TDrawCellEvent:
			v.(TDrawCellEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				*getRectPtr(3),
				TGridDrawState(getVal(4)))

			//type TFixedCellClickEvent func(sender IObject, aCol, aRow int32)
		case TFixedCellClickEvent:
			v.(TFixedCellClickEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TGetEditEvent func(sender IObject, aCol, aRow int32, value *string)
		case TGetEditEvent:
			str := DStrToGoStr(getPtrVal(3))
			v.(TGetEditEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				&str)
			setPtrVal(3, GoStrToDStr(str))

			//type TSelectCellEvent func(sender IObject, aCol, aRow int32, canSelect *bool)
		case TSelectCellEvent:
			v.(TSelectCellEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				getBoolPtr(3))

			//type TSetEditEvent func(sender IObject, aCol, aRow int32, value string)
		case TSetEditEvent:
			v.(TSetEditEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				DStrToGoStr(getVal(3)))

			// ---- headercontrol
			//type TDrawSectionEvent func(headerControl *THeaderControl, section *THeaderSection, aRect TRect, pressed bool)
		case TDrawSectionEvent:
			v.(TDrawSectionEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)),
				*getRectPtr(2),
				getVal(3) != 0)

			//type TSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TSectionNotifyEvent:
			v.(TSectionNotifyEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)))

			//type TSectionTrackEvent func(headerControl *THeaderControl, section *THeaderSection, width int32, state TSectionTrackState)
		case TSectionTrackEvent:
			v.(TSectionTrackEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)),
				int32(getVal(2)),
				TSectionTrackState(getVal(3)))

			//type TSectionDragEvent func(sender IObject, fromSection, toSection *THeaderSection, allowDrag *bool)
		case TSectionDragEvent:
			v.(TSectionDragEvent)(
				AsObject(getVal(0)),
				AsHeaderSection(getVal(1)),
				AsHeaderSection(getVal(2)),
				getBoolPtr(3))

			//type TCustomSectionNotifyEvent func(headerControl *THeaderControl, section *THeaderSection)
		case TCustomSectionNotifyEvent:
			v.(TCustomSectionNotifyEvent)(
				AsHeaderControl(getVal(0)),
				AsHeaderSection(getVal(1)))

		//type TMouseActivateEvent func(sender IObject, button TMouseButton, shift TShiftState, x, y int32, hitTest int32, mouseActivate *TMouseActivate)
		case TMouseActivateEvent:
			v.(TMouseActivateEvent)(
				AsObject(getVal(0)),
				TMouseButton(getVal(1)),
				TShiftState(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				int32(getVal(5)),
				(*TMouseActivate)(getPtr(6)))

			//type TLBGetDataEvent func(control *TWinControl, index int32, data *string)
		case TLBGetDataEvent:
			str := DStrToGoStr(getPtrVal(2))
			v.(TLBGetDataEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				&str)
			setPtrVal(2, GoStrToDStr(str))

			//type TLBGetDataObjectEvent func(control *TWinControl, index int32, dataObject IObject)
		case TLBGetDataObjectEvent:
			v.(TLBGetDataObjectEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				AsObject(getVal(2))) // 这个参数要改，先这样

			//type TLBFindDataEvent func(control *TWinControl, findString string) int32
		case TLBFindDataEvent:
			result := v.(TLBFindDataEvent)(
				AsWinControl(getVal(0)),
				DStrToGoStr(getVal(1)))
			*getI32Ptr(2) = result

			//type TMeasureItemEvent func(control *TWinControl, index int32, height *int32)
		case TMeasureItemEvent:
			v.(TMeasureItemEvent)(
				AsWinControl(getVal(0)),
				int32(getVal(1)),
				getI32Ptr(2))

			//type TLVChangingEvent func(sender IObject, item *TListItem, change TItemChange, allowChange *bool)
		case TLVChangingEvent:
			v.(TLVChangingEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				TItemChange(getVal(2)),
				getBoolPtr(3))

			//type TLVDataEvent func(sender IObject, item *TListItem)
		case TLVDataEvent:
			v.(TLVDataEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			//type TLVDataFindEvent func(sender IObject, find TItemFind, findString string, findPosition TPoint, findData TCustomData, startIndex int32,
			//	direction TSearchDirection, warp bool, index *int32)
		case TLVDataFindEvent:
			v.(TLVDataFindEvent)(
				AsObject(getVal(0)),
				TItemFind(getVal(1)),
				DStrToGoStr(getVal(2)),
				*getPointPtr(3),
				TCustomData(getVal(4)),
				int32(getVal(5)),
				TSearchDirection(getVal(6)),
				DBoolToGoBool(getVal(7)),
				getI32Ptr(8))

			//type TLVDeletedEvent func(sender IObject, item *TListItem)
		case TLVDeletedEvent:
			v.(TLVDeletedEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)))

			//type TLVEditingEvent func(sender IObject, item *TListItem, allowEdit *bool)
		case TLVEditingEvent:
			v.(TLVEditingEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				getBoolPtr(2))

			//type TLVEditedEvent func(sender IObject, item *TListItem, s *string)
		case TLVEditedEvent:
			str := DStrToGoStr(getPtrVal(2))
			v.(TLVEditedEvent)(
				AsObject(getVal(0)),
				AsListItem(getVal(1)),
				&str)
			setPtrVal(2, GoStrToDStr(str))

			//type TMenuMeasureItemEvent func(sender IObject, aCanvas *TCanvas, width, height *int32)
		case TMenuMeasureItemEvent:
			v.(TMenuMeasureItemEvent)(
				AsObject(getVal(0)),
				AsCanvas(getVal(1)),
				getI32Ptr(2),
				getI32Ptr(3))

			//type TTabChangingEvent func(sender IObject, allowChange *bool)
		case TTabChangingEvent:
			v.(TTabChangingEvent)(
				AsObject(getVal(0)),
				getBoolPtr(1))

			//type TTVChangingEvent func(sender IObject, node *TTreeNode, allowChange *bool)
		case TTVChangingEvent:
			v.(TTVChangingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				getBoolPtr(2))

			//type TTVCollapsingEvent func(sender IObject, node *TTreeNode, allowCollapse *bool)
		case TTVCollapsingEvent:
			v.(TTVCollapsingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				getBoolPtr(2))

			//type TTVEditedEvent func(sender IObject, node *TTreeNode, s *string)
		case TTVEditedEvent:
			str := DStrToGoStr(getPtrVal(2))
			v.(TTVEditedEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				&str)
			setPtrVal(2, GoStrToDStr(str))

			//type TTVEditingEvent func(sender IObject, node *TTreeNode, allowEdit *bool)
		case TTVEditingEvent:
			v.(TTVEditingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				getBoolPtr(2))

			//type TTVExpandingEvent func(sender IObject, node *TTreeNode, allowExpansion *bool)
		case TTVExpandingEvent:
			v.(TTVExpandingEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				getBoolPtr(2))

			//type TTVHintEvent func(sender IObject, node *TTreeNode, hint *string)
		case TTVHintEvent:
			str := DStrToGoStr(getPtrVal(2))
			v.(TTVHintEvent)(
				AsObject(getVal(0)),
				AsTreeNode(getVal(1)),
				&str)
			setPtrVal(2, GoStrToDStr(str))

			//type TUDChangingEvent func(sender IObject, allowChange *bool)
		case TUDChangingEvent:
			v.(TUDChangingEvent)(
				AsObject(getVal(0)),
				getBoolPtr(1))

			//type TCreatingListErrorEvent func(sender IObject, winErrorCode uint32, errorDescription string, handled *bool)
		case TCreatingListErrorEvent:
			v.(TCreatingListErrorEvent)(
				AsObject(getVal(0)),
				uint32(getVal(1)),
				DStrToGoStr(getVal(2)),
				getBoolPtr(3))

		//--

		//type TLVCustomDrawEvent func(sender *TListView, aRect TRect, defaultDraw *bool)
		case TLVCustomDrawEvent:
			v.(TLVCustomDrawEvent)(
				AsListView(getVal(0)),
				*getRectPtr(1),
				getBoolPtr(2))

			//type TLVCustomDrawItemEvent func(sender *TListView, item *TListItem, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawItemEvent:
			v.(TLVCustomDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				TCustomDrawState(getVal(2)),
				getBoolPtr(3))

			//type TLVCustomDrawSubItemEvent func(sender *TListView, item *TListItem, subItem int32, state TCustomDrawStage, defaultDraw *bool)
		case TLVCustomDrawSubItemEvent:
			v.(TLVCustomDrawSubItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				int32(getVal(2)),
				TCustomDrawState(getVal(3)),
				getBoolPtr(4))

			//type TLVDrawItemEvent func(sender *TListView, item *TListItem, rect TRect, state TOwnerDrawState)
		case TLVDrawItemEvent:
			v.(TLVDrawItemEvent)(
				AsListView(getVal(0)),
				AsListItem(getVal(1)),
				*getRectPtr(2),
				TOwnerDrawState(getVal(3)))

		//type TLVDataHintEvent = func(sender IObject, startIndex, endIndex int32)
		case TLVDataHintEvent:
			v.(TLVDataHintEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

			//type TTVCustomDrawEvent func(sender *TTreeView, aRect TRect, defaultDraw *bool)
		case TTVCustomDrawEvent:
			v.(TTVCustomDrawEvent)(
				AsTreeView(getVal(0)),
				*getRectPtr(1),
				getBoolPtr(2))

			//type TTVCustomDrawItemEvent func(sender *TTreeView, node *TTreeNode, state TCustomDrawStage, defaultDraw *bool)
		case TTVCustomDrawItemEvent:
			v.(TTVCustomDrawItemEvent)(
				AsTreeView(getVal(0)),
				AsTreeNode(getVal(1)),
				TCustomDrawState(getVal(2)),
				getBoolPtr(3))

			// type TWebTitleChangeEvent func(sender IObject, text string)
		case TWebTitleChangeEvent:
			v.(TWebTitleChangeEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)))

		case TWebJSExternalEvent:
			str := DStrToGoStr(getPtrVal(3))
			v.(TWebJSExternalEvent)(
				AsObject(getVal(0)),
				DStrToGoStr(getVal(1)),
				DStrToGoStr(getVal(2)),
				&str)
			setPtrVal(3, GoStrToDStr(str))

			//type TTaskDlgClickEvent func(sender IObject, modalResult TModalResult, canClose *bool)
		case TTaskDlgClickEvent:
			v.(TTaskDlgClickEvent)(
				AsObject(getVal(0)),
				TModalResult(getVal(1)),
				getBoolPtr(2))

			//type TTaskDlgTimerEvent func(sender IObject, tickCount uint32, reset *bool)
		case TTaskDlgTimerEvent:
			v.(TTaskDlgTimerEvent)(
				AsObject(getVal(0)),
				uint32(getVal(1)),
				getBoolPtr(2))

		// type TAlignPositionEvent func(sender *TWinControl, control *TControl, newLeft, newTop, newWidth, newHeight *int32, alignRect *TRect, alignInfo TAlignInfo)
		case TAlignPositionEvent:
			v.(TAlignPositionEvent)(
				AsWinControl(getVal(0)),
				AsControl(getVal(1)),
				getI32Ptr(2),
				getI32Ptr(3),
				getI32Ptr(4),
				getI32Ptr(5),
				getRectPtr(6),
				*(*TAlignInfo)(getPtr(7)))

		// type TCheckGroupClicked func(sender IObject, index int32)
		case TCheckGroupClicked:
			v.(TCheckGroupClicked)(
				AsObject(getVal(0)),
				int32(getVal(1)))

		//type TOnSelectEvent func(sender IObject, aCol, aRow int32)
		case TOnSelectEvent:
			v.(TOnSelectEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)))

		//type TToggledCheckboxEvent func(sender IObject, aCol, aRow int32, aState TCheckBoxState)
		case TToggledCheckboxEvent:
			v.(TToggledCheckboxEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				TCheckBoxState(getVal(3)))

		//type TOnCompareCells func(sender IObject, ACol, ARow, BCol, BRow int32, result *int32)
		case TOnCompareCells:
			v.(TOnCompareCells)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)),
				int32(getVal(4)),
				getI32Ptr(5))

		//type TGetCellHintEvent func(sender IObject, ACol, ARow int32, hintText *string)
		case TGetCellHintEvent:
			str := DStrToGoStr(getPtrVal(3))
			v.(TGetCellHintEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				&str)
			setPtrVal(3, GoStrToDStr(str))

		//type TGetCheckboxStateEvent func(sender IObject, ACol, ARow int32, value *TCheckBoxState)
		case TGetCheckboxStateEvent:
			v.(TGetCheckboxStateEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				(*TCheckBoxState)(getPtr(3)))

		//type TSetCheckboxStateEvent func(sender IObject, ACol, ARow int32, Value TCheckBoxState)
		case TSetCheckboxStateEvent:
			v.(TSetCheckboxStateEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				TCheckBoxState(getVal(3)))

		//type THdrEvent func(sender IObject, isColumn bool, index int32)
		case THdrEvent:
			v.(THdrEvent)(
				AsObject(getVal(0)),
				DBoolToGoBool(getVal(1)),
				int32(getVal(2)))

		//type THeaderSizingEvent func(sender IObject, isColumn bool, aIndex, aSize int32)
		case THeaderSizingEvent:
			v.(THeaderSizingEvent)(
				AsObject(getVal(0)),
				DBoolToGoBool(getVal(1)),
				int32(getVal(2)),
				int32(getVal(3)))

		//type TSelectEditorEvent func(sender IObject, aCol, aRow int32, editor **TWinControl)
		case TSelectEditorEvent:
			obj := AsWinControl(getPtrVal(3))
			v.(TSelectEditorEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				&obj)
			if obj != nil {
				setPtrVal(3, obj.instance)
			}

		//type TUserCheckBoxBitmapEvent func(sender IObject, aCol, aRow int32, CheckedState TCheckBoxState, aBitmap **TBitmap)
		case TUserCheckBoxBitmapEvent:
			obj := AsBitmap(getPtrVal(4))
			v.(TUserCheckBoxBitmapEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				TCheckBoxState(getVal(1)),
				&obj)
			if obj != nil {
				setPtrVal(4, obj.instance)
			}

			//type TValidateEntryEvent func(sender IObject, aCol, aRow int32, oldValue string, newValue *string)
		case TValidateEntryEvent:
			str := DStrToGoStr(getPtrVal(4))
			v.(TValidateEntryEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				DStrToGoStr(getVal(3)),
				&str)
			setPtrVal(4, GoStrToDStr(str))
			//type TOnPrepareCanvasEvent = func(sender IObject, aCol, aRow int32, aState TGridDrawState)
		case TOnPrepareCanvasEvent:
			v.(TOnPrepareCanvasEvent)(
				AsObject(getVal(0)),
				int32(getVal(1)),
				int32(getVal(2)),
				TGridDrawState(getVal(3)))
		default:
		}
	}
	return 0
}
