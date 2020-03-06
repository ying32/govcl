package bass

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/win"
)

func cstr(str string) uintptr {
	return uintptr(unsafe.Pointer(&win.UTF8ToANSI(str)[0]))
	//return uintptr(unsafe.Pointer(&([]byte(str + "\x00")[0])))
}
