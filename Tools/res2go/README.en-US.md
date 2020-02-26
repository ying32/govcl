### res2go  

* [中文](README.md)  
* English    

----

`res2go is a Lazarus/Delphi resource window to go tool, can automatically resolve the lfm, dfm component name, component type and event name. Parse window information in lpr, dpr file.`   

Command Line:  
```
usage: res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-outmain true] [-outres true] [-scale]
  -path       The project path to be converted can be empty. The default is the current directory.
  -outpath    Output directory, can be empty, the default is the current directory.
  -outmain    Whether to output "main.go", this is parsing lpr or dpr file, the default is true.
  -outres     Outputs a Windows default resource file, if it does not exist, the default is true.
  -scale      The window scale option, the default is false.
  -encrypt    Using the encrypted format of the *.gfm file, the default is false.  
  -usestr     adding this parameter will output the bytes as characters, the default is false. 
  -origfn     The generated .go file uses the original delphi/lazarus unit name, the default is false.
  -pause      After the end, pause according to the option, for example: -pause "ew", indicating that there is a fault or warning, you can choose "e", "w", "a" e=error, w=warning, a=ignore other options, always display.
  -pkgname    Specifies the name of the generated go file package. The default is main.
  -watch      Monitor files in the "-path" directory and convert if there are changes.
  -h -help    Show help.
  -v -version Show Version.
```


---- 

### Integrated into the Delphi/Lazarus IDE  

* Delphi IDE

Open the IDE: Menu -> Tools -> Configure Tools -> Add   

```
Title              The name displayed on the menu  
Program            res2go program full file name (including path) 
Working directory  Working directory, no need to fill  
Parameters         Command line arguments (the code will be generated in the gocode directory under the current project directory after running): -path "$PATH($PROJECT)" -outpath "$PATH($PROJECT)/gocode" -pause "ew"    
```

* Lazarus IDE  

Open the IDE: Menu -> Tools -> Configure External Tools -> Add  

```
Title              The name displayed on the menu     
Program Filename   res2go program full file name (including path) 
Parameters         Command line arguments (the code will be generated in the gocode directory under the current project directory after running): -path "$Path($ProjFile())" -outpath "$Path($ProjFile())/gocode" -pause "ew"   
Working Directory  Working directory, no need to fill     

Lazarus can also fill in the shortcut keys and set them in the Key group.  
```