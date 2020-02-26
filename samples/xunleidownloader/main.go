// +build windows,386

// 这个项目只是一个测试例程，随便写下
// 相关迅雷sdk可到迅雷官网下载，或者到 github.com/ying32/xldl 里面下载

package main

import (
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/xldl"
)

func main() {

	if !xldl.InitXLEngine() {
		vcl.ShowMessage("初始引擎失败！")
		return
	}
	defer xldl.UnInitXLEngine()

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(&MainForm)

	vcl.Application.Run()

}
