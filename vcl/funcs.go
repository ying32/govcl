//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package vcl

import (
	"reflect"

	"fmt"

	"github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// ShowMessage
//
// 显示一个消息框。
//
// Show a message box.
func ShowMessage(msg string) {
	api.DShowMessage(msg)
}

// ShowMessageFmt
//
// 显示一个格式化文本的消息框。
//
// Show a message box with formatted text.
//go:noinline
func ShowMessageFmt(format string, args ...interface{}) {
	ShowMessage(fmt.Sprintf(format, args...))
}

// MessageDlg
//
// 消息框，Buttons为按钮样式，详见types.TMsgDlgButtons。
//
// Message box, Buttons is the button style. For details, see types.TMsgDlgButtons.
func MessageDlg(Msg string, DlgType TMsgDlgType, Buttons ...uint8) int32 {
	return api.DMessageDlg(Msg, DlgType, NewSet(Buttons...), 0)
}

// CheckPtr
//
// 检测接口是否被实例化，如果已经实例化则返回实例指针。
//
// Checks if the interface is instantiated, and returns an instance pointer if it has been instantiated.
//go:noinline
func CheckPtr(value interface{}) uintptr {
	switch value.(type) {
	case IObject:
		if reflect.ValueOf(value).Pointer() == 0 {
			return 0
		}
		return value.(IObject).Instance()
	}
	return 0
}

// SelectDirectory1
//
// 选择目录对话框。
//
// Select Directory dialog.
func SelectDirectory1(options TSelectDirOpts) (bool, string) {
	return api.DSelectDirectory1(options)
}

// SelectDirectory2
//
// 选择目录对话框，一般options默认是SdNewUI，parent默认为nil。
//
// Select the directory dialog box. the default options are SdNewUI, and the default parent is nil.
func SelectDirectory2(caption, root string, showHidden bool) (bool, string) {
	return api.DSelectDirectory2(caption, root, showHidden)
}

// ThreadSync
//
// 主线程中执行。
//
// Executed in the main thread.
func ThreadSync(fn TThreadProc) {
	api.DSynchronize(fn, 1)
}

// InputBox
//
// 输入框。
//
// Input box.
func InputBox(aCaption, aPrompt, aDefault string) string {
	return api.DInputBox(aCaption, aPrompt, aDefault)
}

// InputQuery
//
// 输入框。
//
// Input box.
func InputQuery(aCaption, aPrompt string, value *string) bool {
	return api.DInputQuery(aCaption, aPrompt, value)
}

// PasswordBox
//
// 密码输入框。
//
// Password input box.
func PasswordBox(aCaption, aPrompt string) string {
	return api.DPasswordBox(aCaption, aPrompt)
}

func InputCombo(aCaption, aPrompt string, aList []string) int32 {
	obj := NewStringList()
	defer obj.Free()
	obj.AddStrings2(aList)
	return InputCombo2(aCaption, aPrompt, obj)
}

func InputCombo2(aCaption, aPrompt string, aList IStrings) int32 {
	return api.DInputCombo(aCaption, aPrompt, CheckPtr(aList))
}

func InputComboEx(aCaption, aPrompt string, aList []string, allowCustomText bool) string {
	obj := NewStringList()
	defer obj.Free()
	obj.AddStrings2(aList)
	return InputComboEx2(aCaption, aPrompt, obj, allowCustomText)
}

func InputComboEx2(aCaption, aPrompt string, aList IStrings, allowCustomText bool) string {
	return api.DInputComboEx(aCaption, aPrompt, CheckPtr(aList), allowCustomText)
}

// RunApp
//
// 简化运行。
//
// simplify running.
func RunApp(values ...interface{}) {
	Application.Initialize()
	Application.SetMainFormOnTaskBar(true)
	for i := 0; i < len(values); i++ {
		switch values[i].(type) {
		case func():
			values[i].(func())()
		default:
			Application.CreateForm(values[i])
		}
	}
	Application.Run()
}

// EqualsObject
//
// 比较两个对象
func EqualsObject(obj1, obj2 IObject) bool {
	return CheckPtr(obj1) == CheckPtr(obj2)
}

func FindControl(handle HWND) *TWinControl {
	return AsWinControl(api.DFindControl(handle))
}

func FindLCLControl(screenPos TPoint) *TControl {
	return AsControl(api.DFindLCLControl(screenPos))
}

func FindOwnerControl(handle HWND) *TWinControl {
	return AsWinControl(api.DFindOwnerControl(handle))
}

func FindControlAtPosition(position TPoint, allowDisabled bool) *TControl {
	return AsControl(api.DFindControlAtPosition(position, allowDisabled))
}

func FindLCLWindow(screenPos TPoint, allowDisabled bool) *TWinControl {
	return AsWinControl(api.DFindLCLWindow(screenPos, allowDisabled))
}

func FindDragTarget(position TPoint, allowDisabled bool) *TControl {
	return AsControl(api.DFindDragTarget(position, allowDisabled))
}
