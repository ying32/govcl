package rtl

import (
	"os"
	"strings"
)

// 初始化
func init() {
	SysLocale.FarEast = true
	SysLocale.MiddleEast = true
	lang := os.Getenv("LANG")
	if lang != "" {
		lang = strings.Replace(lang, "_", "-", -1)
		if strings.Contains(lang, ".") {
			SysLocale.DefaultLCID = LocaleIDFromName(lang[:strings.Index(lang, ".")])
		} else {
			SysLocale.DefaultLCID = LocaleIDFromName(lang)
		}
	}
}
