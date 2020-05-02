* [中文](README.md)   
* English     

----

### liblcl 

liblcl Based on Lazarus 2.0.8 version FPC 3.0.4, the specific installation method of each platform refers to the official website installation instructions.   

The liblcl project source code is located in the "govcl\UILibSources\liblcl" directory. `liblcl for Windows, Linux, MacOS`.   

Compile steps:    

* 1. Download and install [Lazarus](https://www.lazarus-ide.org/index.php?page=downloads) of the corresponding platform.
* 2. [Third Party Control](liblcl/3rd-party/README.en-US.md) required for installation.
* 3. double-click liblcl.lpi.  
* 4. Menu -> Project -> Project Options -> Compiler Options -> Build modes Switch the relevant ring mode. The current valid mode is as follows:  
   * Win32  
   * Win64  
   * Linux64  
   * Linux32
   * LinuxARM           
   * MacOS64(cocoa) Make sure under Tools->Options that "Compiler Executable" is set to "/usr/local/bin/fpc" to get 64 bit apps(32bit Lazarus).   
  
* 5. menu -> Run-> Build or Shift + F9 (Must use Build)   

The compiled binary can be viewed in the following directory:      

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`
