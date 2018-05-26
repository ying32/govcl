package rtl

type TArchitecture uint32

const (
	ArIntelX86 TArchitecture = iota + 0
	ArIntelX64
	ArARM32
	ArARM64
)

type TPlatform uint32

const (
	PfWindows TPlatform = iota + 0
	PfMacOS
	PfiOS
	PfAndroid
	PfWinRT
	PfLinux
)

type TOSVersion struct {
	Name             string
	Build            int
	Major            int
	Minor            int
	ServicePackMajor int
	ServicePackMinor int
	Architecture     TArchitecture
	Platform         TPlatform

	// linux: PrettyName, LibCVersionMajor, LibCVersionMinor
	PrettyName       string
	LibCVersionMajor int
	LibCVersionMinor int

	fmtVerString string
}

var OSVersion TOSVersion

func (v *TOSVersion) CheckMajor(AMajor int) bool {
	return v.Major >= AMajor
}

func (v *TOSVersion) CheckMajorMinor(AMajor, AMinor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor >= AMinor)
}

func (v *TOSVersion) CheckMajorMinorServicePackMajor(AMajor, AMinor, AServicePackMajor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor > AMinor) ||
		((v.Major == AMajor && v.Minor == AMinor) && (v.ServicePackMajor >= AServicePackMajor))
}

func (v *TOSVersion) ToString() string {
	return v.fmtVerString
}
