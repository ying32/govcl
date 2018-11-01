## 一个跨平台的Golang GUI库，底层绑定自Delphi VCL和Lazarus LCL  

* 中文   
* [English](README.md)   

----

**从1.1.24版本开始govcl将最低要求go1.9。**  

----

[![license](https://img.shields.io/badge/开源协议-Apache%20License%202.0-green.svg)](https://github.com/ying32/govcl/blob/master/LICENSE)
![Recommended Golang Version](https://img.shields.io/badge/最低Go版本-1.9.0-green.svg)
[![screenshots](https://img.shields.io/badge/例程截图-查看-green.svg)](https://github.com/ying32/govcl/tree/master/Screenshot)
[![Chinese Wiki](https://img.shields.io/badge/维基-中文WIKI-green.svg)](https://gitee.com/ying32/govcl/wikis/pages)
[![Chinese Chat](https://img.shields.io/badge/QQ群-点击加入：263106281-red.svg)](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq)  

![Support Platform](https://img.shields.io/badge/支持的平台-win--32%20%7C%20win--64%20%7C%20linux--64%20%7C%20osx--32-green.svg)  
**注: linux和macOS由于底层使用了lcl库，则部分组件、属性和方法无效。**

**如果想要支持linux arm及linux 32bit则需要自己编译对应的liblcl二进制。**   


GUI库二进制下载：   
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.1.23.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.1.23/Librarys-1.1.23.zip)  
**注：压缩包内包含的“libvcl”库二进制（libvcl.dll、libvclx64.dll）仅供预览和测试使用。正式使用请自行编译“libvcl”源代码，具体编译方法参考[UILIbSrcources](UILibSources/README.zh-CN.md)中的说明。**  


res2go工具下载  
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.5-blue.svg)](Tools/res2go)  
**注：用Delphi/Lazarus设计界面，用Golang写代码。**    


VCL样式文件    
[![VCL style files](https://img.shields.io/badge/downloads-VCL%20style%20files-blue.svg)](https://github.com/ying32/govcl/releases/download/v1.1.20/vcl-styles.zip)  

---
### 使用方法
> go get github.com/ying32/govcl  

* 方法一：  

```golang
package main

import (
   "github.com/ying32/govcl/vcl"
)

var (
   mainForm *vcl.TForm
)

func main() {
    vcl.Application.Initialize()
    mainForm = vcl.Application.CreateForm()
    mainForm.SetCaption("Hello")
    mainForm.EnabledMaximize(false)
    mainForm.ScreenCenter()
    vcl.Application.Run()
}
```  

* 方法二：  

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
    vcl.Application.CreateForm(&mainForm)
    vcl.Application.Run()
}


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

```

* 方法三：  

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
**方法三需要配合UI设计器或者res2go工具使用。**  


----

### 注意:  

**当使用"liblcl"库时，是以兼容"libvcl"库形式运行的，所以有部分组件和组件的方法、属性及事件不可用。**  

----


### API文档

* [Delphi VCL组件文档WIKI](http://docwiki.embarcadero.com/RADStudio/Tokyo/en/Category:VCL_Reference)  
* [Lazarus LCL组件文档WIKI](http://wiki.freepascal.org/LCL_Components)  
* [Windows API文档](https://msdn.microsoft.com/zh-cn/library/ms123401.aspx)

----

* Windows: 复制"libvcl.dll"或者"libvclx64.dll"或者“liblcl.dll”到当前exe目录或系统环境路径下。    

* Linux: 复制"liblcl.so"可执行文件目录下(也可复制liblcl.so到`/usr/lib/`目录中，作为公共库使用)。  

* MacOS: 复制"liblcl.dylib"可执行文件目录下（MacOS下注意：需要自行创建info.plist文件），或者参考：[MacOS上应用打包](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS) 