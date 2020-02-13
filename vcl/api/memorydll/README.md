### 简介

Windows上32/64bit内存加载DLL库。    


32位上使用go原生代码编写。  

64位上使用cgo编译第三方库    
直接使用cgo方式调用，因为不想再将C++的代码完全翻译成Go的了，太累了，也觉得没必要。   
https://github.com/fancycode/MemoryModule   