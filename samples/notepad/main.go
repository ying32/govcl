package main

import (
	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
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
