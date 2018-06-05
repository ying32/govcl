package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.SetFormScaled(true)
	mainForm := vcl.Application.CreateForm()
	mainForm.SetWidth(700)
	mainForm.SetHeight(500)
	mainForm.WorkAreaCenter()
	mainForm.SetCaption("九九乘法表")

	grid := vcl.NewStringGrid(mainForm)
	grid.SetParent(mainForm)
	grid.SetAlign(types.AlClient)

	// 这里设置的是作为字段，像头一样的
	//grid.SetFixedCols(1)
	//grid.SetFixedRows(1)

	// 10x10
	grid.SetColCount(10)
	grid.SetRowCount(10)

	var i, j int32
	// 横第一行
	for i = 0; i < grid.ColCount(); i++ {
		grid.SetCells(i, 0, fmt.Sprintf("%d", i))
	}
	// 竖第一行
	for j = 0; j < grid.RowCount(); j++ {
		grid.SetCells(0, j, fmt.Sprintf("%d", j))
	}
	// 九九乘法表
	for i = 1; i <= 9; i++ {
		for j = i; j <= 9; j++ {
			grid.SetCells(i, j, fmt.Sprintf("%dx%d=%d", i, j, i*j))
		}
	}

	vcl.Application.Run()
}
