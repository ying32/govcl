//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//----------------------------------------
// 加载文件或者内存中的窗口资源文件功能
// 需要配合窗口设计器使用
//----------------------------------------
package api

var (
	resFormLoadFromStream       = libvcl.NewProc("ResFormLoadFromStream")
	resFormLoadFromFile         = libvcl.NewProc("ResFormLoadFromFile")
	resFormLoadFromResourceName = libvcl.NewProc("ResFormLoadFromResourceName")
)

// ResFormLoadFromStream
func ResFormLoadFromStream(obj, root uintptr) {
	resFormLoadFromStream.Call(obj, root)
}

func ResFormLoadFromFile(filename string, root uintptr) {
	resFormLoadFromFile.Call(GoStrToDStr(filename), root)
}

func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	resFormLoadFromResourceName.Call(instance, GoStrToDStr(resName), root)
}
