// +build windows

package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.Run()
}
