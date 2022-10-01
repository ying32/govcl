package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

// windows : stdcall;  其他平台： cdecl
//var runLoopReceivedProc = syscall.NewCallback(runLoop)

// data 为 Application 实例
/*func runLoop(data uintptr) uintptr {
	// 这里不能产生异常，否则程序会崩溃

	return 0
	//vcl.Application.HandleMessage()
	//return 1 // 返回1表示不由Lazarus处理后面的，必须要加上Application.HandleMessage()，一般情况下最好不自己处理
}*/

func main() {
	//vcl.Application.SetRunLoopReceived(runLoopReceivedProc)
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
