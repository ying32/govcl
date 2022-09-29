//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package api

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/api/dllimports"
)

func MemoryStream_Read(obj uintptr, count int32) (int32, []byte) {
	if count <= 0 {
		return 0, nil
	}
	bs := make([]byte, count)
	r := defSyscallN(dllimports.MEMORYSTREAM_READ, obj, uintptr(unsafe.Pointer(&bs[0])), uintptr(count))
	return int32(r), bs
}

func MemoryStream_Write(obj uintptr, buffer []byte) int32 {
	return int32(defSyscallN(dllimports.MEMORYSTREAM_WRITE, obj, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer))))
}
