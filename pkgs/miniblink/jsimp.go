// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl，因为没有完整测试，所以不保证100%正确

package miniblink

import (
	"github.com/ying32/dylib/floatpatch"

	"unsafe"
)

var (
	_jsBindFunction            = wkedll.NewProc("jsBindFunction")
	_jsBindGetter              = wkedll.NewProc("jsBindGetter")
	_jsBindSetter              = wkedll.NewProc("jsBindSetter")
	_wkeJsBindFunction         = wkedll.NewProc("wkeJsBindFunction")
	_wkeJsBindGetter           = wkedll.NewProc("wkeJsBindGetter")
	_wkeJsBindSetter           = wkedll.NewProc("wkeJsBindSetter")
	_jsArgCount                = wkedll.NewProc("jsArgCount")
	_jsArgType                 = wkedll.NewProc("jsArgType")
	_jsArg                     = wkedll.NewProc("jsArg")
	_jsTypeOf                  = wkedll.NewProc("jsTypeOf")
	_jsIsNumber                = wkedll.NewProc("jsIsNumber")
	_jsIsString                = wkedll.NewProc("jsIsString")
	_jsIsBoolean               = wkedll.NewProc("jsIsBoolean")
	_jsIsObject                = wkedll.NewProc("jsIsObject")
	_jsIsFunction              = wkedll.NewProc("jsIsFunction")
	_jsIsUndefined             = wkedll.NewProc("jsIsUndefined")
	_jsIsNull                  = wkedll.NewProc("jsIsNull")
	_jsIsArray                 = wkedll.NewProc("jsIsArray")
	_jsIsTrue                  = wkedll.NewProc("jsIsTrue")
	_jsIsFalse                 = wkedll.NewProc("jsIsFalse")
	_jsToInt                   = wkedll.NewProc("jsToInt")
	_jsToFloat                 = wkedll.NewProc("jsToFloat")
	_jsToDouble                = wkedll.NewProc("jsToDouble")
	_jsToBoolean               = wkedll.NewProc("jsToBoolean")
	_jsArrayBuffer             = wkedll.NewProc("jsArrayBuffer")
	_jsGetArrayBuffer          = wkedll.NewProc("jsGetArrayBuffer")
	_jsToTempString            = wkedll.NewProc("jsToTempString")
	_jsToTempStringW           = wkedll.NewProc("jsToTempStringW")
	_jsToV8Value               = wkedll.NewProc("jsToV8Value")
	_jsInt                     = wkedll.NewProc("jsInt")
	_jsFloat                   = wkedll.NewProc("jsFloat")
	_jsDouble                  = wkedll.NewProc("jsDouble")
	_jsBoolean                 = wkedll.NewProc("jsBoolean")
	_jsUndefined               = wkedll.NewProc("jsUndefined")
	_jsNull                    = wkedll.NewProc("jsNull")
	_jsTrue                    = wkedll.NewProc("jsTrue")
	_jsFalse                   = wkedll.NewProc("jsFalse")
	_jsString                  = wkedll.NewProc("jsString")
	_jsStringW                 = wkedll.NewProc("jsStringW")
	_jsEmptyObject             = wkedll.NewProc("jsEmptyObject")
	_jsEmptyArray              = wkedll.NewProc("jsEmptyArray")
	_jsObject                  = wkedll.NewProc("jsObject")
	_jsFunction                = wkedll.NewProc("jsFunction")
	_jsGetData                 = wkedll.NewProc("jsGetData")
	_jsGet                     = wkedll.NewProc("jsGet")
	_jsSet                     = wkedll.NewProc("jsSet")
	_jsGetAt                   = wkedll.NewProc("jsGetAt")
	_jsSetAt                   = wkedll.NewProc("jsSetAt")
	_jsGetKeys                 = wkedll.NewProc("jsGetKeys")
	_jsIsJsValueValid          = wkedll.NewProc("jsIsJsValueValid")
	_jsIsValidExecState        = wkedll.NewProc("jsIsValidExecState")
	_jsDeleteObjectProp        = wkedll.NewProc("jsDeleteObjectProp")
	_jsGetLength               = wkedll.NewProc("jsGetLength")
	_jsSetLength               = wkedll.NewProc("jsSetLength")
	_jsGlobalObject            = wkedll.NewProc("jsGlobalObject")
	_jsGetWebView              = wkedll.NewProc("jsGetWebView")
	_jsEval                    = wkedll.NewProc("jsEval")
	_jsEvalW                   = wkedll.NewProc("jsEvalW")
	_jsEvalExW                 = wkedll.NewProc("jsEvalExW")
	_jsCall                    = wkedll.NewProc("jsCall")
	_jsCallGlobal              = wkedll.NewProc("jsCallGlobal")
	_jsGetGlobal               = wkedll.NewProc("jsGetGlobal")
	_jsSetGlobal               = wkedll.NewProc("jsSetGlobal")
	_jsGC                      = wkedll.NewProc("jsGC")
	_jsAddRef                  = wkedll.NewProc("jsAddRef")
	_jsReleaseRef              = wkedll.NewProc("jsReleaseRef")
	_jsGetLastErrorIfException = wkedll.NewProc("jsGetLastErrorIfException")
	_jsThrowException          = wkedll.NewProc("jsThrowException")
	_jsGetCallstack            = wkedll.NewProc("jsGetCallstack")
)

func jsBindFunction(name string, fn JsNativeFunction, argCount uint) {
	_jsBindFunction.Call(CAStr(name), uintptr(fn), uintptr(argCount))
}

func jsBindGetter(name string, fn JsNativeFunction) {
	_jsBindGetter.Call(CAStr(name), uintptr(fn))
}

func jsBindSetter(name string, fn JsNativeFunction) {
	_jsBindSetter.Call(CAStr(name), uintptr(fn))
}

func wkeJsBindFunction(name string, fn WkeJsNativeFunction, param uintptr, argCount uint) {
	_wkeJsBindFunction.Call(CAStr(name), uintptr(fn), param, uintptr(argCount))
}

func wkeJsBindGetter(name string, fn WkeJsNativeFunction, param uintptr) {
	_wkeJsBindGetter.Call(CAStr(name), uintptr(fn), param)
}

func wkeJsBindSetter(name string, fn WkeJsNativeFunction, param uintptr) {
	_wkeJsBindSetter.Call(CAStr(name), uintptr(fn), param)
}

func jsArgCount(es JsExecState) int {
	r, _, _ := _jsArgCount.Call(uintptr(es))
	return int(r)
}

func jsArgType(es JsExecState, argIdx int) JsType {
	r, _, _ := _jsArgType.Call(uintptr(es), uintptr(argIdx))
	return JsType(r)
}

func jsArg(es JsExecState, argIdx int) JsValue {
	r, r2, _ := _jsArg.Call(uintptr(es), uintptr(argIdx))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsTypeOf(v JsValue) JsType {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsTypeOf.Call(v1, v2)
	} else {
		r, _, _ = _jsTypeOf.Call(uintptr(v))
	}
	return JsType(r)
}

func jsIsNumber(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsNumber.Call(v1, v2)
	} else {
		r, _, _ = _jsIsNumber.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsString(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsString.Call(v1, v2)
	} else {
		r, _, _ = _jsIsString.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsBoolean(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsBoolean.Call(v1, v2)
	} else {
		r, _, _ = _jsIsBoolean.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsObject(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsObject.Call(v1, v2)
	} else {
		r, _, _ = _jsIsObject.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsFunction(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsFunction.Call(v1, v2)
	} else {
		r, _, _ = _jsIsFunction.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsUndefined(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsUndefined.Call(v1, v2)
	} else {
		r, _, _ = _jsIsUndefined.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsNull(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsNull.Call(v1, v2)
	} else {
		r, _, _ = _jsIsNull.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsArray(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsArray.Call(v1, v2)
	} else {
		r, _, _ = _jsIsArray.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsTrue(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsTrue.Call(v1, v2)
	} else {
		r, _, _ = _jsIsTrue.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsIsFalse(v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsIsFalse.Call(v1, v2)
	} else {
		r, _, _ = _jsIsFalse.Call(uintptr(v))
	}
	return GoBool(r)
}

func jsToInt(es JsExecState, v JsValue) int {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToInt.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToInt.Call(uintptr(es), uintptr(v))
	}
	return int(r)
}

func jsToFloat(es JsExecState, v JsValue) float32 {
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		_jsToFloat.Call(uintptr(es), v1, v2)
	} else {
		_jsToFloat.Call(uintptr(es), uintptr(v))
	}
	return floatpatch.Getfloat32()
}

func jsToDouble(es JsExecState, v JsValue) float64 {
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		_jsToDouble.Call(uintptr(es), v1, v2)
	} else {
		_jsToDouble.Call(uintptr(es), uintptr(v))
	}
	return floatpatch.Getfloat64()
}

func jsToBoolean(es JsExecState, v JsValue) bool {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToBoolean.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToBoolean.Call(uintptr(es), uintptr(v))
	}
	return GoBool(r)
}

func jsArrayBuffer(es JsExecState, buffer string, size uint) JsValue {
	r, r2, _ := _jsArrayBuffer.Call(uintptr(es), CAStr(buffer), uintptr(size))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsGetArrayBuffer(es JsExecState, value JsValue) *WkeMemBuf {
	var r uintptr
	if is386 {
		value1, value2 := UInt64To(uint64(value))
		r, _, _ = _jsGetArrayBuffer.Call(uintptr(es), value1, value2)
	} else {
		r, _, _ = _jsGetArrayBuffer.Call(uintptr(es), uintptr(value))
	}
	return (*WkeMemBuf)(unsafe.Pointer(r))
}

func jsToTempString(es JsExecState, v JsValue) string {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToTempString.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToTempString.Call(uintptr(es), uintptr(v))
	}
	return GoAStr(r)
}

func jsToTempStringW(es JsExecState, v JsValue) string {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToTempStringW.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToTempStringW.Call(uintptr(es), uintptr(v))
	}
	return GoWStr(r)
}

func jsToV8Value(es JsExecState, v JsValue) uintptr {
	var r uintptr
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		r, _, _ = _jsToV8Value.Call(uintptr(es), v1, v2)
	} else {
		r, _, _ = _jsToV8Value.Call(uintptr(es), uintptr(v))
	}
	return r
}

func jsInt(n int) JsValue {
	r, r2, _ := _jsInt.Call(uintptr(n))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsFloat(f float32) JsValue {
	r, r2, _ := _jsFloat.Call(uintptr(f))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsDouble(d float64) JsValue {
	r, r2, _ := _jsDouble.Call(uintptr(d))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsBoolean(b bool) JsValue {
	r, r2, _ := _jsBoolean.Call(CBool(b))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsUndefined() JsValue {
	r, r2, _ := _jsUndefined.Call()
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsNull() JsValue {
	r, r2, _ := _jsNull.Call()
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsTrue() JsValue {
	r, r2, _ := _jsTrue.Call()
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsFalse() JsValue {
	r, r2, _ := _jsFalse.Call()
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsString(es JsExecState, str string) JsValue {
	r, r2, _ := _jsString.Call(uintptr(es), CAStr(str))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsStringW(es JsExecState, str string) JsValue {
	r, r2, _ := _jsStringW.Call(uintptr(es), CWStr(str))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsEmptyObject(es JsExecState) JsValue {
	r, r2, _ := _jsEmptyObject.Call(uintptr(es))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsEmptyArray(es JsExecState) JsValue {
	r, r2, _ := _jsEmptyArray.Call(uintptr(es))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsObject(es JsExecState, obj *JsData) JsValue {
	r, r2, _ := _jsObject.Call(uintptr(es), uintptr(unsafe.Pointer(obj)))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsFunction(es JsExecState, obj *JsData) JsValue {
	r, r2, _ := _jsFunction.Call(uintptr(es), uintptr(unsafe.Pointer(obj)))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsGetData(es JsExecState, object JsValue) *JsData {
	var r uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, _, _ = _jsGetData.Call(uintptr(es), object1, object2)
	} else {
		r, _, _ = _jsGetData.Call(uintptr(es), uintptr(object))
	}
	return (*JsData)(unsafe.Pointer(r))
}

func jsGet(es JsExecState, object JsValue, prop string) JsValue {
	var r, r2 uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, r2, _ = _jsGet.Call(uintptr(es), object1, object2, CAStr(prop))
		return JsValue(ToUInt64(r, r2))
	} else {
		r, r2, _ = _jsGet.Call(uintptr(es), uintptr(object), CAStr(prop))
		return JsValue(r)
	}
}

func jsSet(es JsExecState, object JsValue, prop string, v JsValue) {
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		v1, v2 := UInt64To(uint64(v))
		_jsSet.Call(uintptr(es), object1, object2, CAStr(prop), v1, v2)
	} else {
		_jsSet.Call(uintptr(es), uintptr(object), CAStr(prop), uintptr(v))
	}
}

func jsGetAt(es JsExecState, object JsValue, index int) JsValue {
	var r, r2 uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, r2, _ = _jsGetAt.Call(uintptr(es), object1, object2, uintptr(index))
		return JsValue(ToUInt64(r, r2))
	} else {
		r, r2, _ = _jsGetAt.Call(uintptr(es), uintptr(object), uintptr(index))
		return JsValue(r)
	}
}

func jsSetAt(es JsExecState, object JsValue, index int, v JsValue) {
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		v1, v2 := UInt64To(uint64(v))
		_jsSetAt.Call(uintptr(es), object1, object2, uintptr(index), v1, v2)
	} else {
		_jsSetAt.Call(uintptr(es), uintptr(object), uintptr(index), uintptr(v))
	}
}

func jsGetKeys(es JsExecState, object JsValue) *JsKeys {
	var r uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, _, _ = _jsGetKeys.Call(uintptr(es), object1, object2)
	} else {
		r, _, _ = _jsGetKeys.Call(uintptr(es), uintptr(object))
	}
	return (*JsKeys)(unsafe.Pointer(r))
}

func jsIsJsValueValid(es JsExecState, object JsValue) bool {
	var r uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, _, _ = _jsIsJsValueValid.Call(uintptr(es), object1, object2)
	} else {
		r, _, _ = _jsIsJsValueValid.Call(uintptr(es), uintptr(object))
	}
	return GoBool(r)
}

func jsIsValidExecState(es JsExecState) bool {
	r, _, _ := _jsIsValidExecState.Call(uintptr(es))
	return GoBool(r)
}

func jsDeleteObjectProp(es JsExecState, object JsValue, prop string) {
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		_jsDeleteObjectProp.Call(uintptr(es), object1, object2, CAStr(prop))
	} else {
		_jsDeleteObjectProp.Call(uintptr(es), uintptr(object), CAStr(prop))
	}
}

func jsGetLength(es JsExecState, object JsValue) int {
	var r uintptr
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		r, _, _ = _jsGetLength.Call(uintptr(es), object1, object2)
	} else {
		r, _, _ = _jsGetLength.Call(uintptr(es), uintptr(object))
	}
	return int(r)
}

func jsSetLength(es JsExecState, object JsValue, length int) {
	if is386 {
		object1, object2 := UInt64To(uint64(object))
		_jsSetLength.Call(uintptr(es), object1, object2, uintptr(length))
	} else {
		_jsSetLength.Call(uintptr(es), uintptr(object), uintptr(length))
	}
}

func jsGlobalObject(es JsExecState) JsValue {
	r, r2, _ := _jsGlobalObject.Call(uintptr(es))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsGetWebView(es JsExecState) WkeWebView {
	r, _, _ := _jsGetWebView.Call(uintptr(es))
	return WkeWebView(r)
}

func jsEval(es JsExecState, str string) JsValue {
	r, r2, _ := _jsEval.Call(uintptr(es), CAStr(str))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsEvalW(es JsExecState, str string) JsValue {
	r, r2, _ := _jsEvalW.Call(uintptr(es), CWStr(str))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsEvalExW(es JsExecState, str string, isInClosure bool) JsValue {
	r, r2, _ := _jsEvalExW.Call(uintptr(es), CWStr(str), CBool(isInClosure))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsCall(es JsExecState, aFunc JsValue, thisObject JsValue, args *JsValue, argCount int) JsValue {
	var r, r2 uintptr
	if is386 {
		aFunc1, aFunc2 := UInt64To(uint64(aFunc))
		thisObject1, thisObject2 := UInt64To(uint64(thisObject))
		r, r2, _ = _jsCall.Call(uintptr(es), aFunc1, aFunc2, thisObject1, thisObject2, uintptr(unsafe.Pointer(args)), uintptr(argCount))
		return JsValue(ToUInt64(r, r2))
	} else {
		r, r2, _ = _jsCall.Call(uintptr(es), uintptr(aFunc), uintptr(thisObject), uintptr(unsafe.Pointer(args)), uintptr(argCount))
		return JsValue(r)
	}
}

func jsCallGlobal(es JsExecState, aFunc JsValue, args *JsValue, argCount int) JsValue {
	var r, r2 uintptr
	if is386 {
		aFunc1, aFunc2 := UInt64To(uint64(aFunc))
		r, r2, _ = _jsCallGlobal.Call(uintptr(es), aFunc1, aFunc2, uintptr(unsafe.Pointer(args)), uintptr(argCount))
		return JsValue(ToUInt64(r, r2))
	} else {
		r, r2, _ = _jsCallGlobal.Call(uintptr(es), uintptr(aFunc), uintptr(unsafe.Pointer(args)), uintptr(argCount))
		return JsValue(r)
	}
}

func jsGetGlobal(es JsExecState, prop string) JsValue {
	r, r2, _ := _jsGetGlobal.Call(uintptr(es), CAStr(prop))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsSetGlobal(es JsExecState, prop string, v JsValue) {
	if is386 {
		v1, v2 := UInt64To(uint64(v))
		_jsSetGlobal.Call(uintptr(es), CAStr(prop), v1, v2)
	} else {
		_jsSetGlobal.Call(uintptr(es), CAStr(prop), uintptr(v))
	}
}

func jsGC() {
	_jsGC.Call()
}

func jsAddRef(es JsExecState, val JsValue) bool {
	var r uintptr
	if is386 {
		val1, val2 := UInt64To(uint64(val))
		r, _, _ = _jsAddRef.Call(uintptr(es), val1, val2)
	} else {
		r, _, _ = _jsAddRef.Call(uintptr(es), uintptr(val))
	}
	return GoBool(r)
}

func jsReleaseRef(es JsExecState, val JsValue) bool {
	var r uintptr
	if is386 {
		val1, val2 := UInt64To(uint64(val))
		r, _, _ = _jsReleaseRef.Call(uintptr(es), val1, val2)
	} else {
		r, _, _ = _jsReleaseRef.Call(uintptr(es), uintptr(val))
	}
	return GoBool(r)
}

func jsGetLastErrorIfException(es JsExecState) uintptr {
	r, _, _ := _jsGetLastErrorIfException.Call(uintptr(es))
	return r
}

func jsThrowException(es JsExecState, exception string) JsValue {
	r, r2, _ := _jsThrowException.Call(uintptr(es), CAStr(exception))
	if is386 {
		return JsValue(ToUInt64(r, r2))
	}
	return JsValue(r)
}

func jsGetCallstack(es JsExecState) string {
	r, _, _ := _jsGetCallstack.Call(uintptr(es))
	return GoAStr(r)
}
