// +build linux,cgo darwin,cgo

package vcl

// extern void* doCallbackProc(void* f, void* args, long argcount);
// static void* doGetAddr() {
//    return &doCallbackProc;
// }
import "C"

import (
	"unsafe"
)

//export doCallbackProc
func doCallbackProc(f unsafe.Pointer, args unsafe.Pointer, argcount C.long) unsafe.Pointer {
	callbackProc(uintptr(f), uintptr(args), int(argcount))
	return unsafe.Pointer(uintptr(0))
}

var (
	callbackStdcall = uintptr(C.doGetAddr())
)
