package main

import (
	"fmt"
	"runtime"

	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types/colors"

	"time"

	"math/rand"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

type TMainFrom struct {
	*vcl.TForm
	ListView    *vcl.TListView
	stateImages *vcl.TImageList
	isWindows   bool
}

type TTempItem struct {
	Caption string
	Sub1    string
	Sub2    string
	Sub3    string
	Sub4    string
	Sub5    string
	Checked bool
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
	f.isWindows = runtime.GOOS == "windows"
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

	// 要显示状态图标就得添加
	f.ListView.SetCheckboxes(true)

	// windows下OwnerData不能显示checkbox
	// linux和macOS在OwnerData下支持显示CheckBox
	if f.isWindows {
		// 这里模拟显示
		f.stateImages = vcl.NewImageList(f)
		bmpFileName := "checkbox.png"
		if rtl.FileExists(bmpFileName) {
			pic := vcl.NewPicture()
			pic.LoadFromFile(bmpFileName)
			f.stateImages.AddSliced(pic.Bitmap(), 1, 2)
			pic.Free()
		}
		f.ListView.SetStateImages(f.stateImages)
		f.ListView.SetOnMouseDown(f.OnListView1MouseDown)
	}

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
		tempData[i].Checked = false
	}
	ns := time.Now().UnixNano() - t // 1e-6
	fmt.Println("t:", ns, "ns, ", ns/1e6, "ms")
	f.ListView.Items().SetCount(int32(len(tempData))) //   必须主动的设置Virtual List的行数

	// 解决subitem数据为空时，不能连续绘制背景色
	if f.isWindows {
		f.ListView.SetOnAdvancedCustomDrawItem(f.onAdvancedCustomDrawItem)
	} else {
		f.ListView.SetOnItemChecked(f.onListView1ItemChecked)
	}

}

func (f *TMainFrom) onListView1ItemChecked(sender vcl.IObject, item *vcl.TListItem) {
	tempData[item.Index()].Checked = item.Checked()
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
	if f.isWindows {
		if data.Checked {
			item.SetStateIndex(1)
		} else {
			item.SetStateIndex(0)
		}
	} else {
		item.SetChecked(data.Checked)
	}

	item.SetCaption(data.Caption)
	item.SubItems().Add(data.Sub1)
	item.SubItems().Add("") //data.Sub2)
	item.SubItems().Add(data.Sub3)
	item.SubItems().Add(data.Sub4)
	item.SubItems().Add(data.Sub5)
}

func (f *TMainFrom) OnListView1MouseDown(sender vcl.IObject, button types.TMouseButton, shift types.TShiftState, x, y int32) {
	if f.ListView.Checkboxes() && x <= 16 { //16= f.stateImages.Width
		item := f.ListView.GetItemAt(x, y)
		if item != nil {
			idx := item.Index()
			r := item.DisplayRect(types.DrIcon)
			if y >= r.Top && y <= r.Bottom {
				tempData[idx].Checked = !tempData[idx].Checked
				// 不知道为啥idx=0时要repaint，但Repaint效率不如Invalidate
				if idx == 0 {
					f.ListView.Repaint()
				} else {
					f.ListView.Invalidate()
				}
			}
		}
	}
}
