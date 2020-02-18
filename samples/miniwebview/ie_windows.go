// +build windows

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/win"

	"github.com/ying32/govcl/vcl"
)

func SetIEVersion(webview *vcl.TMiniWebview) {
	if win.IsAdministrator() {
		v := webview.GetIEVersion()
		fmt.Println("IE Version:", v)
		if v >= 7 {
			webview.SetIEVersion(v)
		}
	} else {
		fmt.Println("需要Administrator权限才可使用。")
	}
}
