package main

import (
	"fmt"
	"strings"
	"time"
	"unsafe"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/rtl"
	"github.com/ying32/xldl"
)

//::private::
type TMainFormFields struct {
	dloader *xldl.XLDownloader
}

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
	f.dloader = xldl.NewXLDownloader(rtl.ExtractFilePath(vcl.Application.ExeName()))

	f.EditdlURL.SetText("https://qd.myapp.com/myapp/qqteam/pcqq/PCQQ2020.exe")
}

func (f *TMainForm) OnFormDestroy(sender vcl.IObject) {

}

func (f *TMainForm) OnBtnAddClick(sender vcl.IObject) {
	if !f.indexOfItem(MainForm.EditdlURL.Text()) {
		dlURL := MainForm.EditdlURL.Text()
		item := MainForm.LVDlList.Items().Add()
		item.SetCaption(fmt.Sprintf("%d", MainForm.LVDlList.Items().Count()))
		item.SubItems().Add(dlURL)
		item.SubItems().Add("0M")
		item.SubItems().Add("0%")
		item.SubItems().Add("0KB/s")
		item.SubItems().Add("下载中")
		f.startTask(dlURL, item)
	}
}

func (f *TMainForm) OnBtnStartAllClick(sender vcl.IObject) {
	for _, v := range f.dloader.Tasks {
		v.Start()
	}
}

func (f *TMainForm) OnBtnStopAllClick(sender vcl.IObject) {
	for _, v := range f.dloader.Tasks {
		v.Stop()
	}
}

func (f *TMainForm) OnActStartExecute(sender vcl.IObject) {
	item := f.LVDlList.Selected()
	if item != nil && item.IsValid() && item.Data() != nil {
		(*xldl.XLTask)(unsafe.Pointer(item.Data())).Start()
	}
}

func (f *TMainForm) OnActStopExecute(sender vcl.IObject) {
	item := f.LVDlList.Selected()
	if item != nil && item.IsValid() && item.Data() != nil {
		(*xldl.XLTask)(unsafe.Pointer(item.Data())).Stop()
	}
}

func (f *TMainForm) OnApplicationException(sender vcl.IObject, e *vcl.Exception) {

}

func (f *TMainForm) indexOfItem(s string) bool {
	var i int32
	for i = 0; i < f.LVDlList.Items().Count(); i++ {
		if f.LVDlList.Items().Item(i).SubItems().Strings(0) == s {
			return true
		}
	}
	return false
}

func (f *TMainForm) getURLFileName(s string) string {
	i := strings.LastIndex(s, "/")
	if i > 0 {
		return s[i+1:]
	}
	return ""
}

func (f *TMainForm) startTask(dlURL string, item *vcl.TListItem) {
	task := f.dloader.AddTask(dlURL, f.getURLFileName(dlURL))
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

func (f *TMainForm) OnActDeleteExecute(sender vcl.IObject) {
	item := f.LVDlList.Selected()
	if item != nil && item.IsValid() && item.Data() != nil {
		task := (*xldl.XLTask)(unsafe.Pointer(item.Data()))
		task.Stop()
		task.Delete()
		task.DeleteTempFile()
		item.Delete()
	}
}
