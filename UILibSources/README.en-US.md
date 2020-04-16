* [中文](README.md)   
* English     

----  

### liblcl 

liblcl Based on Lazarus 2.0.8 version FPC 3.0.4, the specific installation method of each platform refers to the official website installation instructions.   

The liblcl project source code is located in the "govcl\UILibSources\liblcl" directory. `liblcl for Windows, Linux, MacOS`.   

Compile steps:    

**Note: If necessary, you can run the "govcl\Tools\LazarusPatch" tool to patch the Lazarus source code.**   

* 1. Install Lazarus 2.0.8 64bit version and i386 expansion package  
* 2. double-click lcl.lpi  
* 3. Menu -> Project -> Project Options -> Compiler Options -> Build modes Switch the relevant ring mode. The current valid mode is as follows:  
   * Win32  
   * Win64  
   * Linux64  
   * Linux32
   * LinuxARM           
   * MacOS32(carbon)
   * MacOS64(cocoa) Make sure under Tools->Options that "Compiler Executable" is set to "/usr/local/bin/fpc" to get 64 bit apps(32bit Lazarus).   
* 4. menu -> Run-> Compile (or Build)  

**Note: If you want to compile ARM and Linux 32-bit liblcl, you need to install the corresponding [Lazarus](http://www.lazarus-ide.org/) and FPC on the corresponding platform. Also consider installing [CodeTyphon](http://www.pilotlogic.com/sitejoom/index.php/codetyphon) for cross-compilation, but it is quite troublesome. Under MacOS, carbon and cocoa are different UI interfaces, but cocoa is unstable, and many components have problems. Generally, carbon components can be used.**    

The compiled binary can be viewed in the following directory:      

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**Note: The source code has no third-party dependencies, and can be compiled directly after installing the relevant version of Delphi.**  
