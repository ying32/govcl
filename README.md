## 一个跨平台的Golang GUI库，底层绑定自Delphi VCL和Lazarus LCL  

* 中文   
* [English](README.en-US.md)   

----

**从1.2.0版本开始govcl将最低要求go1.9。**  

----

[![license](https://img.shields.io/badge/开源协议-Apache%20License%202.0-green.svg)](https://github.com/ying32/govcl/blob/master/LICENSE)
![Minimum Go version](https://img.shields.io/badge/最低Go版本-1.9.0-green.svg)
[![screenshots](https://img.shields.io/badge/例程截图-查看-green.svg)](https://github.com/ying32/govcl/tree/master/Screenshot)
[![Chinese Wiki](https://img.shields.io/badge/维基-中文WIKI-green.svg)](https://gitee.com/ying32/govcl/wikis/pages)
[![Chinese Chat](https://img.shields.io/badge/QQ群-点击加入：263106281-red.svg)](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq)
[![Update log](https://img.shields.io/badge/更新日志-查看-blue.svg)](https://github.com/ying32/govcl/wiki/%E6%9B%B4%E6%96%B0%E6%97%A5%E5%BF%97(What's-new))
[![Support govcl](https://img.shields.io/badge/支持govcl-赞助作者-blueviolet.svg)](https://z-kit.cc/sponsor.html)  

![Support Platform](https://img.shields.io/badge/支持的平台-Windows%20%7C%20Linux%20%7C%20Mac%20OS-green.svg)  
**注: linux和macOS由于底层使用了lcl库，则部分组件、属性和方法无效。**

**如果你想要支持linux arm及linux 32bit则需要自己编译对应的liblcl二进制。**   


预编译GUI库二进制下载：     
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.2.7.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.2.7/Librarys-1.2.7.zip)  
**注：压缩包内包含的“libvcl”库二进制（libvcl.dll、libvclx64.dll）仅供预览和测试使用。正式使用请自行编译“libvcl”源代码，具体编译方法参考[UILIbSrcources](UILibSources/README.md)中的说明。**  


res2go工具下载（[文档](Tools/res2go)）  
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.17-blue.svg)](https://github.com/ying32/govcl/releases/download/v1.2.7/res2go-1.0.17.zip)  
**注：用Delphi/Lazarus设计界面，用Golang写代码。**     
  
[govcl扩展包](https://github.com/ying32/exts)  
**注：扩展包里面的inc或者pas文件需要自己放到libvcl/liblcl目录下编译。**    

[govcl视频教程(第三方)](https://video.0-w.cc/videos/1)   

---
### 使用方法  

#### 步骤一：获取govcl代码  

> go get -u github.com/ying32/govcl  

#### 步骤二：编写代码    

* 方法一(使用Delphi/Lazarus或者GoVCLDesigner设计界面。推荐)：  

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
   // 如果你使用自定义的syso文件则不要引用此包
   _ "github.com/ying32/govcl/vcl/exts/winappres"
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
**方法一需要配合UI设计器或者res2go工具使用。**  


* 方法二(纯代码，仿照Delphi类的方式，可自动绑定事件。)：  

```golang
package main


import (
   "github.com/ying32/govcl/vcl"
   // 如果你使用自定义的syso文件则不要引用此包
   _ "github.com/ying32/govcl/vcl/exts/winappres"
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
    // 创建完后关联子组件事件
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
    //f.Btn1.SetName("Btn1")
    f.Btn1.SetParent(f)
    f.Btn1.SetBounds(10, 10, 88, 28)
    f.Btn1.SetCaption("Button1")
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}

```


* 方法三(纯代码。不推荐方式)：  

```golang
package main

import (
   "github.com/ying32/govcl/vcl"
   // 如果你使用自定义的syso文件则不要引用此包
   _ "github.com/ying32/govcl/vcl/exts/winappres"
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


#### 步骤三：复制对应的二进制    

* Windows: 根据编译的二进制是32还是64位的，复制对应的"libvcl.dll"或者"libvclx64.dll"或者“liblcl.dll”到当前exe目录或系统环境路径下。 
  * Go环境变量： `GOARCH = amd64 386` `GOOS = windows` `CGO_ENABLED=0`   

* Linux: 复制"liblcl.so"可执行文件目录下(也可复制liblcl.so到`/usr/lib/`目录中，作为公共库使用)。  
  * Go环境变量： `GOARCH = amd64` `GOOS = linux` `CGO_ENABLED=1`

* MacOS: 复制"liblcl.dylib"可执行文件目录下（MacOS下注意：需要自行创建info.plist文件），或者参考：[MacOS上应用打包](https://gitee.com/ying32/govcl/wikis/pages?sort_id=410056&doc_id=102420)   
  * Go环境变量： `GOARCH = amd64` `GOOS = darwin` `CGO_ENABLED=1`  


----

### 注意:  

**当使用"liblcl"库时，是以兼容"libvcl"库形式运行的，所以有部分组件和组件的方法、属性及事件不可用。**  

----

**特别注意：所有UI组件都是非线程/协程安全的，当在goroutine中使用时，请使用`vcl.ThreadSync`来同步更新到UI上。**  


### API文档

* [Delphi VCL组件文档WIKI](http://docwiki.embarcadero.com/RADStudio/Tokyo/en/Category:VCL_Reference)  
* [Lazarus LCL组件文档WIKI](http://wiki.freepascal.org/LCL_Components)  
* [Windows API文档](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)
