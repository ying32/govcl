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
	resFormLoadFromStream       = newDLLProc("ResFormLoadFromStream")
	resFormLoadFromFile         = newDLLProc("ResFormLoadFromFile")
	resFormLoadFromResourceName = newDLLProc("ResFormLoadFromResourceName")
)

// ResFormLoadFromStream
//
// 从流中加载窗口资源
func ResFormLoadFromStream(obj, root uintptr) {
	resFormLoadFromStream.Call(obj, root)
}

// ResFormLoadFromFile
//
// 从文件中加载窗口资源
func ResFormLoadFromFile(filename string, root uintptr) {
	resFormLoadFromFile.Call(PascalStr(filename), root)
}

// ResFormLoadFromResourceName
//
// 从指定资源中加载窗口资源
func ResFormLoadFromResourceName(instance uintptr, resName string, root uintptr) {
	resFormLoadFromResourceName.Call(instance, PascalStr(resName), root)
}
