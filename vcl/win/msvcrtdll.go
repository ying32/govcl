//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package win

import (
	"syscall"
	"unsafe"
)

var (
	// msvcrt.dll
	msvcrtdll = syscall.NewLazyDLL("msvcrt.dll")

	_memcpy  = msvcrtdll.NewProc("memcpy")
	_memset  = msvcrtdll.NewProc("memset")
	_malloc  = msvcrtdll.NewProc("malloc")
	_realloc = msvcrtdll.NewProc("realloc")
	_free    = msvcrtdll.NewProc("free")
	_system  = msvcrtdll.NewProc("system")
)

// Memcpy 内存拷贝
func Memcpy(dest, src, count uintptr) uintptr {
	r, _, _ := _memcpy.Call(dest, src, count)
	return r
}

// Memset 将指定内存区值默认为val中的，一般用于清0也是是 Memset(ptr, 0, size)
func Memset(dest uintptr, val int, count uintptr) uintptr {
	r, _, _ := _memset.Call(dest, uintptr(val), count)
	return r
}

// Malloc 分配内存，此为手动管理，分配后要使用Free释放
func Malloc(size uintptr) uintptr {
	r, _, _ := _malloc.Call(size)
	return r
}

// Realloc 重新分配内存，此为手动管理，分配后要使用Free释放
func Realloc(P, NewSize uintptr) uintptr {
	r, _, _ := _realloc.Call(P, NewSize)
	return r
}

// Free 释放内存
func Free(pBlock uintptr) {
	_free.Call(pBlock)
}

// System 系统命令，现在貌似不管用，以前CGO还管用的呢？Delphi也正常啊，好吧，不知道了。
func System(aCommand string) int {
	if aCommand == "" {
		return -1
	}
	ptr := uintptr(unsafe.Pointer(&UTF8ToANSI(aCommand)[0]))
	r, _, _ := _system.Call(ptr)
	return int(r)
}
