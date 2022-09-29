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
	"unsafe"

	"github.com/ying32/govcl/pkgs/libname"

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

// Load liblcl
func loadUILib() *dylib.LazyDLL {
	libName := getDLLName()
	// 如果支持运行时释放，则使用此种方法
	if support, newDLLPath := checkAndReleaseDLL(); support {
		libName = newDLLPath
	} else {
		if libname.LibName != "" {
			libName = libname.LibName
		}
	}
	lib := dylib.NewLazyDLL(libName)
	err := lib.Load()
	if err != nil {
		panic(err)
	}
	return lib
}

func closeLib() {
	if uiLib != nil {
		uiLib.Close()
		uiLib = nil
	}
}

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

func syscallGetTextBuf(trap int, obj uintptr, buffer *string, bufSize uintptr) uintptr {
	if buffer == nil || bufSize == 0 {
		return 0
	}
	strPtr := make([]uint8, bufSize+1)
	sLen := syscallN(trap, obj, uintptr(unsafe.Pointer(&strPtr[0])), bufSize)
	if sLen > 0 {
		*buffer = string(strPtr[:sLen])
	}

	/*
		strPtr := getBuff(int32(BufSize))
		result := syscallN(trap, obj, getBuffPtr(strPtr), BufSize)
		getTextBuf(strPtr, Buffer, int(result))
	*/
	return sLen
}

func defSyscallN(trap int, args ...uintptr) uintptr {
	r1, _, _ := dllimports.GetImportDefFunc(uiLib, trap).Call(args...)
	return r1
}

func newDLLProc(name string) *dylib.LazyProc {
	return uiLib.NewProc(name)
}
