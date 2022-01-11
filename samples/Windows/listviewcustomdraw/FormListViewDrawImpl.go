// 代码由简易GoVCL IDE自动生成。
// 不要更改此文件名
// 在这里写你的事件。

package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl"
)

//::private::
type TFormListViewDrawFields struct {
}

func (f *TFormListViewDraw) OnFormCreate(sender vcl.IObject) {
	// 加载信息
	trainData, err := parseFromFile("testtraindata.json")
	if err == nil {
		// 设计器设计的ListView
		fullResFormListViewInstance(trainData)
		// 代码创建的ListView
		codeBuildistViewInstance(trainData)
	} else {
		fmt.Println("err:", err)
	}
}
