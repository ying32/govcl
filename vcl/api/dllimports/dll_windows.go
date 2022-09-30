//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package dllimports

import (
	"syscall"
)

func NewDLL(name string) (DLL, error) {
	h, err := syscall.LoadLibrary(name)
	return DLL(h), err
}

func (h DLL) Release() error {
	if h != 0 {
		return syscall.FreeLibrary(syscall.Handle(h))
	}
	return nil
}

func (h DLL) GetProcAddr(name string) (ProcAddr, error) {
	addr, err := syscall.GetProcAddress(syscall.Handle(h), name)
	return ProcAddr(addr), err
}

func (p ProcAddr) Call(args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	return SyscallN(uintptr(p), args...)
}

// SyscallN 调用一个函数，最大12个参数。
func SyscallN(addr uintptr, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
	if addr == 0 {
		return 0, 0, syscall.ERROR_PROC_NOT_FOUND
	}
	switch len(args) {
	case 0:
		return syscall.Syscall(addr, 0, 0, 0, 0)
	case 1:
		return syscall.Syscall(addr, 1, args[0], 0, 0)
	case 2:
		return syscall.Syscall(addr, 2, args[0], args[1], 0)
	case 3:
		return syscall.Syscall(addr, 3, args[0], args[1], args[2])
	case 4:
		return syscall.Syscall6(addr, 4, args[0], args[1], args[2], args[3], 0, 0)
	case 5:
		return syscall.Syscall6(addr, 5, args[0], args[1], args[2], args[3], args[4], 0)
	case 6:
		return syscall.Syscall6(addr, 6, args[0], args[1], args[2], args[3], args[4], args[5])
	case 7:
		return syscall.Syscall9(addr, 7, args[0], args[1], args[2], args[3], args[4], args[5], args[6], 0, 0)
	case 8:
		return syscall.Syscall9(addr, 8, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], 0)
	case 9:
		return syscall.Syscall9(addr, 9, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8])
	case 10:
		return syscall.Syscall12(addr, 10, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], 0, 0)
	case 11:
		return syscall.Syscall12(addr, 11, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], 0)
	case 12:
		return syscall.Syscall12(addr, 12, args[0], args[1], args[2], args[3], args[4], args[5], args[6], args[7], args[8], args[9], args[10], args[11])
	default:
		return 0, 0, 87 //ERROR_INVALID_PARAMETER = 87
	}
}
