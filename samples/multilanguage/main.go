package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/multilang"
	"github.com/ying32/govcl/vcl/rtl"
)

func main() {

	// 首先设置lib中资源
	multilang.RegisterLibResouces(rtl.GetLibResouceItems(), rtl.ModifyLibResouce)
	multilang.InitDefaultLang()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.Run()

}
