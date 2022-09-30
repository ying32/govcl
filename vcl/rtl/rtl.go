//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import (
	"path/filepath"
	"strings"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

var (
	// SysLocale
	//
	// 本地化相关
	//
	// localization.
	SysLocale types.TSysLocale
)

// Move
//
// FreePascal中的内存操作，不过这里传入的是指针
//
// Memory operations in FreePascal, but pointers are passed here.
func Move(src, dest uintptr, llen int) {
	api.DMove(src, dest, llen)
}

// StrLen
//
// FreePascal的字符串长度。
//
// FreePascal string length.
func StrLen(str uintptr) int {
	return api.DStrLen(str)
}

// GetStringArrOf
//
// 从一个FreePascal字符串数组获取成员
func GetStringArrOf(p uintptr, index int) string {
	return api.DGetStringArrOf(p, index)
}

//----------------------------Delphi/Lazarus集合操作-------------------------------------------------------

// Include
// Deprecated: use value.Include.
func Include(r uint32, val ...uint8) uint32 {
	return uint32(types.TSet(r).Include(val...))
}

// Exclude
// Deprecated: use value.Exclude.
func Exclude(r uint32, val ...uint8) uint32 {
	return uint32(types.TSet(r).Exclude(val...))
}

// InSets
// Deprecated: use value.In.
func InSets(r uint32, s uint32) bool {
	return types.TSet(r).In(s)
}

//-------------------------------------------------------------------------------------------------------

// TextToShortCut
//
// 将字符串转为ShortCut类型
//
// Convert string to ShortCut type.
func TextToShortCut(val string) types.TShortCut {
	return api.DTextToShortCut(val)
}

// ShortCutToText
//
// 将ShortCut类型转为字符串
//
// Convert ShortCut type to string.
func ShortCutToText(val types.TShortCut) string {
	return api.DShortCutToText(val)
}

// SysOpen 打开，windows下调用ShellExecute
/*
	// windows
	rtl.SysOpen("http://www.xxx.com")
	rtl.SysOpen("c:\")
	rtl.SysOpen("c:\xxx.exe")

	// linux or macOS
	rtl.SysOpen("https://wwww.xxx.com")
	rtl.SysOpen("file:///xxx.png");
*/
func SysOpen(filename string) {
	api.DSysOpen(filename)
}

// ExtractFilePath
//
// 提取文件名的路径，带“\”的
func ExtractFilePath(filename string) string {
	return api.DExtractFilePath(filename)
}

// FileExists
//
// 判断文件是否存在
func FileExists(filename string) bool {
	return api.DFileExists(filename)
}

// ExtractFileExt
//
// 提取文件扩展名
func ExtractFileExt(path string) string {
	return filepath.Ext(path)
}

// ExtractFileName
//
// 获取一个文件名
func ExtractFileName(path string) string {
	return filepath.Base(path)
}

// GetFileNameWithoutExt
//
// 获取一个无扩展的文件名
func GetFileNameWithoutExt(path string) string {
	filename := ExtractFileName(path)
	return filename[:len(filename)-len(ExtractFileExt(filename))]
}

// ExtractFilePath
//
// 提取文件名路径
//func ExtractFilePath(path string) string {
//	filename := GetFileName(path)
//	return path[:len(path)-len(filename)]
//	//return filepath.Dir(path) + string(filepath.Separator)
//}

// Combine
//
// 路径合并
func Combine(path, name string) string {
	if path != "" && !strings.HasSuffix(path, PathSeparator) {
		path += PathSeparator
	}
	if name != "" && strings.HasPrefix(name, PathSeparator) {
		name = name[1:]
	}
	return path + name
}

//func FileExists(path string) bool  {
//	_, err := os.Stat(path)
//	if err == nil {
//		return true
//	}
//	if os.IsNotExist(err) {
//		return false
//	}
//	return false
//}

// ------------------- SetProperty

// SetPropertyValue
//
// 设置对象属性值
//
// Set object property value
func SetPropertyValue(instance uintptr, propName, value string) {
	api.DSetPropertyValue(instance, propName, value)
}

// SetPropertySecValue
//
// 设置对象二级属性值
//
// Set the secondary attribute value of the object
func SetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	api.DSetPropertySecValue(instance, propName, secPropName, value)
}

// GetLibResourceCount
//
// 获取Lib中资源字符数组总数
func GetLibResourceCount() int32 {
	return api.DGetLibResourceCount()
}

// GetLibResourceItem
//
// 从指定索引中获取字符资源项目
func GetLibResourceItem(aIndex int32) types.TLibResource {
	return api.DGetLibResourceItem(aIndex)
}

// GetLibResourceItems
//
// 获取资源字符串数组
func GetLibResourceItems() []types.TLibResource {
	ret := make([]types.TLibResource, GetLibResourceCount())
	for i := 0; i < len(ret); i++ {
		ret[i] = GetLibResourceItem(int32(i))
	}
	return ret
}

// ModifyLibResource
//
// 修改指定资源字符串
func ModifyLibResource(aPtr uintptr, aValue string) {
	api.DModifyLibResource(aPtr, aValue)
}

// LibVersion
//
// 获取库的版本，共8位，2位2位的，如：$01020100 表示 1.2.1.0
func LibVersion() uint32 {
	return api.DLibVersion()
}

func ShiftStateToWord(shift types.TShiftState) uint32 {
	// 这里不直接使用win包的常量，是考虑要跨平台使用
	const (
		MOD_ALT      = 1
		MOD_CONTROL  = 2
		MOD_SHIFT    = 4
		MOD_WIN      = 8
		MOD_NOREPEAT = 0x4000
	)
	var result uint32
	if shift.In(types.SsShift) {
		result += MOD_SHIFT
	}
	if shift.In(types.SsCtrl) {
		result += MOD_CONTROL
	}
	if shift.In(types.SsAlt) {
		result += MOD_ALT
	}
	// lcl没不知道哪个值
	//if shift.In(types.SsCommand) {
	//	result += MOD_WIN
	//}
	return result
}

// LibAbout
//
// liblcl的关于信息
func LibAbout() string {
	return api.DLibAbout()
}

// MainThreadId
//
// 返回主线程ID
//
// Return the main thread id.
func MainThreadId() uintptr {
	return api.DMainThreadId()
}

// CurrentThreadId
//
// 返回当前线程iD
//
// Return the current thread id.
func CurrentThreadId() uintptr {
	return api.DCurrentThreadId()
}

// InitGoDll
//
// 一般无用，主要用来在go生成的dll中使用liblcl。
func InitGoDll(aMainThreadId uintptr) {
	api.DInitGoDll(aMainThreadId)
}
