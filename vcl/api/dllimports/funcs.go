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

	"github.com/ying32/dylib"
)

var (
	refs1 sync.Mutex
	refs2 sync.Mutex
)

func GetImportFunc(uiLib *dylib.LazyDLL, index int) *dylib.LazyProc {
	item := dllImports[index]
	if item.proc == nil {
		refs1.Lock()
		defer refs1.Unlock()
		item.proc = uiLib.NewProc(item.name)
		dllImports[index] = item
	}
	return item.proc
}

func GetImportDefFunc(uiLib *dylib.LazyDLL, index int) *dylib.LazyProc {
	item := dllImportDefs[index]
	if item.proc == nil {
		refs2.Lock()
		defer refs2.Unlock()
		item.proc = uiLib.NewProc(item.name)
		dllImportDefs[index] = item
	}
	return item.proc
}
