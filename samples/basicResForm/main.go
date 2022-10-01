package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

//// windows : stdcall;  其他： cdecl
//var callproc = syscall.NewCallback(processdata)
//
//func processdata(data uintptr) uintptr {
//	// data 为 Application 实例
//	fmt.Println("收到事件了")
//	return 0 // 反回值实际未有作用
//}

func main() {
	//vcl.Application.SetRunLoopReceived(callproc)
	vcl.RunApp(&Form1, &Form2)

	//vcl.Application.SetScaled(true)
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

}
