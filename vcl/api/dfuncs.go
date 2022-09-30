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

	"github.com/ying32/govcl/vcl/api/dllimports"

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

func PascalBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
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
	defSyscallN(dllimports.SETEVENTCALLBACK, ptr)
}

func SetMessageCallback(ptr uintptr) {
	defSyscallN(dllimports.SETMESSAGECALLBACK, ptr)
}

func SetThreadSyncCallback(ptr uintptr) {
	defSyscallN(dllimports.SETTHREADSYNCCALLBACK, ptr)
}

func SetRequestCallCreateParamsCallback(ptr uintptr) {
	defSyscallN(dllimports.SETREQUESTCALLCREATEPARAMSCALLBACK, ptr)
}

func SetRemoveEventCallback(ptr uintptr) {
	defSyscallN(dllimports.SETREMOVEEVENTCALLBACK, ptr)
}

func DGetStringArrOf(p uintptr, index int) string {
	return GoStr(defSyscallN(dllimports.DGETSTRINGARROF, p, uintptr(index)))
}

func DStrLen(p uintptr) int {
	return int(defSyscallN(dllimports.DSTRLEN, p))
}

func DMove(src, dest uintptr, nLen int) {
	defSyscallN(dllimports.DMOVE, src, dest, uintptr(nLen))
}

func DShowMessage(s string) {
	defSyscallN(dllimports.DSHOWMESSAGE, PascalStr(s))
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	return int32(defSyscallN(dllimports.DMESSAGEDLG, PascalStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx)))
}

func DTextToShortCut(val string) TShortCut {
	return TShortCut(defSyscallN(dllimports.DTEXTTOSHORTCUT, PascalStr(val)))
}

func DShortCutToText(val TShortCut) string {
	return GoStr(defSyscallN(dllimports.DSHORTCUTTOTEXT, uintptr(val)))
}

func DSysOpen(filename string) {
	defSyscallN(dllimports.DSYSOPEN, PascalStr(filename))
}

func DExtractFilePath(filename string) string {
	return GoStr(defSyscallN(dllimports.DEXTRACTFILEPATH, PascalStr(filename)))
}

func DFileExists(filename string) bool {
	return GoBool(defSyscallN(dllimports.DFILEEXISTS, PascalStr(filename)))
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(dllimports.DSELECTDIRECTORY1, uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	var ptr uintptr
	v := GoBool(defSyscallN(dllimports.DSELECTDIRECTORY2, PascalStr(caption), PascalStr(root), PascalBool(showHidden), uintptr(unsafe.Pointer(&ptr))))
	if v {
		return true, GoStr(ptr)
	}
	return false, ""
}

func DSynchronize(fn func(), useMsg uintptr) {
	threadSync.Lock()
	defer threadSync.Unlock()
	threadSyncFn = fn
	defSyscallN(dllimports.DSYNCHRONIZE, useMsg)
	threadSyncFn = nil
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	return GoStr(defSyscallN(dllimports.DINPUTBOX, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(aDefault)))
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r := defSyscallN(dllimports.DINPUTQUERY, PascalStr(aCaption), PascalStr(aPrompt), PascalStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = GoStr(strPtr)
	}
	return GoBool(r)
}

func DPasswordBox(aCaption, aPrompt string) string {
	return GoStr(defSyscallN(dllimports.DPASSWORDBOX, PascalStr(aCaption), PascalStr(aPrompt)))
}

func DInputCombo(aCaption, aPrompt string, aList uintptr) int32 {
	return int32(defSyscallN(dllimports.DINPUTCOMBO, PascalStr(aCaption), PascalStr(aPrompt), aList))
}

func DInputComboEx(aCaption, aPrompt string, aList uintptr, allowCustomText bool) string {
	return GoStr(defSyscallN(dllimports.DINPUTCOMBOEX, PascalStr(aCaption), PascalStr(aPrompt), aList, PascalBool(allowCustomText)))
}

func DSysLocale(aInfo *TSysLocale) {
	defSyscallN(dllimports.DSYSLOCALE, uintptr(unsafe.Pointer(aInfo)))
}

func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	defSyscallN(dllimports.DCREATEURLSHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aURL))
}

func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	return GoBool(defSyscallN(dllimports.DCREATESHORTCUT, PascalStr(aDestPath), PascalStr(aShortCutName), PascalStr(aSrcFileName),
		PascalStr(aIconFileName), PascalStr(aDescription), PascalStr(aCmdArgs)))
}

func DSetPropertyValue(instance uintptr, propName, value string) {
	defSyscallN(dllimports.DSETPROPERTYVALUE, instance, PascalStr(propName), PascalStr(value))
}

func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	defSyscallN(dllimports.DSETPROPERTYSECVALUE, instance, PascalStr(propName), PascalStr(secPropName), PascalStr(value))
}

func DGUIDToString(guid TGUID) string {
	return GoStr(defSyscallN(dllimports.DGUIDTOSTRING, uintptr(unsafe.Pointer(&guid))))
}

func DStringToGUID(str string) (guid TGUID) {
	defSyscallN(dllimports.DSTRINGTOGUID, PascalStr(str), uintptr(unsafe.Pointer(&guid)))
	return
}

func DCreateGUID() (guid TGUID) {
	defSyscallN(dllimports.DCREATEGUID, uintptr(unsafe.Pointer(&guid)))
	return
}

func DGetLibResourceCount() int32 {
	return int32(defSyscallN(dllimports.DGETLIBRESOURCECOUNT))
}

func DGetLibResourceItem(aIndex int32) (ret TLibResource) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	defSyscallN(dllimports.DGETLIBRESOURCEITEM, uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = GoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResource(aPtr uintptr, aValue string) {
	defSyscallN(dllimports.DMODIFYLIBRESOURCE, aPtr, PascalStr(aValue))
}

func DLibStringEncoding() TStringEncoding {
	return TStringEncoding(defSyscallN(dllimports.DLIBSTRINGENCODING))
}

func DLibVersion() uint32 {
	return uint32(defSyscallN(dllimports.DLIBVERSION))
}

func DLibAbout() string {
	return GoStr(defSyscallN(dllimports.DLIBABOUT))
}

func DMainThreadId() uintptr {
	return defSyscallN(dllimports.DMAINTHREADID)
}

func DCurrentThreadId() uintptr {
	return defSyscallN(dllimports.DCURRENTTHREADID)
}

func DInitGoDll(aMainThreadId uintptr) {
	defSyscallN(dllimports.DINITGODLL, aMainThreadId)
}

func DFindControl(handle HWND) uintptr {
	return defSyscallN(dllimports.DFINDCONTROL, handle)
}

func DFindLCLControl(screenPos TPoint) uintptr {
	return defSyscallN(dllimports.DFINDLCLCONTROL, uintptr(unsafe.Pointer(&screenPos)))
}

func DFindOwnerControl(handle HWND) uintptr {
	return defSyscallN(dllimports.DFINDOWNERCONTROL, handle)
}

func DFindControlAtPosition(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDCONTROLATPOSITION, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}

func DFindLCLWindow(screenPos TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDLCLWINDOW, uintptr(unsafe.Pointer(&screenPos)), PascalBool(allowDisabled))
}

func DFindDragTarget(position TPoint, allowDisabled bool) uintptr {
	return defSyscallN(dllimports.DFINDDRAGTARGET, uintptr(unsafe.Pointer(&position)), PascalBool(allowDisabled))
}
