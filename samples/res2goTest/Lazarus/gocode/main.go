// 由res2go自动生成。
package main

import (
    "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.SetFormScaled(true)
    vcl.Application.Initialize()
    vcl.Application.CreateForm(mainFormBytes, &MainForm)
    vcl.Application.CreateForm(aboutBytes, &About)
    vcl.Application.Run()
}
