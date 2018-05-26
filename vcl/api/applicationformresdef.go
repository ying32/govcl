//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 本文件内部函数不在开源范围内
// 需要配合窗口设计器使用
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------
package api

var (
	resFormLoadFromStream       = libvcl.NewProc("ResFormLoadFromStream")
	resFormLoadFromFile         = libvcl.NewProc("ResFormLoadFromFile")
	resFormLoadFromResourceName = libvcl.NewProc("ResFormLoadFromResourceName")
)

// ResFormLoadFromStream与ResFormLoadFromFile不在开源范围内
func ResFormLoadFromStream(obj, root uintptr) {
	resFormLoadFromStream.Call(obj, root)
}

func ResFormLoadFromFile(filename string, root uintptr) {
	resFormLoadFromFile.Call(GoStrToDStr(filename), root)
}

func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	resFormLoadFromResourceName.Call(instance, GoStrToDStr(resName), root)
}
