package main

import (
	_ "github.com/ying32/govcl/pkgs/macapp"
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form1)
	vcl.Application.CreateForm(&Form2)

	vcl.Application.Run()

}
