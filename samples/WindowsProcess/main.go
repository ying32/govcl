// +build windows

package main

import (
	"unsafe"

	"path/filepath"

	"fmt"

	"gitee.com/ying32/govcl/vcl"
	"gitee.com/ying32/govcl/vcl/rtl"
	"gitee.com/ying32/govcl/vcl/types"
	"gitee.com/ying32/govcl/vcl/win"
)

func main() {

	vcl.Application.SetIconResId(3)
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)

	mainForm := vcl.Application.CreateForm()
	mainForm.SetCaption("Windows Process")
	mainForm.SetPosition(types.PoScreenCenter)
	mainForm.SetWidth(600)
	mainForm.SetHeight(400)
	// 双缓冲
	mainForm.SetDoubleBuffered(true)

	imgList := vcl.NewImageList(mainForm)
	imgList.SetWidth(24)
	imgList.SetHeight(24)

	lv1 := vcl.NewListView(mainForm)
	lv1.SetParent(mainForm)
	lv1.SetAlign(types.AlClient)
	lv1.SetRowSelect(true)
	lv1.SetReadOnly(true)
	lv1.SetViewStyle(types.VsReport)
	lv1.SetGridLines(true)
	lv1.SetLargeImages(imgList)
	lv1.SetSmallImages(imgList)

	col := lv1.Columns().Add()
	col.SetWidth(60)

	col = lv1.Columns().Add()
	col.SetCaption("进程名")
	col.SetAutoSize(true)

	col = lv1.Columns().Add()
	col.SetCaption("PID")
	col.SetAutoSize(true)

	lv1.Clear()
	imgList.Clear()
	fullListView(lv1, imgList)

	vcl.Application.Run()
}

func fullListView(lv *vcl.TListView, imgList *vcl.TImageList) {
	var fProcessEntry32 win.TProcessEntry32
	fProcessEntry32.DwSize = uint32(unsafe.Sizeof(fProcessEntry32))

	fSnapShotHandle := win.CreateToolhelp32SnapShot(win.TH32CS_SNAPPROCESS, 0)
	continueLoop := win.Process32First(fSnapShotHandle, &fProcessEntry32)
	lv.Items().BeginUpdate()
	defer lv.Items().EndUpdate()

	ico := vcl.NewIcon()
	defer ico.Free()
	var index int32
	for continueLoop {
		item := lv.Items().Add()
		exeFileName := win.GoStr(fProcessEntry32.SzExeFile[:])
		item.SubItems().Add(filepath.Base(exeFileName))
		item.SubItems().Add(fmt.Sprintf("%.4X", fProcessEntry32.Th32ProcessID))
		hIcon := win.ExtractIcon(rtl.MainInstance(), exeFileName, 0)
		if hIcon != 0 {
			ico.SetHandle(hIcon)
			imgList.AddIcon(ico)
			item.SetImageIndex(index)
			index++
		}

		continueLoop = win.Process32Next(fSnapShotHandle, &fProcessEntry32)
	}
	if fSnapShotHandle != 0 {
		win.CloseHandle(fSnapShotHandle)
	}
}
