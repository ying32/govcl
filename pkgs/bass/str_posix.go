// +build !windows

package bass

import "unsafe"

func cstr(str string) uintptr {
	return uintptr(unsafe.Pointer(&([]byte(str + "\x00")[0])))
}
