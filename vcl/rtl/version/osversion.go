//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package version

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

// CheckMajor 检测系统主版本号
func (v *TOSVersion) CheckMajor(AMajor int) bool {
	return v.Major >= AMajor
}

// CheckMajorMinor 检测系统主版本和子版本号
func (v *TOSVersion) CheckMajorMinor(AMajor, AMinor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor >= AMinor)
}

func (v *TOSVersion) CheckMajorMinorServicePackMajor(AMajor, AMinor, AServicePackMajor int) bool {
	return v.Major > AMajor || (v.Major == AMajor && v.Minor > AMinor) ||
		((v.Major == AMajor && v.Minor == AMinor) && (v.ServicePackMajor >= AServicePackMajor))
}

// ToString 版本信息
func (v *TOSVersion) ToString() string {
	return v.fmtVerString
}
