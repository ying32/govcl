//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	. "github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// 从资源中创建Form，不使用Application.CreateForm
func CreateResForm(owner IComponent, fields ...interface{}) {
	resObjtBuild(1, owner, 0, fields...)
}

// 居于当前屏幕中心
func (f *TForm) ScreenCenter() {
	f.SetLeft((Screen.Width() - f.Width()) / 2)
	f.SetTop((Screen.Height() - f.Height()) / 2)
}

// 窗口居于工作区中心，工作区为当前屏幕 - 任务栏空间
func (f *TForm) WorkAreaCenter() {
	f.SetLeft((Screen.WorkAreaWidth() - f.Width()) / 2)
	f.SetTop((Screen.WorkAreaHeight() - f.Height()) / 2)
}

// 启用/禁用 标题栏最大化按钮
func (f *TForm) EnabledMaximize(val bool) {
	Form_EnabledMaximize(f.instance, val)
}

// 启用/禁用 标题栏最小化按钮
func (f *TForm) EnabledMinimize(val bool) {
	Form_EnabledMinimize(f.instance, val)
}

// 启用/禁用 标题栏系统菜单
func (f *TForm) EnabledSystemMenu(val bool) {
	Form_EnabledSystemMenu(f.instance, val)
}

// 窗口PPI缩放
func (f *TForm) ScaleForPPI(val int32) {
	Form_ScaleForPPI(f.instance, val)
}

func (f *TForm) ScaleControlsForDpi(val int32) {
	Form_ScaleControlsForDpi(f.instance, val)
}

// 设置允许拖放文件
func (f *TForm) SetAllowDropFiles(val bool) {
	Form_SetAllowDropFiles(f.instance, val)
}

// AllowDropFiles 获取允许拖放文件
func (f *TForm) AllowDropFiles() bool {
	return Form_GetAllowDropFiles(f.instance)
}

// 窗口文件拖放事件
func (f *TForm) SetOnDropFiles(fn TDropFilesEvent) {
	Form_SetOnDropFiles(f.instance, fn)
}

// 窗口销毁事件
func (f *TForm) SetOnDestroy(fn TNotifyEvent) {
	Form_SetOnDestroy(f.instance, fn)
}

// 约束窗口大小事件
func (f *TForm) SetOnConstrainedResize(fn TConstrainedResizeEvent) {
	Form_SetOnConstrainedResize(f.instance, fn)
}

// 窗口被取消激活事件（失去焦点）
func (f *TForm) SetOnDeactivate(fn TNotifyEvent) {
	Form_SetOnDeactivate(f.instance, fn)
}

// 窗口激活事件（获取焦点）
func (f *TForm) SetOnActivate(fn TNotifyEvent) {
	Form_SetOnActivate(f.instance, fn)
}

// 样式已改变事件，为解决样式切换后的问题特意加上的
func (f *TForm) SetOnStyleChanged(fn TNotifyEvent) {
	Form_SetOnStyleChanged(f.instance, fn)
}

// 窗口消息过程
func (f *TForm) SetOnWndProc(fn TWndProcEvent) {
	Form_SetOnWndProc(f.instance, fn)
}

// 设置窗口显示在任务栏上
func (f *TForm) SetShowInTaskBar(val types.TShowInTaskbar) {
	Form_SetShowInTaskBar(f.instance, val)
}

// 获取当前窗口是否显示在任务栏上
func (f *TForm) ShowInTaskBar() types.TShowInTaskbar {
	return Form_ShowInTaskBar(f.instance)
}

// 这个方法主要是用于当不使用资源窗口创建时用，这个方法要用于设置了Width, Height或者ClientWidth、ClientHeight之后
func (f *TForm) ScaleSelf() {
	if globalFormScaled {
		f.SetClientWidth(int32(float64(f.ClientWidth()) * (float64(Screen.PixelsPerInch()) / 96.0)))
		f.SetClientHeight(int32(float64(f.ClientHeight()) * (float64(Screen.PixelsPerInch()) / 96.0)))
	}
}

// Delphi Loaded方法中的
/// <summary>
/// Checks if there is a change in dpi and perform the necessary changes to scale all
/// the controls for the new dpi
/// </summary>
func (f *TForm) ScaleForCurrentDpi() {
	Form_ScaleForCurrentDpi(f.instance)
}

// OnWndProc必须要调用的， 内部为  inherited WndProc(msg)
func (f *TForm) InheritedWndProc(msg *types.TMessage) {
	Form_InheritedWndProc(f.instance, msg)
}
