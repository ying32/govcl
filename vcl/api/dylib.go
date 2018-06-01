package api

import (
	"runtime"

	"fmt"

	"github.com/ying32/govcl/vcl/dylib"
)

// 全局导入库
var libvcl = loadUILib()

// 判断是否加载lcl库，一般用于windows上
var IsloadedLcl = false

const (
	// 各个平台对应的lib名
	libvcldll    = "libvcl.dll"
	libvclx64dll = "libvclx64.dll"
	liblcldll    = "liblcl.dll"
	liblclso     = "liblcl.so"
	liblcldylib  = "liblcl.dylib"
)

// 加载库
func loadUILib() *dylib.LazyDLL {
	libname := ""
	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "amd64" {
			libname = libvclx64dll
		} else {
			libname = libvcldll
		}
	case "linux":
		libname = liblclso
		IsloadedLcl = true

	case "darwin":
		libname = liblcldylib
		IsloadedLcl = true

	}
	lib := dylib.NewLazyDLL(libname)
	// 这里做个判断，当libvcl.dll或者libvclx64.dll加载失败时尝试加载liblcl.dll
	// 这样做主要为以后考虑，对于某些人来说怕什么的来说，可以使用非Delphi的组件
	err := lib.Load()
	if err != nil && runtime.GOOS == "windows" && (libname == libvcldll || libname == libvclx64dll) {
		fmt.Println(fmt.Sprintf("\"%s\" does not exist, trying to load liblcl.dll.", libname))
		lib = dylib.NewLazyDLL(liblcldll)
		IsloadedLcl = true
	}
	return lib
}
