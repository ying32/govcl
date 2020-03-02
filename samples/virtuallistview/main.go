package main

import (
	"fmt"

	"time"

	"math/rand"

	"github.com/ying32/govcl/vcl"
	_ "github.com/ying32/govcl/vcl/exts/winappres"
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
	fmt.Println("t:", ns, "ns, ", ns/1E6, "ms")
	f.ListView.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

}

func (f *TMainFrom) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
	data := tempData[int(item.Index())]
	item.SetCaption(data.Caption)
	item.SubItems().Add(data.Sub1)
	item.SubItems().Add(data.Sub2)
	item.SubItems().Add(data.Sub3)
	item.SubItems().Add(data.Sub4)
	item.SubItems().Add(data.Sub5)
}
