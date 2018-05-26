package main

import (
    "gitee.com/ying32/govcl/vcl"
)

func main() {

    vcl.Application.SetIconResId(3) // 具体资源id根据rsrc.exe编译的为准
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateFormFromBytes(form1Bytes, &Form1)
    vcl.Application.CreateFormFromBytes(form2Bytes, &Form2)
    vcl.Application.CreateFormFromBytes(form3Bytes, &Form3)
    vcl.Application.CreateFormFromBytes(form4Bytes, &Form4)
    vcl.Application.CreateFormFromBytes(form5Bytes, &Form5)
    vcl.Application.CreateFormFromBytes(form6Bytes, &Form6)
    vcl.Application.Run()

}

