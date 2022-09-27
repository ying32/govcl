//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package dllimports

import (
	"sync"
	"sync/atomic"
	"unsafe"

	"github.com/ying32/dylib"
)

var (
	refs1 sync.Mutex
)

func GetImportFunc(uiLib *dylib.LazyDLL, index int) *dylib.LazyProc {
	item := dllImports[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.proc))) == nil {
		if item.proc == nil {
			refs1.Lock()
			defer refs1.Unlock()
			item.proc = uiLib.NewProc(item.name)
			dllImports[index] = item
		}
	}
	return item.proc
}
