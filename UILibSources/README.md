* [中文](README.zh-CN.md)   
* English     

----  

### libvcl

libvcl based on Delphi 10.2.1.  

The libvcl project source is located in the "UILibSources\libvcl" directory. `libvcl is limited to windows only.      

Compile steps:    

* 1. Install Delphi 10.2.1.  
* 2. Double-click vcl.dproj or vcl.dpr to open the project in the IDE.  
* 3. Select Build Configurations as Release in the File list on the right of Project Manager.  
* 4. Select Target Platforms (32-bit Windows or 64-bit Windows) in the File list on the right Project Manager.  
* 5. press Ctrl + F9 (or directly in the Release configuration, select Make or Build).  

The compiled binary can be viewed in the following directory:     

> ..\libvcl.dll  
> ..\x64\libvclx64.dll    

**Note: The source code has no third-party dependencies, and can be compiled directly after installing the relevant version of Delphi.**    


----

### liblcl 

liblcl Based on Lazarus 1.8.0 version FPC 3.0.4, the specific installation method of each platform refers to the official website installation instructions.  


Compile steps:     

* 1. Install Lazarus 1.8.0 64bit version and i386 expansion package  
* 2. double-click lcl.lpi  
* 3. Menu -> Project -> Project Options -> Compiler Options -> Build modes Switch the relevant ring mode, the current effective mode is the following four:  
    * ReleaseWindows64  
    * ReleaseWindows32  
    * ReleaseLinux64  
    * ReleaseMacOS32  
* 4. menu -> Run-> Compile (or Build)  

 
The compiled binary can be viewed in the following directory:      

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**Note: The source code has no third-party dependencies, and can be compiled directly after installing the relevant version of Delphi.**  
