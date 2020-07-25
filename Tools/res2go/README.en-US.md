### res2go  

* [中文](README.md)  
* English    

----

`res2go is a Lazarus resource window to go tool, can automatically resolve the lfm, dfm component name, component type and event name. Parse window information in lpr, dpr file.`   

Command Line:  
```
usage: res2go [-path "C:\project\"] [-outpath "C:\xxx\"] [-watch]
  -path       The project path to be converted can be empty. The default is the current directory.
  -outpath    Output directory, can be empty, the default is the current directory.
  -origfn     The generated .go file uses the original delphi/lazarus unit name, the default is false.
  -pkgname    Specifies the name of the generated go file package. The default is main.
  -watch      Monitor files in the "-path" directory and convert if there are changes.
  -h -help    Show help.
  -v -version Show Version.
```


---- 

### Integrated into the Lazarus IDE  

* Lazarus IDE  

Open the IDE: Menu -> Tools -> Configure External Tools -> Add  

```
Title              The name displayed on the menu     
Program Filename   res2go program full file name (including path) 
Parameters         Command line arguments (the code will be generated in the gocode directory under the current project directory after running): -path "$Path($ProjFile())" -outpath "$Path($ProjFile())/gocode"
Working Directory  Working directory, no need to fill     

Check "Scan output for FPC messages"  

Lazarus can also fill in the shortcut keys and set them in the Key group.  
```