### res2go  

* 中文    
* [English](README.en-US.md)  

----

`res2go是一个将Lazarus/Delphi资源窗口转go工具，可自动解析lfm、dfm中的组件名、组件类型、事件名称。解析lpr、dpr文件中窗口信息。`  

命令行：  

```
用法：res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path       待转换的工程路径，可为空，默认以当前目录为准。
  -outpath    输出目录，可为空，默认为当前目录。
  -outmain    是否输出“main.go”，此为解析lpr或者dpr文件，默认为true。
  -outres     输出一个Windows默认资源文件，如果存在则不创建，默认为true。
  -scale      缩放窗口选项，默认为false，默认为不缩放。  
  -encrypt    使用加密格式的*.gfm文件，默认为false。
  -usestr     加上此参数会以字符形式输出字节，默认为true。 
  -origfn     生成的.go文件使用原始的delphi/lazarus单元名，默认为false。  
  -pause      结束后根据选项暂停，比如： -pause "ew"，表示有错或者警告，可选为“e”,“w”,“a” e=错误，w=警告，a=忽略其它选项，总是显示。
  -pkgname    指定生成的go文件包名，默认为main。
  -watch      监视“-path”目录的文件，如果有改变则进行转换。
  -h -help    显示帮助。
  -v -version 显示版本号。
```

---- 

### 集成到Delphi/Lazarus IDE内 

* Delphi IDE

打开IDE： 菜单 -> Tools -> Configure Tools -> Add, 显示了Tool Properties窗口   

```
Title              菜单栏显示的名字  
Program            res2go程序全文件名（含路径）  
Working directory  工作目录，可不填  
Parameters         命令行参数（填这句，运行后会在当前工程目录下的gocode生成代码）： -path "$PATH($PROJECT)" -outpath "$PATH($PROJECT)/gocode" -pause "ew"   
```

* Lazarus IDE  

打开IDE： 菜单 -> Tools -> Configure External Tools -> Add, 显示了Edit Tool窗口  

```
Title              菜单栏显示的名字   
Program Filename   res2go程序全文件名（含路径）   
Parameters 命令行参数（填这句，运行后会在当前工程目录下的gocode生成代码）： -path "$Path($ProjFile())" -outpath "$Path($ProjFile())/gocode" -pause "ew"     
Working directory  工作目录，可不填   

Lazarus 还可以额外填写快捷键，在Key分组里面设置。  
```