// +build windows

// miniblink及wke头文件导入
// 由ying32翻译，应用于govcl

package miniblink

import (
	"fmt"
	"runtime"
	"syscall"
)

var (
	wkedll = LoadMiniBlinkDLL()
	is386  = runtime.GOARCH == "386"
)

func SetFPMask()

func LoadMiniBlinkDLL() *syscall.LazyDLL {
	return syscall.NewLazyDLL(fmt.Sprintf("node-%s.dll", runtime.GOARCH))
}
