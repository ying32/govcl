
* 中文   
* [English](README.en-US.md)   

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

> ..\libvcl.dll  
> ..\x64\libvclx64.dll    
 
**注：源码无第三方依赖库，安装好相关版本Delphi后直接编译即可。**


----

### liblcl 

liblcl 基于Lazarus 2.0.6版本 FPC 3.0.4，具体各个平台的安装方式自行参考官网安装说明, 其它版本编译出现的问题请自行解决相关。。

libvcl工程源码位于"govcl\UILibSources\liblcl"目录。`liblcl适用于Win32、Win64、Linux64、MacOS32`     

编译步骤：  

* 1、安装好Lazarus 2.0.6 64bit版本及i386扩展包   
* 2、双击lcl.lpi  
* 3、菜单->Project->Project Options -> Compiler Options -> Build modes 切换相关编环模式，当前有效模式为以下几种种：   
   * Win32  
   * Win64  
   * Linux64  
   * Linux32
   * LinuxARM           
   * MacOS32(carbon)
   * MacOS64(cocoa) 确保在Tools-> Options下将“Compiler Executable”设置为“/usr/local/bin/fpc”以获得64位应用程序。 
* 4、菜单->Run->Compile(或者Build)  

**注意： 如你想要编译ARM和Linux 32位的liblcl则需要在对应平台安装相应的[Lazarus](http://www.lazarus-ide.org/)和FPC。也可考虑安装[CodeTyphon](http://www.pilotlogic.com/sitejoom/index.php/codetyphon)进行交叉编译，不过挺麻烦。**

> Lazarus在编译时可选择不同接口的UI库，位于：菜单->Project->Project Options -> Compiler Options -> Additions and Overrides 右边：Set "LCLWidegetType", 可选值有：gtk、gtk2、gtk3、qt、qt5、win32、wince、fpgui、nogui、carbon、cocoa、customdraw、mui， 虽然有这么多可选的，但也不是什么都能编译过的，相关可参考lazarus官方的文档。 Windows默认接口为win32, Linux默认接口为gtk2，MacOS默认接口为carbon，cocoa是在现苹果使用的图形接口，还支持64位的。 像gtk、Qt这些跨平台的是可以进行3个平台用的，具体看你自己了，我这只作一个小小的解释说明。    

生成的文件位于：  

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`

**注：源码无第三方依赖库，安装好相关版本Lazarus后直接编译即可。**  