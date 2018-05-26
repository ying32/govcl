
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "gitee.com/ying32/govcl/vcl/api"
    . "gitee.com/ying32/govcl/vcl/types"
)

type TScreen struct {
    IComponent
    instance uintptr
}

func NewScreen(owner IComponent) *TScreen {
    s := new(TScreen)
    s.instance = Screen_Create(CheckPtr(owner))
    return s
}

func ScreenFromInst(inst uintptr) *TScreen {
    s := new(TScreen)
    s.instance = inst
    return s
}

func ScreenFromObj(obj IObject) *TScreen {
    s := new(TScreen)
    s.instance = CheckPtr(obj)
    return s
}

func (s *TScreen) Free() {
    if s.instance != 0 {
        Screen_Free(s.instance)
        s.instance = 0
    }
}

func (s *TScreen) Instance() uintptr {
    return s.instance
}

func (s *TScreen) IsValid() bool {
    return s.instance != 0
}

func (s *TScreen) Realign() {
    Screen_Realign(s.instance)
}

func (s *TScreen) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Screen_FindComponent(s.instance, AName))
}

func (s *TScreen) GetNamePath() string {
    return Screen_GetNamePath(s.instance)
}

func (s *TScreen) HasParent() bool {
    return Screen_HasParent(s.instance)
}

func (s *TScreen) Assign(Source IObject) {
    Screen_Assign(s.instance, CheckPtr(Source))
}

func (s *TScreen) ClassName() string {
    return Screen_ClassName(s.instance)
}

func (s *TScreen) Equals(Obj IObject) bool {
    return Screen_Equals(s.instance, CheckPtr(Obj))
}

func (s *TScreen) GetHashCode() int32 {
    return Screen_GetHashCode(s.instance)
}

func (s *TScreen) ToString() string {
    return Screen_ToString(s.instance)
}

func (s *TScreen) ActiveForm() *TForm {
    return FormFromInst(Screen_GetActiveForm(s.instance))
}

func (s *TScreen) CustomFormCount() int32 {
    return Screen_GetCustomFormCount(s.instance)
}

func (s *TScreen) CursorCount() int32 {
    return Screen_GetCursorCount(s.instance)
}

func (s *TScreen) Cursor() TCursor {
    return Screen_GetCursor(s.instance)
}

func (s *TScreen) SetCursor(value TCursor) {
    Screen_SetCursor(s.instance, value)
}

func (s *TScreen) FocusedForm() *TForm {
    return FormFromInst(Screen_GetFocusedForm(s.instance))
}

func (s *TScreen) SetFocusedForm(value IControl) {
    Screen_SetFocusedForm(s.instance, CheckPtr(value))
}

func (s *TScreen) MonitorCount() int32 {
    return Screen_GetMonitorCount(s.instance)
}

func (s *TScreen) DesktopRect() TRect {
    return Screen_GetDesktopRect(s.instance)
}

func (s *TScreen) DesktopHeight() int32 {
    return Screen_GetDesktopHeight(s.instance)
}

func (s *TScreen) DesktopLeft() int32 {
    return Screen_GetDesktopLeft(s.instance)
}

func (s *TScreen) DesktopTop() int32 {
    return Screen_GetDesktopTop(s.instance)
}

func (s *TScreen) DesktopWidth() int32 {
    return Screen_GetDesktopWidth(s.instance)
}

func (s *TScreen) WorkAreaRect() TRect {
    return Screen_GetWorkAreaRect(s.instance)
}

func (s *TScreen) WorkAreaHeight() int32 {
    return Screen_GetWorkAreaHeight(s.instance)
}

func (s *TScreen) WorkAreaLeft() int32 {
    return Screen_GetWorkAreaLeft(s.instance)
}

func (s *TScreen) WorkAreaTop() int32 {
    return Screen_GetWorkAreaTop(s.instance)
}

func (s *TScreen) WorkAreaWidth() int32 {
    return Screen_GetWorkAreaWidth(s.instance)
}

func (s *TScreen) Fonts() *TStrings {
    return StringsFromInst(Screen_GetFonts(s.instance))
}

func (s *TScreen) FormCount() int32 {
    return Screen_GetFormCount(s.instance)
}

func (s *TScreen) Imes() *TStrings {
    return StringsFromInst(Screen_GetImes(s.instance))
}

func (s *TScreen) DefaultIme() string {
    return Screen_GetDefaultIme(s.instance)
}

func (s *TScreen) Height() int32 {
    return Screen_GetHeight(s.instance)
}

func (s *TScreen) PixelsPerInch() int32 {
    return Screen_GetPixelsPerInch(s.instance)
}

func (s *TScreen) PrimaryMonitor() *TMonitor {
    return MonitorFromInst(Screen_GetPrimaryMonitor(s.instance))
}

func (s *TScreen) Width() int32 {
    return Screen_GetWidth(s.instance)
}

func (s *TScreen) ComponentCount() int32 {
    return Screen_GetComponentCount(s.instance)
}

func (s *TScreen) ComponentIndex() int32 {
    return Screen_GetComponentIndex(s.instance)
}

func (s *TScreen) SetComponentIndex(value int32) {
    Screen_SetComponentIndex(s.instance, value)
}

func (s *TScreen) Owner() *TComponent {
    return ComponentFromInst(Screen_GetOwner(s.instance))
}

func (s *TScreen) Name() string {
    return Screen_GetName(s.instance)
}

func (s *TScreen) SetName(value string) {
    Screen_SetName(s.instance, value)
}

func (s *TScreen) Tag() int {
    return Screen_GetTag(s.instance)
}

func (s *TScreen) SetTag(value int) {
    Screen_SetTag(s.instance, value)
}

func (s *TScreen) Cursors(Index int32) HICON {
    return Screen_GetCursors(s.instance, Index)
}

func (s *TScreen) SetCursors(Index int32, value HICON) {
    Screen_SetCursors(s.instance, Index, value)
}

func (s *TScreen) Monitors(Index int32) *TMonitor {
    return MonitorFromInst(Screen_GetMonitors(s.instance, Index))
}

func (s *TScreen) Forms(Index int32) *TForm {
    return FormFromInst(Screen_GetForms(s.instance, Index))
}

func (s *TScreen) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Screen_GetComponents(s.instance, AIndex))
}

