package main

import (
	"github.com/ying32/govcl/pkgs/miniblink"
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {

	miniblink.Init()
	defer miniblink.Finalize()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.CreateForm(&HTMLForm)
	vcl.Application.CreateForm(&JsForm)
	vcl.Application.Run()

}
