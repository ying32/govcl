package main

import "github.com/ying32/govcl/vcl"
import _ "github.com/ying32/govcl/pkgs/winappres"

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&mainForm, true)
	vcl.Application.Run()
}
