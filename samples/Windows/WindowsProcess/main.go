// +build windows

package main

import (
	"unsafe"

	"github.com/ying32/govcl/vcl/types"

	"github.com/ying32/govcl/vcl/rtl"

	"path/filepath"

	"fmt"

	_ "github.com/ying32/govcl/pkgs/winappres"
	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/win"
)

type TMainForm struct {
	*vcl.TForm
	imgList *vcl.TImageList
	lv1     *vcl.TListView
}

var mainForm *TMainForm

func (f *TMainForm) OnFormCreate(object vcl.IObject) {
	f.SetCaption("Windows Process")
	f.SetWidth(600)
	f.SetHeight(400)
	// 双缓冲
	f.SetDoubleBuffered(true)

	f.ScreenCenter()

	f.initComponents()
}

func (f *TMainForm) initComponents() {

	f.imgList = vcl.NewImageList(f)
	f.imgList.SetWidth(24)
	f.imgList.SetHeight(24)

	f.lv1 = vcl.NewListView(f)
	f.lv1.SetParent(f)
	f.lv1.SetAlign(types.AlClient)
	f.lv1.SetRowSelect(true)
	f.lv1.SetReadOnly(true)
	f.lv1.SetViewStyle(types.VsReport)
	f.lv1.SetGridLines(true)
	f.lv1.SetLargeImages(f.imgList)
	f.lv1.SetSmallImages(f.imgList)

	col := f.lv1.Columns().Add()
	col.SetWidth(60)

	col = f.lv1.Columns().Add()
	col.SetCaption("进程名")
	col.SetAutoSize(true)

	col = f.lv1.Columns().Add()
	col.SetCaption("PID")
	col.SetAutoSize(true)

	f.lv1.Clear()
	f.imgList.Clear()
	f.fullListView()

}

func (f *TMainForm) fullListView() {
	var fProcessEntry32 win.TProcessEntry32
	fProcessEntry32.DwSize = uint32(unsafe.Sizeof(fProcessEntry32))

	fSnapShotHandle := win.CreateToolhelp32SnapShot(win.TH32CS_SNAPPROCESS, 0)
	continueLoop := win.Process32First(fSnapShotHandle, &fProcessEntry32)
	f.lv1.Items().BeginUpdate()
	defer f.lv1.Items().EndUpdate()

	ico := vcl.NewIcon()
	ico.SetTransparent(true)
	defer ico.Free()
	for continueLoop {
		item := f.lv1.Items().Add()
		exeFileName := win.GoStr(fProcessEntry32.SzExeFile[:])
		item.SubItems().Add(filepath.Base(exeFileName))
		item.SubItems().Add(fmt.Sprintf("%.4X", fProcessEntry32.Th32ProcessID))

		hProcess := win.OpenProcess(win.PROCESS_QUERY_INFORMATION|win.PROCESS_VM_READ, false, fProcessEntry32.Th32ProcessID)
		if hProcess > 0 {
			fullFileName, _ := win.GetModuleFileNameEx(hProcess, 0)
			fmt.Println(fullFileName)

			hIcon := win.ExtractIcon(rtl.MainInstance(), fullFileName, 0)
			if hIcon != 0 {
				ico.SetHandle(hIcon)
				index := f.imgList.AddIcon(ico)
				item.SetImageIndex(index)
			}
			win.CloseHandle(hProcess)
		}

		continueLoop = win.Process32Next(fSnapShotHandle, &fProcessEntry32)
	}
	if fSnapShotHandle != 0 {
		win.CloseHandle(fSnapShotHandle)
	}
}

func main() {
	vcl.RunApp(&mainForm)
}
