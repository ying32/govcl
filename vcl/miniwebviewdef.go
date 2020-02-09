// +build windows

package vcl

import (
	"runtime"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/win"
)

// Windows下设置TMiiWebview的IE内核版本，好像是需管理员权限才可以写注册表
func (m *TMiniWebview) SetIEVersion(version int32) {
	// 7000   Webpages containing standards-based !DOCTYPE directives are displayed in IE7 Standards mode.
	// 8000   Webpages containing standards-based !DOCTYPE directives are displayed in IE8 mode.
	// 9000   Internet Explorer 9. Webpages containing standards-based !DOCTYPE directives are displayed in IE9 mode.
	// 10000  Internet Explorer 10.
	// 11000  Internet Explorer 11. Webpages containing standards-based !DOCTYPE directives are displayed in IE11 Standards mode.
	// 默认的就是7
	// HKEY_CURRENT_USER\Software\Microsoft\Internet Explorer\Main\FeatureControl\FEATURE_BROWSER_EMULATION
	if version >= 7 && version <= 11 {
		var access uint32 = win.KEY_ALL_ACCESS
		if runtime.GOARCH == "amd64" {
			access |= win.KEY_WOW64_64KEY
		}
		reg := NewRegistry(access)
		defer reg.Free()
		reg.SetRootKey(win.HKEY_CURRENT_USER)
		if reg.OpenKey("Software\\Microsoft\\Internet Explorer\\Main\\FeatureControl\\FEATURE_BROWSER_EMULATION", false) {
			defer reg.CloseKey()
			reg.WriteInteger(rtl.ExtractFileName(Application.ExeName()), version*1000)
		}
	}
}
