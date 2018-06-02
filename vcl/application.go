
//----------------------------------------
// 代码由GenlibVcl工具自动生成。
// Copyright © ying32. All Rights Reserved.
//
//----------------------------------------


package vcl


import (
	. "github.com/ying32/govcl/vcl/api"
    . "github.com/ying32/govcl/vcl/types"
)

type TApplication struct {
    IComponent
    instance uintptr
}

func NewApplication(owner IComponent) *TApplication {
    a := new(TApplication)
    a.instance = Application_Create(CheckPtr(owner))
    return a
}

func ApplicationFromInst(inst uintptr) *TApplication {
    a := new(TApplication)
    a.instance = inst
    return a
}

func ApplicationFromObj(obj IObject) *TApplication {
    a := new(TApplication)
    a.instance = CheckPtr(obj)
    return a
}

func (a *TApplication) Free() {
    if a.instance != 0 {
        Application_Free(a.instance)
        a.instance = 0
    }
}

func (a *TApplication) Instance() uintptr {
    return a.instance
}

func (a *TApplication) IsValid() bool {
    return a.instance != 0
}

func TApplicationClass() TClass {
    return Application_StaticClassType()
}

func (a *TApplication) ActivateHint(CursorPos TPoint) {
    Application_ActivateHint(a.instance, CursorPos)
}

func (a *TApplication) BringToFront() {
    Application_BringToFront(a.instance)
}

func (a *TApplication) CancelHint() {
    Application_CancelHint(a.instance)
}

func (a *TApplication) HandleMessage() {
    Application_HandleMessage(a.instance)
}

func (a *TApplication) HideHint() {
    Application_HideHint(a.instance)
}

func (a *TApplication) Minimize() {
    Application_Minimize(a.instance)
}

func (a *TApplication) ModalStarted() {
    Application_ModalStarted(a.instance)
}

func (a *TApplication) ModalFinished() {
    Application_ModalFinished(a.instance)
}

func (a *TApplication) NormalizeAllTopMosts() {
    Application_NormalizeAllTopMosts(a.instance)
}

func (a *TApplication) NormalizeTopMosts() {
    Application_NormalizeTopMosts(a.instance)
}

func (a *TApplication) ProcessMessages() {
    Application_ProcessMessages(a.instance)
}

func (a *TApplication) Restore() {
    Application_Restore(a.instance)
}

func (a *TApplication) RestoreTopMosts() {
    Application_RestoreTopMosts(a.instance)
}

func (a *TApplication) Terminate() {
    Application_Terminate(a.instance)
}

func (a *TApplication) MessageBox(Text string, Caption string, Flags int32) int32 {
    return Application_MessageBox(a.instance, Text , Caption , Flags)
}

func (a *TApplication) FindComponent(AName string) *TComponent {
    return ComponentFromInst(Application_FindComponent(a.instance, AName))
}

func (a *TApplication) GetNamePath() string {
    return Application_GetNamePath(a.instance)
}

func (a *TApplication) HasParent() bool {
    return Application_HasParent(a.instance)
}

func (a *TApplication) Assign(Source IObject) {
    Application_Assign(a.instance, CheckPtr(Source))
}

func (a *TApplication) DisposeOf() {
    Application_DisposeOf(a.instance)
}

func (a *TApplication) ClassType() TClass {
    return Application_ClassType(a.instance)
}

func (a *TApplication) ClassName() string {
    return Application_ClassName(a.instance)
}

func (a *TApplication) InstanceSize() int32 {
    return Application_InstanceSize(a.instance)
}

func (a *TApplication) InheritsFrom(AClass TClass) bool {
    return Application_InheritsFrom(a.instance, AClass)
}

func (a *TApplication) Equals(Obj IObject) bool {
    return Application_Equals(a.instance, CheckPtr(Obj))
}

func (a *TApplication) GetHashCode() int32 {
    return Application_GetHashCode(a.instance)
}

func (a *TApplication) ToString() string {
    return Application_ToString(a.instance)
}

func (a *TApplication) DialogHandle() HWND {
    return Application_GetDialogHandle(a.instance)
}

func (a *TApplication) SetDialogHandle(value HWND) {
    Application_SetDialogHandle(a.instance, value)
}

func (a *TApplication) ExeName() string {
    return Application_GetExeName(a.instance)
}

func (a *TApplication) Hint() string {
    return Application_GetHint(a.instance)
}

func (a *TApplication) SetHint(value string) {
    Application_SetHint(a.instance, value)
}

func (a *TApplication) HintColor() TColor {
    return Application_GetHintColor(a.instance)
}

func (a *TApplication) SetHintColor(value TColor) {
    Application_SetHintColor(a.instance, value)
}

func (a *TApplication) HintHidePause() int32 {
    return Application_GetHintHidePause(a.instance)
}

func (a *TApplication) SetHintHidePause(value int32) {
    Application_SetHintHidePause(a.instance, value)
}

func (a *TApplication) HintPause() int32 {
    return Application_GetHintPause(a.instance)
}

func (a *TApplication) SetHintPause(value int32) {
    Application_SetHintPause(a.instance, value)
}

func (a *TApplication) HintShortCuts() bool {
    return Application_GetHintShortCuts(a.instance)
}

func (a *TApplication) SetHintShortCuts(value bool) {
    Application_SetHintShortCuts(a.instance, value)
}

func (a *TApplication) HintShortPause() int32 {
    return Application_GetHintShortPause(a.instance)
}

func (a *TApplication) SetHintShortPause(value int32) {
    Application_SetHintShortPause(a.instance, value)
}

func (a *TApplication) Icon() *TIcon {
    return IconFromInst(Application_GetIcon(a.instance))
}

func (a *TApplication) SetIcon(value *TIcon) {
    Application_SetIcon(a.instance, CheckPtr(value))
}

func (a *TApplication) IsMetropolisUI() bool {
    return Application_GetIsMetropolisUI(a.instance)
}

func (a *TApplication) MainForm() *TForm {
    return FormFromInst(Application_GetMainForm(a.instance))
}

func (a *TApplication) MainFormHandle() HWND {
    return Application_GetMainFormHandle(a.instance)
}

func (a *TApplication) MainFormOnTaskBar() bool {
    return Application_GetMainFormOnTaskBar(a.instance)
}

func (a *TApplication) SetMainFormOnTaskBar(value bool) {
    Application_SetMainFormOnTaskBar(a.instance, value)
}

func (a *TApplication) BiDiMode() TBiDiMode {
    return Application_GetBiDiMode(a.instance)
}

func (a *TApplication) SetBiDiMode(value TBiDiMode) {
    Application_SetBiDiMode(a.instance, value)
}

func (a *TApplication) ShowHint() bool {
    return Application_GetShowHint(a.instance)
}

func (a *TApplication) SetShowHint(value bool) {
    Application_SetShowHint(a.instance, value)
}

func (a *TApplication) ShowMainForm() bool {
    return Application_GetShowMainForm(a.instance)
}

func (a *TApplication) SetShowMainForm(value bool) {
    Application_SetShowMainForm(a.instance, value)
}

func (a *TApplication) Title() string {
    return Application_GetTitle(a.instance)
}

func (a *TApplication) SetTitle(value string) {
    Application_SetTitle(a.instance, value)
}

func (a *TApplication) SetOnException(fn TExceptionEvent) {
    Application_SetOnException(a.instance, fn)
}

func (a *TApplication) SetOnHelp(fn THelpEvent) {
    Application_SetOnHelp(a.instance, fn)
}

func (a *TApplication) SetOnHint(fn TNotifyEvent) {
    Application_SetOnHint(a.instance, fn)
}

func (a *TApplication) SetOnMessage(fn TMessageEvent) {
    Application_SetOnMessage(a.instance, fn)
}

func (a *TApplication) SetOnMinimize(fn TNotifyEvent) {
    Application_SetOnMinimize(a.instance, fn)
}

func (a *TApplication) SetOnRestore(fn TNotifyEvent) {
    Application_SetOnRestore(a.instance, fn)
}

func (a *TApplication) SetOnShortCut(fn TShortCutEvent) {
    Application_SetOnShortCut(a.instance, fn)
}

func (a *TApplication) Handle() HWND {
    return Application_GetHandle(a.instance)
}

func (a *TApplication) SetHandle(value HWND) {
    Application_SetHandle(a.instance, value)
}

func (a *TApplication) ComponentCount() int32 {
    return Application_GetComponentCount(a.instance)
}

func (a *TApplication) ComponentIndex() int32 {
    return Application_GetComponentIndex(a.instance)
}

func (a *TApplication) SetComponentIndex(value int32) {
    Application_SetComponentIndex(a.instance, value)
}

func (a *TApplication) Owner() *TComponent {
    return ComponentFromInst(Application_GetOwner(a.instance))
}

func (a *TApplication) Name() string {
    return Application_GetName(a.instance)
}

func (a *TApplication) SetName(value string) {
    Application_SetName(a.instance, value)
}

func (a *TApplication) Tag() int {
    return Application_GetTag(a.instance)
}

func (a *TApplication) SetTag(value int) {
    Application_SetTag(a.instance, value)
}

func (a *TApplication) Components(AIndex int32) *TComponent {
    return ComponentFromInst(Application_GetComponents(a.instance, AIndex))
}

