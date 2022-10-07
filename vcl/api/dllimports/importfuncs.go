//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package dllimports

import (
	"fmt"
	"sync/atomic"
	"unsafe"
)

type importTable struct {
	name string
	addr ProcAddr
}

func internalGetImportFunc(uiLib DLL, table []importTable, index int) ProcAddr {
	item := table[index]
	if atomic.LoadPointer((*unsafe.Pointer)(unsafe.Pointer(&item.addr))) == nil {
		var err error
		item.addr, err = uiLib.GetProcAddr(item.name)
		if err != nil {
			fmt.Println(err)
			return 0
		}
		atomic.StorePointer((*unsafe.Pointer)(unsafe.Pointer(&table[index].addr)), unsafe.Pointer(item.addr))
	}
	return item.addr
}

func GetImportFunc(uiLib DLL, index int) ProcAddr {
	return internalGetImportFunc(uiLib, dllImports, index)
}

func GetImportDefFunc(uiLib DLL, index int) ProcAddr {
	return internalGetImportFunc(uiLib, dllImportDefs, index)
}
