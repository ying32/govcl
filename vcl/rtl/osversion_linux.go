// 移植来自Delphi

package rtl

import (
	"bufio"
	"fmt"
	"io"
	"os"
	"runtime"
	"strconv"
	"strings"
	"unsafe"
)

/*

#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <sys/utsname.h>
#include <gnu/libc-version.h>

  char* GetutsNameRelease() {
    struct utsname Name;
	if(uname(&Name) == 0) {
	   char* str = (char*)malloc(_UTSNAME_RELEASE_LENGTH);
	   strncpy(str, Name.release, _UTSNAME_RELEASE_LENGTH);
	   return str;
	}
    return NULL;
  }

  char* GetutsNameVersion() {
    struct utsname Name;
	if(uname(&Name) == 0) {
	   char* str = (char*)malloc(_UTSNAME_VERSION_LENGTH);
	   strncpy(str, Name.version, _UTSNAME_VERSION_LENGTH);
	   return str;
	}
    return NULL;
  }

*/
import "C"

func getUTSNameRelease() string {
	str := C.GetutsNameRelease()
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func getUTSNameVersion() string {
	str := C.GetutsNameVersion()
	defer C.free(unsafe.Pointer(str))
	return C.GoString(str)
}

func getGNULibCVersion() string {
	return C.GoString(C.gnu_get_libc_version())
}

func parseUname() {
	str := getUTSNameRelease()
	if str != "" {
		// Example: '4.8.0-34-generic'
		// 4: FMajor
		// 8: FMinor,
		// 0: FServicePackMajor
		// 34: FServicePackMinor
		str = strings.Replace(str, "-", ".", -1)
		strArr := strings.Split(str, ".")
		if len(strArr) >= 1 {
			OSVersion.Major, _ = strconv.Atoi(strArr[0])
		}
		if len(strArr) >= 2 {
			OSVersion.Minor, _ = strconv.Atoi(strArr[1])
		}
		if len(strArr) >= 3 {
			OSVersion.ServicePackMajor, _ = strconv.Atoi(strArr[2])
		}
		if len(strArr) >= 4 {
			OSVersion.ServicePackMinor, _ = strconv.Atoi(strArr[3])
		}
	}
	// Example: '#36-Ubuntu SMP Wed Dec 21 17:24:18 UTC 2016'
	OSVersion.Name = getUTSNameVersion()
}

const (
	osReleaseFileName = "/etc/os-release"
	prettyName        = "PRETTY_NAME="
	sVersionStr       = "%s %s (Version %d.%d.%d)"
)

func parseOSRelease() {
	f, err := os.Open(osReleaseFileName)
	if err != nil {
		return
	}
	defer f.Close()
	rd := bufio.NewReader(f)
	for {
		line, err := rd.ReadString('\n')
		if err != nil || err == io.EOF {
			break
		}
		// Example: 'PRETTY_NAME="Ubuntu 16.04.1 LTS"'
		if strings.HasPrefix(line, prettyName) {
			OSVersion.PrettyName = line[strings.Index(line, "=")+2 : strings.LastIndex(line, "\"")]
			break
		}
	}
}

func parseLibcVersion() {
	str := getGNULibCVersion()
	items := strings.Split(str, ".")
	if len(items) >= 1 {
		OSVersion.LibCVersionMajor, _ = strconv.Atoi(items[0])
	}
	if len(items) >= 2 {
		OSVersion.LibCVersionMinor, _ = strconv.Atoi(items[1])
	}
}

func initOSVersion() {
	OSVersion.Platform = PfLinux
	switch runtime.GOARCH {
	case "386":
		OSVersion.Architecture = ArIntelX86
	case "amd64":
		OSVersion.Architecture = ArIntelX64
	case "arm":
		OSVersion.Architecture = ArARM32
	case "arm64":
		OSVersion.Architecture = ArARM64
	default:
	}
	parseUname()
	parseOSRelease()
	parseLibcVersion()
	OSVersion.fmtVerString = fmt.Sprintf(sVersionStr, OSVersion.PrettyName, OSVersion.Name, OSVersion.Major, OSVersion.Minor, OSVersion.ServicePackMajor)
}
