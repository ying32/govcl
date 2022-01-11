
中文   | [English](README.md)    

----

### GoVCL
 
跨平台的Golang GUI库, 核心绑定自 [Lazarus](https://www.lazarus-ide.org/) 创建的通用跨平台GUI库 [liblcl](https://github.com/ying32/liblcl) 。    

**GoVCL是一个原生GUI库，不是基于HTML，更不是DirectUI库，一切以实用为主。**    

**中文全称：`Go语言可视化组件库`；英文全称：`Go Language Visual Component Library`**    

*govcl最低要求go1.9.2。*    

[截图查看](https://z-kit.cc/screenshot.html) | 
[中文文档](https://gitee.com/ying32/govcl/wikis/pages) | 
[更新日志](https://z-kit.cc/changelog.html) | 
[视频教程(非官方)](https://video.0-w.cc/videos/1) | 
[赞助GoVCL](https://z-kit.cc/sponsor.html)   

----

### 支持的平台  
**Windows** | **Linux** | **macOS**  

> 如果你想要支持linux arm及linux 32bit则需要自己编译对应的liblcl二进制。  

----

### 预编译GUI库二进制下载（[源代码](https://github.com/ying32/liblcl)）       
[![liblcl](https://img.shields.io/github/downloads/ying32/govcl/latest/liblcl-2.2.0.zip.svg)](https://github.com/ying32/govcl/releases/download/v2.2.0/liblcl-2.2.0.zip)  


### res2go IDE插件源代码（[源代码](https://github.com/ying32/res2go-ide-plugin)）  

**使用方法: [安装方法](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2645001&doc_id=102420)**   

> 注：用Lazarus设计界面，用Golang写代码。    
  
---
### 使用方法  

#### 步骤一：获取govcl代码  

> go get -u github.com/ying32/govcl  

*注：也可用go module方式拉取，在go.mod中配置如：`github.com/ying32/govcl v2.2.0+incompatible`。*  

#### 步骤二：编写代码    

* 方法一(使用Lazarus设计界面。推荐)：  

```golang
package main


import (
   // 如果你使用自定义的syso文件则不要引用此包
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
    aboutForm.Show()
}

// -- TAboutForm

func (f *TAboutForm) OnFormCreate(sender vcl.IObject) {
 
}

func (f *TAboutForm) OnBtn1Click(sender vcl.IObject) {
    vcl.ShowMessage("Hello!")
}
```
**方法一需要配合res2go工具使用。**  


* 方法二(纯代码，仿照FreePascal类的方式)：  

```golang
package main


import (
   // 如果你使用自定义的syso文件则不要引用此包
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

#### 步骤三：复制对应的二进制    

* Windows: 根据编译的二进制是32还是64位的，复制对应的`liblcl.dll`到当前可执行文件目录或系统环境路径下。 
  * Go环境变量： `GOARCH = amd64 386` `GOOS = windows` `CGO_ENABLED=0`   

* Linux: 复制`liblcl.so`当前可执行文件目录下(也可复制liblcl.so到`/usr/lib/`(32bit liblcl)或者`/usr/lib/x86_64-linux-gnu/`(64bit liblcl)目录中，作为公共库使用)。  
  * Go环境变量： `GOARCH = amd64` `GOOS = linux` `CGO_ENABLED=1`

* MacOS: 复制`liblcl.dylib`当前可执行文件目录下（MacOS下注意：需要自行创建info.plist文件），或者参考：[MacOS上应用打包](https://gitee.com/ying32/govcl/wikis/pages?sort_id=410056&doc_id=102420)   
  * Go环境变量： `GOARCH = amd64` `GOOS = darwin` `CGO_ENABLED=1`  


注：这里的`当前可执行文件目录`指的是你当前编译的项目生成的可执行文件位置。

----

### 注意:  

**特别注意：所有UI组件都是非线程/协程安全的，当在goroutine中使用时，请使用[vcl.ThreadSync](https://gitee.com/ying32/govcl/wikis/pages?sort_id=976890&doc_id=102420)来同步更新到UI上。**    

**特别注意2：如果你使用go>=1.15编译Windows可执行文件，则必须则必须使用`-buildmode=exe`编译选项，不然会有错误。**  

### API文档

* [Lazarus LCL组件WIKI](http://wiki.freepascal.org/LCL_Components)
* [Windows API文档](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)  


----

![jetbrains](https://z-kit.cc/assets/images/jetbrains.png)   
[鸣谢jetbrains](https://www.jetbrains.com/?from=govcl)  
