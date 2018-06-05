package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/tools"
	"github.com/ying32/govcl/vcl/rtl"
)

func main() {
	// mac下发布时去掉，只在测试时使用
	tools.RunWithMacOSApp()

	vcl.Application.SetFormScaled(true)
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		vcl.ShowMessage(e.Message())
	})
	//   Form1.gfm
	vcl.Application.CreateForm(rtl.ExtractFilePath(vcl.Application.ExeName())+"Form1.gfm", &Form1)
	// 文件加载方式
	//	vcl.Application.CreateForm("Form2.gfm", &Form2)
	// 字节加载方式
	vcl.Application.CreateForm(form2Bytes, &Form2)

	vcl.Application.Run()

}
