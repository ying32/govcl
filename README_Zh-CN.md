# govcl

## 目录
* [项目介绍](#项目介绍)
* [重要说明](#重要说明)
* [UI设计器](https://gitee.com/ying32/govcl/wikis/pages?title=%E5%89%8D%E8%A8%80&parent=%E8%AE%BE%E8%AE%A1%E5%99%A8%E5%B8%AE%E5%8A%A9)
* [FAQ](https://gitee.com/ying32/govcl/wikis/pages)
* [使用方法](#使用方法)
* [支持的组件列表](https://gitee.com/ying32/govcl/wikis/pages?title=%E6%94%AF%E6%8C%81%E7%9A%84%E7%BB%84%E4%BB%B6%E5%88%97%E8%A1%A8&parent=)
* [截图预览](https://gitee.com/ying32/govcl/tree/master/Screenshot)
* [点击链接加入群【govcl交流】](https://jq.qq.com/?_wv=1027&k=5Sv7Qiq) 群号：263106281

---
### 项目介绍

* 1、由于现有第三方的Go UI库不是太庞大就是用的不习惯，或者组件太少。就萌生了自己写一个UI库的想法  
     Delphi(Lazarus)有些许多优秀的VCL(LCL)组件，不拿来使用太可惜了。所以就索性做了一套。  
     目前支持`Win32`、`Win64`、`Linux64`、`MacOS32`(对于`Linux64`、`MacOS32`提供**有限的组件、属性及函数方法**的支持)。      
 
* 2、项目现在支持VCL(LCL)标准控件中的大部分，足以满足日常操作了，具体见[支持的组件列表](https://gitee.com/ying32/govcl/wikis/pages?title=%E6%94%AF%E6%8C%81%E7%9A%84%E7%BB%84%E4%BB%B6%E5%88%97%E8%A1%A8&parent=)。  
     事件方面也支持部分，参见：[支持的事件](https://gitee.com/ying32/govcl/wikis/pages?title=%E6%94%AF%E6%8C%81%E7%9A%84%E4%BA%8B%E4%BB%B6&parent=)    
 
---
### 重要说明
* 码云上的代码暂时不再更新了，全部迁移到这。     
 
* 推荐Go版 `Go Version >= 1.9.0` 。   

* 希望大家有问题的话通过[Issues](https://github.com/ying32/govcl/issues)来进行反馈，反馈错误的话最好能带有相关错误的截图之类的， 而不是通过评论来提问。wiki也可关注下，有些问题在会里面作解答。  

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


 