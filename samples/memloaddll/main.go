package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/api"
)

func main() {
	// Free
	defer api.FeeMemoryDLL()
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		vcl.ShowMessage(e.Message())
	})
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.CreateForm(form2Bytes, &Form2)
	vcl.Application.Run()

}
