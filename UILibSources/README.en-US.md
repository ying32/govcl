* [中文](README.md)   
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

liblcl Based on Lazarus 2.0.4 version FPC 3.0.4, the specific installation method of each platform refers to the official website installation instructions.  


Compile steps:     

* 1. Install Lazarus 2.0.4 64bit version and i386 expansion package  
* 2. double-click lcl.lpi  
* 3. Menu -> Project -> Project Options -> Compiler Options -> Build modes Switch the relevant ring mode. The current valid mode is as follows:  
   * Win32  
   * Win64  
   * Linux64  
   * Linux32
   * LinuxARM           
   * MacOS32(carbon)
   * MacOS64(cocoa) Make sure under Tools->Options that "Compiler Executable" is set to "/usr/local/bin/fpc" to get 64 bit apps.   
* 4. menu -> Run-> Compile (or Build)  

**Note: If you want to compile ARM and Linux 32-bit liblcl, you need to install the corresponding [Lazarus](http://www.lazarus-ide.org/) and FPC on the corresponding platform. Also consider installing [CodeTyphon](http://www.pilotlogic.com/sitejoom/index.php/codetyphon) for cross-compilation, but it is quite troublesome. Under MacOS, carbon and cocoa are different UI interfaces, but cocoa is unstable, and many components have problems. Generally, carbon components can be used.**    

The compiled binary can be viewed in the following directory:      

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**Note: The source code has no third-party dependencies, and can be compiled directly after installing the relevant version of Delphi.**  
