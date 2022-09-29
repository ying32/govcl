//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package dllimports

import (
	"sync/atomic"
	"unsafe"

	"github.com/ying32/dylib"
)

type importTable struct {
	name string
	proc *dylib.LazyProc
}

//func internalGetImportFunc(uiLib *dylib.LazyDLL, table []importTable, index int) *dylib.LazyProc {
//	item := table[index]
//	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.proc))) == nil {
//		item.proc = uiLib.NewProc(item.name)
//		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&table[index].proc)), unsafe.Pointer(item.proc))
//	}
//	return item.proc
//}

func GetImportFunc(uiLib *dylib.LazyDLL, index int) *dylib.LazyProc {
	item := dllImports[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.proc))) == nil {
		item.proc = uiLib.NewProc(item.name)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllImports[index].proc)), unsafe.Pointer(item.proc))
	}
	return item.proc
}

func GetImportDefFunc(uiLib *dylib.LazyDLL, index int) *dylib.LazyProc {
	item := dllImportDefs[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.proc))) == nil {
		item.proc = uiLib.NewProc(item.name)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllImportDefs[index].proc)), unsafe.Pointer(item.proc))
	}
	return item.proc
}
