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
	"fmt"
	"runtime"
	"syscall"
	"unsafe"
)

func NewDLL(name string) DLL {

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
	err := h.dlError()
	if err != nil {
		panic(err)
	}
	return h
}

func (h DLL) dlError() error {
	err := C.dlerror()
	if err != nil {
		return errors.New(C.GoString(err))
	}
	return nil
}

func (h DLL) Release() error {
	if h != 0 {
		C.dlclose(unsafe.Pointer(h))
		return h.dlError()
	}
	return nil
}

func (h DLL) GetProcAddr(name string) ProcAddr {
	cName := C.CString(name)
	defer C.free(unsafe.Pointer(cName))
	proc := C.dlsym(unsafe.Pointer(h), cName)
	if err := h.dlError(); err != nil {
		fmt.Println(err)
		return 0
	}
	return ProcAddr(proc)
}

func toPtr(p ProcAddr) unsafe.Pointer {
	return unsafe.Pointer(p)
}

func toAPtr(p uintptr) unsafe.Pointer {
	return unsafe.Pointer(p)
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err error) {
	if p == 0 {
		return 0, 0, syscall.EINVAL
	}
	var ret C.uint64_t
	switch len(args) {
	case 0:
		ret = C.Syscall0(toPtr(p))
	case 1:
		ret = C.Syscall1(toPtr(p), toAPtr(args[0]))
	case 2:
		ret = C.Syscall2(toPtr(p), toAPtr(args[0]), toAPtr(args[1]))
	case 3:
		ret = C.Syscall3(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]))
	case 4:
		ret = C.Syscall4(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]))
	case 5:
		ret = C.Syscall5(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]))
	case 6:
		ret = C.Syscall6(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]))
	case 7:
		ret = C.Syscall7(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]))
	case 8:
		ret = C.Syscall8(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]), toAPtr(args[7]))
	case 9:
		ret = C.Syscall9(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]), toAPtr(args[7]), toAPtr(args[8]))
	case 10:
		ret = C.Syscall10(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]), toAPtr(args[7]), toAPtr(args[8]), toAPtr(args[9]))
	case 11:
		ret = C.Syscall11(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]), toAPtr(args[7]), toAPtr(args[8]), toAPtr(args[9]), toAPtr(args[10]))
	case 12:
		ret = C.Syscall12(toPtr(p), toAPtr(args[0]), toAPtr(args[1]), toAPtr(args[2]), toAPtr(args[3]), toAPtr(args[4]), toAPtr(args[5]), toAPtr(args[6]), toAPtr(args[7]), toAPtr(args[8]), toAPtr(args[9]), toAPtr(args[10]), toAPtr(args[11]))

	default:
		return 0, 0, errors.New("the number of invalid parameters")
	}
	return uintptr(ret), uintptr(ret >> 32), nil
}
