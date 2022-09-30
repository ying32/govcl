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
)

type importTable struct {
	name string
	addr ProcAddr
}

//func internalGetImportFunc(uiLib DLL, table []importTable, index int) *dylib.LazyProc {
//	item := table[index]
//	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.proc))) == nil {
//		item.proc = uiLib.GetProcAddr(item.name)
//		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&table[index].proc)), unsafe.Pointer(item.proc))
//	}
//	return item.proc
//}

func GetImportFunc(uiLib DLL, index int) ProcAddr {
	item := dllImports[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.addr))) == nil {
		item.addr = uiLib.GetProcAddr(item.name)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllImports[index].addr)), unsafe.Pointer(item.addr))
	}
	return item.addr
}

func GetImportDefFunc(uiLib DLL, index int) ProcAddr {
	item := dllImportDefs[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.addr))) == nil {
		item.addr = uiLib.GetProcAddr(item.name)
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&dllImportDefs[index].addr)), unsafe.Pointer(item.addr))
	}
	return item.addr
}
