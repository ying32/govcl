package vcl

import (
	. "gitee.com/ying32/govcl/vcl/api"
)

func (f *TForm) ScreenCenter() {
	f.SetLeft((Screen.Width() - f.Width()) / 2)
	f.SetTop((Screen.Height() - f.Height()) / 2)
}

func (f *TForm) WorkAreaCenter() {
	f.SetLeft((Screen.WorkAreaWidth() - f.Width()) / 2)
	f.SetTop((Screen.WorkAreaHeight() - f.Height()) / 2)
}

func (f *TForm) EnabledMaximize(val bool) {
	Form_EnabledMaximize(f.instance, val)
}

func (f *TForm) EnabledMinimize(val bool) {
	Form_EnabledMinimize(f.instance, val)
}

func (f *TForm) EnabledSystemMenu(val bool) {
	Form_EnabledSystemMenu(f.instance, val)
}

func (f *TForm) ScaleForPPI(val int32) {
	Form_ScaleForPPI(f.instance, val)
}

func (f *TForm) ScaleControlsForDpi(val int32) {
	Form_ScaleControlsForDpi(f.instance, val)
}

func (f *TForm) SetAllowDropFiles(val bool) {
	Form_SetAllowDropFiles(f.instance, val)
}

func (f *TForm) AllowDropFiles() bool {
	return Form_GetAllowDropFiles(f.instance)
}

func (f *TForm) SetOnDropFiles(fn TDropFilesEvent) {
	Form_SetOnDropFiles(f.instance, fn)
}

func (f *TForm) SetOnDestroy(fn TNotifyEvent) {
	Form_SetOnDestroy(f.instance, fn)
}

func (f *TForm) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	Form_SetOnConstrainedResize(f.instance, fn)
}

func (f *TForm) SetOnDeactivate(fn TNotifyEvent) {
	Form_SetOnDeactivate(f.instance, fn)
}

func (f *TForm) SetOnActivate(fn TNotifyEvent) {
	Form_SetOnActivate(f.instance, fn)
}
