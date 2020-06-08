// +build windows

package main

import (
	"fmt"
	"math"
	"os"
	"path/filepath"

	"github.com/ying32/govcl/vcl/types/colors"

	"github.com/ying32/govcl/vcl/win"

	"time"

	"math/rand"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainFrom struct {
	*vcl.TForm
	ListView *vcl.TListView
	Icons    *vcl.TImageList
	tempIco  *vcl.TIcon
}

type TTempItem struct {
	IconIndex int32
	Sub1      string
	Sub2      string
	Progress  int
	Sub4      string
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

	f.SetDoubleBuffered(true)
	f.SetHeight(600)
	f.SetWidth(800)
	f.ScreenCenter()
	rand.Seed(time.Now().Unix())

	f.tempIco = vcl.NewIcon()

	f.Icons = vcl.NewImageList(f)
	f.Icons.SetHeight(16)
	f.Icons.SetWidth(16)

	f.ListView = vcl.NewListView(f)
	f.ListView.SetAlign(types.AlClient)
	f.ListView.SetParent(f)
	f.ListView.SetViewStyle(types.VsReport)
	f.ListView.SetOwnerData(true)
	f.ListView.SetGridLines(true)
	f.ListView.SetReadOnly(true)
	f.ListView.SetRowSelect(true)
	f.ListView.SetSmallImages(f.Icons)
	//f.ListView.SetOnData(f.OnListView1Data)
	f.ListView.SetOnAdvancedCustomDrawSubItem(f.OnListView1AdvancedCustomDrawSubItem)

	addCol := func(name string, width int32) {
		col := f.ListView.Columns().Add()
		col.SetCaption(name)
		col.SetWidth(width)
		col.SetAlignment(types.TaCenter)
	}

	addCol("", 50)
	addCol("子项1", 100)
	addCol("子项2", 100)
	addCol("进度", 200)
	addCol("M", 100)

	ico := vcl.NewIcon()
	_ = filepath.Walk("./icons", func(path string, info os.FileInfo, err error) error {
		ext := filepath.Ext(info.Name())
		if ext == ".ico" {
			ico.LoadFromFile(path)
			f.Icons.AddIcon(ico)
		}
		return nil
	})
	ico.Free()

	// 产生100w条数据
	tempData = make([]TTempItem, 1000)
	t := time.Now().UnixNano()
	for i := 0; i < len(tempData); i++ {
		tempData[i].IconIndex = rand.Int31n(f.Icons.Count() - 1)
		tempData[i].Sub1 = fmt.Sprintf("子项1:%d", rand.Intn(1000000))
		tempData[i].Sub2 = fmt.Sprintf("子项2:%d", rand.Intn(1000000))
		tempData[i].Progress = rand.Intn(100)
		tempData[i].Sub4 = fmt.Sprintf("%dM", rand.Intn(100))
	}
	ns := time.Now().UnixNano() - t // 1e-6
	fmt.Println("t:", ns, "ns, ", ns/1e6, "ms")
	f.ListView.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

}

func (f *TMainFrom) OnFormDestroy(sender vcl.IObject) {
	if f.tempIco != nil {
		f.tempIco.Free()
	}
}

//func (f *TMainFrom) OnListView1Data(sender vcl.IObject, item *vcl.TListItem) {
//	data := tempData[int(item.Index())]
//	item.SetCaption(data.Caption)
//	item.SubItems().Add(data.Sub1)
//	item.SubItems().Add(data.Sub2)
//	item.SubItems().Add("")
//	item.SubItems().Add(data.Sub4)
//}

func (f *TMainFrom) GetSubItemRect(hwndLV types.HWND, iItem, iSubItem int32) (ret types.TRect) {
	win.ListView_GetSubItemRect(hwndLV, iItem, iSubItem, win.LVIR_LABEL, &ret)
	return
}

func (f *TMainFrom) OnListView1AdvancedCustomDrawSubItem(sender *vcl.TListView, item *vcl.TListItem, subItem int32, state types.TCustomDrawState, stage types.TCustomDrawStage, defaultDraw *bool) {
	if len(tempData) == 0 {
		return
	}
	canvas := sender.Canvas()

	boundRect := item.DisplayRect(types.DrBounds)
	if state.In(types.CdsFocused) {
		canvas.Brush().SetColor(0x00C5F1FF)

	} else {
		canvas.Brush().SetColor(sender.Color())
	}
	canvas.FillRect(boundRect)
	data := tempData[item.Index()]
	drawFlags := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
	var i int32
	for i = 0; i < sender.Columns().Count(); i++ {
		r := f.GetSubItemRect(sender.Handle(), item.Index(), i)
		switch i {
		case 0:
			var hw int32 = 16

			f.Icons.GetIcon(data.IconIndex, f.tempIco)
			if !f.tempIco.Empty() {
				canvas.Draw(r.Right/2-hw/2, r.Top+(r.Bottom-r.Top-hw)/2, f.tempIco)
			}

		case 1:
			canvas.TextRect3(&r, data.Sub1, drawFlags)
		case 2:
			canvas.TextRect3(&r, data.Sub2, drawFlags)
		case 3:

			canvas.Brush().SetColor(colors.ClSkyblue)
			canvas.FrameRect(r)
			r.Inflate(-2, -2)

			right := r.Right - int32(math.Ceil(float64(r.Width())/100.0)*(100.0-float64(data.Progress)))
			if right < 0 {
				right = 0
			}
			progRect := types.TRect{r.Left, r.Top, right, r.Bottom}

			canvas.FillRect(progRect)

			//flags := types.NewSet(types.TfCenter, types.TfSingleLine, types.TfVerticalCenter)
			canvas.Brush().SetStyle(types.BsClear)
			win.SetBkMode(canvas.Handle(), win.TRANSPARENT)
			canvas.TextRect3(&r, fmt.Sprintf("%d%%", data.Progress), drawFlags)

		case 4:

			canvas.TextRect3(&r, data.Sub4, drawFlags)
		}

	}

}
