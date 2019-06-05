// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl

package miniblink

import (
	"runtime"

	"github.com/ying32/govcl/vcl/dylib"
)

var (
	wkedll = dylib.NewLazyDLL("node.dll") //"wke.dll")
	is386  = runtime.GOARCH == "386"
)
