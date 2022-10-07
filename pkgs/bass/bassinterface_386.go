package bass

import (
	"math"

	"github.com/ying32/dylib/floatpatch"
)

func ToUInt64(r1, r2 uintptr) uint64 {
	ret := uint64(r2)
	ret = uint64(ret<<32) + uint64(r1)
	return ret
}

func UInt64To(val uint64) (uintptr, uintptr) {
	return uintptr(uint32(val)), uintptr(uint32(val >> 32))
}

func BASS_StreamCreateFile(mem bool, file string, offset, length uint64, flags uint32) HSTREAM {
	offset1, offset2 := UInt64To(offset)
	length1, length2 := UInt64To(length)
	r, _, _ := _BASS_StreamCreateFile.Call(cbool(mem), cstr(file), offset1, offset2, length1, length2, uintptr(flags))
	return HSTREAM(r)
}

func BASS_ChannelSetPosition(handle HSTREAM, pos uint64, mode uint32) bool {
	pos1, pos2 := UInt64To(pos)
	r, _, _ := _BASS_ChannelSetPosition.Call(uintptr(handle), pos1, pos2, uintptr(mode))
	return r != 0
}

func BASS_ChannelBytes2Seconds(handle HSTREAM, pos uint64) float64 {
	pos1, pos2 := UInt64To(pos)
	_BASS_ChannelBytes2Seconds.Call(uintptr(handle), pos1, pos2)
	return floatpatch.Getfloat64()
}

func BASS_ChannelSeconds2Bytes(handle HSTREAM, pos float64) uint64 {
	pos1, pos2 := UInt64To(math.Float64bits(pos))
	r1, r2, _ := _BASS_ChannelSeconds2Bytes.Call(uintptr(handle), pos1, pos2)
	return ToUInt64(r1, r2)
}
