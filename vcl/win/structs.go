// +build windows

package win

type TRGBQuad struct {
	RgbBlue     uint8
	RgbGreen    uint8
	RgbRed      uint8
	RgbReserved uint8
}

type TBitmapInfoHeader struct {
	BiSize          uint32
	BiWidth         uint32
	BiHeight        uint32
	BiPlanes        uint16
	BiBitCount      uint16
	BiCompression   uint32
	BiSizeImage     uint32
	BiXPelsPerMeter uint32
	BiYPelsPerMeter uint32
	BiClrUsed       uint32
	BiClrImportant  uint32
}

type TBitmapInfo struct {
	BmiHeader TBitmapInfoHeader
	BmiColors [1]TRGBQuad
}

type TBlendFunction struct {
	BlendOp             uint8
	BlendFlags          uint8
	SourceConstantAlpha uint8
	AlphaFormat         uint8
}

type TSystemInfo struct {

	//0: (
	//dwOemId: DWORD);
	//1: (
	ProcessorArchitecture     uint16
	Reserved                  uint16
	PageSize                  uint32
	MinimumApplicationAddress uintptr
	MaximumApplicationAddress uintptr
	ActiveProcessorMask       uintptr
	NumberOfProcessors        uint32
	ProcessorType             uint32
	AllocationGranularity     uint32
	ProcessorLevel            uint16
	ProcessorRevision         uint16
}

type TSecurityAttributes struct {
	nLength              uint32
	lpSecurityDescriptor uintptr
	bInheritHandle       bool // BOOL
}
