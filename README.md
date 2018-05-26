# govcl

* [中文](https://github.com/ying32/govcl/blob/master/README_Zh-CN.md)   
* [English](https://github.com/ying32/govcl/blob/master/README.md)   


# Do not use this repository code. This is a backup.  

## golang binding Delphi VCL/Lazarus LCL

Support win32, win64, linux64, macOS32 (linux and macOS only part of the components, properties, events and methods are valid)    

**[See more screenshots](https://github.com/ying32/govcl/tree/master/Screenshot)**  


`Go Version >= 1.9.0`  

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
### FAQ

Q: Why not submit the code on github.com?  
A: Visit site too slow.  


Q: Why is there no English WIKI?   
A: My English is bad. You can use Google Translate.    


Q:  What if I want to use the code of this repository and not want to use gitee repository?  
A:  Do not use the "go get" command. First create the directory "%GOPATH%/src/gitee.com/ying32". Then use the "git clone" command in this directory.  

---  
### Chinese wiki  

[Chinese wiki](https://gitee.com/ying32/govcl/wikis/pages)  


---  
### Other 

**All binaries are moved to the attachment and [download](https://gitee.com/ying32/govcl/attach_files) .**   

* Windows: copy "bin\win32\libvcl.dll" or "bin\win64\libvclx64.dll" to the current exe directory or system environment path. If you don't want to use Delphi binary, you can copy the corresponding LCL library binaries under the "bin\liblcl.dll\" directory. Note: the use of LCL is limited in components, events, properties, and methods.

* Linux: copy "bin\linux64\liblcl.so" executable file directory (also can be copied from the liblcl.so to `/usr/lib/` directory as a Public Library).

* MacOS: copy "bin\MacOS32\liblcl.dylib" executable file directory. [APP package](https://gitee.com/ying32/govcl/wikis/pages?title=APP%E6%89%93%E5%8C%85&parent=FAQ%2FMac-OS).

---  
