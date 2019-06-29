package msrdp

import (
	"unsafe"

	. "github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

//--------------------------- TMsRdpClient9NotSafeForScripting ---------------------------

func MsRdpClient9NotSafeForScripting_Create(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_Create.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_Free(obj uintptr) {
	msRdpClient9NotSafeForScripting_Free.Call(obj)
}

func MsRdpClient9NotSafeForScripting_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32) {
	msRdpClient9NotSafeForScripting_SetBounds.Call(obj, uintptr(ALeft), uintptr(ATop), uintptr(AWidth), uintptr(AHeight))
}

func MsRdpClient9NotSafeForScripting_CanFocus(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_CanFocus.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_ContainsControl(obj uintptr, Control uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_ContainsControl.Call(obj, Control)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_ControlAtPos.Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled), GoBoolToDBool(AllowWinControls), GoBoolToDBool(AllLevels))
	return ret
}

func MsRdpClient9NotSafeForScripting_DisableAlign(obj uintptr) {
	msRdpClient9NotSafeForScripting_DisableAlign.Call(obj)
}

func MsRdpClient9NotSafeForScripting_EnableAlign(obj uintptr) {
	msRdpClient9NotSafeForScripting_EnableAlign.Call(obj)
}

func MsRdpClient9NotSafeForScripting_FindChildControl(obj uintptr, ControlName string) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_FindChildControl.Call(obj, GoStrToDStr(ControlName))
	return ret
}

func MsRdpClient9NotSafeForScripting_FlipChildren(obj uintptr, AllLevels bool) {
	msRdpClient9NotSafeForScripting_FlipChildren.Call(obj, GoBoolToDBool(AllLevels))
}

func MsRdpClient9NotSafeForScripting_Focused(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_Focused.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_HandleAllocated(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_HandleAllocated.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_InsertControl(obj uintptr, AControl uintptr) {
	msRdpClient9NotSafeForScripting_InsertControl.Call(obj, AControl)
}

func MsRdpClient9NotSafeForScripting_Invalidate(obj uintptr) {
	msRdpClient9NotSafeForScripting_Invalidate.Call(obj)
}

func MsRdpClient9NotSafeForScripting_PaintTo(obj uintptr, DC HDC, X int32, Y int32) {
	msRdpClient9NotSafeForScripting_PaintTo.Call(obj, uintptr(DC), uintptr(X), uintptr(Y))
}

func MsRdpClient9NotSafeForScripting_RemoveControl(obj uintptr, AControl uintptr) {
	msRdpClient9NotSafeForScripting_RemoveControl.Call(obj, AControl)
}

func MsRdpClient9NotSafeForScripting_Realign(obj uintptr) {
	msRdpClient9NotSafeForScripting_Realign.Call(obj)
}

func MsRdpClient9NotSafeForScripting_Repaint(obj uintptr) {
	msRdpClient9NotSafeForScripting_Repaint.Call(obj)
}

func MsRdpClient9NotSafeForScripting_ScaleBy(obj uintptr, M int32, D int32) {
	msRdpClient9NotSafeForScripting_ScaleBy.Call(obj, uintptr(M), uintptr(D))
}

func MsRdpClient9NotSafeForScripting_ScrollBy(obj uintptr, DeltaX int32, DeltaY int32) {
	msRdpClient9NotSafeForScripting_ScrollBy.Call(obj, uintptr(DeltaX), uintptr(DeltaY))
}

func MsRdpClient9NotSafeForScripting_SetFocus(obj uintptr) {
	msRdpClient9NotSafeForScripting_SetFocus.Call(obj)
}

func MsRdpClient9NotSafeForScripting_Update(obj uintptr) {
	msRdpClient9NotSafeForScripting_Update.Call(obj)
}

func MsRdpClient9NotSafeForScripting_UpdateControlState(obj uintptr) {
	msRdpClient9NotSafeForScripting_UpdateControlState.Call(obj)
}

func MsRdpClient9NotSafeForScripting_BringToFront(obj uintptr) {
	msRdpClient9NotSafeForScripting_BringToFront.Call(obj)
}

func MsRdpClient9NotSafeForScripting_ClientToScreen(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	msRdpClient9NotSafeForScripting_ClientToScreen.Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_ClientToParent(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	msRdpClient9NotSafeForScripting_ClientToParent.Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_Dragging(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_Dragging.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_HasParent(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_HasParent.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_Hide(obj uintptr) {
	msRdpClient9NotSafeForScripting_Hide.Call(obj)
}

func MsRdpClient9NotSafeForScripting_Perform(obj uintptr, Msg uint32, WParam uintptr, LParam int) int {
	ret, _, _ := msRdpClient9NotSafeForScripting_Perform.Call(obj, uintptr(Msg), WParam, uintptr(LParam))
	return int(ret)
}

func MsRdpClient9NotSafeForScripting_Refresh(obj uintptr) {
	msRdpClient9NotSafeForScripting_Refresh.Call(obj)
}

func MsRdpClient9NotSafeForScripting_ScreenToClient(obj uintptr, Point TPoint) TPoint {
	var ret TPoint
	msRdpClient9NotSafeForScripting_ScreenToClient.Call(obj, uintptr(unsafe.Pointer(&Point)), uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_ParentToClient(obj uintptr, Point TPoint, AParent uintptr) TPoint {
	var ret TPoint
	msRdpClient9NotSafeForScripting_ParentToClient.Call(obj, uintptr(unsafe.Pointer(&Point)), AParent, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_SendToBack(obj uintptr) {
	msRdpClient9NotSafeForScripting_SendToBack.Call(obj)
}

func MsRdpClient9NotSafeForScripting_Show(obj uintptr) {
	msRdpClient9NotSafeForScripting_Show.Call(obj)
}

func MsRdpClient9NotSafeForScripting_GetTextBuf(obj uintptr, Buffer string, BufSize int32) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTextBuf.Call(obj, GoStrToDStr(Buffer), uintptr(BufSize))
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetTextLen(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTextLen.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetTextBuf(obj uintptr, Buffer string) {
	msRdpClient9NotSafeForScripting_SetTextBuf.Call(obj, GoStrToDStr(Buffer))
}

func MsRdpClient9NotSafeForScripting_FindComponent(obj uintptr, AName string) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_FindComponent.Call(obj, GoStrToDStr(AName))
	return ret
}

func MsRdpClient9NotSafeForScripting_GetNamePath(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetNamePath.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_Assign(obj uintptr, Source uintptr) {
	msRdpClient9NotSafeForScripting_Assign.Call(obj, Source)
}

func MsRdpClient9NotSafeForScripting_DisposeOf(obj uintptr) {
	msRdpClient9NotSafeForScripting_DisposeOf.Call(obj)
}

func MsRdpClient9NotSafeForScripting_ClassType(obj uintptr) TClass {
	ret, _, _ := msRdpClient9NotSafeForScripting_ClassType.Call(obj)
	return TClass(ret)
}

func MsRdpClient9NotSafeForScripting_ClassName(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_ClassName.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_InstanceSize(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_InstanceSize.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_InheritsFrom(obj uintptr, AClass TClass) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_InheritsFrom.Call(obj, uintptr(AClass))
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_Equals(obj uintptr, Obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_Equals.Call(obj, Obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_GetHashCode(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetHashCode.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_ToString(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_ToString.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_GetControlInterface(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetControlInterface.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_GetDefaultInterface(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDefaultInterface.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_GetVersion(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetVersion.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_GetAdvancedSettings9(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAdvancedSettings9.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_GetAnchors(obj uintptr) TAnchors {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAnchors.Call(obj)
	return TAnchors(ret)
}

func MsRdpClient9NotSafeForScripting_SetAnchors(obj uintptr, value TAnchors) {
	msRdpClient9NotSafeForScripting_SetAnchors.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetTabStop(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTabStop.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetTabStop(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetTabStop.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetAlign(obj uintptr) TAlign {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAlign.Call(obj)
	return TAlign(ret)
}

func MsRdpClient9NotSafeForScripting_SetAlign(obj uintptr, value TAlign) {
	msRdpClient9NotSafeForScripting_SetAlign.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetDragCursor(obj uintptr) TCursor {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDragCursor.Call(obj)
	return TCursor(ret)
}

func MsRdpClient9NotSafeForScripting_SetDragCursor(obj uintptr, value TCursor) {
	msRdpClient9NotSafeForScripting_SetDragCursor.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetDragMode(obj uintptr) TDragMode {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDragMode.Call(obj)
	return TDragMode(ret)
}

func MsRdpClient9NotSafeForScripting_SetDragMode(obj uintptr, value TDragMode) {
	msRdpClient9NotSafeForScripting_SetDragMode.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetParentShowHint(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetParentShowHint.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetParentShowHint(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetParentShowHint.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetPopupMenu(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetPopupMenu.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_SetPopupMenu(obj uintptr, value uintptr) {
	msRdpClient9NotSafeForScripting_SetPopupMenu.Call(obj, value)
}

func MsRdpClient9NotSafeForScripting_GetShowHint(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetShowHint.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetShowHint(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetShowHint.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetTabOrder(obj uintptr) TTabOrder {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTabOrder.Call(obj)
	return TTabOrder(ret)
}

func MsRdpClient9NotSafeForScripting_SetTabOrder(obj uintptr, value TTabOrder) {
	msRdpClient9NotSafeForScripting_SetTabOrder.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetVisible(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetVisible.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetVisible(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetVisible.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_SetOnDragDrop(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnDragDrop.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_SetOnDragOver(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnDragOver.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_SetOnEndDrag(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnEndDrag.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_SetOnEnter(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnEnter.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_SetOnExit(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnExit.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_GetServer(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetServer.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetServer(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetServer.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetUserName(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetUserName.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetUserName(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetUserName.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetDisconnectedText(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDisconnectedText.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetDisconnectedText(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetDisconnectedText.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetConnectingText(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetConnectingText.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetConnectingText(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetConnectingText.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetDesktopWidth(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDesktopWidth.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetDesktopWidth(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetDesktopWidth.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetDesktopHeight(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDesktopHeight.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetDesktopHeight(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetDesktopHeight.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetStartConnected(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetStartConnected.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetStartConnected(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetStartConnected.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetColorDepth(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetColorDepth.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetColorDepth(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetColorDepth.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetFullScreen(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetFullScreen.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetFullScreen(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetFullScreen.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetDockClientCount(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDockClientCount.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetDockSite(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDockSite.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetDockSite(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetDockSite.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetDoubleBuffered(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDoubleBuffered.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetDoubleBuffered(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetAlignDisabled(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAlignDisabled.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_GetMouseInClient(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetMouseInClient.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_GetVisibleDockClientCount(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetVisibleDockClientCount.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetBrush(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetBrush.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_GetControlCount(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetControlCount.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetHandle(obj uintptr) HWND {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetHandle.Call(obj)
	return HWND(ret)
}

func MsRdpClient9NotSafeForScripting_GetParentDoubleBuffered(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetParentDoubleBuffered.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetParentDoubleBuffered(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetParentDoubleBuffered.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetParentWindow(obj uintptr) HWND {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetParentWindow.Call(obj)
	return HWND(ret)
}

func MsRdpClient9NotSafeForScripting_SetParentWindow(obj uintptr, value HWND) {
	msRdpClient9NotSafeForScripting_SetParentWindow.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetUseDockManager(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetUseDockManager.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetUseDockManager(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetUseDockManager.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetEnabled(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetEnabled.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetEnabled(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetEnabled.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetAction(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAction.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_SetAction(obj uintptr, value uintptr) {
	msRdpClient9NotSafeForScripting_SetAction.Call(obj, value)
}

func MsRdpClient9NotSafeForScripting_GetBiDiMode(obj uintptr) TBiDiMode {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetBiDiMode.Call(obj)
	return TBiDiMode(ret)
}

func MsRdpClient9NotSafeForScripting_SetBiDiMode(obj uintptr, value TBiDiMode) {
	msRdpClient9NotSafeForScripting_SetBiDiMode.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetBoundsRect(obj uintptr) TRect {
	var ret TRect
	msRdpClient9NotSafeForScripting_GetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_SetBoundsRect(obj uintptr, value TRect) {
	msRdpClient9NotSafeForScripting_SetBoundsRect.Call(obj, uintptr(unsafe.Pointer(&value)))
}

func MsRdpClient9NotSafeForScripting_GetClientHeight(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetClientHeight.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetClientHeight(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetClientHeight.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetClientOrigin(obj uintptr) TPoint {
	var ret TPoint
	msRdpClient9NotSafeForScripting_GetClientOrigin.Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_GetClientRect(obj uintptr) TRect {
	var ret TRect
	msRdpClient9NotSafeForScripting_GetClientRect.Call(obj, uintptr(unsafe.Pointer(&ret)))
	return ret
}

func MsRdpClient9NotSafeForScripting_GetClientWidth(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetClientWidth.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetClientWidth(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetClientWidth.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetControlState(obj uintptr) TControlState {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetControlState.Call(obj)
	return TControlState(ret)
}

func MsRdpClient9NotSafeForScripting_SetControlState(obj uintptr, value TControlState) {
	msRdpClient9NotSafeForScripting_SetControlState.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetControlStyle(obj uintptr) TControlStyle {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetControlStyle.Call(obj)
	return TControlStyle(ret)
}

func MsRdpClient9NotSafeForScripting_SetControlStyle(obj uintptr, value TControlStyle) {
	msRdpClient9NotSafeForScripting_SetControlStyle.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetExplicitLeft(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetExplicitLeft.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetExplicitTop(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetExplicitTop.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetExplicitWidth(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetExplicitWidth.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetExplicitHeight(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetExplicitHeight.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetFloating(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetFloating.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_GetParent(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetParent.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_SetParent(obj uintptr, value uintptr) {
	msRdpClient9NotSafeForScripting_SetParent.Call(obj, value)
}

func MsRdpClient9NotSafeForScripting_GetStyleElements(obj uintptr) TStyleElements {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetStyleElements.Call(obj)
	return TStyleElements(ret)
}

func MsRdpClient9NotSafeForScripting_SetStyleElements(obj uintptr, value TStyleElements) {
	msRdpClient9NotSafeForScripting_SetStyleElements.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_SetOnGesture(obj uintptr, fn interface{}) {
	msRdpClient9NotSafeForScripting_SetOnGesture.Call(obj, addEventToMap(fn))
}

func MsRdpClient9NotSafeForScripting_GetAlignWithMargins(obj uintptr) bool {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetAlignWithMargins.Call(obj)
	return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_SetAlignWithMargins(obj uintptr, value bool) {
	msRdpClient9NotSafeForScripting_SetAlignWithMargins.Call(obj, GoBoolToDBool(value))
}

func MsRdpClient9NotSafeForScripting_GetLeft(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetLeft.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetLeft(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetLeft.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetTop(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTop.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetTop(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetTop.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetWidth(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetWidth.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetWidth(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetWidth.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetHeight(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetHeight.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetHeight(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetHeight.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetCursor(obj uintptr) TCursor {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetCursor.Call(obj)
	return TCursor(ret)
}

func MsRdpClient9NotSafeForScripting_SetCursor(obj uintptr, value TCursor) {
	msRdpClient9NotSafeForScripting_SetCursor.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetHint(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetHint.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetHint(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetHint.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetMargins(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetMargins.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_SetMargins(obj uintptr, value uintptr) {
	msRdpClient9NotSafeForScripting_SetMargins.Call(obj, value)
}

func MsRdpClient9NotSafeForScripting_GetCustomHint(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetCustomHint.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_SetCustomHint(obj uintptr, value uintptr) {
	msRdpClient9NotSafeForScripting_SetCustomHint.Call(obj, value)
}

func MsRdpClient9NotSafeForScripting_GetComponentCount(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetComponentCount.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_GetComponentIndex(obj uintptr) int32 {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetComponentIndex.Call(obj)
	return int32(ret)
}

func MsRdpClient9NotSafeForScripting_SetComponentIndex(obj uintptr, value int32) {
	msRdpClient9NotSafeForScripting_SetComponentIndex.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetOwner(obj uintptr) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetOwner.Call(obj)
	return ret
}

func MsRdpClient9NotSafeForScripting_GetName(obj uintptr) string {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetName.Call(obj)
	return DStrToGoStr(ret)
}

func MsRdpClient9NotSafeForScripting_SetName(obj uintptr, value string) {
	msRdpClient9NotSafeForScripting_SetName.Call(obj, GoStrToDStr(value))
}

func MsRdpClient9NotSafeForScripting_GetTag(obj uintptr) int {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetTag.Call(obj)
	return int(ret)
}

func MsRdpClient9NotSafeForScripting_SetTag(obj uintptr, value int) {
	msRdpClient9NotSafeForScripting_SetTag.Call(obj, uintptr(value))
}

func MsRdpClient9NotSafeForScripting_GetDockClients(obj uintptr, Index int32) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetDockClients.Call(obj, uintptr(Index))
	return ret
}

func MsRdpClient9NotSafeForScripting_GetControls(obj uintptr, Index int32) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetControls.Call(obj, uintptr(Index))
	return ret
}

func MsRdpClient9NotSafeForScripting_GetComponents(obj uintptr, AIndex int32) uintptr {
	ret, _, _ := msRdpClient9NotSafeForScripting_GetComponents.Call(obj, uintptr(AIndex))
	return ret
}

func MsRdpClient9NotSafeForScripting_StaticClassType() TClass {
	r, _, _ := msRdpClient9NotSafeForScripting_StaticClassType.Call()
	return TClass(r)
}
