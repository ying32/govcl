package api

import (
    "unsafe"
)

// MemoryStream_Read 还需要待测试
func MemoryStream_Read(obj uintptr, count int32) (int32, []byte) {
   if count <= 0 {
       return 0, nil
   }
   bs := make([]byte, count)
   r, _, _ := memoryStream_Read.Call(obj, uintptr(unsafe.Pointer(&bs[0])), uintptr(count))
   return int32(r), bs
}

func MemoryStream_Write(obj uintptr, buffer []byte) int32 {
    r, _, _ := memoryStream_Write.Call(obj, uintptr(unsafe.Pointer(&buffer[0])), uintptr(len(buffer)))
    return int32(r)
}

