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
	"unsafe"

	"fmt"

	"github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

// CN: 显示一个消息框。
// EN: Show a message box.
func ShowMessage(msg string) {
	api.DShowMessage(msg)
}

// CN: 显示一个格式化文本的消息框。
// EN: Show a message box with formatted text.
//go:noinline
func ShowMessageFmt(format string, args ...interface{}) {
	ShowMessage(fmt.Sprintf(format, args...))
}

// CN: 消息框，Buttons为按钮样式，详见types.TMsgDlgButtons。
// EN: Message box, Buttons is the button style. For details, see types.TMsgDlgButtons.
func MessageDlg(Msg string, DlgType TMsgDlgType, Buttons ...uint8) int32 {
	return api.DMessageDlg(Msg, DlgType, NewSet(Buttons...), 0)
}

// CN: 检测接口是否被实例化，如果已经实例化则返回实例指针。
// EN: Checks if the interface is instantiated, and returns an instance pointer if it has been instantiated.
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

// CN: As操作的简化。
// EN: Simplification of As operation.
//go:noinline
func getInstance(value interface{}) (uintptr, unsafe.Pointer) {
	var ptr uintptr
	switch value.(type) {
	case uintptr:
		// CN: 一个对象来自已经存在的对象实例指针
		// EN: an object from a pointer to an existing object instance
		ptr = value.(uintptr)
	case unsafe.Pointer:
		// CN: 一个对象来自不安全的地址。注意：使用此函数可能造成一些不明情况，慎用。
		// EN: An object from an unsafe address. Note: Using this function may cause some unknown situations. Use it with caution.
		ptr = uintptr(value.(unsafe.Pointer))
	case IObject:
		// CN: 一个对象来自已经存在的对象实例。
		// EN: An object from an existing object instance.
		ptr = CheckPtr(value)
	default:
		// 尝试
		ptr = getUIntPtr(value)
	}
	return ptr, unsafe.Pointer(ptr)
}

func getUIntPtr(v interface{}) uintptr {
	switch v.(type) {
	case int:
		return uintptr(v.(int))
	case uint:
		return uintptr(v.(uint))
	case int8:
		return uintptr(v.(int8))
	case uint8:
		return uintptr(v.(uint8))
	case int16:
		return uintptr(v.(int16))
	case uint16:
		return uintptr(v.(uint16))
	case int32:
		return uintptr(v.(int32))
	case uint32:
		return uintptr(v.(uint32))
	case int64:
		return uintptr(v.(int64))
	case uint64:
		return uintptr(v.(uint64))
	}
	return 0
}

// CN: 选择目录。
// EN: Select directory.
func SelectDirectory1(options TSelectDirOpts) (bool, string) {
	return api.DSelectDirectory1(options)
}

// CN: 选择目录，一般options默认是SdNewUI，parent默认为nil。
// EN: Select directory, options defaults to SdNewUI, parent defaults to nil.
func SelectDirectory2(caption, root string, options TSelectDirExtOpts, parent IObject) (bool, string) {
	return api.DSelectDirectory2(caption, root, options, CheckPtr(parent))
}

// CN: 选择目录， options默认是SdNewUI，parent默认为nil。
// EN: Select directory, options defaults to SdNewUI, parent defaults to nil.
func SelectDirectory3(caption, root string, options ...uint8) (bool, string) {
	opts := NewSet(options...)
	if len(options) == 0 {
		opts = opts.Include(SdNewUI)
	}
	return SelectDirectory2(caption, root, opts, nil)
}

// CN: 主线程中执行。
// EN: Executed in the main thread.
func ThreadSync(fn TThreadProc) {
	api.DSynchronize(fn, 1)
}

// CN: 输入框。
// EN: Input box.
func InputBox(aCaption, aPrompt, aDefault string) string {
	return api.DInputBox(aCaption, aPrompt, aDefault)
}

// CN: 输入框。
// EN: Input box.
func InputQuery(aCaption, aPrompt string, value *string) bool {
	return api.DInputQuery(aCaption, aPrompt, value)
}

// CN: 简化运行。
// EN: simplify running.
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

// CN: 当前是否使用LCL库。
// EN: Whether it is currently an LCL library.
// Deprecated
func LclLoaded() bool {
	return true
}

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
