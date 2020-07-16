package main

import (
	"fmt"

	"github.com/ying32/govcl/vcl/types/colors"

	"time"

	"math/rand"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainFrom struct {
	*vcl.TForm
	ListView *vcl.TListView
}

type TTempItem struct {
	Caption string
	Sub1    string
	Sub2    string
	Sub3    string
	Sub4    string
	Sub5    string
}

var (
	MainFrom *TMainFrom
	tempData []TTempItem
)

func main() {
	vcl.Application.Initialize()
	vcl.Application.CreateForm(&MainFrom)
	vcl.Application.Run()
}

func (f *TMainFrom) OnFormCreate(sender vcl.IObject) {
	fmt.Println("OnCreate")
	f.SetWidth(800)
	f.SetHeight(600)
	f.ScreenCenter()
	f.SetDoubleBuffered(true)

	f.ListView = vcl.NewListView(f)
	f.ListView.SetAlign(types.AlClient)
	f.ListView.SetParent(f)
	f.ListView.SetViewStyle(types.VsReport)
	f.ListView.SetOwnerData(true)
	f.ListView.SetGridLines(true)
	f.ListView.SetReadOnly(true)
	f.ListView.SetRowSelect(true)
	f.ListView.SetOnData(f.OnListView1Data)

	col := f.ListView.Columns().Add()
	col.SetCaption("行号")
	col.SetWidth(100)

	col = f.ListView.Columns().Add()
	col.SetCaption("子项1")
	col.SetWidth(100)

	col = f.ListView.Columns().Add()
	col.SetCaption("子项2")
	col.SetWidth(100)

	col = f.ListView.Columns().Add()
	col.SetCaption("子项3")
	col.SetWidth(100)

	col = f.ListView.Columns().Add()
	col.SetCaption("子项4")
	col.SetWidth(100)

	col = f.ListView.Columns().Add()
	col.SetCaption("子项5")
	col.SetWidth(100)

	// 产生100w条数据
	tempData = make([]TTempItem, 1000000)
	t := time.Now().UnixNano()
	for i := 0; i < len(tempData); i++ {
		tempData[i].Caption = fmt.Sprintf("%d", i+1)
		tempData[i].Sub1 = fmt.Sprintf("子项1:%d", rand.Intn(1000000))
		tempData[i].Sub2 = fmt.Sprintf("子项2:%d", rand.Intn(1000000))
		tempData[i].Sub3 = fmt.Sprintf("子项3:%d", rand.Intn(1000000))
		tempData[i].Sub4 = fmt.Sprintf("子项4:%d", rand.Intn(1000000))
		tempData[i].Sub5 = fmt.Sprintf("子项5:%d", rand.Intn(1000000))
	}
	ns := time.Now().UnixNano() - t // 1e-6
	fmt.Println("t:", ns, "ns, ", ns/1e6, "ms")
	f.ListView.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

	// 解决subitem数据为空时，不能连续绘制背景色
	f.ListView.SetOnAdvancedCustomDrawItem(f.onAdvancedCustomDrawItem)

}

func (f *TMainFrom) onAdvancedCustomDrawItem(sender *vcl.TListView, item *vcl.TListItem, state types.TCustomDrawState, Stage types.TCustomDrawStage, defaultDraw *bool) {
	if state.In(types.CdsSelected) && Stage == types.CdPrePaint {
		r := item.DisplayRect(types.DrBounds)
		canvas := sender.Canvas()
		canvas.Brush().SetColor(types.TColor(colors.RGB(0, 120, 215)))
		canvas.FillRect(r)
	}
}

func (f *TMainFrom) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	data := tempData[int(item.Index())]
	item.SetCaption(data.Caption)
	item.SubItems().Add(data.Sub1)
	item.SubItems().Add("") //data.Sub2)
	item.SubItems().Add(data.Sub3)
	item.SubItems().Add(data.Sub4)
	item.SubItems().Add(data.Sub5)
}
