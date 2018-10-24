package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
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

// SetOnStyleChanged 样式已改变事件，为解决样式切换后的问题特意加上的
func (f *TForm) SetOnStyleChanged(fn TNotifyEvent) {
	Form_SetOnStyleChanged(f.instance, fn)
}

// ScaleSelf 这个方法主要是用于当不使用资源窗口创建时用，这个方法要用于设置了Width, Height或者ClientWidth、ClientHeight之后
func (f *TForm) ScaleSelf() {
	if globalFormScaled {
		f.SetClientWidth(int32(float64(f.ClientWidth()) * (float64(Screen.PixelsPerInch()) / 96.0)))
		f.SetClientHeight(int32(float64(f.ClientHeight()) * (float64(Screen.PixelsPerInch()) / 96.0)))
	}
}
