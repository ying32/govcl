// +build windows

package main

import (
	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"github.com/tryor/gdiplus"
	"github.com/tryor/winapi"
)

var (
	gpToken winapi.ULONG_PTR
)

// 用了一个第三方的go gdi+库， 看样子是个国人的作品

func main() {

	_, err := gdiplus.Startup(&gpToken, nil, nil)
	if err != nil {
		if err != nil {
			vcl.ShowMessage("初始GDI+失败，错误:" + err.Error())
		}
	} else {
		fmt.Println("gdi+实始成功")
		defer gdiplus.Shutdown(gpToken)
	}

	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		// 在这里自行处理VCL中的异常
	})
	vcl.Application.SetIconResId(3) // 具体资源id根据rsrc.exe编译的为准
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateFormFromBytes(gdipFormBytes, &GdipForm)
	vcl.Application.Run()

}
