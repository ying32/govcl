## A cross-platform Golang GUI library. Use Delphi VCL and Lazarus LCL for binding.

* [中文](README.zh-CN.md)   
* English  

----

**The 1.2.0 version begins with govcl with a minimum requirement of go1.9.**   

[](https://github.com/ying32/govcl/wiki/%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97(What's-new))   

----
[![License](https://img.shields.io/badge/license-Apache%20License%202.0-green.svg)](https://github.com/ying32/govcl/blob/master/LICENSE)
![Minimum Go version](https://img.shields.io/badge/Minimum%20Go%20version-1.9.0-green.svg)
[![Screenshots](https://img.shields.io/badge/screenshots-view-green.svg)](https://github.com/ying32/govcl/tree/master/Screenshot)
[![Update log](https://img.shields.io/badge/Update%20log-view-blue.svg)](https://github.com/ying32/govcl/wiki/%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97(What's-new))  

![Support Platform](https://img.shields.io/badge/Platform-win--32%20%7C%20win--64%20%7C%20linux--64%20%7C%20osx--32-green.svg)  
**Note: linux and macOS only part of the components, properties, events and methods are valid.**  

**If you want to support linux arm and linux 32bit, you need to compile the corresponding liblcl binary.**   


Librarys binaries  
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.2.1.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.2.1/Librarys-1.2.1.zip)  
**Note: The "libvcl" library binary (libvcl. dll, libvclx64. dll) contained in the compression package is only for preview and test purposes. Please compile the "libvcl" source code for official use. please refer to the instructions in [UILIbSrcources](UILibSources/README.md).**  


res2go Tool  
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.10-blue.svg)](Tools/res2go)  
**Note: Designed in Delphi/Lazarus, code written in Golang.**  

### usage: 

> go get github.com/ying32/govcl    

* Method 1: 

```golang
package main

import (
   "github.com/ying32/govcl/vcl"
)

func main() {
    vcl.Application.Initialize()
    mainForm := vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.ScreenCenter()
    btn := vcl.NewButton(mainForm)
    btn.SetParent(mainForm)
    btn.SetCaption("Hello")
    btn.SetOnClick(func(sender vcl.IObject) {
        vcl.ShowMessage("Hello!")
    })
    vcl.Application.Run()
}
```  

* Method 2:  

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    Btn1     *vcl.TButton
}

type TAboutForm struct {
    *vcl.TForm
    Btn1    *vcl.TButton
}

var (
    mainForm *TMainForm
    aboutForm *TAboutForm
)

func main() {
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(&mainForm)
   // Bind subcomponent events after creation.
    vcl.Application.CreateForm(&aboutForm, true)
    vcl.Application.Run()
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
    f.SetCaption("Hello")
    f.Btn1 = vcl.NewButton(f)
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
    f.Btn1.SetOnClick(f.OnButtonClick)  
}

func (f *TMainForm) OnButtonClick(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}


// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
    f.SetCaption("Hello")
    f.Btn1 = vcl.NewButton(f)
    f.Btn1.SetName("Btn1")
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}
```

* Method 3: 

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
)

type TMainForm struct {
    *vcl.TForm
    Btn1     *vcl.TButton
}

var (
    mainForm *TMainForm
)

func main() {
    vcl.Application.Initialize()
    vcl.Application.SetMainFormOnTaskBar(true)
    vcl.Application.CreateForm(mainFormBytes, &mainForm)
    vcl.Application.Run()
}


func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
    
}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}
```
**Method 3 needs to be used in conjunction with the UI designer or the res2go tool.**  

---   
### FAQ

Q: Why is there no English WIKI?   
A: My English is bad. You can try using Google Translate [Chinese WIKI](https://gitee.com/ying32/govcl/wikis/pages).    
 
---  
### Note:  

**When using the "liblcl" library, it is run in a compatible "libvcl" library, so some methods and properties of components and components are not available.**  

---

### API document

* [Delphi VCL component document  WIKI](http://docwiki.embarcadero.com/RADStudio/Tokyo/en/Category:VCL_Reference)  
* [Lazarus LCL component document  WIKI](http://wiki.freepascal.org/LCL_Components)  
* [Windows API document](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)

----

* Windows: Copy "libvcl.dll" or "libvclx64.dll" or "liblcl.dll" to the current exe directory or system environment path.

* Linux: Copy the "liblcl.so" executable directory (you can also copy liblcl.so to the `/usr/lib/` directory and use it as a public library).

* MacOS: Copy the "liblcl.dylib" executable directory (Note for MacOS: you need to create the info.plist file yourself), or refer to: [App packaging on MacOS](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS)

---  
