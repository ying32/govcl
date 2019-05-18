package main

import (
	"runtime"

	"github.com/ying32/govcl/vcl"
)

func main() {

	runtime.GOMAXPROCS(runtime.NumCPU()*2 - 1)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(mainFormBytes, &MainForm)
	vcl.Application.Run()

}
