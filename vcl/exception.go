
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type Exception struct {
    IObject
    instance uintptr
}

func ExceptionFromInst(inst uintptr) *Exception {
    e := new(Exception)
    e.instance = inst
    return e
}

func ExceptionFromObj(obj IObject) *Exception {
    e := new(Exception)
    e.instance = CheckPtr(obj)
    return e
}

func (e *Exception) Instance() uintptr {
    return e.instance
}

func (e *Exception) IsValid() bool {
    return e.instance != 0
}

func ExceptionClass() TClass {
    return Exception_StaticClassType()
}

func (e *Exception) ToString() string {
    return Exception_ToString(e.instance)
}

func (e *Exception) DisposeOf() {
    Exception_DisposeOf(e.instance)
}

func (e *Exception) ClassType() TClass {
    return Exception_ClassType(e.instance)
}

func (e *Exception) ClassName() string {
    return Exception_ClassName(e.instance)
}

func (e *Exception) InstanceSize() int32 {
    return Exception_InstanceSize(e.instance)
}

func (e *Exception) InheritsFrom(AClass TClass) bool {
    return Exception_InheritsFrom(e.instance, AClass)
}

func (e *Exception) Equals(Obj IObject) bool {
    return Exception_Equals(e.instance, CheckPtr(Obj))
}

func (e *Exception) GetHashCode() int32 {
    return Exception_GetHashCode(e.instance)
}

func (e *Exception) BaseException() *Exception {
    return ExceptionFromInst(Exception_GetBaseException(e.instance))
}

func (e *Exception) InnerException() *Exception {
    return ExceptionFromInst(Exception_GetInnerException(e.instance))
}

func (e *Exception) Message() string {
    return Exception_GetMessage(e.instance)
}

func (e *Exception) SetMessage(value string) {
    Exception_SetMessage(e.instance, value)
}

func (e *Exception) StackTrace() string {
    return Exception_GetStackTrace(e.instance)
}

func (e *Exception) StackInfo() uintptr {
    return Exception_GetStackInfo(e.instance)
}

