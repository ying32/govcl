package main

import (
	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

/*

   这里是使用代码来创建ListView，演示与UI设计器中构建中的ListView一致。部分窗口PageControl依然使用UI设计器的
*/

func codeBuildistViewInstance(trainData *TTrainSearchResultData) {
	// 这里使用了部分UI设计构建的UI
	sheet := vcl.NewTabSheet(FormListViewDraw)
	sheet.SetCaption("代码构建ListView自绘演示")
	sheet.SetPageControl(FormListViewDraw.PageMain)

	lv := vcl.NewListView(FormListViewDraw)
	lv.SetParent(sheet)
	lv.SetAlign(types.AlClient)
	// 这里偷下懒，使用原有的ImageList
	lv.SetSmallImages(FormListViewDraw.ImageList1)
	lv.SetGridLines(true)
	lv.SetGroupView(true)
	lv.SetReadOnly(true)
	lv.SetRowSelect(true)
	lv.SetViewStyle(types.VsReport)
	lv.Font().SetName("微软雅黑")
	lv.Font().SetSize(7)

	addCol := func(caption string, width int32) {
		col := lv.Columns().Add()
		col.SetCaption(caption)
		col.SetWidth(width)
		col.SetAlignment(types.TaCenter)
	}

	addCol("车次", 90)
	addCol("出发站", 85)
	addCol("到达站", 85)
	addCol("历时", 90)
	addCol("商务座", 60)
	addCol("特等座", 60)
	addCol("一等座", 60)
	addCol("二等座", 60)
	addCol("高级软卧", 65)
	addCol("软卧", 60)
	addCol("硬卧", 60)
	addCol("软座", 60)
	addCol("硬座", 60)
	addCol("无座", 60)
	addCol("其它", 60)
	addCol("备注", 120)

	addGroup := func(caption string) {
		g := lv.Groups().Add()
		g.SetFooterAlign(types.TaCenter)
		g.SetHeaderAlign(types.TaCenter)
		g.SetHeader(caption)
		//state := g.State() //默认为0
		state := types.TListGroupStateSet(rtl.Include(0, types.LgsCollapsible))
		g.SetState(state)
		g.SetTitleImage(-1)
	}
	addGroup("高铁/城际")
	addGroup("动车")
	addGroup("直达")
	addGroup("特快")
	addGroup("快速")
	addGroup("其它")

	fullListViewDataAndSetEvent(lv, trainData)
}
