// +build windows
// +build 386
// +build memorydll

// 编译时加上 -tags memorydll

// 内存加载libvcl.dll或者liblcl.dll，只支持win32。
// 注：内存加载dll容易被杀毒软件杀掉。。。
package api

import (
	"fmt"

	"github.com/ying32/govcl/vcl/dylib/memorydll"
	"github.com/ying32/govcl/vcl/win"
)

func loadUILib() *memorydll.LazyDLL {
	dllBytes, ok := win.ResouceToBytes(win.GetSelfModuleHandle(), "GOVCLLIB", win.RT_RCDATA)
	if !ok {
		fmt.Println(fmt.Sprintf("\"GOVCLLIB\"资源不存在。\r\n(\"GOVCLLIB\" resource does not exist.)"))
		return nil
	}
	lib := memorydll.NewMemoryDLL(dllBytes)
	if lib.Load() != nil {
		fmt.Println(fmt.Sprintf("无法加载dll，未知原因。\r\n(Unable to load dll, unknown reason.)"))
		return nil
	}
	IsloadedLcl = getLibType(lib) == LtLCL
	fmt.Println("IsloadedLcl:", IsloadedLcl)
	return lib
}

func getLibType(lib *memorydll.LazyDLL) TLibType {
	proc := lib.NewProc("DGetLibType")
	r, _, _ := proc.Call()
	return TLibType(r)
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibVcl() *memorydll.LazyDLL {
	return libvcl
}

// 释放内存，Memory DLL的最好调用下
func FeeMemoryDLL() {
	if libvcl != nil {
		libvcl.Close()
	}
}
