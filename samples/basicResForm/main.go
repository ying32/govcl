package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {

	//vcl.Application.SetFormScaled(true)
	//vcl.Application.Initialize()
	//vcl.Application.SetMainFormOnTaskBar(true)
	//vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
	//	vcl.ShowMessage(e.Message())
	//})
	////   Form1.gfm
	//vcl.Application.CreateForm(&Form1)
	//// 字节加载方式
	//vcl.Application.CreateForm(&Form2)
	//
	//vcl.Application.Run()
	vcl.RunApp(&Form1, &Form2)

}
