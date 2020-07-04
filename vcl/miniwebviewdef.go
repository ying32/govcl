//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build windows

package vcl

import (
	"runtime"
	"strconv"
	"strings"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/win"
)

// Windows下设置TMiiWebview的IE内核版本，需管理员权限才可以写注册表
//
// Set the IE kernel version of TMiiWebview under Windows, you need administrator privileges.
func (m *TMiniWebview) SetIEVersion(version int) {
	// 7000   Webpages containing standards-based !DOCTYPE directives are displayed in IE7 Standards mode.
	// 8000   Webpages containing standards-based !DOCTYPE directives are displayed in IE8 mode.
	// 9000   Internet Explorer 9. Webpages containing standards-based !DOCTYPE directives are displayed in IE9 mode.
	// 10000  Internet Explorer 10.
	// 11000  Internet Explorer 11. Webpages containing standards-based !DOCTYPE directives are displayed in IE11 Standards mode.
	// 默认的就是7
	// HKEY_CURRENT_USER\Software\Microsoft\Internet Explorer\Main\FeatureControl\FEATURE_BROWSER_EMULATION
	if version >= 7 && version <= 11 {
		var access uint32 = win.KEY_ALL_ACCESS
		// 当前是32位的程序运行在x64下面
		//if win.IsWow64() {
		if runtime.GOARCH == "amd64" {
			access |= win.KEY_WOW64_64KEY
		}
		//}
		reg := NewRegistry(access)
		defer reg.Free()
		reg.SetRootKey(win.HKEY_CURRENT_USER)
		if reg.OpenKey("Software\\Microsoft\\Internet Explorer\\Main\\FeatureControl\\FEATURE_BROWSER_EMULATION", false) {
			defer reg.CloseKey()
			reg.WriteInteger(rtl.ExtractFileName(Application.ExeName()), int32(version*1000))
		}
	}
}

// Windows下读取IE内核版本，需管理员权限才可以写注册表
//
// Read the IE kernel version under Windows, you need administrator rights to write the registry.
func (m *TMiniWebview) GetIEVersion() int {
	reg := NewRegistryAllAccess()
	defer reg.Free()
	reg.SetRootKey(win.HKEY_LOCAL_MACHINE)
	if reg.OpenKey("SOFTWARE\\Microsoft\\Internet Explorer", false) {
		defer reg.CloseKey()
		verStr := ""
		if reg.ValueExists("svcVersion") {
			verStr = reg.ReadString("svcVersion")
		} else if reg.ValueExists("Version") {
			verStr = reg.ReadString("Version")
		}
		if verStr != "" {
			vArr := strings.Split(verStr, ".")
			if len(vArr) >= 1 {
				v, _ := strconv.Atoi(vArr[0])
				return v
			}
		}
	}
	return 0
}
