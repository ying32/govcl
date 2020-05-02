
* 中文   
* [English](README.en-US.md)   

----

### liblcl 

liblcl 基于Lazarus 2.0.8版本 FPC 3.0.4，具体各个平台的安装方式自行参考官网安装说明, 其它版本编译出现的问题请自行解决相关。。

liblcl工程源码位于"govcl\UILibSources\liblcl"目录。`liblcl适用于Windows、Linux、MacOS`。       

编译步骤：  

* 1、下载并安装对应平台的[Lazarus](https://www.lazarus-ide.org/index.php?page=downloads)。
* 2、安装所需要的[第三方控件](liblcl/3rd-party/README.md)。
* 3、双击liblcl.lpi  
* 4、菜单->Project->Project Options -> Compiler Options -> Build modes 切换相关编环模式，当前有效模式为以下几种种：   
   * Win32  
   * Win64  
   * Linux64  
   * Linux32
   * LinuxARM           
   * MacOS64(cocoa) 确保在Tools-> Options下将“Compiler Executable”设置为“/usr/local/bin/fpc”以获得64位应用程序(32bit Lazarus)。 
* 5、菜单 -> Run -> Build 或者 Shift + F9 (必须使用Build)  

**注意： 如你想要编译ARM和Linux 32位的liblcl则需要在对应平台安装相应的[Lazarus](http://www.lazarus-ide.org/)和FPC。也可考虑安装[CodeTyphon](http://www.pilotlogic.com/sitejoom/index.php/codetyphon)进行交叉编译，不过挺麻烦。**

生成的文件位于：  

> Windows: `"..\..\..\..\..\bin\liblcl"`     
> Linux: `"../bin/liblcl"`  
> MacOS: `"../../../../../bin/liblcl"`
