//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"runtime"
	"sync"

	"github.com/ying32/dylib"
)

var (

	// 专用于判断是否为Windows系统
	IsWindows = runtime.GOOS == "windows"

	// 全局导入库
	libvcl = loadUILib()

	// 导出的DLL，考虑到导入的函数太多了，导致go无法编译通过
	// 只能动态作，这样可能牺牲一点性能吧，但文件大小会减小几M左右吧。
	functions sync.Map
)

var (
	platformExtNames = map[string]string{
		"windows": ".dll",
		"linux":   ".so",
		"darwin":  ".dylib",
	}
)

func getDLLName() string {
	libName := "liblcl"
	if ext, ok := platformExtNames[runtime.GOOS]; ok {
		return libName + ext
	}
	return libName
}

func getLazyProc(name string) *dylib.LazyProc {
	if val, ok := functions.Load(name); !ok {
		proc := libvcl.NewProc(name)
		functions.Store(name, proc)
		return proc
	} else {
		return val.(*dylib.LazyProc)
	}
}
