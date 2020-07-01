* [中文](README.md)   
* English  

----

### GoVCL

Cross-platform Golang GUI library, The core binding is [liblcl](https://github.com/ying32/liblcl), a common cross-platform GUI library created by [Lazarus](https://www.lazarus-ide.org/).  

**Full name: `Go Language Visual Component Library`**    

**The current version no longer supports `Delphi/VCL`. The last branch that supports the vcl version: [last-vcl-support](https://github.com/ying32/govcl/tree/last-vcl-support/).**   

*govcl minimum requirement is go1.9.*    

[Screenshots](https://z-kit.cc/en/screenshot.html) | 
[WIKI(Chinese)](https://gitee.com/ying32/govcl/wikis/pages) | 
[What's-new(Chinese)](https://z-kit.cc/changelog.html) | 
[GoVCL video tutorial (third party)](https://video.0-w.cc/videos/1) | 
[Sponsor govcl](https://z-kit.cc/en/sponsor.html)  

----

### Support Platform    
**Windows** | **Linux** | **macOS**  

> If you want to support linux arm and linux 32bit, you need to compile the corresponding liblcl binary.   

### Pre-compiled GUI library binary download ([source code](UILibSources))     
[![liblcl](https://img.shields.io/github/downloads/ying32/govcl/latest/liblcl-2.0.3.1.zip.svg)](https://github.com/ying32/govcl/releases/download/v2.0.3/liblcl-2.0.3.1.zip)   


* Dev version pre-compiled GUI binary   
[liblcl-2.0.4-dev.zip](http://z-kit.cc/downloads/liblcl-2.0.4-dev.zip)  

### res2go Tool([doc, source code](Tools/res2go))    

**Need to compile yourself: [Compile Guide](Tools/res2go/src/README.en-US.md)**   

> Note: Designed in Lazarus, code written in Golang.  
 
### usage: 

#### Step 1: Get the govcl code  

> go get -u github.com/ying32/govcl    

#### Step 2: Write the code

* Method 1(Use Lazarus to design the GUI. recommend): 

```golang
package main


import (
   // Do not reference this package if you use custom syso files
   _ "github.com/ying32/govcl/pkgs/winappres"
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


* Method 2(Pure code, imitating the way of FreePascal class):  

```golang
package main


import (
   // Do not reference this package if you use custom syso files
   _ "github.com/ying32/govcl/pkgs/winappres"
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
    vcl.RunApp(&mainForm, &aboutForm)
}

// -- TMainForm

func (f *TMainForm) OnFormCreate(sender vcl.IObject) {
    f.SetCaption("MainForm")
    f.Btn1 = vcl.NewButton(f)
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
    f.Btn1.SetOnClick(f.OnBtn1Click)  
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
    f.Btn1.SetOnClick(f.OnBtn1Click)  
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

* Windows: Depending on whether the compiled binary is 32 or 64 bits, copy the corresponding `liblcl.dll` to the current executable file directory or system environment path.  
  * Go environment variable: `GOARCH = amd64 386` `GOOS = windows` `CGO_ENABLED=0`    

* Linux: Copy `liblcl.so` under the current executable file directory (you can also copy `liblcl.so` to `/usr/lib/` (32bit liblcl) or `/usr/lib/x86_64-linux-gnu/` (64bit liblcl) directory , Used as a public library).  
  * Go environment variable: `GOARCH = amd64` `GOOS = linux` `CGO_ENABLED=1`  

* MacOS: Copy `liblcl.dylib` to the current executable file directory (note under MacOS: you need to create info.plist file yourself), or refer to: [App packaging on MacOS](https://gitee.com/ying32/govcl/wikis/pages?sort_id=410056&doc_id=102420)  
  * Go environment variable: `GOARCH = amd64` `GOOS = darwin` `CGO_ENABLED=1`  

Note: The "current executable file directory" here refers to the location of the executable file generated by your currently compiled project.

---   

**Special Note: All UI components are non-threaded/non-coroutine safe. When used in goroutine, use [vcl.ThreadSync](https://gitee.com/ying32/govcl/wikis/pages?sort_id=976890&doc_id=102420) to synchronize updates to the UI.**

---

### FAQ

Q: Why is there no English WIKI?   
A: My English is bad. You can try using Google Translate [Chinese WIKI](https://gitee.com/ying32/govcl/wikis/pages).    
 
---  

### API document

* [Lazarus LCL component document  WIKI](http://wiki.freepascal.org/LCL_Components)  
* [Windows API document](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)

----
