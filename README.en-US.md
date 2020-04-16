* [中文](README.md)   
* English  

----

**The current version no longer supports `Delphi/VCL`. The last branch that supports the vcl version: [last-vcl-support](https://github.com/ying32/govcl/tree/last-vcl-support/).**  

## Cross-platform Golang GUI library. Use Lazarus LCL for binding.

> The govcl version >=1.2.0 must require the go version >=1.9.0.  

[Screenshots](https://github.com/ying32/govcl/tree/master/Screenshot) | 
[What's-new(Chinese)](https://z-kit.cc/changelog.html) | 
[GoVCL video tutorial (third party)](https://video.0-w.cc/videos/1) | 
[Sponsor govcl](https://z-kit.cc/sponsor.html)  

----

### Support Platform    
**Windows** | **Linux** | **macOS**  

> If you want to support linux arm and linux 32bit, you need to compile the corresponding liblcl binary.   

### Pre-compiled GUI library binary download     
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-2.0.0.zip.svg)](https://github.com/ying32/govcl/releases/download/v2.0.0/Librarys-2.0.0.zip)   

### res2go Tool([doc](Tools/res2go))    
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.19-blue.svg)](https://github.com/ying32/govcl/releases/download/v1.2.9/res2go-1.0.19.zip)  
> Note: Designed in Lazarus, code written in Golang.  

[GoVCL extension package](https://github.com/ying32/exts)  
> Note: The inc or pas file in the extension package needs to be compiled into the liblcl directory.     

 
### usage: 

#### Step 1: Get the govcl code  

> go get -u github.com/ying32/govcl    

#### Step 2: Write the code

* Method 1(Use Lazarus to design the GUI. recommend): 

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
   // Do not reference this package if you use custom syso files
   _ "github.com/ying32/govcl/pkgs/winappres"
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
    vcl.Application.CreateForm(&aboutForm)
    vcl.Application.Run()
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
    
}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}

// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
 
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}
```
**Method 1 needs to be used in conjunction with the res2go tool.**  


* Method 2(Pure code, imitating the way of FreePascal class, can automatically bind events.):  

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
   // Do not reference this package if you use custom syso files
   _ "github.com/ying32/govcl/pkgs/winappres"
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
    vcl.RunApp(&mainForm, &aboutForm)
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
    f.SetCaption("MainForm")
    f.Btn1 = vcl.NewButton(f)
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
    //f.Btn1.SetOnClick(f.OnBtn1Click)  
}

func (f *TMainForm) OnBtn1Click(sender vcl.IObject) {
    aboutForm.Show()
}


// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
    f.SetCaption("About")
    f.Btn1 = vcl.NewButton(f)
    //f.Btn1.SetName("Btn1")
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}
``` 

* Method 3  
```go
// Not recommended, so no examples are given.
```

#### Step 3: Copy the corresponding binary   

* Windows: According to whether the compiled binary is 32 or 64 bit, copy the corresponding `liblcl.dll` to the current exe directory or system environment path.  
  * Go environment variable: `GOARCH = amd64 386` `GOOS = windows` `CGO_ENABLED=0`    

* Linux: Copy the `liblcl.so` executable directory (you can also copy liblcl.so to the `/usr/lib/`(32bit liblcl) or `/usr/lib/x86_64-linux-gnu/`(64bit liblcl) directory and use it as a public library).  
  * Go environment variable: `GOARCH = amd64` `GOOS = linux` `CGO_ENABLED=1`  

* MacOS: Copy the `liblcl.dylib` executable directory (Note for MacOS: you need to create the info.plist file yourself), or refer to: [App packaging on MacOS](https://gitee.com/ying32/govcl/wikis/pages?sort_id=410056&doc_id=102420)  
  * Go environment variable: `GOARCH = amd64` `GOOS = darwin` `CGO_ENABLED=1`  

---   

**Special Note: All UI components are non-threaded/non-coroutine safe. When used in goroutine, use `vcl.ThreadSync` to synchronize updates to the UI.**

---

### FAQ

Q: Why is there no English WIKI?   
A: My English is bad. You can try using Google Translate [Chinese WIKI](https://gitee.com/ying32/govcl/wikis/pages).    
 
---  

### API document

* [Lazarus LCL component document  WIKI](http://wiki.freepascal.org/LCL_Components)  
* [Windows API document](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)

----
