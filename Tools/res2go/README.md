### res2go  

* ---------------------------------------------中文---------------------------------------------  

`res2go是一个将Lazarus/Delphi资源窗口转go工具，可自动解析lfm、dfm中的组件名、组件类型、事件名称。解析lpr、dpr文件中窗口信息。`  

命令行：  
```
用法：res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path      待转换的工程路径，可为空，默认以当前目录为准。
  -outpath   输出目录，可为空，默认为当前目录。
  -outmain   是否输出“main.go”，此为解析lpr或者dpr文件，默认输出。
  -outres    输出一个Windows默认资源文件，如果存在则不创建，默认输出。
  -scale     缩放窗口选项，默认为不缩放。
  -h -help   显示帮助。
```


* ---------------------------------------------English---------------------------------------------  

`res2go is a Lazarus/Delphi resource window to go tool, can automatically resolve the lfm, dfm component name, component type and event name. Parse window information in lpr, dpr file.`   


Command Line:  
```
usage: res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path      The project path to be converted can be empty. The default is the directory.
  -outpath   Output directory, can be empty, the default is the current directory.
  -outmain   Whether to output "main.go", this is parsing lpr or dpr file, the default output.
  -outres    Outputs a Windows default resource file, if it does not exist, the default output.
  -scale     The window scale option, the default is false.
  -h -help   Show help.
```