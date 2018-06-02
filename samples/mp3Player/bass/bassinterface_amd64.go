package bass

func ToUInt64(r1, r2 uintptr) uint64 {
	return uint64(r1)
}

//func ToUInt64(r1, r2 uintptr) uint64 {
//	ret := uint64(r2)
//	ret = uint64(ret<<32) + uint64(r1)
//	return ret
//}

//func UInt64To(val uint64) (uintptr, uintptr) {
//	return uintptr(val), 0
//}

func BASS_StreamCreateFile(mem bool, file string, offset, length uint64, flags uint32) HSTREAM {
	r, _, _ := _BASS_StreamCreateFile.Call(cbool(mem), cstr(file), uintptr(offset), uintptr(length), uintptr(flags))
	return HSTREAM(r)
}

func BASS_ChannelSetPosition(handle HSTREAM, pos uint64, mode uint32) bool {
	r, _, _ := _BASS_ChannelSetPosition.Call(uintptr(handle), uintptr(pos), uintptr(mode))
	return r != 0
}

//func BASS_ChannelBytes2Seconds(handle HSTREAM, pos uint64) float64 {
//	r, r2, _ := _BASS_ChannelBytes2Seconds.Call(uintptr(handle), uintptr(pos))
//	var result float64
//	*(*uint64)(unsafe.Pointer(&result)) = uint64(r)
//	fmt.Println("result:", r, r2, result, ", pos:", pos)
//	return result
//}

//func BASS_ChannelSeconds2Bytes(handle HSTREAM, pos float64) uint64 {
//	r, _, _ := _BASS_ChannelSeconds2Bytes.Call(uintptr(handle), uintptr(pos))
//	return uint64(r)
//}
