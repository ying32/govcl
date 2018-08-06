
* [中文](README.zh-CN.md)   
* [English](README.md)   

----  

### libvcl

libvcl 基于Delphi 10.2.1版本, 其它版本编译出现的问题请自行解决相关。

libvcl工程源码位于"govcl\UILibSources\libvcl"目录。`libvcl限仅于windows`     

编译步骤：  


* 1、安装好Delphi 10.2.1。  
* 2、双击vcl.dproj或者vcl.dpr在IDE中打开工程。  
* 3、在右边Project Manager的File列表中选择Build Configurations为Release。
* 4、在右边Project Manager的File列表中选择Target Platforms (32-bit Windows或者64-bit Windows)。  
* 5、按下Ctrl+F9（或直接在Release配置右键，选择Make或者Build）。  

因为配置中设置了相关编译后的操作，编译完后的二进制可以在以下目录查看  

> %GOPATH%\bin\libvcl.dll  
> %GOPATH%\bin\libvclx64.dll    
> %GOPATH%\src\github.com\ying32\bin\win32\libvcl.dll  
> %GOPATH%\src\github.com\ying32\bin\win64\libvclx64.dll    


**注：源码无第三方依赖库，安装好相关版本Delphi后直接编译即可。**


----

### liblcl 

liblcl 基于Lazarus 1.8.0版本 FPC 3.0.4，具体各个平台的安装方式自行参考官网安装说明, 其它版本编译出现的问题请自行解决相关。。

libvcl工程源码位于"govcl\UILibSources\liblcl"目录。`liblcl适用于Win32、Win64、Linux64、MacOS32`     

编译步骤：  

* 1、安装好Lazarus 1.8.0 64bit版本及i386扩展包   
* 2、双击lcl.lpi  
* 3、菜单->Project->Project Options -> Compiler Options -> Build modes 切换相关编环模式，当前有效模式为以下4种：   
  * ReleaseWindows64  
  * ReleaseWindows32  
  * ReleaseLinux64  
  * ReleaseMacOS32  
* 4、菜单->Run->Compile(或者Build)  

> Lazarus在编译时可选择不同接口的UI库，位于：菜单->Project->Project Options -> Compiler Options -> Additions and Overrides 右边：Set "LCLWidegetType", 可选值有：gtk、gtk2、gtk3、qt、qt5、win32、wince、fpgui、nogui、carbon、cocoa、customdraw、mui， 虽然有这么多可选的，但也不是什么都能编译过的，相关可参考lazarus官方的文档。 Windows默认接口为win32, Linux默认接口为gtk2，MacOS默认接口为carbon，cocoa是在现苹果使用的图形接口，还支持64位的，不过不稳定，有些组件都使用不了，还是使用老的carbon吧，32位就32位呗。 像gtk、Qt这些跨平台的是可以进行3个平台用的，具体看你自己了，我这只作一个小小的解释说明。    

生成的文件位于：  

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**注：源码无第三方依赖库，安装好相关版本Lazarus后直接编译即可。**  