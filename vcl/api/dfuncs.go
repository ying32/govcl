//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"sync"
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

var (
	// 防止GC的表
	events              = sync.Map{}
	eventNextId uintptr = 0 // 初始id
	//events = list.New()

	// ThreadSync
	threadSync   sync.Mutex
	threadSyncFn func()

	// sync
	refs sync.Mutex
)

func GoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

// Deprecated: use GoBool.
func DBoolToGoBool(val uintptr) bool {
	return GoBool(val)
}

func PascalBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

// Deprecated: use PascalBool.
func GoBoolToDBool(val bool) uintptr {
	return PascalBool(val)
}

// typedef struct { void *type; void *value; } GoInterface;
type interfacePtr struct {
	tpy uintptr
	val uintptr
}

func interfaceNotNil(v interface{}) bool {
	ptr := (*interfacePtr)(unsafe.Pointer(&v))
	return ptr != nil && ptr.tpy != 0 && ptr.val != 0
}

//func PtrToElementPtr(v unsafe.Pointer) *list.Element {
//	if v == nil {
//		return nil
//	}
//	return (*list.Element)(v)
//}
//
//func PtrToElementValue(v unsafe.Pointer) interface{} {
//	element := PtrToElementPtr(v)
//	if element != nil {
//		return element.Value
//	}
//	return nil
//}
//
//func RemoveEventElement(v *list.Element) bool {
//	if v != nil {
//		refs.Lock()
//		defer refs.Unlock()
//		events.Remove(v)
//		return true
//	}
//	return false
//}
//
//func MakeEventDataPtr(fn interface{}) uintptr {
//
//	refs.Lock()
//	defer refs.Unlock()
//	if interfaceNotNil(fn) {
//		return uintptr(unsafe.Pointer(events.PushBack(fn)))
//	}
//	return 0
//}

func PtrToElementValue(v unsafe.Pointer) interface{} {
	if fn, ok := events.Load(uintptr(v)); ok {
		return fn
	}
	return nil
}

func RemoveEventElement(v unsafe.Pointer) bool {
	refs.Lock()
	defer refs.Unlock()
	if v != nil {
		events.Delete(uintptr(v))
		return true
	}
	return false
}

// MakeEventDataPtr
//  本想直接返回一个Go的指针，岂料Linux下Go强制不允许这样操作，一定要做就必须先设置环境变量`GODEBUG=cgocheck=0`关闭，但用户肯定不可能每次弄个啊
//  关于不允许传入Go指针到C指针的原因： https://chai2010.cn/advanced-go-programming-book/ch2-cgo/ch2-07-memory.html
func MakeEventDataPtr(fn interface{}) uintptr {
	refs.Lock()
	defer refs.Unlock()
	if interfaceNotNil(fn) {
		eventNextId++
		events.Store(eventNextId, fn)
		return eventNextId
	}
	return 0
}

func ThreadSyncCallbackFn() func() {
	return threadSyncFn
}

func SetEventCallback(ptr uintptr) {
	setEventCallback.Call(ptr)
}

func SetMessageCallback(ptr uintptr) {
	setMessageCallback.Call(ptr)
}

func SetThreadSyncCallback(ptr uintptr) {
	setThreadSyncCallback.Call(ptr)
}

func SetRequestCallCreateParamsCallback(ptr uintptr) {
	setRequestCallCreateParamsCallback.Call(ptr)
}

func SetRemoveEventCallback(ptr uintptr) {
	setRemoveEventCallback.Call(ptr)
}

func DGetStringArrOf(p uintptr, index int) string {
	r, _, _ := dGetStringArrOf.Call(p, uintptr(index))
	return GoStr(r)
}

func DStrLen(p uintptr) int {
	ret, _, _ := dStrLen.Call(p)
	return int(ret)
}

func DMove(src, dest uintptr, llen int) {
	dMove.Call(src, dest, uintptr(llen))
}

func DShowMessage(s string) {
	dShowMessage.Call(PascalStr(s))
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	ret, _, _ := dMessageDlg.Call(PascalStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx))
	return int32(ret)
}

func DTextToShortCut(val string) TShortCut {
	ret, _, _ := dTextToShortCut.Call(PascalStr(val))
	return TShortCut(ret)
}

func DShortCutToText(val TShortCut) string {
	ret, _, _ := dShortCutToText.Call(uintptr(val))
	return GoStr(ret)
}

func DSysOpen(filename string) {
	dSysOpen.Call(PascalStr(filename))
}

func DExtractFilePath(filename string) string {
	r, _, _ := dExtractFilePath.Call(PascalStr(filename))
	return GoStr(r)
}

func DFileExists(filename string) bool {
	r, _, _ := dFileExists.Call(PascalStr(filename))
	return GoBool(r)
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory1.Call(uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0)
	v := GoBool(r)
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory2.Call(PascalStr(caption), PascalStr(root), PascalBool(showHidden), uintptr(unsafe.Pointer(&ptr)))
	v := GoBool(r)
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSynchronize(fn func(), useMsg uintptr) {
	threadSync.Lock()
	defer threadSync.Unlock()
	threadSyncFn = fn
	dSynchronize.Call(useMsg)
	threadSyncFn = nil
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	r, _, _ := dInputBox.Call(PascalStr(aCaption), PascalStr(aPrompt), PascalStr(aDefault))
	return GoStr(r)
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r, _, _ := dInputQuery.Call(PascalStr(aCaption), PascalStr(aPrompt), PascalStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = GoStr(strPtr)
	}
	return GoBool(r)
}

func DPasswordBox(aCaption, aPrompt string) string {
	r, _, _ := dPasswordBox.Call(PascalStr(aCaption), PascalStr(aPrompt))
	return GoStr(r)
}

func DInputCombo(aCaption, aPrompt string, aList uintptr) int32 {
	r, _, _ := dInputCombo.Call(PascalStr(aCaption), PascalStr(aPrompt), aList)
	return int32(r)
}

func DInputComboEx(aCaption, aPrompt string, aList uintptr, allowCustomText bool) string {
	r, _, _ := dInputComboEx.Call(PascalStr(aCaption), PascalStr(aPrompt), aList, PascalBool(allowCustomText))
	return GoStr(r)
}

func DSysLocale(aInfo *TSysLocale) {
	dSysLocale.Call(uintptr(unsafe.Pointer(aInfo)))
}

func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	dCreateURLShortCut.Call(PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aURL))
}

func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	r, _, _ := dCreateShortCut.Call(PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aSrcFileName),
		PascalStr(aIconFileName), PascalStr(aDescription), PascalStr(aCmdArgs))
	return GoBool(r)
}

func DSetPropertyValue(instance uintptr, propName, value string) {
	dSetPropertyValue.Call(instance, PascalStr(propName), PascalStr(value))
}

func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	dSetPropertySecValue.Call(instance, PascalStr(propName), PascalStr(secPropName), PascalStr(value))
}

func DGUIDToString(guid TGUID) string {
	r, _, _ := dGUIDToString.Call(uintptr(unsafe.Pointer(&guid)))
	return GoStr(r)
}

func DStringToGUID(str string) TGUID {
	var guid TGUID
	dStringToGUID.Call(PascalStr(str), uintptr(unsafe.Pointer(&guid)))
	return guid
}

func DCreateGUID() TGUID {
	var guid TGUID
	dCreateGUID.Call(uintptr(unsafe.Pointer(&guid)))
	return guid
}

func DGetLibResourceCount() int32 {
	r, _, _ := dGetLibResourceCount.Call()
	return int32(r)
}

func DGetLibResourceItem(aIndex int32) (ret TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	dGetLibResourceItem.Call(uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = GoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResource(aPtr uintptr, aValue string) {
	dModifyLibResource.Call(aPtr, PascalStr(aValue))
}

func DLibStringEncoding() TStringEncoding {
	r, _, _ := dLibStringEncoding.Call()
	return TStringEncoding(r)
}

func DLibVersion() uint32 {
	r, _, _ := dLibVersion.Call()
	return uint32(r)
}

func DLibAbout() string {
	r, _, _ := dLibAbout.Call()
	return GoStr(r)
}

func DMainThreadId() uintptr {
	r, _, _ := dMainThreadId.Call()
	return r
}

func DCurrentThreadId() uintptr {
	r, _, _ := dCurrentThreadId.Call()
	return r
}

func DInitGoDll(aMainThreadId uintptr) {
	dInitGoDll.Call(aMainThreadId)
}

func DFindControl(handle HWND) uintptr {
	r, _, _ := dFindControl.Call(handle)
	return r
}

func DFindLCLControl(screenPos TPoint) uintptr {
	r, _, _ := dFindLCLControl.Call(uintptr(unsafe.Pointer(&screenPos)))
	return r
}

func DFindOwnerControl(handle HWND) uintptr {
	r, _, _ := dFindOwnerControl.Call(handle)
	return r
}

func DFindControlAtPosition(position TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindControlAtPosition.Call(uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
	return r
}

func DFindLCLWindow(screenPos TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindLCLWindow.Call(uintptr(unsafe.Pointer(&screenPos)), PascalBool(allowDisabled))
	return r
}

func DFindDragTarget(position TPoint, allowDisabled bool) uintptr {
	r, _, _ := dFindDragTarget.Call(uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
	return r
}
