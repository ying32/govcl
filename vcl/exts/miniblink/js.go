// +build windows

package miniblink

// JsExecState
func (es JsExecState) ArgCount() int {
	return jsArgCount(es)
}

func (es JsExecState) ArgType(argIdx int) JsType {
	return jsArgType(es, argIdx)
}

func (es JsExecState) Arg(argIdx int) JsValue {
	return jsArg(es, argIdx)
}

func (es JsExecState) ToInt(v JsValue) int {
	return jsToInt(es, v)
}

func (es JsExecState) ToFloat(v JsValue) float32 {
	return jsToFloat(es, v)
}

func (es JsExecState) ToDouble(v JsValue) float64 {
	return jsToDouble(es, v)
}

func (es JsExecState) ToBoolean(v JsValue) bool {
	return jsToBoolean(es, v)
}

func (es JsExecState) ArrayBuffer(buffer string, size uint) JsValue {
	return jsArrayBuffer(es, buffer, size)
}

func (es JsExecState) GetArrayBuffer(value JsValue) *WkeMemBuf {
	return jsGetArrayBuffer(es, value)
}

func (es JsExecState) ToTempString(v JsValue) string {
	return jsToTempString(es, v)
}

func (es JsExecState) ToTempStringW(v JsValue) string {
	return jsToTempStringW(es, v)
}

func (es JsExecState) ToV8Value(v JsValue) uintptr {
	return jsToV8Value(es, v)
}

func (es JsExecState) String(str string) JsValue {
	return jsString(es, str)
}

func (es JsExecState) StringW(str string) JsValue {
	return jsStringW(es, str)
}

func (es JsExecState) EmptyObject() JsValue {
	return jsEmptyObject(es)
}

func (es JsExecState) EmptyArray() JsValue {
	return jsEmptyArray(es)
}

func (es JsExecState) Object(obj *JsData) JsValue {
	return jsObject(es, obj)
}

func (es JsExecState) Function(obj *JsData) JsValue {
	return jsFunction(es, obj)
}

func (es JsExecState) GetData(object JsValue) *JsData {
	return jsGetData(es, object)
}

func (es JsExecState) Get(object JsValue, prop string) JsValue {
	return jsGet(es, object, prop)
}

func (es JsExecState) Set(object JsValue, prop string, v JsValue) {
	jsSet(es, object, prop, v)
}

func (es JsExecState) GetAt(object JsValue, index int) JsValue {
	return jsGetAt(es, object, index)
}

func (es JsExecState) SetAt(object JsValue, index int, v JsValue) {
	jsSetAt(es, object, index, v)
}

func (es JsExecState) GetKeys(object JsValue) *JsKeys {
	return jsGetKeys(es, object)
}

func (es JsExecState) IsJsValueValid(object JsValue) bool {
	return jsIsJsValueValid(es, object)
}

func (es JsExecState) IsValidExecState() bool {
	return jsIsValidExecState(es)
}

func (es JsExecState) DeleteObjectProp(object JsValue, prop string) {
	jsDeleteObjectProp(es, object, prop)
}

func (es JsExecState) GetLength(object JsValue) int {
	return jsGetLength(es, object)
}

func (es JsExecState) SetLength(object JsValue, length int) {
	jsSetLength(es, object, length)
}

func (es JsExecState) GlobalObject() JsValue {
	return jsGlobalObject(es)
}

func (es JsExecState) GetWebView() WkeWebView {
	return jsGetWebView(es)
}

func (es JsExecState) Eval(str string) JsValue {
	return jsEval(es, str)
}

func (es JsExecState) EvalW(str string) JsValue {
	return jsEvalW(es, str)
}

func (es JsExecState) EvalExW(str string, isInClosure bool) JsValue {
	return jsEvalExW(es, str, isInClosure)
}

func (es JsExecState) Call(aFunc JsValue, thisObject JsValue, args *JsValue, argCount int) JsValue {
	return jsCall(es, aFunc, thisObject, args, argCount)
}

func (es JsExecState) CallGlobal(aFunc JsValue, args *JsValue, argCount int) JsValue {
	return jsCallGlobal(es, aFunc, args, argCount)
}

func (es JsExecState) GetGlobal(prop string) JsValue {
	return jsGetGlobal(es, prop)
}

func (es JsExecState) SetGlobal(prop string, v JsValue) {
	jsSetGlobal(es, prop, v)
}

func (es JsExecState) AddRef(val JsValue) bool {
	return jsAddRef(es, val)
}

func (es JsExecState) ReleaseRef(val JsValue) bool {
	return jsReleaseRef(es, val)

}

func (es JsExecState) GetLastErrorIfException() uintptr {
	return jsGetLastErrorIfException(es)
}

func (es JsExecState) ThrowException(exception string) JsValue {
	return jsThrowException(es, exception)
}

func (es JsExecState) GetCallstack() string {
	return jsGetCallstack(es)
}

// JsValue
// 输出
func (j JsValue) TypeOf() JsType {
	return jsTypeOf(j)
}

func (j JsValue) IsNumber() bool {
	return jsIsNumber(j)
}

func (j JsValue) IsString() bool {
	return jsIsString(j)
}

func (j JsValue) IsBoolean() bool {
	return jsIsBoolean(j)
}

func (j JsValue) IsObject() bool {
	return jsIsObject(j)
}

func (j JsValue) IsFunction() bool {
	return jsIsFunction(j)
}

func (j JsValue) IsUndefined() bool {
	return jsIsUndefined(j)
}

func (j JsValue) IsNull() bool {
	return jsIsNull(j)
}

func (j JsValue) IsArray() bool {
	return jsIsArray(j)
}

func (j JsValue) IsTrue() bool {
	return jsIsTrue(j)
}

func (j JsValue) IsFalse() bool {
	return jsIsFalse(j)
}

// 输入
func (j *JsValue) Int(n int) {
	*j = jsInt(n)
}

func (j *JsValue) Float(f float32) {
	*j = jsFloat(f)
}

func (j *JsValue) Double(d float64) {
	*j = jsDouble(d)
}

func (j *JsValue) Boolean(b bool) {
	*j = jsBoolean(b)
}

// 绝对
func (j *JsValue) Undefined() {
	*j = jsUndefined()
}

func (j *JsValue) Null() {
	*j = jsNull()
}

func (j *JsValue) True() {
	*j = jsTrue()
}

func (j *JsValue) False() {
	*j = jsFalse()
}

// GC
func (j JsValue) GC() {
	jsGC()
}
