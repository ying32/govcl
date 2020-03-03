package main

import (
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()
}
