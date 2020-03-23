package main

import (
	"github.com/ying32/govcl/vcl"

	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl/types"
)

func main() {

	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("DropFiles测试")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.EnabledMaximize(false)
	mainForm.SetWidth(300)
	mainForm.SetHeight(200)

	// UAC下无法接收一些消息，需要使用私有函数启用
	windowsUACMessageFilter(mainForm.Handle())

	fmt.Println("allowDropFiles1:", mainForm.AllowDropFiles())
	mainForm.SetAllowDropFiles(true)
	fmt.Println("allowDropFiles2:", mainForm.AllowDropFiles())
	mainForm.SetOnDropFiles(func(sender vcl.IObject, aFileNames []string) {
		fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
		for i, s := range aFileNames {
			fmt.Println("index:", i, ", filename:", s)
		}
	})
	vcl.Application.Run()
}
