//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !windows
// +build cgo

package vcl

//// #cgo darwin CFLAGS: -mmacosx-version-min=10.10 -DMACOSX_DEPLOYMENT_TARGET=10.10
// #cgo darwin CFLAGS: -mmacosx-version-min=10.10
// #cgo darwin LDFLAGS: -mmacosx-version-min=10.10
// #include <stdint.h>
//
// extern void* doEventCallbackProc(uintptr_t f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(uintptr_t f, void* msg);
// static void* doGetMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
//
// extern void* doThreadSyncCallbackProc();
// static void* doGetThreadSyncCallbackAddr() {
//    return &doThreadSyncCallbackProc;
// }
//
// extern void* doRequestCallCreateParamsCallbackProc(uintptr_t ptr, void* params);
// static void* doRequestCallCreateParamsCallbackAddr() {
//    return &doRequestCallCreateParamsCallbackProc;
// }
//
// extern void* doRemoveEventCallbackProc(uintptr_t ptr);
// static void* doRemoveEventCallbackAddr() {
//    return &doRemoveEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f C.uintptr_t, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(uintptr(f), uintptr(args), int(argcount))
	return nullptr
}

//export doMessageCallbackProc
func doMessageCallbackProc(f C.uintptr_t, msg unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(uintptr(f), uintptr(msg))
	return nullptr
}

//export doThreadSyncCallbackProc
func doThreadSyncCallbackProc() unsafe.Pointer {
	threadSyncCallbackProc()
	return nullptr
}

//export doRequestCallCreateParamsCallbackProc
func doRequestCallCreateParamsCallbackProc(ptr C.uintptr_t, params unsafe.Pointer) unsafe.Pointer {
	requestCallCreateParamsCallbackProc(uintptr(ptr), uintptr(params))
	return nullptr
}

//export doRemoveEventCallbackProc
func doRemoveEventCallbackProc(ptr C.uintptr_t) unsafe.Pointer {
	removeEventCallbackProc(uintptr(ptr))
	return nullptr
}

var (
	eventCallback                   = uintptr(C.doGetEventCallbackAddr())
	messageCallback                 = uintptr(C.doGetMessageCallbackAddr())
	threadSyncCallback              = uintptr(C.doGetThreadSyncCallbackAddr())
	requestCallCreateParamsCallback = uintptr(C.doRequestCallCreateParamsCallbackAddr())
	removeEventCallback             = uintptr(C.doRemoveEventCallbackAddr())
)
