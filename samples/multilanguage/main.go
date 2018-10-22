package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/exts/multilang"
)

func main() {

	multilang.RegisterLangChnaged(func() {
		// 这里还没想好怎么更新库中的常量，光是Windows还好做，涉及到跨平台这个得想想了
		fmt.Println("语言已改变")
	})

	//multilang.AppNodeName
	multilang.InitDefaultLang()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.Run()

}
