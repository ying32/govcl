package api

import (
	"reflect"
	"unsafe"

	. "github.com/ying32/govcl/vcl/types"
)

type TGoParam struct {
	Type  uint8
	Value uintptr
}

var (
	CallbackMap = map[uintptr]interface{}{}
)

func DBoolToGoBool(val uintptr) bool {
	if val != 0 {
		return true
	}
	return false
}

func GoBoolToDBool(val bool) uintptr {
	if val {
		return 1
	}
	return 0
}

func addEventToMap(f interface{}) uintptr {
	p := reflect.ValueOf(f).Pointer()
	CallbackMap[p] = f
	return p
}

func SetEventCallback(ptr uintptr) {
	setEventCallback.Call(ptr)
}

func DGetParam(index int, ptr uintptr) TGoParam {
	p := TGoParam{}
	dGetParam.Call(uintptr(index), ptr, uintptr(unsafe.Pointer(&p)))
	return p
}

func DGetStringArrOf(p uintptr, index int) string {
	r, _, _ := dGetStringArrOf.Call(p, uintptr(index))
	return DStrToGoStr(r)
}

func DStrLen(p uintptr) int {
	ret, _, _ := dStrLen.Call(p)
	return int(ret)
}

func DMove(src, dest uintptr, llen int) {
	dMove.Call(src, dest, uintptr(llen))
}

func DShowMessage(s string) {
	dShowMessage.Call(GoStrToDStr(s))
}

func DGetMainInstance() uintptr {
	ret, _, _ := dGetMainInstance.Call()
	return ret
}

func DMessageDlg(Msg string, DlgType TMsgDlgType, Buttons TMsgDlgButtons, HelpCtx int32) int32 {
	ret, _, _ := dMessageDlg.Call(GoStrToDStr(Msg), uintptr(DlgType), uintptr(Buttons), uintptr(HelpCtx))
	return int32(ret)
}

func DSetReportMemoryLeaksOnShutdown(v bool) {
	dSetReportMemoryLeaksOnShutdown.Call(GoBoolToDBool(v))
}

func DTextToShortCut(val string) TShortCut {
	ret, _, _ := dTextToShortCut.Call(GoStrToDStr(val))
	return TShortCut(ret)
}

func DShortCutToText(val TShortCut) string {
	ret, _, _ := dShortCutToText.Call(uintptr(val))
	return DStrToGoStr(ret)
}

func DSysOpen(filename string) {
	dSysOpen.Call(GoStrToDStr(filename))
}

func DExtractFilePath(filename string) string {
	r, _, _ := dExtractFilePath.Call(GoStrToDStr(filename))
	return DStrToGoStr(r)
}

func DFileExists(filename string) bool {
	r, _, _ := dFileExists.Call(GoStrToDStr(filename))
	return DBoolToGoBool(r)
}

func DSelectDirectory1(options TSelectDirOpts) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory1.Call(uintptr(unsafe.Pointer(&ptr)), uintptr(options), 0)
	v := DBoolToGoBool(r)
	if v {
		return true, DStrToGoStr(ptr)
	}
	return false, ""
}

func DSelectDirectory2(caption, root string, options TSelectDirExtOpts, parent uintptr) (bool, string) {
	var ptr uintptr
	r, _, _ := dSelectDirectory2.Call(GoStrToDStr(caption), GoStrToDStr(root), uintptr(unsafe.Pointer(&ptr)),
		uintptr(options), parent)
	v := DBoolToGoBool(r)
	if v {
		return true, DStrToGoStr(ptr)
	}
	return false, ""
}

func DSynchronize(fn interface{}) {
	dSynchronize.Call(addEventToMap(fn))
}

func DInputBox(aCaption, aPrompt, aDefault string) string {
	r, _, _ := dInputBox.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), GoStrToDStr(aDefault))
	return DStrToGoStr(r)
}

func DInputQuery(aCaption, aPrompt string, value *string) bool {
	if value == nil {
		return false
	}
	var strPtr uintptr
	r, _, _ := dInputQuery.Call(GoStrToDStr(aCaption), GoStrToDStr(aPrompt), GoStrToDStr(*value), uintptr(unsafe.Pointer(&strPtr)))
	if strPtr != 0 {
		*value = DStrToGoStr(strPtr)
	}
	return DBoolToGoBool(r)
}

// DSysLocaled
func DSysLocale(aInfo *TSysLocale) {
	dSysLocale.Call(uintptr(unsafe.Pointer(aInfo)))
}

// Shortcut
//DCreateURLShortCut
func DCreateURLShortCut(aDestPath, aShortCutName, aURL string) {
	dCreateURLShortCut.Call(GoStrToDStr(aDestPath), GoStrToDStr(aShortCutName), GoStrToDStr(aURL))
}

//DCreateShortCut
func DCreateShortCut(aDestPath, aShortCutName, aSrcFileName, aIconFileName, aDescription, aCmdArgs string) bool {
	r, _, _ := dCreateShortCut.Call(GoStrToDStr(aDestPath), GoStrToDStr(aShortCutName), GoStrToDStr(aSrcFileName),
		GoStrToDStr(aIconFileName), GoStrToDStr(aDescription), GoStrToDStr(aCmdArgs))
	return DBoolToGoBool(r)
}

// SetProperty
// DSetPropertyValue
func DSetPropertyValue(instance uintptr, propName, value string) {
	dSetPropertyValue.Call(instance, GoStrToDStr(propName), GoStrToDStr(value))
}

// DSetPropertySecValue
func DSetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	dSetPropertySecValue.Call(instance, GoStrToDStr(propName), GoStrToDStr(secPropName), GoStrToDStr(value))
}

// guid
// DGUIDToString
func DGUIDToString(guid TGUID) string {
	r, _, _ := dGUIDToString.Call(uintptr(unsafe.Pointer(&guid)))
	return DStrToGoStr(r)
}

// DStringToGUID
func DStringToGUID(str string) TGUID {
	var guid TGUID
	dStringToGUID.Call(GoStrToDStr(str), uintptr(unsafe.Pointer(&guid)))
	return guid
}

// DCreateGUID
func DCreateGUID() TGUID {
	var guid TGUID
	dCreateGUID.Call(uintptr(unsafe.Pointer(&guid)))
	return guid
}

// LibResouces
func DGetLibResouceCount() int32 {
	r, _, _ := dGetLibResouceCount.Call()
	return int32(r)
}

func DGetLibResouceItem(aIndex int32) (ret TLibResouce) {
	item := struct {
		Name     uintptr
		ValuePtr uintptr
	}{}
	dGetLibResouceItem.Call(uintptr(aIndex), uintptr(unsafe.Pointer(&item)))
	ret.Name = DStrToGoStr(item.Name)
	ret.Ptr = item.ValuePtr
	return
}

func DModifyLibResouce(aPtr uintptr, aValue string) {
	dModifyLibResouce.Call(aPtr, GoStrToDStr(aValue))
}
