//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/locales/zh_CN"
)

func main() {
	vcl.RunApp(&MainForm, &NewConnection)
}
