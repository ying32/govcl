// +build windows

// 正体改自Delphi TOSVersion
package rtl

import (
	"fmt"

	"unsafe"

	"gitee.com/ying32/govcl/vcl/win"
)

const (
	sVersionStr          = "%s (Version %d.%d, Build %d, %s)"
	sSPVersionStr        = "%s Service Pack %d (Version %d.%d, Build %d, %s)"
	sVersion32           = "32-bit Edition"
	sVersion64           = "64-bit Edition"
	sWindows             = "Windows"
	sWindowsVista        = "Windows Vista"
	sWindowsServer2008   = "Windows Server 2008"
	sWindows7            = "Windows 7"
	sWindowsServer2008R2 = "Windows Server 2008 R2"
	sWindows2000         = "Windows 2000"
	sWindowsXP           = "Windows XP"
	sWindowsServer2003   = "Windows Server 2003"
	sWindowsServer2003R2 = "Windows Server 2003 R2"
	sWindowsServer2012   = "Windows Server 2012"
	sWindowsServer2012R2 = "Windows Server 2012 R2"
	sWindowsServer2016   = "Windows Server 2016"
	sWindows8            = "Windows 8"
	sWindows8Point1      = "Windows 8.1"
	sWindows10           = "Windows 10"
)

func iifStr(b bool, aTrue, aFalse string) string {
	if b {
		return aTrue
	}
	return aFalse
}

func IsWindowsServer() bool {

	var osvi win.TOSVersionInfoEx
	osvi.ProductType = win.VER_NT_WORKSTATION

	var dwlConditionMask uint64

	dwlConditionMask = win.VerSetConditionMask(0, win.VER_PRODUCT_TYPE, win.VER_EQUAL)
	return win.VerifyVersionInfo(&osvi, win.VER_PRODUCT_TYPE, dwlConditionMask) == false
}

func GetProductVersion(AFileName string, AMajor, AMinor, ABuild *uint32) bool {
	var wnd uint32
	infoSize := win.GetFileVersionInfoSize(AFileName, &wnd)
	if infoSize != 0 {
		verBuf := make([]byte, infoSize)
		bufPtr := uintptr(unsafe.Pointer(&verBuf[0]))
		if win.GetFileVersionInfo(AFileName, wnd, infoSize, bufPtr) {
			var verSize uint32
			var fI *win.TVSFixedFileInfo
			if win.VerQueryValue(bufPtr, "\\", (*uintptr)(unsafe.Pointer(&fI)), &verSize) {
				*AMajor = fI.ProductVersionMS >> 16
				*AMinor = uint32(uint16(fI.ProductVersionMS))
				*ABuild = fI.ProductVersionLS >> 16
				return true
			}
		}
	}
	return false
}

func GetNetWkstaMajorMinor(MajorVersion, MinorVersion *uint32) bool {

	var LBuf *win.WKSTA_INFO_100
	result := win.NetWkstaGetInfo100("", 100, &LBuf) == win.NERR_Success
	if result {
		*MajorVersion = LBuf.Wki100_ver_major
		*MinorVersion = LBuf.Wki100_ver_minor
	} else {
		*MajorVersion = 0
		*MinorVersion = 0
	}
	if LBuf != nil {
		win.NetApiBufferFree(LBuf)
	}
	return result
}

func initOSVersion() {
	var verInfo win.TOSVersionInfoEx
	win.GetVersionEx(&verInfo)

	OSVersion.Platform = PfWindows
	OSVersion.Major = int(verInfo.MajorVersion)
	OSVersion.Minor = int(verInfo.MinorVersion)
	OSVersion.Build = int(verInfo.BuildNumber)
	OSVersion.ServicePackMajor = int(verInfo.ServicePackMajor)
	OSVersion.ServicePackMinor = int(verInfo.ServicePackMinor)

	var sysInfo win.TSystemInfo
	if OSVersion.CheckMajorMinor(5, 1) {
		win.GetNativeSystemInfo(&sysInfo)
	} // GetNativeSystemInfo not supported on Windows 2000

	OSVersion.Architecture = ArIntelX86
	if sysInfo.ProcessorArchitecture == win.PROCESSOR_ARCHITECTURE_AMD64 {
		OSVersion.Architecture = ArIntelX64
	}

	var majorNum, minorNum, buildNum uint32
	if OSVersion.Major > 6 || (OSVersion.Major == 6 && OSVersion.Minor > 1) {
		if GetProductVersion("kernel32.dll", &majorNum, &minorNum, &buildNum) {
			OSVersion.Major = int(majorNum)
			OSVersion.Minor = int(minorNum)
			OSVersion.Build = int(buildNum)
		} else if GetNetWkstaMajorMinor(&majorNum, &minorNum) {
			OSVersion.Major = int(majorNum)
			OSVersion.Minor = int(minorNum)
		}
	}

	OSVersion.Name = sWindows
	switch OSVersion.Major {
	case 10:
		switch OSVersion.Minor {
		case 0:
			if !IsWindowsServer() {
				OSVersion.Name = sWindows10
			} else {
				OSVersion.Name = sWindowsServer2016
			}
		}

	case 6:
		switch OSVersion.Minor {
		case 0:
			if verInfo.ProductType == win.VER_NT_WORKSTATION {
				OSVersion.Name = sWindowsVista
			} else {
				OSVersion.Name = sWindowsServer2008
			}

		case 1:
			if verInfo.ProductType == win.VER_NT_WORKSTATION {
				OSVersion.Name = sWindows7
			} else {
				OSVersion.Name = sWindowsServer2008R2
			}

		case 2:
			if verInfo.ProductType == win.VER_NT_WORKSTATION {
				OSVersion.Name = sWindows8
			} else {
				OSVersion.Name = sWindowsServer2012
			}
		case 3:
			if !IsWindowsServer() {
				OSVersion.Name = sWindows8Point1
			} else {
				OSVersion.Name = sWindowsServer2012R2
			}
		}

	case 5:
		switch OSVersion.Minor {
		case 0:
			OSVersion.Name = sWindows2000
		case 1:
			OSVersion.Name = sWindowsXP
		case 2:
			if verInfo.ProductType == win.VER_NT_WORKSTATION && sysInfo.ProcessorArchitecture == win.PROCESSOR_ARCHITECTURE_AMD64 {
				OSVersion.Name = sWindowsXP
			} else {
				if win.GetSystemMetrics(win.SM_SERVERR2) == 0 {
					OSVersion.Name = sWindowsServer2003
				} else {
					OSVersion.Name = sWindowsServer2003R2
				}
			}
		}
	}

	if OSVersion.ServicePackMinor != 0 {
		OSVersion.fmtVerString = fmt.Sprintf(sSPVersionStr,
			OSVersion.Name, OSVersion.Major, OSVersion.Minor, OSVersion.Build, OSVersion.ServicePackMinor,
			iifStr(OSVersion.Architecture == ArIntelX64, sVersion64, sVersion32))
	} else {
		OSVersion.fmtVerString = fmt.Sprintf(sVersionStr,
			OSVersion.Name, OSVersion.Major, OSVersion.Minor, OSVersion.Build,
			iifStr(OSVersion.Architecture == ArIntelX64, sVersion64, sVersion32))
	}

}
