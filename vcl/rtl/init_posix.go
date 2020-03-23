//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !windows

package rtl

import (
	"os"
	"strings"
)

// 初始化
func init() {
	SysLocale.FarEast = true
	SysLocale.MiddleEast = true
	// 这个环境变量在MacOS下只有通过命令行启动的才能获取得到。
	// Linux下不知道有哪些不能获取到的。
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
