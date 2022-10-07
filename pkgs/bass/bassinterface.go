package bass

import (
	"math"
	"runtime"

	"unsafe"

	"github.com/ying32/dylib"
	"github.com/ying32/govcl/vcl/types"
)

var (
	libbass = dylib.NewLazyDLL(dylibName())

	_BASS_GetVersion       = libbass.NewProc("BASS_GetVersion")
	_BASS_Init             = libbass.NewProc("BASS_Init")
	_BASS_Free             = libbass.NewProc("BASS_Free")
	_BASS_StreamCreateFile = libbass.NewProc("BASS_StreamCreateFile")
	_BASS_StreamFree       = libbass.NewProc("BASS_StreamFree")
	_BASS_ChannelPlay      = libbass.NewProc("BASS_ChannelPlay")
	_BASS_ChannelStop      = libbass.NewProc("BASS_ChannelStop")
	_BASS_ChannelPause     = libbass.NewProc("BASS_ChannelPause")

	// 这个因为返回的是double，暂时没找到可解决的办法，只能自己伪造一个了，伪造的函数是猜测的
	// 因为不知道怎么才好吧
	_BASS_ChannelBytes2Seconds = libbass.NewProc("BASS_ChannelBytes2Seconds")
	_BASS_ChannelGetLength     = libbass.NewProc("BASS_ChannelGetLength")
	_BASS_ChannelGetPosition   = libbass.NewProc("BASS_ChannelGetPosition")
	_BASS_ChannelSetAttribute  = libbass.NewProc("BASS_ChannelSetAttribute")
	_BASS_ChannelGetAttribute  = libbass.NewProc("BASS_ChannelGetAttribute")
	_BASS_ErrorGetCode         = libbass.NewProc("BASS_ErrorGetCode")
	_BASS_ChannelSetPosition   = libbass.NewProc("BASS_ChannelSetPosition")
	_BASS_ChannelSeconds2Bytes = libbass.NewProc("BASS_ChannelSeconds2Bytes")
)

const (
	BASSVERSION     = 0x204 // API version
	BASSVERSIONTEXT = "2.4"
)

func dylibName() string {
	if runtime.GOOS == "windows" {
		return "bass.dll"
	} else if runtime.GOOS == "linux" {
		return "libbass.so"
	} else if runtime.GOOS == "darwin" {
		return "libbass.dylib"
	}
	return ""
}

type HSTREAM uint32

func cbool(b bool) uintptr {
	if b {
		return 1
	}
	return 0
}

func BASS_GetVersion() uint32 {
	r, _, _ := _BASS_GetVersion.Call()
	return uint32(r)
}

func BASS_Init(device int, freq, flags uint32, win types.HWND, dsguid uintptr) bool {
	r, _, _ := _BASS_Init.Call(uintptr(device), uintptr(freq), uintptr(flags), uintptr(win), uintptr(dsguid))
	return r != 0
}

func BASS_Free() bool {
	r, _, _ := _BASS_Free.Call()
	return r != 0
}

func BASS_StreamFree(handle HSTREAM) bool {
	r, _, _ := _BASS_StreamFree.Call(uintptr(handle))
	return r != 0
}

func BASS_ChannelPlay(handle HSTREAM, restart bool) bool {
	r, _, _ := _BASS_ChannelPlay.Call(uintptr(handle), cbool(restart))
	return r != 0
}

func BASS_ChannelStop(handle HSTREAM) bool {
	r, _, _ := _BASS_ChannelStop.Call(uintptr(handle))
	return r != 0
}

func BASS_ChannelPause(handle HSTREAM) bool {
	r, _, _ := _BASS_ChannelPause.Call(uintptr(handle))
	return r != 0
}

func BASS_ChannelGetLength(handle HSTREAM, mode uint32) uint64 {
	r1, r2, _ := _BASS_ChannelGetLength.Call(uintptr(handle), uintptr(mode))
	return ToUInt64(r1, r2)
}

func BASS_ChannelGetPosition(handle HSTREAM, mode uint32) uint64 {
	r1, r2, _ := _BASS_ChannelGetPosition.Call(uintptr(handle), uintptr(mode))
	return ToUInt64(r1, r2)
}

func BASS_ChannelSetAttribute(handle HSTREAM, attrib uint32, value float32) bool {
	r, _, _ := _BASS_ChannelSetAttribute.Call(uintptr(handle), uintptr(attrib), uintptr(math.Float32bits(value)))
	return r != 0
}

func BASS_ChannelGetAttribute(handle HSTREAM, attrib uint32, value *float32) bool {
	r, _, _ := _BASS_ChannelGetAttribute.Call(uintptr(handle), uintptr(attrib), uintptr(unsafe.Pointer(value)))
	return r != 0
}

func BASS_ErrorGetCode() int {
	r, _, _ := _BASS_ErrorGetCode.Call()
	return int(r)
}

// 伪造的，经过简单的测试出一个值

//func BASS_ChannelBytes2Seconds(handle HSTREAM, pos uint64) float64 {
//	return float64(pos) / 176400.0
//}

//func BASS_ChannelSeconds2Bytes(handle HSTREAM, pos float64) uint64 {
//	return uint64(pos * 176400)
//}
