//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

/*
//go:build !windows
// +build !windows
// +build cgo
*/

package dllimports

//
///*
//	#include <dlfcn.h>
//	#include <limits.h>
//	#include <stdlib.h>
//	#include <stdint.h>
//	#include <stdio.h>
//
//    static uint64_t Syscall0(void* addr) {
//		return ((uint64_t(*)())addr)();
//	}
//
//    static uint64_t Syscall1(void* addr, void* p1) {
//		return ((uint64_t(*)(void*))addr)(p1);
//	}
//
//    static uint64_t Syscall2(void* addr, void* p1, void* p2) {
//		return ((uint64_t(*)(void*,void*))addr)(p1, p2);
//	}
//
//    static uint64_t Syscall3(void* addr, void* p1, void* p2, void* p3) {
//		return ((uint64_t(*)(void*,void*,void*))addr)(p1, p2, p3);
//	}
//
//    static uint64_t Syscall4(void* addr, void* p1, void* p2, void* p3, void* p4) {
//		return ((uint64_t(*)(void*,void*,void*,void*))addr)(p1, p2, p3, p4);
//	}
//
//    static uint64_t Syscall5(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5);
//	}
//
//    static uint64_t Syscall6(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6);
//	}
//
//    static uint64_t Syscall7(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*, void*))addr)(p1, p2, p3, p4, p5, p6, p7);
//	}
//
//    static uint64_t Syscall8(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8);
//	}
//
//    static uint64_t Syscall9(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9);
//	}
//
//    static uint64_t Syscall10(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10);
//	}
//
//    static uint64_t Syscall11(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10, void *p11) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11);
//	}
//
//    static uint64_t Syscall12(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12) {
//		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11,p12);
//	}
//
//*/
//import "C"
//
//import (
//	"syscall"
//	"unsafe"
//)
//
//func toPtr(a uintptr) unsafe.Pointer {
//	return unsafe.Pointer(a)
//}
//
//func SyscallN(trap uintptr, args ...uintptr) (r1, r2 uintptr, err syscall.Errno) {
//	if trap == 0 {
//		return 0, 0, syscall.EINVAL
//	}
//	var ret C.uint64_t
//	switch len(args) {
//	case 0:
//		ret = C.Syscall0(toPtr(trap))
//	case 1:
//		ret = C.Syscall1(toPtr(trap), toPtr(args[0]))
//	case 2:
//		ret = C.Syscall2(toPtr(trap), toPtr(args[0]), toPtr(args[1]))
//	case 3:
//		ret = C.Syscall3(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]))
//	case 4:
//		ret = C.Syscall4(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]))
//	case 5:
//		ret = C.Syscall5(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]))
//	case 6:
//		ret = C.Syscall6(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]))
//	case 7:
//		ret = C.Syscall7(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]))
//	case 8:
//		ret = C.Syscall8(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]))
//	case 9:
//		ret = C.Syscall9(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]))
//	case 10:
//		ret = C.Syscall10(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]))
//	case 11:
//		ret = C.Syscall11(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]), toPtr(args[10]))
//	case 12:
//		ret = C.Syscall12(toPtr(trap), toPtr(args[0]), toPtr(args[1]), toPtr(args[2]), toPtr(args[3]), toPtr(args[4]), toPtr(args[5]), toPtr(args[6]), toPtr(args[7]), toPtr(args[8]), toPtr(args[9]), toPtr(args[10]), toPtr(args[11]))
//
//	default:
//		//	panic("Call " + p.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
//	}
//	return uintptr(ret), uintptr(ret >> 32), 0
//}
