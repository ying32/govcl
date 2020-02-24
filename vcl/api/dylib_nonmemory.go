// +build !memorydll


//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package api

import (
	"os"
	"os/exec"
	"path/filepath"
	"runtime"

	"fmt"

	"github.com/ying32/dylib"
)

const (
	// 各个平台对应的lib名
	libvcldll    = "libvcl.dll"
	libvclx64dll = "libvclx64.dll"
	liblcldll    = "liblcl.dll"
	liblclso     = "liblcl.so"
	liblcldylib  = "liblcl.dylib"
)

// Windows下专用的
func windowsDLLExists(name string) bool {
	file, _ := exec.LookPath(os.Args[0])
	dllFileName := filepath.Dir(file) + "\\" + name
	_, err := os.Stat(dllFileName)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

// 加载库
func loadUILib() *dylib.LazyDLL {
	libName := ""
	switch runtime.GOOS {
	case "windows":
		if runtime.GOARCH == "amd64" {
			libName = libvclx64dll
		} else {
			libName = libvcldll
		}
		// 如果exe目录下存在libvclx64.dll且为x64系统
		if runtime.GOARCH == "amd64" && windowsDLLExists(libvclx64dll) {
			libName = libvclx64dll
		} else
		// 如果当前存在
		if runtime.GOARCH == "386" && windowsDLLExists(libvcldll) {
			libName = libvcldll
		} else
		// 如果当下目录存在liblcl则加载
		if windowsDLLExists(liblcldll) {
			libName = liblcldll
		}
	case "linux":
		libName = liblclso
	case "darwin":
		libName = liblcldylib
	}
	fmt.Println("LoadLibrary:", libName)
	lib := dylib.NewLazyDLL(libName)
	// 这里做个判断，当libvcl.dll或者libvclx64.dll加载失败时尝试加载liblcl.dll
	// 这样做主要为以后考虑，对于某些人来说怕什么的来说，可以使用非Delphi的组件
	err := lib.Load()
	if err != nil && runtime.GOOS == "windows" && (libName == libvcldll || libName == libvclx64dll) {
		fmt.Println(fmt.Sprintf("\"%s\"不存在，尝试加载“liblcl.dll”。\r\n(\"%s\" does not exist, trying to load \"liblcl.dll\".)", libName, libName))
		lib = dylib.NewLazyDLL(liblcldll)
	}

	IsloadedLcl = getLibType(lib) == LtLCL
	fmt.Println(guiLibTypeString())

	return lib
}

func getLibType(lib *dylib.LazyDLL) TLibType {
	proc := lib.NewProc("DGetLibType")
	r, _, _ := proc.Call()
	return TLibType(r)
}

// 获取dll库实例，用于在外扩展第三方组件的。移动来自dfuncs.go
func GetLibVcl() *dylib.LazyDLL {
	return libvcl
}
