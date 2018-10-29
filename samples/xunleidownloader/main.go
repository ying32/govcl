// +build windows,386

// 这个项目只是一个测试例程，随便写下
// 相关迅雷sdk可到迅雷官网下载，或者到 github.com/ying32/xldl 里面下载

package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/govcl/vcl/types"
	"github.com/ying32/xldl"
)

var (
	dloader     *xldl.XLDownloader
	mainFormhWd types.HWND
)

func main() {

	if !xldl.InitXLEngine() {
		vcl.ShowMessage("初始引擎失败！")
		return
	}
	defer xldl.UnInitXLEngine()
	dloader = xldl.NewXLDownloader(rtl.ExtractFilePath(vcl.Application.ExeName()))

	vcl.Application.SetOnException(func(sender vcl.IObject, e *vcl.Exception) {
		// 在这里自行处理VCL中的异常
	})
	vcl.Application.SetIconResId(3) // 具体资源id根据rsrc.exe编译的为准
	vcl.Application.Initialize()
	vcl.Application.SetMainFormOnTaskBar(true)
	vcl.Application.CreateForm(mainFormBytes, &MainForm)

	mainFormhWd = MainForm.Handle()
	MainForm.EditdlURL.SetText("http://sw.bos.baidu.com/sw-search-sp/software/19de58890ffb8/QQ_8.6.18804.0_setup.exe")
	MainForm.BtnAdd.SetOnClick(func(vcl.IObject) {
		if !indexOfItem(MainForm.EditdlURL.Text()) {
			dlURL := MainForm.EditdlURL.Text()
			item := MainForm.LVDlList.Items().Add()
			item.SetCaption(fmt.Sprintf("%d", MainForm.LVDlList.Items().Count()))
			item.SubItems().Add(dlURL)
			item.SubItems().Add("0M")
			item.SubItems().Add("0%")
			item.SubItems().Add("0KB/s")
			item.SubItems().Add("下载中")
			startTask(dlURL, item)
		}
	})

	MainForm.BtnStopAll.SetOnClick(func(vcl.IObject) {
		for _, v := range dloader.Tasks {
			v.Stop()
		}
	})

	MainForm.BtnStartAll.SetOnClick(func(vcl.IObject) {
		for _, v := range dloader.Tasks {
			v.Start()
		}
	})

	MainForm.ActStart.SetOnExecute(func(vcl.IObject) {
		item := MainForm.LVDlList.Selected()
		if item != nil && item.IsValid() && item.Data() != nil {
			(*xldl.XLTask)(unsafe.Pointer(item.Data())).Start()
		}
	})

	MainForm.ActStop.SetOnExecute(func(vcl.IObject) {
		item := MainForm.LVDlList.Selected()
		if item != nil && item.IsValid() && item.Data() != nil {
			(*xldl.XLTask)(unsafe.Pointer(item.Data())).Stop()
		}
	})

	vcl.Application.Run()

}

func indexOfItem(s string) bool {
	var i int32
	for i = 0; i < MainForm.LVDlList.Items().Count(); i++ {
		if MainForm.LVDlList.Items().Item(i).SubItems().Strings(0) == s {
			return true
		}
	}
	return false
}

func getURLFileName(s string) string {
	i := strings.LastIndex(s, "/")
	if i > 0 {
		return s[i+1:]
	}
	return ""
}

func startTask(dlURL string, item *vcl.TListItem) {
	task := dloader.AddTask(dlURL, getURLFileName(dlURL))
	if task.Start() {
		item.SetData(unsafe.Pointer(task))
		go func() {
			for {
				info, ret := task.Info()
				if ret {
					vcl.ThreadSync(func() {
						item.SubItems().SetStrings(1, fmt.Sprintf("%.2fM", float64(info.TotalSize)/1024.0/1024.0))
						item.SubItems().SetStrings(2, fmt.Sprintf("%.2f%%", info.Percent*100))
						item.SubItems().SetStrings(3, fmt.Sprintf("%dKB/s", info.Speed/1024))
						switch info.Stat {
						case xldl.TSC_ERROR:
							item.SubItems().SetStrings(4, "错误")
							break
						case xldl.TSC_PAUSE:
							item.SubItems().SetStrings(4, "暂停")
						case xldl.TSC_DOWNLOAD:
							item.SubItems().SetStrings(4, "下载中")
						case xldl.TSC_COMPLETE:
							item.SubItems().SetStrings(4, "下载完成")
							break
						case xldl.TSC_STARTPENDING:

						case xldl.TSC_STOPPENDING:

						}
					})
				}
				time.Sleep(time.Millisecond * 500)
			}
		}()
	}
}
