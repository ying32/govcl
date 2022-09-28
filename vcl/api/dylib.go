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

	"github.com/ying32/govcl/vcl/api/dllimports"

	"github.com/ying32/dylib"
)

var (

	// 全局导入库
	uiLib = loadUILib()
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

func syscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetImportFunc(uiLib, trap).Call(args...)
	return r1
}

func syscallGetTextBuf(trap int, obj uintptr, Buffer *string, BufSize uintptr) uintptr {
	if Buffer == nil || BufSize == 0 {
		return 0
	}
	strPtr := getBuff(int32(BufSize))
	result := syscallN(trap, obj, getBuffPtr(strPtr), BufSize)
	getTextBuf(strPtr, Buffer, int(result))
	return result
}

func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetImportDefFunc(uiLib, trap).Call(args...)
	return r1
}

func newDLLProc(name string) *dylib.LazyProc {
	return uiLib.NewProc(name)
}
