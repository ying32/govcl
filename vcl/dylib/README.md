### 简介

dylib是从govcl的api包中分离出来，可以方便其它使用。
dylib实现现了一个通用的共享库调用，可以在windows、linux、macOS下不需要任何改变实现通用。
Window下使用syscall.NewLazyDLL加载dll，linux与macOS下使用dlopen加载，然后使用dlsym来获取,
通过封装，使用接口与Windows下的一模一样，如此达到通用，最多实现12个参数的过程。

在被调用的共享库中如果导出了“MySyscall”函数，则会使用此函数来call， 一般导出此函数的目的是为了在共享库中捕获异常，不然一但共享库中出现异常
程序就挂了。

```pascal
// 一个Delphi的"MySyscall"实现例程：

type
  TSyscall0 = function: UInt64; stdcall;
  TSyscall1 = function(A1: Pointer): UInt64; stdcall;
  TSyscall2 = function(A1, A2: Pointer): UInt64; stdcall;
  TSyscall3 = function(A1, A2, A3: Pointer): UInt64; stdcall;
  TSyscall4 = function(A1, A2, A3, A4: Pointer): UInt64; stdcall;
  TSyscall5 = function(A1, A2, A3, A4, A5: Pointer): UInt64; stdcall;
  TSyscall6 = function(A1, A2, A3, A4, A5, A6: Pointer): UInt64; stdcall;
  TSyscall7 = function(A1, A2, A3, A4, A5, A6, A7: Pointer): UInt64; stdcall;
  TSyscall8 = function(A1, A2, A3, A4, A5, A6, A7, A8: Pointer): UInt64; stdcall;
  TSyscall9 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9: Pointer): UInt64; stdcall;
  TSyscall10 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10: Pointer): UInt64; stdcall;
  TSyscall11 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11: Pointer): UInt64; stdcall;
  TSyscall12 = function(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;

function MySyscall(AProc: Pointer; ALen: NativeInt; A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12: Pointer): UInt64; stdcall;
begin
  Result := 0;
  if AProc = nil then
    Exit;
  try
    case ALen of
      0: Result := TSyscall0(AProc)();
      1: Result := TSyscall1(AProc)(A1);
      2: Result := TSyscall2(AProc)(A1, A2);
      3: Result := TSyscall3(AProc)(A1, A2, A3);
      4: Result := TSyscall4(AProc)(A1, A2, A3, A4);
      5: Result := TSyscall5(AProc)(A1, A2, A3, A4, A5);
      6: Result := TSyscall6(AProc)(A1, A2, A3, A4, A5, A6);
      7: Result := TSyscall7(AProc)(A1, A2, A3, A4, A5, A6, A7);
      8: Result := TSyscall8(AProc)(A1, A2, A3, A4, A5, A6, A7, A8);
      9: Result := TSyscall9(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9);
      10: Result := TSyscall10(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10);
      11: Result := TSyscall11(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11);
      12: Result := TSyscall12(AProc)(A1, A2, A3, A4, A5, A6, A7, A8, A9, A10, A11, A12);
    else
      Exit;
    end;
  except
    on E: Exception do
    begin
      Writeln('Syscall', ALen, ' Error: ', E.Message);
    end;
  end;
end;

```


### 使用方法

```go

import "gitee.com/ying32/govcl/vcl/dylib"

var (
    lib = dylib.NewLazyDLL("xxx.dll") // 或者 dylib.NewLazyDLL("xxx.so") 或者 dylib.NewLazyDLL("xxx.dylib")
    _Func1 = lib.NewProc("Func1") 
    _Func2 = lib.NewProc("Func2") 
    _Func3 = lib.NewProc("Func3") 
)

// 普通的
func Func1(a1, a2 int) int {
    r, _, _ := _Func1.Call(uintptr(a1), uintptr(a2))
    return int(r)
}

// 浮点类型的, 切记不要直接在共享库中直接返回浮点类型的，要返回记得使用指针参数进行传递
func Func2() float32 {
    var f float32
    _Func2.Call(uintptr(unsafe.Pointer(&f)))
    return f
}

// 如果是外部的共享库返回float32或者float64则另使用补丁方式，暂时不支持arm
// import "github.com/ying32/govcl/dylib/floatpatch"
// float32
func Func2() float32 {
    _Func2.Call()  
    return floatpatch.Getfloat32() // 注，在此函数前面不要调用其他函数了，不然就获取不到结果
}

// float64
func Func2() float64 {
    _Func2.Call()  
    return floatpatch.Getfloat64() // 注，在此函数前面不要调用其他函数了，不然就获取不到结果
}


// 结构类型的, 切记不要直接在共享库中直接返回结构类型，要返回记得使用指针参数进行传递
// 参数的传递也要使用指针类型传递
type TPoint struct {
   X int32
   Y int32
}

func Func3(p1 TPoint) TPoint {
    var pret TPoint
    _Func3.Call(uintptr(unsafe.Pointer(&p1)), uintptr(unsafe.Pointer(&pret)))
    return pret
}



```  


### 需要特别注意的事情   

* 1、不要在共享库中返回float类的，不然会被go强转，如需使用可选择使用参数指针传递（注：如是外部无法修改的共享库，则使用上面的补丁方式）。    
* 2、在参数或者返回值类型中需要注意： 在x86库中应该尽量避免直接返回或者使用size>=8的类型，如需使用最好选择参数指针替代，以保持x86跟x64一致性。  