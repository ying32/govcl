package main

import (
	"gitee.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateFormFromBytes(mainFormBytes, &MainForm)
	vcl.Application.Run()

}
