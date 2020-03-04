//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

/*
   要自定义加载dll或者so之类的，需要按照go初始包的规则来做，
   Go是按文件名排序的，也就是说要创建一个自定义的包，这个包的文件名必须比以下这个包的
   github.com/ying32/govcl/vcl 的文件名小，要能在main包里排在这个之前

   例：
     package main

     import "yourpackage" // 必须在自动排序要能排在下面包之前方能使用。
     import "github.com/ying32/govcl/vcl"
*/

package libname

var (
	LibName string
)
