* 中文   
* [English](README.en-US.md)   

----

**自定义加载指定位置的libvcl或者liblcl**

> 要自定义加载dll或者so之类的，需要按照go初始包的规则来做。
> Go是按文件名排序的，也就是说要创建一个自定义的包，这个包的文件名必须比以下这个包的
> github.com/ying32/govcl/vcl 的文件名小，要能在main包里排在这个之前

* 例：
```go
package main

import "yourpackage" // 必须在自动排序后，排在vcl包之前方能生效。
import "github.com/ying32/govcl/vcl"

```
