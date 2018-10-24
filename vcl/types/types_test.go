package types

import (
	"fmt"
	"runtime"
	"testing"
	"unsafe"
)

func TestSize(t *testing.T) {
	var val TWMKey
	if runtime.GOARCH == "386" {
		if unsafe.Sizeof(val) == 16 {
			t.Log("OK")
		} else {
			t.Fatal("size: ", unsafe.Sizeof(val))
		}
	} else if runtime.GOARCH == "amd64" {
		if unsafe.Sizeof(val) == 32 {
			t.Log("OK")
		} else {
			t.Fatal("size: ", unsafe.Sizeof(val))
		}
	}

	var guid TGUID
	fmt.Println("guid ssizeOf: ", unsafe.Sizeof(guid))

}
