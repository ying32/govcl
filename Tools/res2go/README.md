### res2go  

* 中文    
* [English](README.en-US.md)  

----

**注意：res2go命令行工具会在将来某个时候废弃，请改用[res2go IDE插件](https://gitee.com/ying32/govcl/wikis/pages?sort_id=2645001&doc_id=102420)，目前在QQ群内测试。**  

----

`res2go是一个将Lazarus资源窗口转go工具，可自动解析lfm、dfm中的组件名、组件类型、事件名称。解析lpr、dpr文件中窗口信息。`  

命令行：  

```
用法：res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-watch]
  -path       待转换的工程路径，可为空，默认以当前目录为准。
  -outpath    输出目录，可为空，默认为当前目录。 
  -origfn     生成的.go文件使用原始的delphi/lazarus单元名，默认为false。  
  -pkgname    指定生成的go文件包名，默认为main。
  -watch      监视“-path”目录的文件，如果有改变则进行转换。
  -h -help    显示帮助。
  -v -version 显示版本号。
```

---- 

### 集成到Lazarus IDE内 

* Lazarus IDE  

打开IDE： 菜单 -> Tools -> Configure External Tools -> Add, 显示了Edit Tool窗口  

```
Title              菜单栏显示的名字   
Program Filename   res2go程序全文件名（含路径）   
Parameters 命令行参数（填这句，运行后会在当前工程目录下的gocode生成代码）： -path "$Path($ProjFile())" -outpath "$Path($ProjFile())/gocode"
Working directory  工作目录，可不填   

选中"Scan output for FPC messages"  

Lazarus 还可以额外填写快捷键，在Key分组里面设置。  
```