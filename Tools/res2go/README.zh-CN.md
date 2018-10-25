### res2go  

* 中文    
* [English](README.md)  

----

`res2go是一个将Lazarus/Delphi资源窗口转go工具，可自动解析lfm、dfm中的组件名、组件类型、事件名称。解析lpr、dpr文件中窗口信息。`  

命令行：  

```
用法：res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path       待转换的工程路径，可为空，默认以当前目录为准。
  -outpath    输出目录，可为空，默认为当前目录。
  -outmain    是否输出“main.go”，此为解析lpr或者dpr文件，默认输出。
  -outres     输出一个Windows默认资源文件，如果存在则不创建，默认输出。
  -outbytes   将gfm文件以字节形式保存至go文件中，默认输出。
  -scale      缩放窗口选项，默认为不缩放。  
  -encrypt    使用加密格式的*.gfm文件，默认为true。  
  -h -help    显示帮助。
  -v -version 显示版本号。
```

