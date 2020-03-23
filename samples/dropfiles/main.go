package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"

	_ "github.com/ying32/govcl/pkgs/winappres"
)

type TMainForm struct {
	*vcl.TForm
}

var mainForm *TMainForm

func main() {
	vcl.RunApp(&mainForm)
}

func (f *TMainForm) OnFormCreate(object vcl.IObject) {
	f.SetCaption("drop files")
	f.SetWidth(300)
	f.SetHeight(200)
	f.ScreenCenter()
	f.EnabledMaximize(false)

	// allow drop file
	f.SetAllowDropFiles(true)

	// windows10没生效，有待研究
	windowsUACMessageFilter(f.Handle())
}

func (f *TMainForm) OnFormDropFiles(sender vcl.IObject, aFileNames []string) {
	fmt.Println("当前拖放文件事件执行，文件数：", len(aFileNames))
	for i, s := range aFileNames {
		fmt.Println("index:", i, ", filename:", s)
	}
}
