//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !windows

package dllimports

/*
	#cgo LDFLAGS: -ldl
	#include <dlfcn.h>
	#include <limits.h>
	#include <stdlib.h>
	#include <stdint.h>
	#include <stdio.h>
    #include "posix_syscall.h"
*/
import "C"
import (
	"errors"
	"runtime"
	"syscall"
	"unsafe"
)

func NewDLL(name string) (DLL, error) {

	cPath := (*C.char)(C.malloc(C.PATH_MAX + 1))
	defer C.free(unsafe.Pointer(cPath))

	// macOS下强制加载执行文件路径下的liblcl
	if runtime.GOOS == "darwin" {
		name = "@executable_path/" + name
	}
	cRelName := C.CString(name)
	defer C.free(unsafe.Pointer(cRelName))

	var h DLL
	if C.realpath(cRelName, cPath) == nil {
		h = DLL(C.dlopen(cRelName, C.RTLD_LAZY|C.RTLD_GLOBAL))
	} else {
		h = DLL(C.dlopen(cPath, C.RTLD_LAZY|C.RTLD_GLOBAL))
	}
	return h, dlError()
}

func dlError() error {
	err := C.dlerror()
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

func (h DLL) Release() error {
	if h != 0 {
		C.dlclose(unsafe.Pointer(h))
		return dlError()
	}
	return nil
}

func (h DLL) GetProcAddr(name string) (ProcAddr, error) {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	return ProcAddr(C.dlsym(unsafe.Pointer(h), cName)), dlError()
}

func toPtr(p uintptr) C.uintptr_t {
	return C.uintptr_t(p)
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err error) {
	return SyscallN(uintptr(p), args...)
}

// SyscallN 调用一个函数，最大12个参数。
func SyscallN(addr uintptr, args ...uintptr) (r1, r2 uintptr, err error) {
	if addr == 0 {
		return 0, 0, syscall.EINVAL
	}
	var ret C.uint64_t
	switch len(args) {
	case 0:
		ret = C.Syscall0(toPtr(addr))
	case 1:
		ret = C.Syscall1(toPtr(addr), toPtr(args[0]))
	case 2:
		ret = C.Syscall2(toPtr(addr), toPtr(args[0]), toPtr(args[1]))
	case 3:
		ret = C.Syscall3(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]))
	case 4:
		ret = C.Syscall4(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]))
	case 5:
		ret = C.Syscall5(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]))
	case 6:
		ret = C.Syscall6(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]))
	case 7:
		ret = C.Syscall7(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]))
	case 8:
		ret = C.Syscall8(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]))
	case 9:
		ret = C.Syscall9(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]))
	case 10:
		ret = C.Syscall10(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]))
	case 11:
		ret = C.Syscall11(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]), toPtr(args[10]))
	case 12:
		ret = C.Syscall12(toPtr(addr), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]), toPtr(args[10]), toPtr(args[11]))

	default:
		return 0, 0, errors.New("the number of invalid parameters")
	}
	return uintptr(ret), uintptr(ret >> 32), nil
}
