## A cross-platform Golang GUI library. Use Delphi VCL and Lazarus LCL for binding.

* [中文](README.zh-CN.md)   
* [English](README.md)   

----

[![license](https://img.shields.io/badge/license-Apache%20License%202.0-green.svg)](https://github.com/ying32/govcl/blob/master/LICENSE)
![Recommended Golang Version](https://img.shields.io/badge/recommended%20golang%20version->=1.9.0-green.svg)
[![screenshots](https://img.shields.io/badge/screenshots-view-green.svg)](https://github.com/ying32/govcl/tree/master/Screenshot)
[![Chinese Wiki](https://img.shields.io/badge/wiki-中文WIKI(Chinese%20WIKI)-green.svg)](https://gitee.com/ying32/govcl/wikis/pages)
[![Chinese Chat](https://img.shields.io/badge/QQ群-点击加入：263106281-red.svg)](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq)  

![Support Platform](https://img.shields.io/badge/Platform-win--32%20%7C%20win--64%20%7C%20linux--64%20%7C%20osx--32-green.svg)  
**Note: linux and macOS only part of the components, properties, events and methods are valid**   


Librarys binaries  
[![Librarys](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.1.21.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.1.21/Librarys-1.1.21.zip)  
**Note: The attachment contains liblcl three platform binary, libvcl does not currently provide, you need to compile. For the compilation steps, please refer to the instructions in [UILIbSrcources](UILibSources/README.md)**  


res2go Tool  
[![res2go](https://img.shields.io/badge/downloads-res2go%201.0.4.zip-blue.svg)](https://github.com/ying32/govcl/blob/master/Tools/res2go)  
**Note: Designed in Delphi/Lazarus, code written in Golang.**  


VCL style files  
[![VCL style files](https://img.shields.io/badge/downloads-VCL%20style%20files-blue.svg)](https://github.com/ying32/govcl/releases/download/v1.1.20/vcl-styles.zip)  

### usage: 

> go get github.com/ying32/govcl    

```go
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
---   
### FAQ

Q: Why is there no English WIKI?   
A: My English is bad. You can try using Google Translate Chinese WIKI.    
 
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
