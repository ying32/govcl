package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/miniblink"
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
