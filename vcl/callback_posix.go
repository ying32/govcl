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
//
// extern void* doEventCallbackProc(void* f, void* args, long argcount);
// static void* doGetEventCallbackAddr() {
//    return &doEventCallbackProc;
// }
//
// extern void* doMessageCallbackProc(void* f, void* msg);
// static void* doGetMessageCallbackAddr() {
//    return &doMessageCallbackProc;
// }
//
// extern void* doThreadSyncCallbackProc();
// static void* doGetThreadSyncCallbackAddr() {
//    return &doThreadSyncCallbackProc;
// }
//
// extern void* doRequestCallCreateParamsCallbackProc(void* ptr, void* params);
// static void* doRequestCallCreateParamsCallbackAddr() {
//    return &doRequestCallCreateParamsCallbackProc;
// }
//
// extern void* doRemoveEventCallbackProc(void* ptr);
// static void* doRemoveEventCallbackAddr() {
//    return &doRemoveEventCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doEventCallbackProc
func doEventCallbackProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	eventCallbackProc(f, uintptr(args), int(argcount))
	return nullptr
}

//export doMessageCallbackProc
func doMessageCallbackProc(f unsafe.Pointer, msg unsafe.Pointer) unsafe.Pointer {
	messageCallbackProc(f, uintptr(msg))
	return nullptr
}

//export doThreadSyncCallbackProc
func doThreadSyncCallbackProc() unsafe.Pointer {
	threadSyncCallbackProc()
	return nullptr
}

//export doRequestCallCreateParamsCallbackProc
func doRequestCallCreateParamsCallbackProc(ptr unsafe.Pointer, params unsafe.Pointer) unsafe.Pointer {
	requestCallCreateParamsCallbackProc(uintptr(ptr), uintptr(params))
	return nullptr
}

//export doRemoveEventCallbackProc
func doRemoveEventCallbackProc(ptr unsafe.Pointer) unsafe.Pointer {
	removeEventCallbackProc(ptr)
	return nullptr
}

var (
	eventCallback                   = uintptr(C.doGetEventCallbackAddr())
	messageCallback                 = uintptr(C.doGetMessageCallbackAddr())
	threadSyncCallback              = uintptr(C.doGetThreadSyncCallbackAddr())
	requestCallCreateParamsCallback = uintptr(C.doRequestCallCreateParamsCallbackAddr())
	removeEventCallback             = uintptr(C.doRemoveEventCallbackAddr())
)
