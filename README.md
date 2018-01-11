# govcl

## golang binding delphi vcl(lazarus lcl)

Support win32, win64, linux64, macOS32 (linux and macOS only part of the components, properties, events and methods are valid)  

### usage: 

> go get gitee.com/ying32/govcl    

```go
package main

import (
   "gitee.com/ying32/govcl/vcl"
)

var (
   mainForm *vcl.TForm
)

func main() {
    vcl.Application.SetIconResId(3)
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}

```

* Windows: copy "bin\win32\libvcl.dll" or "bin\win64\libvclx64.dll" to the current exe directory or system environment path. If you don't want to use Delphi binary, you can copy the corresponding LCL library binaries under the "bin\liblcl.dll\" directory. Note: the use of LCL is limited in components, events, properties, and methods.

* Linux: copy "bin\linux64\liblcl.so" executable file directory (also can be copied from the liblcl.so to `/usr/lib/` directory as a Public Library).

* MacOS: copy "bin\MacOS32\liblcl.dylib" executable file directory. [apply packaging on MacOS](https://gitee.com/ying32/govcl/wikis/pages?title=MacOS%E4%B8%8A%E5%BA%94%E7%94%A8%E6%89%93%E5%8C%85&parent=FQA).

---
Q: Why not submit the code on github.com?  
A: Visit github in China is very bad, so choose China's domestic git repository.
---

### Chinese wiki  

[Chinese wiki](https://gitee.com/ying32/govcl/wikis/Home)  

---

* [Windows](#Windows)
* [Linux](#Linux)
* [MacOS](#MacOS)


## Windows

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1.png)    
![截图2](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2.png)      
![绘图](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/draw.png)  
![ListView](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/listview.png)  
![RichEdit](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/richedit.png)  
![标准控件](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/std.png)  
![样式](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/style.png)  

---
## Linux 
### Ubuntu 16.04 

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1_linux.png)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/3_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/4_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/5_linux.jpg)  


---
## MacOS

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/3_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/4_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/5_macOS.jpg)  