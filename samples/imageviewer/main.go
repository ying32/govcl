package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/tools"
)

func main() {
	tools.RunWithMacOSApp()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(mainFormBytes, &MainForm)
	vcl.Application.Run()

}
