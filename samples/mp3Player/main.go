package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/tools"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {

	tools.RunWithMacOSApp()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form1)
	vcl.Application.Run()

}
