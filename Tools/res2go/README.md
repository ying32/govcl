### res2go  

* [中文](README.zh-CN.md)  
* English    

----

`res2go is a Lazarus/Delphi resource window to go tool, can automatically resolve the lfm, dfm component name, component type and event name. Parse window information in lpr, dpr file.`   

Command Line:  
```
usage: res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path       The project path to be converted can be empty. The default is the current directory.
  -outpath    Output directory, can be empty, the default is the current directory.
  -outmain    Whether to output "main.go", this is parsing lpr or dpr file, the default output.
  -outres     Outputs a Windows default resource file, if it does not exist, the default output.
  -outbytes   Save the gfm file as a byte to the go file, the default output.
  -scale      The window scale option, the default is false.
  -encrypt    Using the encrypted format of the *.gfm file, the default is true.  
  -h -help    Show help.
  -v -version Show Version.
```