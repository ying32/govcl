//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows
// +build memorydll
// +build !tempdll

// 指令为：target == windows && memorydll && !tempdll

// 编译时加上 -tags memorydll

// 内存加载liblcl.dll
// 注：内存加载dll容易被杀毒软件杀掉。。。
package api

import (
	"github.com/ying32/govcl/vcl/api/memorydll"
	"github.com/ying32/govcl/vcl/win"
)

func loadUILib() *memorydll.LazyDLL {
	dllBytes, ok := win.ResourceToBytes(win.GetSelfModuleHandle(), "GOVCLLIB", win.RT_RCDATA)
	if !ok {
		panic("\"GOVCLLIB\" resource does not exist.")
		return nil
	}
	lib := memorydll.NewMemoryDLL(dllBytes)
	if lib.Load() != nil {
		panic("Unable to load dll, unknown reason.")
		return nil
	}
	if getLibType(lib) != 1 {
		// 当前已经不再支持VCL库了。如果有需要，请使用最后一个支持VCL版本的代码：https://github.com/ying32/govcl/tree/last-vcl-support。
		panic("The VCL library is no longer supported. If necessary, please use the last code that supports VCL version: https://github.com/ying32/govcl/tree/last-vcl-support.")
	}
	return lib
}

func getLibType(lib *memorydll.LazyDLL) int32 {
	proc := lib.NewProc("DGetLibType")
	r, _, _ := proc.Call()
	return int32(r)
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibVcl() *memorydll.LazyDLL {
	return libvcl
}

// 释放内存，Memory DLL的最好调用下
func FeeMemoryDLL() {
	if libvcl != nil {
		libvcl.Close()
		libvcl = nil
	}
}

func closeLib() {
	FeeMemoryDLL()
}
