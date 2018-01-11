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