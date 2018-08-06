## A cross-platform Golang GUI library. Use Delphi VCL and Lazarus LCL for binding.

## [QQ群号:263106281](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq)   

* [中文](README.zh-CN.md)   
* [English](README.md)   

----

* [中文wiki(Chinese wiki) ](https://gitee.com/ying32/govcl/wikis/pages)  

----  

Support win32, win64, linux64, macOS32 (linux and macOS only part of the components, properties, events and methods are valid)    

---

* Librarys binaries  
[![Releases](https://img.shields.io/github/downloads/ying32/govcl/latest/Librarys-1.1.19.zip.svg)](https://github.com/ying32/govcl/releases/download/v1.1.19/Librarys-1.1.19.zip)

**Note: The attachment contains liblcl three platform binary, libvcl does not currently provide, you need to compile.**  

* res2go Tool  
[![Releases](https://img.shields.io/github/downloads/ying32/govcl/1.0.3/res2go.svg)](https://github.com/ying32/govcl/blob/master/Tools/res2go/res2go.zip)

---

**[See more screenshots](https://github.com/ying32/govcl/tree/master/Screenshot)**  


`Recommended Golang Version >= 1.9.0`  

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
### Other 

**binaries [download](https://github.com/ying32/govcl/releases) .**   

* Windows: copy "bin\win32\libvcl.dll" or "bin\win64\libvclx64.dll" to the current exe directory or system environment path. If you don't want to use Delphi binary, you can copy the corresponding LCL library binaries under the "bin\liblcl.dll\" directory. Note: the use of LCL is limited in components, events, properties, and methods.

* Linux: copy "bin\linux64\liblcl.so" executable file directory (also can be copied from the liblcl.so to `/usr/lib/` directory as a Public Library).

* MacOS: copy "bin\MacOS32\liblcl.dylib" executable file directory. [APP package](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS).

---  
