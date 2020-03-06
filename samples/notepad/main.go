package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		// 在这里自行处理VCL中的异常
	})

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.Run()

}
