// 由res2go自动生成。
package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

func main() {
	vcl.Application.SetFormScaled(true)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&Form11)
	vcl.Application.Run()
}
