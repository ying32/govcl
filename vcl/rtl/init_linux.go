package rtl

import (
	"os"
	"strings"
)

// 初始化
func init() {
	lang := os.Getenv("LANG")
	if lang != "" {
		if strings.Contains(lang, ".") {
			SysLocale.DefaultLCID = LocaleIDFromName(lang[:strings.Index(lang, ".")])
		} else {
			SysLocale.DefaultLCID = LocaleIDFromName(lang)
		}
	}
}
