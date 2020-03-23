//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package version

import "runtime"

func init() {
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
	initOSVersion()
}
