## 一个跨平台的Golang GUI库，底层绑定自Delphi VCL和Lazarus LCL  

* [中文](README.zh-CN.md)   
* [English](README.md)   

----

[![license](https://img.shields.io/badge/license-Apache%20License%202.0-green.svg)](https://github.com/ying32/govcl/blob/master/LICENSE)
![Recommended Golang Version](https://img.shields.io/badge/推荐Goland版本->=1.9.0-green.svg)
[![screenshots](https://img.shields.io/badge/例程截图-查看-green.svg)](https://github.com/ying32/govcl/tree/master/Screenshot)
[![Chinese Wiki](https://img.shields.io/badge/wiki-中文WIKI-green.svg)](https://gitee.com/ying32/govcl/wikis/pages)
[![Chinese Chat](https://img.shields.io/badge/QQ群-点击加入：263106281-red.svg)](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq)  

![Support Platform](https://img.shields.io/badge/支持的平台-win--32%20%7C%20win--64%20%7C%20linux--64%20%7C%20osx--32-green.svg)  
**注意: linux和macOS由于底层使用了lcl库，则部分组件、属性和方法无效。**   


GUI库二进制下载：   
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.1.19.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.1.19/Librarys-1.1.19.zip)  
**注意: 压缩包中下载的二进制只包含liblcl库，libvcl则需要你自己编译，具体编译方法参考UILIbSrcources中的说明。**  


res2go工具下载  
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.4.zip-green.svg)](https://github.com/ying32/govcl/blob/master/Tools/res2go)


---
### 使用方法
> go get github.com/ying32/govcl  

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

**相关二进制已经移到附件中[下载](https://github.com/ying32/govcl/releases)。**  

* Windows: 复制"bin\win32\libvcl.dll"或者"bin\win64\libvclx64.dll"到当前exe目录或系统环境路径下(如果不想使用Delphi的二进制可到“bin\liblcl.dll\”目录下复制对应的lcl库二进制。注：使用lcl在组件，事件，属性及方法上会受到限制)  

* Linux: 复制"bin\linux64\liblcl.so"可执行文件目录下(也可复制liblcl.so到`/usr/lib/`目录中，作为公共库使用)。  

* MacOS: 复制"bin\MacOS32\liblcl.dylib"可执行文件目录下（MacOS下注意：需要自行创建info.plist文件），或者参考：[MacOS上应用打包](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS) 