## 内存加载dll示例，此方法目前只适用于win32，而且被编译到exe内部的liblcl/libvcl不能使用upx进行压缩。。

### 资源定义  resDefine.rc

```
1 MANIFEST "exe.manifest"
MAINICON ICON "icon.ico"
GOVCLLIB RCDATA "libvcl.dll"
```

### 资源编译 buildRes.bat

需要借助mingw中的windres资源编译器。

```
windres.exe -o defaultRes_windows.syso resDefine.rc
```

更多说明参考：https://gitee.com/ying32/govcl/wikis/pages?sort_id=1525083&doc_id=102420