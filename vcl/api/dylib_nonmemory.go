//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !memorydll

package api

import (
	"runtime"

	"github.com/ying32/govcl/pkgs/libname"

	"github.com/ying32/dylib"
)

var (
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

// 加载库
func loadUILib() *dylib.LazyDLL {
	libName := "liblcl"
	if libname.LibName == "" {
		if ext, ok := platformExtNames[runtime.GOOS]; ok {
			libName += ext
		}
	} else {
		libName = libname.LibName
	}
	//fmt.Println("LoadLibrary:", libName)
	lib := dylib.NewLazyDLL(libName)
	err := lib.Load()
	if err != nil {
		panic(err)
	}
	if getLibType(lib) != LtLCL {
		// 当前已经不再支持VCL库了。如果有需要，请使用最后一个支持VCL版本的代码：https://github.com/ying32/govcl/tree/last-vcl-support。
		panic("The VCL library is no longer supported. If necessary, please use the last code that supports VCL version: https://github.com/ying32/govcl/tree/last-vcl-support.")
	}

	return lib
}

func getLibType(lib *dylib.LazyDLL) TLibType {
	proc := lib.NewProc("DGetLibType")
	r, _, _ := proc.Call()
	return TLibType(r)
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibVcl() *dylib.LazyDLL {
	return libvcl
}

func closeLib() {
	if libvcl != nil {
		libvcl.Close()
		libvcl = nil
	}
}
