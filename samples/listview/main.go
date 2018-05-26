package main

import (
	"fmt"
	"math/rand"
	"strings"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/api"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
)

func main() {
	icon := vcl.NewIcon()
	icon.LoadFromResourceID(rtl.MainInstance(), 3)
	defer icon.Free()
	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Hello")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(500)
	mainForm.SetHeight(600)
	// 双缓冲
	mainForm.SetDoubleBuffered(true)

	imgList := vcl.NewImageList(mainForm)
	imgList.AddIcon(icon)

	lv1 := vcl.NewListView(mainForm)
	lv1.SetParent(mainForm)
	lv1.SetAlign(types.AlTop)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetViewStyle(types.VsReport)
	lv1.SetGridLines(true)
	//lv1.SetColumnClick(false)

	col := lv1.Columns().Add()
	col.SetCaption("序号")
	col.SetWidth(100)

	col = lv1.Columns().Add()
	col.SetCaption("项目1")
	col.SetWidth(200)

	lv1.SetOnClick(func(vcl.IObject) {
		if lv1.ItemIndex() != -1 {
			item := lv1.Selected() // lv1.Items().Item(lv1.ItemIndex())
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})

	lv1.SetOnColumnClick(func(sender vcl.IObject, column *vcl.TListColumn) {
		// 按柱头索引排序, lcl兼容版第二个参数永远为 column
		lv1.CustomSort(0, int(column.Index()))
	})

	// 排序事件, lcl自动的，vcl才需要
	if !api.IsloadedLcl {
		lv1.SetOnCompare(func(sender vcl.IObject, item1, item2 *vcl.TListItem, data int32, compare *int32) {
			// lcl data 无效
			if data == 0 {
				*compare = int32(strings.Compare(item1.Caption(), item2.Caption()))
			} else {
				*compare = int32(strings.Compare(item1.SubItems().Strings(data-1), item2.SubItems().Strings(data-1)))
			}
		})
	}

	//	lv1.Clear()
	lv1.Items().BeginUpdate()
	for i := 1; i <= 100; i++ {
		item := lv1.Items().Add()
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i+rand.Int()))
		item.SubItems().Add(fmt.Sprintf("值：%d", i+rand.Int()))
	}
	lv1.Items().EndUpdate()

	// icon样式

	lv2 := vcl.NewListView(mainForm)
	lv2.SetParent(mainForm)
	lv2.SetAlign(types.AlTop)
	lv2.SetRowSelect(true)
	lv2.SetReadOnly(true)
	lv2.SetViewStyle(types.VsIcon)
	//lv2.SetSmallImages(imgList)
	lv2.SetLargeImages(imgList)
	//lv2.SetStateImages(imgList)

	lv2.SetOnClick(func(vcl.IObject) {
		if lv2.ItemIndex() != -1 {
			item := lv2.Selected()
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})
	lv2.Items().BeginUpdate()
	for i := 1; i <= 10; i++ {
		item := lv2.Items().Add()
		item.SetImageIndex(0)
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	lv2.Items().EndUpdate()

	// lv3 Group
	lv3 := vcl.NewListView(mainForm)
	lv3.SetParent(mainForm)
	lv3.SetAlign(types.AlClient)
	lv3.SetRowSelect(true)
	lv3.SetReadOnly(true)
	lv3.SetViewStyle(types.VsReport)
	lv3.SetGridLines(true)

	lv3.SetGroupHeaderImages(imgList)
	lv3.SetGroupView(true)
	col = lv3.Columns().Add()
	col.SetCaption("序号")
	col.SetWidth(100)

	grp := lv3.Groups().Add()
	grp.SetHeader("第一个分组头")
	grp.SetFooter("第一个分组结束")
	grp.SetHeaderAlign(types.TaCenter)
	grp.SetFooterAlign(types.TaCenter)
	grp.SetSubtitle("这是子标题")
	grp.SetTitleImage(0)
	state := grp.State()
	state = types.TListGroupStateSet(rtl.Include(uint32(state), types.LgsCollapsible))
	grp.SetState(state)

	grp = lv3.Groups().Add()
	grp.SetHeader("第二个分组头")
	grp.SetFooter("第二个分组结束")
	grp.SetHeaderAlign(types.TaCenter)
	grp.SetFooterAlign(types.TaCenter)
	grp.SetSubtitle("这是子标题")
	state = grp.State()
	state = types.TListGroupStateSet(rtl.Include(uint32(state), types.LgsCollapsible))
	grp.SetState(state)

	lv3.SetOnClick(func(vcl.IObject) {
		if lv3.ItemIndex() != -1 {
			item := lv3.Selected()
			fmt.Println(item.Caption(), ", ", item.SubItems().Strings(0))
		}
	})
	lv3.Items().BeginUpdate()
	for i := 1; i <= 2; i++ {
		item := lv3.Items().Add()
		item.SetImageIndex(0)
		item.SetGroupID(0)
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	for i := 1; i <= 2; i++ {
		item := lv3.Items().Add()
		item.SetImageIndex(0)
		item.SetGroupID(1)
		// 第一列为Caption属性所管理
		item.SetCaption(fmt.Sprintf("%d", i))
		item.SubItems().Add(fmt.Sprintf("值：%d", i))
	}
	lv3.Items().EndUpdate()

	vcl.Application.Run()
}
