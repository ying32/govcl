// +build windows,386

package btmemorymodule

type TImageOptionalHeader64 struct {
	// Standard fields.
	Magic                   Word
	MajorLinkerVersion      Byte
	MinorLinkerVersion      Byte
	SizeOfCode              DWORD
	SizeOfInitializedData   DWORD
	SizeOfUninitializedData DWORD
	AddressOfEntryPoint     DWORD
	BaseOfCode              DWORD
	// NT additional fields.
	ImageBase                   uintptr //  ULONGLONG
	SectionAlignment            DWORD
	FileAlignment               DWORD
	MajorOperatingSystemVersion Word
	MinorOperatingSystemVersion Word
	MajorImageVersion           Word
	MinorImageVersion           Word
	MajorSubsystemVersion       Word
	MinorSubsystemVersion       Word
	Win32VersionValue           DWORD
	SizeOfImage                 DWORD
	SizeOfHeaders               DWORD
	CheckSum                    DWORD
	Subsystem                   Word
	DllCharacteristics          Word
	SizeOfStackReserve          ULONGLONG
	SizeOfStackCommit           ULONGLONG
	SizeOfHeapReserve           ULONGLONG
	SizeOfHeapCommit            ULONGLONG
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]TImageDataDirectory
}

type TImageNtHeaders64 struct {
	Signature      DWORD
	FileHeader     TImageFileHeader
	OptionalHeader TImageOptionalHeader64
}

type TImageNtHeaders = TImageNtHeaders64
