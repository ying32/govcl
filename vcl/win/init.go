// +build windows

package win

var (
	Win32MajorVersion int
	Win32MinorVersion int
	Win32BuildNumber  int
	Win32Platform     int
	ServicePackMajor  int
	Win32CSDVersion   string
)

func init() {
	initPlatformId()
}

func initPlatformId() {
	var OSVersionInfo TOSVersionInfoEx
	var Major, Minor, ServicePack int
	if GetVersionEx(&OSVersionInfo) {
		if (OSVersionInfo.MajorVersion > 6) || ((OSVersionInfo.MajorVersion == 6) && (OSVersionInfo.MinorVersion > 1)) {
			ServicePack = -1
			FixWindowsVersion(&Major, &Minor, &ServicePack)
			OSVersionInfo.MajorVersion = uint32(Major)
			OSVersionInfo.MinorVersion = uint32(Minor)
			OSVersionInfo.BuildNumber = 0
		}
		Win32Platform = int(OSVersionInfo.PlatformId)
		Win32MajorVersion = int(OSVersionInfo.MajorVersion)
		Win32MinorVersion = int(OSVersionInfo.MinorVersion)
		if Win32Platform == VER_PLATFORM_WIN32_WINDOWS {
			Win32BuildNumber = int(OSVersionInfo.BuildNumber & 0xFFFF)
		} else {
			Win32BuildNumber = int(OSVersionInfo.BuildNumber)
		}
		Win32CSDVersion = GoStr(OSVersionInfo.CSDVersion[:])
	}
}
