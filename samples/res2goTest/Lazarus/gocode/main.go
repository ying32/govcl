// 由res2go自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
)

const Lazarus = true

func main() {
    if rtl.LcLLoaded() {
		vcl.Application.SetFormScaled(true)
	}
    vcl.Application.Initialize()
    vcl.Application.CreateForm(&MainForm)
    vcl.Application.CreateForm(&About)
    vcl.Application.Run()
}
