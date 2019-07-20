package main

import (
	"github.com/ying32/govcl/vcl"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(form1Bytes, &Form1)
	vcl.Application.CreateForm(form2Bytes, &Form2)
	vcl.Application.CreateForm(form3Bytes, &Form3)
	vcl.Application.CreateForm(form4Bytes, &Form4)
	vcl.Application.CreateForm(form5Bytes, &Form5)
	vcl.Application.CreateForm(form6Bytes, &Form6)
	vcl.Application.Run()

}
