// +build windows,386

/*
  此项功能稳定性还有待观查。

  作者：ying32，本为govcl项目中的一个dylib扩展包
  https://github.com/ying32/govcl

  这是一个移植自delphi BTMemoryModule.pas的，暂时只支持32bit dll， 最原始版本的c++代码
  已经支持64bit，暂时还没有想法去做移植。

  原作者信息：
    *- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -*
    * Memory DLL loading code (32bit)                                         *
    *- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -*
    * Delphi BTMemoryModule 0.0.4                                             *
    * Copyright (c) 2005-2010 by Martin Offenwanger / coder@dsplayer.de       *
    * http://www.dsplayer.de                                                  *
    *- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -*
  他也是移植自这里：
    * BTMemoryModule originally is a plain pascal port from c code            *
    *- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -*
    * Original C Code Copyright (c) 2004- 2005 by Joachim Bauch               *
    * mail@joachim-bauch.de                                                   *
    * http://www.joachim-bauch.de/tutorials/loading-a-dll-from-memory/        *
    *- - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - - -*

*/

package btmemorymodule

import (
	"runtime"
	"syscall"
	"unsafe"
)

type TBTMemoryModule struct {
	headers     *TImageNtHeaders
	codeBase    Pointer
	modules     Pointer
	numModules  int
	initialized bool
	isValid     bool
}

func (m *TBTMemoryModule) this() Pointer {
	return Pointer(unsafe.Pointer(m))
}

func (m *TBTMemoryModule) callDLLEntry(fdwReason DWORD) bool {
	DllEntry := m.codeBase + Pointer(m.headers.OptionalHeader.AddressOfEntryPoint)
	r, _, _ := syscall.Syscall(DllEntry, 3, m.codeBase, uintptr(fdwReason), 0)
	return r != 0
}

func (m *TBTMemoryModule) IsValid() bool {
	return m.isValid
}

var (
	lastErrStr string
)

type TImageBaseRelocation struct {
	VirtualAddress DWORD
	SizeOfBlock    DWORD
}

//func GetFieldOffset(Struc, Field interface{}) int {
//	//return UInt64(@Field) - UInt64(@Struc);
//	return 0
//}

func GetHeaderDictionary(module *TBTMemoryModule, idx Integer) *TImageDataDirectory {
	return (*TImageDataDirectory)(unsafe.Pointer(&(module.headers.OptionalHeader.DataDirectory[idx])))
}

func GetImageFirstSection(NtHeader *TImageNtHeaders) *TImageSectionHeader {
	//fieldOffset := Pointer(unsafe.Pointer(&NtHeader.OptionalHeader)) - Pointer(unsafe.Pointer(&(*((*TImageNtHeaders)(unsafe.Pointer(NtHeader))))))
	fieldOffset := Pointer(unsafe.Pointer(&NtHeader.OptionalHeader)) - Pointer(unsafe.Pointer(NtHeader))
	return (*TImageSectionHeader)(unsafe.Pointer(Pointer(unsafe.Pointer(NtHeader)) +
		fieldOffset +
		Pointer(NtHeader.FileHeader.SizeOfOptionalHeader)))
}

func GetImageSnapByOrdinal(Ordinal Pointer) bool {
	return Ordinal&IMAGE_ORDINAL_FLAG32 != 0
}

func GetImageOrdinal(Ordinal Pointer) Word {
	return Word(Ordinal & 0xFFFF)
}

func CopySections(fp_data Pointer, f_old_headers TImageNtHeaders, fp_module *TBTMemoryModule) {
	lp_section := GetImageFirstSection(fp_module.headers)
	for i := 0; i < int(fp_module.headers.FileHeader.NumberOfSections); i++ {
		// section doesn't contain data in the dll itself, but may define
		// uninitialized data
		if lp_section.SizeOfRawData == 0 {
			l_size := f_old_headers.OptionalHeader.SectionAlignment
			if l_size > 0 {
				lp_dest := VirtualAlloc(fp_module.codeBase+Pointer(lp_section.VirtualAddress), SIZE_T(l_size), MEM_COMMIT, PAGE_READWRITE)

				//lp_dest = fp_module.codeBase + Pointer(lp_section.VirtualAddress)
				// NOTE: On 64bit systems we truncate to 32bit here but expand
				// again later when "PhysicalAddress" is used.
				//section.Misc.PhysicalAddress = (DWORD) ((uintptr_t) dest & 0xffffffff);
				lp_section.Misc.PhysicalAddress = DWORD(lp_dest & 0xffffffff)
				Memset(lp_dest, 0, SIZE_T(l_size))
			}
			lp_section = (*TImageSectionHeader)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_section)) + Pointer(unsafe.Sizeof(*lp_section))))
			// Continue with the nex loop
			continue
		}
		// commit memory block and copy data from dll
		lp_dest := VirtualAlloc(fp_module.codeBase+Pointer(lp_section.VirtualAddress), SIZE_T(lp_section.SizeOfRawData), MEM_COMMIT, PAGE_READWRITE)
		Memcpy(lp_dest, fp_data+Pointer(lp_section.PointerToRawData), SIZE_T(lp_section.SizeOfRawData))
		lp_section.Misc.PhysicalAddress = DWORD(lp_dest & 0xffffffff)

		lp_section = (*TImageSectionHeader)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_section)) + Pointer(unsafe.Sizeof(*lp_section))))
	}
}

func PerformBaseRelocation(f_module *TBTMemoryModule, f_delta int) {
	lp_directory := GetHeaderDictionary(f_module, IMAGE_DIRECTORY_ENTRY_BASERELOC)
	if lp_directory.Size > 0 {
		lp_relocation := (*TImageBaseRelocation)(unsafe.Pointer(f_module.codeBase + Pointer(lp_directory.VirtualAddress)))
		for lp_relocation.VirtualAddress > 0 {
			lp_dest := f_module.codeBase + Pointer(lp_relocation.VirtualAddress)
			lp_relInfo := Pointer(unsafe.Pointer(lp_relocation)) + IMAGE_SIZEOF_BASE_RELOCATION
			for i := 0; i < int((lp_relocation.SizeOfBlock-IMAGE_SIZEOF_BASE_RELOCATION)/2); i++ {
				val := (*(*Word)(unsafe.Pointer(lp_relInfo)))
				// the upper 4 bits define the type of relocation
				l_type := val >> 12
				// the lower 12 bits define the offset
				l_offset := val & 0xFFF
				switch l_type {
				case IMAGE_REL_BASED_ABSOLUTE:
					// skip relocation
					break
				case IMAGE_REL_BASED_HIGHLOW:
					// change complete 32 bit address
					v1 := *(*DWORD)(unsafe.Pointer(lp_dest + Pointer(l_offset)))
					*(*DWORD)(unsafe.Pointer(lp_dest + Pointer(l_offset))) = v1 + DWORD(f_delta)
				case IMAGE_REL_BASED_DIR64:
					if runtime.GOARCH == "amd64" {
						v1 := *(*ULONGLONG)(unsafe.Pointer(lp_dest + Pointer(l_offset)))
						*(*ULONGLONG)(unsafe.Pointer(lp_dest + Pointer(l_offset))) = v1 + ULONGLONG(f_delta)
					}
				default:

				}
				lp_relInfo += 2
			}
			lp_relocation = (*TImageBaseRelocation)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_relocation)) + Pointer(lp_relocation.SizeOfBlock)))
		}
	}
}

func BuildImportTable(fp_module *TBTMemoryModule) bool {
	Result := true
	lp_directory := GetHeaderDictionary(fp_module, IMAGE_DIRECTORY_ENTRY_IMPORT)
	if lp_directory.Size > 0 {
		lp_importDesc := (*TImageImportDescriptor)(unsafe.Pointer(fp_module.codeBase + Pointer(lp_directory.VirtualAddress)))
		for !IsBadReadPtr(Pointer(unsafe.Pointer(lp_importDesc)), uintptr(unsafe.Sizeof(*lp_importDesc))) && lp_importDesc.Name != 0 {
			nameAddr := fp_module.codeBase + Pointer(lp_importDesc.Name)
			l_handle := LoadLibrary(nameAddr)
			if l_handle == INVALID_HANDLE_VALUE {
				nameLen := Lstrlen(nameAddr)
				buff := make([]byte, nameLen)
				Memcpy(Pointer(unsafe.Pointer(&buff[0])), nameAddr, SIZE_T(nameLen))
				lastErrStr = "BuildImportTable: can't load library: " + string(buff)
				return false
			}
			if fp_module.modules == 0 {
				fp_module.modules = Malloc(1)
			}
			var newSize SIZE_T
			newSize = SIZE_T(fp_module.numModules+1) * SIZE_T(unsafe.Sizeof(newSize))
			fp_module.modules = Realloc(fp_module.modules, newSize)
			if fp_module.modules == 0 {
				lastErrStr = "BuildImportTable: ReallocMemory failed"
				return false
			}
			// module->modules[module->numModules++] = handle;
			//l_temp := (SizeOf(Cardinal) * (fp_module^.numModules));
			temp := 4 * fp_module.numModules
			fp_module.modules += Pointer(temp)
			*((*Pointer)(unsafe.Pointer(fp_module.modules))) = l_handle
			fp_module.modules -= Pointer(temp)
			fp_module.numModules++ // = fp_module.numModules + 1
			var lp_thunkRef Pointer
			if lp_importDesc.OriginalFirstThunk != 0 {
				lp_thunkRef = fp_module.codeBase + Pointer(lp_importDesc.OriginalFirstThunk)
			} else {
				lp_thunkRef = fp_module.codeBase + Pointer(lp_importDesc.FirstThunk)
			}
			lpfuncRef := fp_module.codeBase + Pointer(lp_importDesc.FirstThunk)
			for *((*Pointer)(unsafe.Pointer(lp_thunkRef))) != 0 {
				thunkRefVal := *((*Pointer)(unsafe.Pointer(lp_thunkRef)))
				if GetImageSnapByOrdinal(thunkRefVal) {
					*((*Pointer)(unsafe.Pointer(lpfuncRef))) = GetProcAddress(l_handle, LPCWSTR(GetImageOrdinal(thunkRefVal)))
				} else {
					var l_thunkData TImageImportByName
					Memcpy(Pointer(unsafe.Pointer(&l_thunkData)), fp_module.codeBase+Pointer(thunkRefVal), SIZE_T(unsafe.Sizeof(l_thunkData)))
					*((*Pointer)(unsafe.Pointer(lpfuncRef))) = GetProcAddress(l_handle, LPCWSTR(unsafe.Pointer(&l_thunkData.Name[0])))
				}
				if *((*Pointer)(unsafe.Pointer(lpfuncRef))) == 0 {
					lastErrStr = "BuildImportTable: GetProcAddress failed"
					FreeLibrary(l_handle)
					Result = false
					break
				}
				var nSize Pointer
				lpfuncRef += unsafe.Sizeof(nSize)
				lp_thunkRef += unsafe.Sizeof(nSize)
			}
			lp_importDesc = (*TImageImportDescriptor)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_importDesc)) + Pointer(unsafe.Sizeof(*lp_importDesc))))
		}
	}
	return Result
}

func GetSectionProtection(f_SC DWORD) DWORD {
	var Result DWORD
	if f_SC&IMAGE_SCN_MEM_NOT_CACHED != 0 {
		Result &= PAGE_NOCACHE
	}
	// E - Execute, R ?Read , W ?Write
	if f_SC&IMAGE_SCN_MEM_EXECUTE != 0 { // E ?
		if f_SC&IMAGE_SCN_MEM_READ != 0 { // ER ?
			if f_SC&IMAGE_SCN_MEM_WRITE != 0 { // ERW ?
				Result = Result | PAGE_EXECUTE_READWRITE
			} else {
				Result = Result | PAGE_EXECUTE_READ
			}
		} else if f_SC&IMAGE_SCN_MEM_WRITE != 0 { // EW?
			Result = Result | PAGE_EXECUTE_WRITECOPY
		} else {
			Result = Result | PAGE_EXECUTE
		}
	} else if f_SC&IMAGE_SCN_MEM_READ != 0 { // R?
		if f_SC&IMAGE_SCN_MEM_WRITE != 0 { // RW?
			Result = Result | PAGE_READWRITE
		} else {
			Result = Result | PAGE_READONLY
		}
	} else if f_SC&IMAGE_SCN_MEM_WRITE != 0 { // W?
		Result = Result | PAGE_WRITECOPY
	} else {
		Result = Result | PAGE_NOACCESS
	}
	return Result
}

func FinalizeSections(fp_module *TBTMemoryModule) {
	lp_section := GetImageFirstSection(fp_module.headers)
	for i := 0; i < int(fp_module.headers.FileHeader.NumberOfSections); i++ {
		if lp_section.Characteristics&IMAGE_SCN_MEM_DISCARDABLE != 0 {
			// section is not needed any more and can safely be freed
			VirtualFree(Pointer(lp_section.Misc.PhysicalAddress), SIZE_T(lp_section.SizeOfRawData), MEM_DECOMMIT)
			lp_section = (*TImageSectionHeader)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_section)) - unsafe.Sizeof(*lp_section)))
			// run next for loop interation...
			continue
		}
		l_protect := GetSectionProtection(lp_section.Characteristics)
		if lp_section.Characteristics&IMAGE_SCN_MEM_NOT_CACHED != 0 {
			l_protect |= PAGE_NOCACHE
		}
		// determine size of region
		l_size := lp_section.SizeOfRawData
		if l_size == 0 {
			if lp_section.Characteristics&IMAGE_SCN_CNT_INITIALIZED_DATA != 0 {
				l_size = fp_module.headers.OptionalHeader.SizeOfInitializedData
			} else {
				if lp_section.Characteristics&IMAGE_SCN_CNT_UNINITIALIZED_DATA != 0 {
					l_size = fp_module.headers.OptionalHeader.SizeOfUninitializedData
				}
			}
			if l_size > 0 {
				var l_oldProtect Pointer
				if !VirtualProtect(Pointer(lp_section.Misc.PhysicalAddress), SIZE_T(lp_section.SizeOfRawData), l_protect, Pointer(unsafe.Pointer(&l_oldProtect))) {
					lastErrStr = "FinalizeSections: VirtualProtect failed"
					return
				}
			}
		}
		lp_section = (*TImageSectionHeader)(unsafe.Pointer(Pointer(unsafe.Pointer(lp_section)) + unsafe.Sizeof(*lp_section)))
	}
}

func BTMemoryLoadLibary(fp_data Pointer, f_size int64) *TBTMemoryModule {
	var result *TBTMemoryModule
	defer func() {
		if err := recover(); err != nil {
			BTMemoryFreeLibrary(result)
		}
	}()
	var l_dos_header TImageDosHeader
	Memcpy(uintptr(unsafe.Pointer(&l_dos_header)), fp_data, SIZE_T(unsafe.Sizeof(l_dos_header)))
	if l_dos_header.e_magic != IMAGE_DOS_SIGNATURE {
		lastErrStr = "BTMemoryLoadLibary: dll dos header is not valid"
		return nil
	}
	var l_old_header TImageNtHeaders
	Memcpy(uintptr(unsafe.Pointer(&l_old_header)), fp_data+Pointer(l_dos_header._lfanew), SIZE_T(unsafe.Sizeof(l_old_header)))
	if l_old_header.Signature != IMAGE_NT_SIGNATURE {
		lastErrStr = "BTMemoryLoadLibary: IMAGE_NT_SIGNATURE is not valid"
		return nil
	}
	// reserve memory for image of library
	l_code := VirtualAlloc(l_old_header.OptionalHeader.ImageBase, SIZE_T(l_old_header.OptionalHeader.SizeOfImage), MEM_RESERVE, PAGE_READWRITE)
	if l_code == 0 {
		// try to allocate memory at arbitrary position
		l_code = VirtualAlloc(0, SIZE_T(l_old_header.OptionalHeader.SizeOfImage), MEM_RESERVE, PAGE_READWRITE)
	}
	if l_code == 0 {
		lastErrStr = "BTMemoryLoadLibary: VirtualAlloc failed"
		return nil
	}
	// alloc space for the result record
	result = (*TBTMemoryModule)(unsafe.Pointer(HeapAlloc(GetProcessHeap(), 0, SIZE_T(unsafe.Sizeof(*result)))))
	//result = new(TBTMemoryModule)
	result.codeBase = l_code
	result.numModules = 0
	result.modules = 0
	result.initialized = false

	// xy: is it correct to commit the complete memory region at once?
	// calling DllEntry raises an exception if we don't...
	VirtualAlloc(l_code, SIZE_T(l_old_header.OptionalHeader.SizeOfImage), MEM_COMMIT, PAGE_READWRITE)
	// commit memory for headers
	l_headers := VirtualAlloc(l_code, SIZE_T(l_old_header.OptionalHeader.SizeOfHeaders), MEM_COMMIT, PAGE_READWRITE)
	// copy PE header to code
	Memcpy(l_headers, fp_data, SIZE_T(l_dos_header._lfanew)+SIZE_T(l_old_header.OptionalHeader.SizeOfHeaders))
	result.headers = (*TImageNtHeaders)(unsafe.Pointer(l_headers + Pointer(l_dos_header._lfanew)))

	// update position
	result.headers.OptionalHeader.ImageBase = l_code
	// copy sections from DLL file block to new memory location
	CopySections(fp_data, l_old_header, result)
	// adjust base address of imported data
	l_locationdelta := int(l_code - Pointer(l_old_header.OptionalHeader.ImageBase))
	if l_locationdelta != 0 {
		PerformBaseRelocation(result, l_locationdelta)
	}
	// load required dlls and adjust function table of imports
	if !BuildImportTable(result) {
		lastErrStr = lastErrStr + " BTMemoryLoadLibary: BuildImportTable failed"
		panic(lastErrStr)
	}
	// mark memory pages depending on section headers and release
	// sections that are marked as "discardable"
	FinalizeSections(result)

	// get entry point of loaded library
	if result.headers.OptionalHeader.AddressOfEntryPoint != 0 {

		lp_DllEntry := l_code + Pointer(result.headers.OptionalHeader.AddressOfEntryPoint)
		if lp_DllEntry == 0 {
			lastErrStr = "BTMemoryLoadLibary: GetDLLEntyPoint failed"
			panic(lastErrStr)
		}
		if !result.callDLLEntry(DLL_PROCESS_ATTACH) {
			lastErrStr = "BTMemoryLoadLibary: Can't attach library"
			panic(lastErrStr)
		}
		result.initialized = true
		lastErrStr = ""
	}
	result.isValid = true
	return result
}

func BTMemoryGetProcAddress(fp_module *TBTMemoryModule, fp_name string) Pointer {
	var idx Integer = -1
	directory := GetHeaderDictionary(fp_module, IMAGE_DIRECTORY_ENTRY_EXPORT)
	if directory.Size == 0 {
		lastErrStr = "BTMemoryGetProcAddress: no export table found"
		return 0
	}
	exports := (*TImageExportDirectory)(unsafe.Pointer(fp_module.codeBase + Pointer(directory.VirtualAddress)))
	if exports.NumberOfNames == 0 || exports.NumberOfFunctions == 0 {
		lastErrStr = "BTMemoryGetProcAddress: DLL doesn't export anything"
		return 0
	}
	// search function name in list of exported names
	nameRef := fp_module.codeBase + Pointer(exports.AddressOfNames)
	ordinal := fp_module.codeBase + Pointer(exports.AddressOfNameOrdinals)
	for l_i := 0; l_i < int(exports.NumberOfNames); l_i++ {
		nameAddr := fp_module.codeBase + Pointer(*(*DWORD)(unsafe.Pointer(nameRef)))
		nameLen := Lstrlen(nameAddr)
		buff := make([]byte, nameLen)
		Memcpy(uintptr(unsafe.Pointer(&buff[0])), nameAddr, SIZE_T(len(buff)))

		if fp_name == string(buff) {
			idx = Integer(*(*Word)(unsafe.Pointer(ordinal)))
			break
		}
		nameRef += 4
		ordinal += 2
	}
	if idx == -1 {
		lastErrStr = "BTMemoryGetProcAddress: exported symbol not found."
		return 0
	}
	if DWORD(idx) > exports.NumberOfFunctions-1 {
		lastErrStr = "BTMemoryGetProcAddress: name <-> ordinal number don't match."
		return 0
	}
	// AddressOfFunctions contains the RVAs to the "real" functions
	v1 := Pointer(exports.AddressOfFunctions) + Pointer(idx*4)
	v2 := *((*Pointer)(unsafe.Pointer(fp_module.codeBase + v1)))
	return fp_module.codeBase + v2
}

// free module
func BTMemoryFreeLibrary(fp_module *TBTMemoryModule) {
	if fp_module != nil {
		if fp_module.initialized {
			fp_module.callDLLEntry(DLL_PROCESS_DETACH)
			fp_module.initialized = false
			fp_module.isValid = false
			// free previously opened libraries
			for i := 0; i < fp_module.numModules; i++ {
				//l_temp := (SizeOf(Cardinal) * (l_i))
				temp := Pointer(4 * i)
				fp_module.modules += temp
				if *(*DWORD)(unsafe.Pointer(fp_module.modules)) != DWORD(INVALID_HANDLE_VALUE) {
					FreeLibrary(HMODULE(*(*Pointer)(unsafe.Pointer(fp_module.modules))))
					fp_module.modules -= temp
				}
			}
			// 后面再看这个怎么处理，
			Free(fp_module.modules)
			if fp_module.codeBase != 0 {
				//	 release memory of library
				VirtualFree(fp_module.codeBase, 0, MEM_RELEASE)
			}
			HeapFree(GetProcessHeap(), 0, fp_module.this())
			fp_module = nil
		}
	}
}

// returns last error
func BTMemoryGetLastError() string {
	return lastErrStr
}
