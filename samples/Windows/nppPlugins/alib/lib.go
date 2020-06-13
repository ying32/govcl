//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package alib

import (
	"path/filepath"

	"github.com/ying32/govcl/pkgs/libname"
	"github.com/ying32/govcl/vcl/win"
)

func ExtractFilePath(path string) string {
	filename := filepath.Base(path)
	return path[:len(path)-len(filename)]
}

func init() {
	fileName, _ := win.GetModuleFileName(win.GetModuleHandle("nppPlugins.dll"))
	libname.LibName = ExtractFilePath(fileName) + "liblcl.dll"
	win.OutputDebugString("libName: ", libname.LibName)
}
