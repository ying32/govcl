* 中文   
* [English](README.en-US.md)   

----

**自动打包成一个macOS下的app格式**

#### 使用方法

如果`main`包中有多个go文件，则创建一个`0.go`文件。然后导入`macapp`包。

```go
import _ "github.com/ying32/govcl/pkgs/macapp"
```
