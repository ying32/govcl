// +build windows,386

package btmemorymodule

import (
	"syscall"
)

type DWORD = uint32
type Word = uint16
type Byte = uint8
type ULONGLONG = uint64
type Pointer = uintptr
type Integer = int32
type THandle = uintptr
type SIZE_T = uintptr
type HMODULE = uintptr
type LPCWSTR = uintptr
type LongInt = uint32

//func (d DWORD) ToPointer() Pointer {
//	return Pointer(d)
//}

const (
	IMAGE_NUMBEROF_DIRECTORY_ENTRIES = 16

	MEM_RELEASE  = 0x8000
	MEM_RESERVE  = 0x2000
	MEM_COMMIT   = 0x1000
	MEM_DECOMMIT = 0x4000

	PAGE_NOACCESS          = 1
	PAGE_READONLY          = 2
	PAGE_READWRITE         = 4
	PAGE_WRITECOPY         = 8
	PAGE_EXECUTE           = 0x10
	PAGE_EXECUTE_READ      = 0x20
	PAGE_EXECUTE_READWRITE = 0x40
	PAGE_EXECUTE_WRITECOPY = 0x80
	PAGE_NOCACHE           = 0x200

	INVALID_HANDLE_VALUE = 0xFFFFFFFF

	DLL_PROCESS_ATTACH = 1
	DLL_THREAD_ATTACH  = 2
	DLL_THREAD_DETACH  = 3
	DLL_PROCESS_DETACH = 0

	IMAGE_DOS_SIGNATURE = 0x5A4D     // MZ
	IMAGE_NT_SIGNATURE  = 0x00004550 // PE00

	IMAGE_REL_BASED_ABSOLUTE = 0  //无具体含义，用于对齐
	IMAGE_REL_BASED_DIR64    = 10 //用于64位PE文件中的地址修正
	IMAGE_REL_BASED_HIGHLOW  = 3  //该项重定位数据需要被修正，基本都是这种情况

	IMAGE_SIZEOF_SHORT_NAME          = 8
	IMAGE_DIRECTORY_ENTRY_EXPORT     = 0
	IMAGE_SIZEOF_BASE_RELOCATION     = 8
	IMAGE_ORDINAL_FLAG32             = 0x80000000
	IMAGE_DIRECTORY_ENTRY_IMPORT     = 1
	IMAGE_DIRECTORY_ENTRY_BASERELOC  = 5
	IMAGE_SCN_LNK_NRELOC_CVFL        = 0x01000000
	IMAGE_SCN_MEM_DISCARDABLE        = 0x02000000
	IMAGE_SCN_MEM_NOT_CACHED         = 0x04000000
	IMAGE_SCN_MEM_NOT_PAGED          = 0x08000000
	IMAGE_SCN_MEM_NOT_SHARED         = 0x10000000
	IMAGE_SCN_MEM_EXECUTE            = 0x20000000
	IMAGE_SCN_MEM_READ               = 0x40000000
	IMAGE_SCN_MEM_WRITE              = 0x80000000
	IMAGE_SCN_CNT_INITIALIZED_DATA   = 0x00000040
	IMAGE_SCN_CNT_UNINITIALIZED_DATA = 0x00000080
	AGE_DIRECTORY_ENTRY_EXPORT       = 0
)

type TImageFileHeader struct {
	Machine              Word
	NumberOfSections     Word
	TimeDateStamp        DWORD
	PointerToSymbolTable DWORD
	NumberOfSymbols      DWORD
	SizeOfOptionalHeader Word
	Characteristics      Word
}

type TImageDataDirectory struct {
	VirtualAddress DWORD
	Size           DWORD
}

type TImageDosHeader struct {
	e_magic    Word     // Magic number
	e_cblp     Word     // Bytes on last page of file
	e_cp       Word     // Pages in file
	e_crlc     Word     // Relocations
	e_cparhdr  Word     // Size of header in paragraphs
	e_minalloc Word     // Minimum extra paragraphs needed
	e_maxalloc Word     // Maximum extra paragraphs needed
	e_ss       Word     // Initial (relative) SS value
	e_sp       Word     // Initial SP value
	e_csum     Word     // Checksum
	e_ip       Word     // Initial IP value
	e_cs       Word     // Initial (relative) CS value
	e_lfarlc   Word     // File address of relocation table
	e_ovno     Word     // Overlay number
	e_res      [4]Word  // Reserved words
	e_oemid    Word     // OEM identifier (for e_oeminfo)
	e_oeminfo  Word     // OEM information; e_oemid specific
	e_res2     [10]Word // Reserved words
	_lfanew    LongInt  // File address of new exe header
}

type TImageImportByName struct {
	Hint Word
	Name [256]Byte //[1]Byte
}

type TImageExportDirectory struct {
	Characteristics       DWORD
	TimeDateStamp         DWORD
	MajorVersion          Word
	MinorVersion          Word
	Name                  DWORD
	Base                  DWORD
	NumberOfFunctions     DWORD
	NumberOfNames         DWORD
	AddressOfFunctions    DWORD
	AddressOfNames        DWORD
	AddressOfNameOrdinals DWORD
}

type TISHMisc struct {
	PhysicalAddress DWORD
	//case Integer of
	//	0: (PhysicalAddress: DWORD);
	//	1: (VirtualSize: DWORD);
}

type TImageSectionHeader struct {
	Name                 [IMAGE_SIZEOF_SHORT_NAME]Byte
	Misc                 TISHMisc
	VirtualAddress       DWORD
	SizeOfRawData        DWORD
	PointerToRawData     DWORD
	PointerToRelocations DWORD
	PointerToLinenumbers DWORD
	NumberOfRelocations  Word
	NumberOfLinenumbers  Word
	Characteristics      DWORD
}

type TImageImportDescriptor struct {
	//case Byte of
	//0: (Characteristics: DWORD);          // 0 for terminating null import descriptor
	//1: (OriginalFirstThunk: DWORD;        // RVA to original unbound IAT (PIMAGE_THUNK_DATA)
	OriginalFirstThunk DWORD
	TimeDateStamp      DWORD // 0 if not bound,
	// -1 if bound, and real date\time stamp
	//     in IMAGE_DIRECTORY_ENTRY_BOUND_IMPORT (new BIND)
	// O.W. date/time stamp of DLL bound to (Old BIND)

	ForwarderChain DWORD // -1 if no forwarders
	Name           DWORD
	FirstThunk     DWORD // RVA to IAT (if bound this IAT has actual addresses)
}

var (
	kernel32dll = syscall.NewLazyDLL("kernel32.dll")

	_HeapFree       = kernel32dll.NewProc("HeapFree")
	_HeapAlloc      = kernel32dll.NewProc("HeapAlloc")
	_GetProcessHeap = kernel32dll.NewProc("GetProcessHeap")
	_VirtualFree    = kernel32dll.NewProc("VirtualFree")
	_FreeLibrary    = kernel32dll.NewProc("FreeLibrary")
	_lstrlen        = kernel32dll.NewProc("lstrlenA")
	_VirtualAlloc   = kernel32dll.NewProc("VirtualAlloc")
	_VirtualProtect = kernel32dll.NewProc("VirtualProtect")
	_IsBadReadPtr   = kernel32dll.NewProc("IsBadReadPtr")
	_LoadLibrary    = kernel32dll.NewProc("LoadLibraryA") // 使用A不使用那个啥
	_GetProcAddress = kernel32dll.NewProc("GetProcAddress")
)

var (
	msvcrtdll = syscall.NewLazyDLL("msvcrt.dll")
	_memcpy   = msvcrtdll.NewProc("memcpy")
	_memset   = msvcrtdll.NewProc("memset")
	_malloc   = msvcrtdll.NewProc("malloc")
	_realloc  = msvcrtdll.NewProc("realloc")
	_free     = msvcrtdll.NewProc("free")
)

func HeapFree(hHeap THandle, dwFlags DWORD, lpMem Pointer) bool {
	r, _, _ := _HeapFree.Call(hHeap, uintptr(dwFlags), lpMem)
	return r != 0
}

func HeapAlloc(hHeap THandle, dwFlags DWORD, dwBytes SIZE_T) Pointer {
	r, _, _ := _HeapAlloc.Call(hHeap, uintptr(dwFlags), dwBytes)
	return r
}

func GetProcessHeap() THandle {
	r, _, _ := _GetProcessHeap.Call()
	return r
}

func VirtualFree(lpAddress Pointer, dwSize SIZE_T, dwFreeType DWORD) bool {
	r, _, _ := _VirtualFree.Call(lpAddress, dwSize, uintptr(dwFreeType))
	return r != 0
}

func FreeLibrary(hLibModule HMODULE) bool {
	r, _, _ := _FreeLibrary.Call(hLibModule)
	return r != 0
}

func Lstrlen(lpString LPCWSTR) Integer {
	r, _, _ := _lstrlen.Call(lpString)
	return Integer(r)
}

func VirtualAlloc(lpvAddress Pointer, dwSize SIZE_T, flAllocationType, flProtect DWORD) Pointer {
	r, _, _ := _VirtualAlloc.Call(lpvAddress, dwSize, uintptr(flAllocationType), uintptr(flProtect))
	return r
}

func VirtualProtect(lpAddress Pointer, dwSize SIZE_T, flNewProtect DWORD, lpflOldProtect Pointer) bool {
	r, _, _ := _VirtualProtect.Call(lpAddress, dwSize, uintptr(flNewProtect), lpflOldProtect)
	return r != 0
}

func IsBadReadPtr(lp Pointer, ucb uintptr) bool {
	r, _, _ := _IsBadReadPtr.Call(lp, ucb)
	return r != 0
}

func LoadLibrary(lpLibFileName LPCWSTR) HMODULE {
	r, _, _ := _LoadLibrary.Call(lpLibFileName)
	return r
}

func GetProcAddress(hModule HMODULE, lpProcName LPCWSTR) Pointer {
	r, _, _ := _GetProcAddress.Call(hModule, lpProcName)
	return r
}

// -----
func Memcpy(dest, src Pointer, count SIZE_T) Pointer {
	r, _, _ := _memcpy.Call(dest, src, count)
	return r
}

func Memset(dest Pointer, val Integer, count SIZE_T) Pointer {
	r, _, _ := _memset.Call(dest, uintptr(val), count)
	return r
}

func Malloc(size SIZE_T) Pointer {
	r, _, _ := _malloc.Call(size)
	return r
}

func Realloc(P Pointer, NewSize SIZE_T) Pointer {
	r, _, _ := _realloc.Call(P, NewSize)
	return r
}

func Free(pBlock Pointer) {
	_free.Call(pBlock)
}
