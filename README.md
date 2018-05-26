# govcl

# Do not use this repository code. This is a backup. Please use "gitee.com/ying32/govcl".  

## golang binding Delphi VCL/Lazarus LCL

Support win32, win64, linux64, macOS32 (linux and macOS only part of the components, properties, events and methods are valid)    

`Go Version >= 1.9.0`  

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

### All binaries are moved to the attachment and [download](https://gitee.com/ying32/govcl/attach_files) . 


* Windows: copy "bin\win32\libvcl.dll" or "bin\win64\libvclx64.dll" to the current exe directory or system environment path. If you don't want to use Delphi binary, you can copy the corresponding LCL library binaries under the "bin\liblcl.dll\" directory. Note: the use of LCL is limited in components, events, properties, and methods.

* Linux: copy "bin\linux64\liblcl.so" executable file directory (also can be copied from the liblcl.so to `/usr/lib/` directory as a Public Library).

* MacOS: copy "bin\MacOS32\liblcl.dylib" executable file directory. [APP package](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS).

---
Q: Why not submit the code on github.com?  
A: Visit site too slow.  


Q: Why is there no English WIKI?   
A: My English is bad.  

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
![图像按钮](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/imagebutton.jpg)  
![basicResForm_Windows](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/basicResForm_windows_vcl.jpg) 
---
## Linux 
### Ubuntu 16.04 

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1_linux.png)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/3_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/4_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/5_linux.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/6_linux.jpg)  

* Deepin  
![basicResForm_Deepin](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/basicResForm_linux_deepin.jpg)  
* Mint  
![basicResForm_Mint](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/basicResForm_linux_mint.jpg)  
* Ubuntu  
![basicResForm_Ubuntu](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/basicResForm_linux_ubuntu.jpg)  
---
## MacOS

![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/1_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/2_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/3_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/4_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/5_macOS.jpg)  
![截图1](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/6_macOS.jpg)  
![basicResForm_MacOS](https://raw.githubusercontent.com/ying32/govcl/master/Screenshot/basicResForm_macos.jpg)