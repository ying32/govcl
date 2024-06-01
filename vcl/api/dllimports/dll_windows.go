// ----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
// ----------------------------------------

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
	return syscall.SyscallN(addr, args...)
}
