// 由res2go自动生成。
package main

import (
	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
)

const Lazarus = true

func main() {

	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainForm)
	vcl.Application.CreateForm(&About)
	vcl.Application.Run()
}
