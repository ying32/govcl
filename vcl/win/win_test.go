//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package win

import (
	"testing"
	"unsafe"
)

func TestAll(t *testing.T) {
	var aa TShellExecuteInfo
	t.Log("TShellExecuteInfo size = ", unsafe.Sizeof(aa)) // 60 | 112

	var bb TSystemInfo
	t.Log("TSystemInfo size=", unsafe.Sizeof(bb)) // 36 | 48
}
