//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"testing"
	"unsafe"
)

func TestAll(t *testing.T) {
	var item FuncItem
	t.Log(unsafe.Sizeof(item))
}
