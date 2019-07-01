这里只是演示使用方法，实际上此包无效。

如果需要使用则需要自己简单维护一套libvcl或者liblcl。   

目前govcl中提供了一项在不修改原govcl和libvcl/liblcl代码的情况下扩展自己的组件。  

[查看演示例子](https://gitee.com/ying32/govcl/tree/master/samples/myext)    

**组件生产工具请加入QQ群后，在群文件中获取。**  

使用方法：  

### go代码部分  
----

1、新建一个go的package，可以是 github.com/yourusername/msrdp 。这里演示的包名为msrdp。  

2、在msrdp目录下新建一个init.go文件，并写入：  
```go
package msrdp

import "github.com/ying32/govcl/vcl"

// 初始
func init() {
    vcl.RegisterExtEventCallback(eventCallback)
}

``` 

3、在msrdp目录下新建一个callback.go文件，并写入： 
```go
package msrdp

import "github.com/ying32/govcl/vcl"

// 要注册的事件回调
func eventCallback(fn interface{}, getVal func(idx int) uintptr) bool {
	switch fn.(type) {
	case TMsRdpClient9NotSafeForScriptingOnDisconnected:
		fn.(TMsRdpClient9NotSafeForScriptingOnDisconnected)(
			vcl.ObjectFromInst(getVal(0)),
			int32(getVal(1)))
	default:

	}
	return false
}

```

4、在msrdp目录下新建一个importdef.go文件，并写入：  
```go
package msrdp

import "github.com/ying32/govcl/vcl/api"

// 默认
var (
    libvcl = api.GetLibVcl()
    addEventToMap = api.GetaddEventToMapFn()
)
```

5、在msrdp目录下新建一个events.go文件，并写入（根据实际需求，如果没有事件也可不写）：  
```go
package msrdp

import (
	. "github.com/ying32/govcl/vcl"
)

type TMsRdpClient9NotSafeForScriptingOnDisconnected func(sender IObject, discReason int32)

```

6、导入你自己的方法，如: `importAuto.go`  
```go
package msrdp

var (
	// TMsRdpClient9NotSafeForScripting
	msRdpClient9NotSafeForScripting_Create                    = libvcl.NewProc("MsRdpClient9NotSafeForScripting_Create")
	msRdpClient9NotSafeForScripting_Free                      = libvcl.NewProc("MsRdpClient9NotSafeForScripting_Free")
	msRdpClient9NotSafeForScripting_SetBounds                 = libvcl.NewProc("MsRdpClient9NotSafeForScripting_SetBounds")
	msRdpClient9NotSafeForScripting_CanFocus                  = libvcl.NewProc("MsRdpClient9NotSafeForScripting_CanFocus")
	msRdpClient9NotSafeForScripting_ContainsControl           = libvcl.NewProc("MsRdpClient9NotSafeForScripting_ContainsControl")
	msRdpClient9NotSafeForScripting_ControlAtPos              = libvcl.NewProc("MsRdpClient9NotSafeForScripting_ControlAtPos")
	msRdpClient9NotSafeForScripting_DisableAlign              = libvcl.NewProc("MsRdpClient9NotSafeForScripting_DisableAlign")
  // .....
  // .....
)
```  

7、编写导入的实现方法，如：`importFuncsAuto.go`  
```go

package msrdp

import (
    "unsafe"
    . "github.com/ying32/govcl/vcl/types"
    . "github.com/ying32/govcl/vcl/api"
)


//--------------------------- TMsRdpClient9NotSafeForScripting ---------------------------

func MsRdpClient9NotSafeForScripting_Create(obj uintptr) uintptr {
    ret, _, _ := msRdpClient9NotSafeForScripting_Create.Call(obj)
    return ret
}

func MsRdpClient9NotSafeForScripting_Free(obj uintptr) {
    msRdpClient9NotSafeForScripting_Free.Call(obj)
}

func MsRdpClient9NotSafeForScripting_SetBounds(obj uintptr, ALeft int32, ATop int32, AWidth int32, AHeight int32)  {
    msRdpClient9NotSafeForScripting_SetBounds.Call(obj, uintptr(ALeft) , uintptr(ATop) , uintptr(AWidth) , uintptr(AHeight) )
}

func MsRdpClient9NotSafeForScripting_CanFocus(obj uintptr) bool {
    ret, _, _ := msRdpClient9NotSafeForScripting_CanFocus.Call(obj)
    return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_ContainsControl(obj uintptr, Control uintptr) bool {
    ret, _, _ := msRdpClient9NotSafeForScripting_ContainsControl.Call(obj, Control )
    return DBoolToGoBool(ret)
}

func MsRdpClient9NotSafeForScripting_ControlAtPos(obj uintptr, Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) uintptr {
    ret, _, _ := msRdpClient9NotSafeForScripting_ControlAtPos.Call(obj, uintptr(unsafe.Pointer(&Pos)), GoBoolToDBool(AllowDisabled) , GoBoolToDBool(AllowWinControls) , GoBoolToDBool(AllLevels) )
    return ret
}
```

8、生成一个伪类，如：`msrdpclient9notsafeforscripting.go`  
```go
package msrdp


import (
    . "github.com/ying32/govcl/vcl/types"
    . "github.com/ying32/govcl/vcl"
    "unsafe"
)

type TMsRdpClient9NotSafeForScripting struct {
    IWinControl
    instance uintptr
    ptr unsafe.Pointer
}

func NewMsRdpClient9NotSafeForScripting(owner IComponent) *TMsRdpClient9NotSafeForScripting {
    m := new(TMsRdpClient9NotSafeForScripting)
    m.instance = MsRdpClient9NotSafeForScripting_Create(CheckPtr(owner))
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

func MsRdpClient9NotSafeForScriptingFromInst(inst uintptr) *TMsRdpClient9NotSafeForScripting {
    m := new(TMsRdpClient9NotSafeForScripting)
    m.instance = inst
    m.ptr = unsafe.Pointer(inst)
    return m
}

func MsRdpClient9NotSafeForScriptingFromObj(obj IObject) *TMsRdpClient9NotSafeForScripting {
    m := new(TMsRdpClient9NotSafeForScripting)
    m.instance = CheckPtr(obj)
    m.ptr = unsafe.Pointer(m.instance)
    return m
}

func MsRdpClient9NotSafeForScriptingFromUnsafePointer(ptr unsafe.Pointer) *TMsRdpClient9NotSafeForScripting {
    m := new(TMsRdpClient9NotSafeForScripting)
    m.instance = uintptr(ptr)
    m.ptr = ptr
    return m
}

func (m *TMsRdpClient9NotSafeForScripting) Free() {
    if m.instance != 0 {
        MsRdpClient9NotSafeForScripting_Free(m.instance)
        m.instance = 0
        m.ptr = unsafe.Pointer(uintptr(0))
    }
}

func (m *TMsRdpClient9NotSafeForScripting) Instance() uintptr {
    return m.instance
}

func (m *TMsRdpClient9NotSafeForScripting) UnsafeAddr() unsafe.Pointer {
    return m.ptr
}

func (m *TMsRdpClient9NotSafeForScripting) IsValid() bool {
    return m.instance != 0
}

func TMsRdpClient9NotSafeForScriptingClass() TClass {
    return MsRdpClient9NotSafeForScripting_StaticClassType()
}

func (m *TMsRdpClient9NotSafeForScripting) SetBounds(ALeft int32, ATop int32, AWidth int32, AHeight int32) {
    MsRdpClient9NotSafeForScripting_SetBounds(m.instance, ALeft , ATop , AWidth , AHeight)
}

func (m *TMsRdpClient9NotSafeForScripting) CanFocus() bool {
    return MsRdpClient9NotSafeForScripting_CanFocus(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) ContainsControl(Control IControl) bool {
    return MsRdpClient9NotSafeForScripting_ContainsControl(m.instance, CheckPtr(Control))
}

func (m *TMsRdpClient9NotSafeForScripting) ControlAtPos(Pos TPoint, AllowDisabled bool, AllowWinControls bool, AllLevels bool) *TControl {
    return ControlFromInst(MsRdpClient9NotSafeForScripting_ControlAtPos(m.instance, Pos , AllowDisabled , AllowWinControls , AllLevels))
}

func (m *TMsRdpClient9NotSafeForScripting) DisableAlign() {
    MsRdpClient9NotSafeForScripting_DisableAlign(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) EnableAlign() {
    MsRdpClient9NotSafeForScripting_EnableAlign(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) FindChildControl(ControlName string) *TControl {
    return ControlFromInst(MsRdpClient9NotSafeForScripting_FindChildControl(m.instance, ControlName))
}

func (m *TMsRdpClient9NotSafeForScripting) FlipChildren(AllLevels bool) {
    MsRdpClient9NotSafeForScripting_FlipChildren(m.instance, AllLevels)
}

func (m *TMsRdpClient9NotSafeForScripting) Focused() bool {
    return MsRdpClient9NotSafeForScripting_Focused(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) HandleAllocated() bool {
    return MsRdpClient9NotSafeForScripting_HandleAllocated(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) InsertControl(AControl IControl) {
    MsRdpClient9NotSafeForScripting_InsertControl(m.instance, CheckPtr(AControl))
}

func (m *TMsRdpClient9NotSafeForScripting) Invalidate() {
    MsRdpClient9NotSafeForScripting_Invalidate(m.instance)
}

func (m *TMsRdpClient9NotSafeForScripting) PaintTo(DC HDC, X int32, Y int32) {
    MsRdpClient9NotSafeForScripting_PaintTo(m.instance, DC , X , Y)
}

func (m *TMsRdpClient9NotSafeForScripting) RemoveControl(AControl IControl) {
    MsRdpClient9NotSafeForScripting_RemoveControl(m.instance, CheckPtr(AControl))
}

func (m *TMsRdpClient9NotSafeForScripting) Realign() {
    MsRdpClient9NotSafeForScripting_Realign(m.instance)
}

// ...
// ...

```

### pascal代码部分  
----

**注：libvcl导出函数调用约定为`stdcall`，而liblcl中则为`extdecl`，`extdecl`实际为一个宏定义，当为Windows时使用`stdcall`当为Linux或者macOS时使用`cdecl`。**  

*pascal部分一共有5个默认的inc文件，以UserDefine开头，分别为：*  

* UserDefineComponents.inc  包含你的组件inc文件，比如：`MsRdpClient9NotSafeForScripting.inc`   
```delphi
// 在这里引入你的组件inc文件。
// 比如：
// {$I MsRdpClient9NotSafeForScripting.inc}
// ...
// ...

{$I MsRdpClient9NotSafeForScripting.inc}
```

* UserDefineComponentsClass.inc  提供给设计器寻找类使用。 
```delphi
// 在这里引入你的组件资源构建时的类的inc文件。

// AddComponentClass(TButton);
// AddComponentClass(TEdit);
// ...
// ...
AddComponentClass(TMsRdpClient9NotSafeForScripting);
```

* UserDefineComponentUses.inc  你的组件引用了哪些单元
```delphi
// 在这里引入你的组件uses单元名inc文件，就是说你使用那个组件需要use哪个单元。
// 比如：
// Forms,
// Buttons,
// ...

MSTSCLib_TLB,

```

* UserDefineEventsDeclaration.inc  你的事件申明部分inc文件。  
```delphi
// 在这里引入你的你的事件申明inc文件。
// 比如：
// {$I MsRdpClient9NotSafeForScriptingEventsImplement.inc}
// ...
// ...
// 那么MsRdpClient9NotSafeForScriptingEventsImplement.inc里面是些啥呢？
// 就是下面这类东西
//
//
// class procedure UpDownOnClick(Sender: TObject; Button: TUDBtnType);
//

{$I MyEventsDeclaration.inc}
```

* UserDefineEventsImplement.inc  你的事件实现部分inc文件。  
```delphi
// 在这里引入你的你的事件实现部分的inc文件。
// 比如：
// {$I MsRdpClient9NotSafeForScriptingEventsDeclaration.inc}
// ...
// ...
// 那么MsRdpClient9NotSafeForScriptingEventsImplement.inc里面是些啥呢？
// 就是下面这类东西

//class procedure TEventClass.UpDownOnClick(Sender: TObject; Button: TUDBtnType);
//begin
//  SendEvent(Sender, @TEventClass.UpDownOnClick, [Sender, Ord(Button)]);
//end;
//
//

{$I MyEventsImplement.inc}
```

*额外的两个事件定义文件(如果你使用了事件功能，就可以按以下来写)：*    

* MyEventsDeclaration.inc 事件明声部分  
```delphi

// TMsRdpClient9NotSafeForScriptingOnDisconnected
class procedure OnDisconnected(ASender: TObject; discReason: Integer);

```

* MyEventsImplement.inc  事件实现部分   
```delphi
// TMsRdpClient9NotSafeForScriptingOnDisconnected
class procedure TEventClass.OnDisconnected(ASender: TObject; discReason: Integer);
begin
  SendEvent(ASender, @TEventClass.OnDisconnected, [ASender, discReason]);
end;


```


----

导出函数，如：`MsRdpClient9NotSafeForScripting.inc`  

```delphi

function MsRdpClient9NotSafeForScripting_Create(AOwner: TComponent): TMsRdpClient9NotSafeForScripting; stdcall;
begin
  Result :=  TMsRdpClient9NotSafeForScripting.Create(AOwner);
end;

procedure MsRdpClient9NotSafeForScripting_Free(AObj: TMsRdpClient9NotSafeForScripting); stdcall;
begin
  AObj.Free;
end;

procedure MsRdpClient9NotSafeForScripting_SetBounds(AObj: TMsRdpClient9NotSafeForScripting; ALeft: Integer; ATop: Integer; AWidth: Integer; AHeight: Integer); stdcall;
begin
  AObj.SetBounds(ALeft, ATop, AWidth, AHeight);
end;

function MsRdpClient9NotSafeForScripting_CanFocus(AObj: TMsRdpClient9NotSafeForScripting): LongBool; stdcall;
begin
  Result :=  AObj.CanFocus;
end;

// ....
// ...


exports
  MsRdpClient9NotSafeForScripting_Create,
  MsRdpClient9NotSafeForScripting_Free,
  MsRdpClient9NotSafeForScripting_SetBounds,
  MsRdpClient9NotSafeForScripting_CanFocus,
  MsRdpClient9NotSafeForScripting_ContainsControl,
  MsRdpClient9NotSafeForScripting_ControlAtPos,
  MsRdpClient9NotSafeForScripting_DisableAlign,
  MsRdpClient9NotSafeForScripting_EnableAlign,
  MsRdpClient9NotSafeForScripting_FindChildControl,
  MsRdpClient9NotSafeForScripting_FlipChildren,
  MsRdpClient9NotSafeForScripting_Focused,

// ...
// ...
```

**以上处理好后将所有inc文件或者pas目录复制到libvcl/liblcl目录下，然后重新编译libvcl/liblcl即可。**  

