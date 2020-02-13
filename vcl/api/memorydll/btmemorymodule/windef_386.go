// +build windows,386

package btmemorymodule

type TImageOptionalHeader32 struct {
	// Standard fields.
	Magic                   Word
	MajorLinkerVersion      Byte
	MinorLinkerVersion      Byte
	SizeOfCode              DWORD
	SizeOfInitializedData   DWORD
	SizeOfUninitializedData DWORD
	AddressOfEntryPoint     DWORD
	BaseOfCode              DWORD
	BaseOfData              DWORD
	// NT additional fields.
	ImageBase                   uintptr // DWORD
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
	SizeOfStackReserve          DWORD
	SizeOfStackCommit           DWORD
	SizeOfHeapReserve           DWORD
	SizeOfHeapCommit            DWORD
	LoaderFlags                 DWORD
	NumberOfRvaAndSizes         DWORD
	DataDirectory               [IMAGE_NUMBEROF_DIRECTORY_ENTRIES]TImageDataDirectory
}

type TImageNtHeaders32 struct {
	Signature      DWORD
	FileHeader     TImageFileHeader
	OptionalHeader TImageOptionalHeader32
}

type TImageNtHeaders = TImageNtHeaders32
