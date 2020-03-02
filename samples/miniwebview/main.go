package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/macapp"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {
	macapp.RunWithMacOSApp()
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()
}
