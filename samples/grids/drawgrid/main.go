package main

import (
	"fmt"
	"math/rand"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.Initialize()

	mainForm := vcl.Application.CreateForm()
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)
	mainForm.WorkAreaCenter()
	mainForm.SetCaption("表格自绘")
	mainForm.ScaleSelf()
	grid := NewPlayControl(mainForm)
	grid.SetParent(mainForm)
	grid.SetAlign(types.AlClient)
	for i := 1; i <= 100; i++ {
		grid.Add(TPlayListItem{fmt.Sprintf("标题%d", i), "张三", 100000 + rand.Int31n(100000), "", ""})

	}
	vcl.Application.Run()
}
