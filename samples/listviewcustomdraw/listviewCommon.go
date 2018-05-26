package main

import (
	"strings"

	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/types"
)

var (
	fSortOrder bool
)

const (
	Color1 = 0x02E0F0F7
	Color2 = 0x02F0EEF7
)

// 这里是两个ListView实例共用代码，不然写得太多了

func fullListViewDataAndSetEvent(lv *vcl.TListView, trainData *TTrainSearchResultData) {
	if trainData != nil {
		// 填充数据到ListView
		var lG, lD, lZ, lT, lK, lQT int
		//FormListViewDraw.ImageList1 这个imagelist内部没有项目的，用来撑大ListItem，设置了W=1, H=40
		//FormListViewDraw.LVTrain.SetGroupView(false)
		//FormListViewDraw.LVTrain.SetGridLines(false)
		lv.Items().BeginUpdate()
		for i, data := range trainData.Data.Datas {

			item := lv.Items().Add()

			groupId := trainCodeToGroupId(data.StationTrainCode)
			switch groupId {
			case 0:
				lG++
			case 1:
				lD++
			case 2:
				lZ++
			case 3:
				lT++
			case 4:
				lK++
			default:
				lQT++
			}

			item.SetGroupID(groupId)
			item.SetCaption(data.StationTrainCode)
			item.SetData(uintptr(i))

			subItem := item.SubItems()
			subItem.Add(fmt.Sprintf("%s\r\n(%s)(%s)", data.FromStationName, data.StartTime, retStartorEnd(data.StartStationTelecode, data.FromStationTelecode, "始")))
			subItem.Add(fmt.Sprintf("%s\r\n(%s)(%s)", data.ToStationName, data.ArriveTime, retStartorEnd(data.EndStationTelecode, data.ToStationTelecode, "终")))
			subItem.Add(fmt.Sprintf("%s\r\n(%s)", getTimeStr(data.LiShiValue), dayDifference(data.DayDifference)))
			subItem.Add(data.SWZNum)
			subItem.Add(data.TZNnum)
			subItem.Add(data.ZYNum)
			subItem.Add(data.ZENum)
			subItem.Add(data.GRNum)
			subItem.Add(data.RWNum)
			subItem.Add(data.YWNum)
			subItem.Add(data.RZNum)
			subItem.Add(data.YZNum)
			subItem.Add(data.WZNum)
			subItem.Add(data.QTNum)
			subItem.Add(data.Note)
		}
		groups := lv.Groups()
		groups.Items(0).SetFooter(fmt.Sprintf("共%d车次 高铁/城际", lG))
		groups.Items(1).SetFooter(fmt.Sprintf("共%d车次 动车", lD))
		groups.Items(2).SetFooter(fmt.Sprintf("共%d车次 直达", lZ))
		groups.Items(3).SetFooter(fmt.Sprintf("共%d车次 特快", lT))
		groups.Items(4).SetFooter(fmt.Sprintf("共%d车次 快速", lK))
		groups.Items(5).SetFooter(fmt.Sprintf("共%d车次 其它", lQT))

		lv.Items().EndUpdate()
	}
	// 相关事件
	lv.SetOnAdvancedCustomDrawSubItem(lvTraiAdvancedCustomDrawSubItem)
	lv.SetOnAdvancedCustomDrawItem(lvTraiAdvancedCustomDrawItem)
	lv.SetOnColumnClick(lvTraiColumnClick)
	lv.SetOnCompare(lvTraiCompare)
	lv.SetOnClick(func(sender vcl.IObject) {
		sel := vcl.ListViewFromObj(sender).Selected()
		if sel.IsValid() {
			fmt.Println("select, index:", sel.Index(), ", caption:", sel.Caption(), ", data:", sel.Data())
		}
	})
}

// 子项目绘制函数
func lvTraiAdvancedCustomDrawSubItem(sender *vcl.TListView, item *vcl.TListItem, subItem int32,
	state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool) {
	canvas := sender.Canvas()
	font := canvas.Font()
	// 10 行后开始绘制，前面用于其它演示
	// 演示数据的使用，，，，， 如果使用了GroupView的话，因为分组排序问题会造成不对的，哈哈哈。。。
	i := int(item.Data()) //item.Index()
	if i > 10 {
		if i%2 == 0 {
			canvas.Brush().SetColor(Color1)
		} else {
			canvas.Brush().SetColor(Color2)
		}
	}
	switch {
	case subItem >= 4 && subItem <= 14:
		s := item.SubItems().Strings(subItem - 1)
		if strings.Contains(s, "*") || strings.Contains(s, "--") || strings.Contains(s, "无") {
			font.SetColor(types.ClSilver)
		} else {
			font.SetColor(types.ClGreen)
		}
	default:
		font.SetColor(types.ClBlack)
	}
}

// 演示隔行换色，如果使用lvTraiAdvancedCustomDrawSubItem来处理子项目，加上相关代码
func lvTraiAdvancedCustomDrawItem(sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool) {
	//*defaultDraw = false

	// 演示数据的使用，，，，， 如果使用了GroupView的话，因为分组排序问题会造成不对的，哈哈哈。。。
	i := int(item.Data()) //item.Index()
	// 10 行后开始绘制，前面用于其它演示
	if i > 10 {
		if i%2 == 0 {
			sender.Canvas().Brush().SetColor(Color1)
		} else {
			sender.Canvas().Brush().SetColor(Color2)
		}
		//sender.Canvas().FillRect(item.DisplayRect(types.DrBounds))
	}
}

// 柱头单击
func lvTraiColumnClick(sender vcl.IObject, column *vcl.TListColumn) {
	fSortOrder = !fSortOrder
	//vcl.ListViewFromObj(sender).AlphaSort()
	vcl.ListViewFromObj(sender).CustomSort(0, int(column.Index()))
}

// 排序
func lvTraiCompare(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
	var s1, s2 string
	if data != 0 {
		s1 = item1.SubItems().Strings(data - 1)
		s2 = item2.SubItems().Strings(data - 1)
	} else {
		s1 = item1.Caption()
		s2 = item2.Caption()
	}
	if fSortOrder {
		*compare = int32(strings.Compare(s1, s2))
	} else {
		*compare = -int32(strings.Compare(s1, s2))
	}
}
