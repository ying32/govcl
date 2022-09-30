//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package dllimports

import (
	"fmt"
	"syscall"
)

func NewDLL(name string) DLL {
	h, err := syscall.LoadLibrary(name)
	if err != nil {
		panic(err)
	}
	return DLL(h)
}

func (h DLL) Release() error {
	if h != 0 {
		return syscall.FreeLibrary(syscall.Handle(h))
	}
	return nil
}

func (h DLL) GetProcAddr(name string) ProcAddr {
	proc, err := syscall.GetProcAddress(syscall.Handle(h), name)
	if err != nil {
		fmt.Println(err)
		return 0
	}
	return ProcAddr(proc)
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	if p == 0 {
		return 0, 0, syscall.ERROR_PROC_NOT_FOUND
	}
	switch len(args) {
	case 0:
		return syscall.Syscall(uintptr(p), 0, 0, 0, 0)
	case 1:
		return syscall.Syscall(uintptr(p), 1, args[0], 0, 0)
	case 2:
		return syscall.Syscall(uintptr(p), 2, args[0], args[1], 0)
	case 3:
		return syscall.Syscall(uintptr(p), 3, args[0], args[1], args[2])
	case 4:
		return syscall.Syscall6(uintptr(p), 4, args[0], args[1], args[2], args[3], 0, 0)
	case 5:
		return syscall.Syscall6(uintptr(p), 5, args[0], args[1], args[2], args[3], args[4], 0)
	case 6:
		return syscall.Syscall6(uintptr(p), 6, args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return syscall.Syscall9(uintptr(p), 7, args[0], args[1], args[2], args[3], args[4], args[5], args[6], 0, 0)
	case 8:
		return syscall.Syscall9(uintptr(p), 8, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], 0)
	case 9:
		return syscall.Syscall9(uintptr(p), 9, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return syscall.Syscall12(uintptr(p), 10, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], 0, 0)
	case 11:
		return syscall.Syscall12(uintptr(p), 11, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], 0)
	case 12:
		return syscall.Syscall12(uintptr(p), 12, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	default:
		return 0, 0, 87 //ERROR_INVALID_PARAMETER = 87
	}

}
