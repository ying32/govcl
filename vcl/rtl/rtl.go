//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package rtl

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

var (
	// SysLocale 本地化相关
	SysLocale types.TSysLocale
)

// Move Delphi中的内存操作，不过这里传入的是指针
func Move(src, dest uintptr, llen int) {
	api.DMove(src, dest, llen)
}

// StrLen String长度，一般用于对delphi或者lazarus的字符长检测
func StrLen(str uintptr) int {
	return api.DStrLen(str)
}

// GetStringArrOf 从一个Delphi/Lazarus字符串数组获取成员
func GetStringArrOf(p uintptr, index int) string {
	return api.DGetStringArrOf(p, index)
}

// IsNil 判断一个接口是否为空
// interface{}数据类型定义为 typedef struct { void *type; void *value; } GoInterface;
// 当type与value值都为nil时则为空。
func IsNil(val interface{}) bool {
	return api.IsNil(val)
}

//----------------------------Delphi/Lazarus集合操作-------------------------------------------------------

// Include Delphi集合加法，val...中存储为位的索引，下标为0
func Include(r uint32, val ...uint8) uint32 {
	return uint32(types.TSet(r).Include(val...))
}

// Exclude Delphi集合减法，val...中存储为位的索引，下标为0
func Exclude(r uint32, val ...uint8) uint32 {
	return uint32(types.TSet(r).Exclude(val...))
}

// InSets Delphi集合类型的判断,类型，然后后面是第几位，下标为0
func InSets(r uint32, s uint32) bool {
	return types.TSet(r).In(s)
}

//-------------------------------------------------------------------------------------------------------

// TextToShortCut 将字符串转为ShortCut类型
func TextToShortCut(val string) types.TShortCut {
	return api.DTextToShortCut(val)
}

// ShortCutToText 将ShortCut类型转为字符串
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

// ExtractFilePath 提取文件名的路径，带“\”的
func ExtractFilePath(filename string) string {
	return api.DExtractFilePath(filename)
}

// FileExists 判断文件是否存在
func FileExists(filename string) bool {
	return api.DFileExists(filename)
}

// Ext 获取文件扩展名
func ExtractFileExt(path string) string {
	return filepath.Ext(path)
}

// GetFileName 获取一个文件名
func ExtractFileName(path string) string {
	return filepath.Base(path)
}

// GetFileNameWithoutExt 获取一个无扩展的文件名
func GetFileNameWithoutExt(path string) string {
	filename := ExtractFileName(path)
	return filename[:len(filename)-len(ExtractFileExt(filename))]
}

// ExtractFilePath 提取文件名路径
//func ExtractFilePath(path string) string {
//	filename := GetFileName(path)
//	return path[:len(path)-len(filename)]
//	//return filepath.Dir(path) + string(filepath.Separator)
//}

// Combine 合并
func Combine(path, name string) string {
	if path != "" && !strings.HasSuffix(path, string(os.PathSeparator)) {
		path += string(os.PathSeparator)
	}
	if name != "" && strings.HasPrefix(name, string(os.PathSeparator)) {
		name = name[1:]
	}
	return path + name
}

// FileExists
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

// LcLLoaded 是否加载的为lcl库，true表是是，false表示不是
func LcLLoaded() bool {
	return api.IsloadedLcl
}

// SetProperty
// SetPropertyValue 设置对象属性
func SetPropertyValue(instance uintptr, propName, value string) {
	api.DSetPropertyValue(instance, propName, value)
}

// SetPropertySecValue 设置对象二级属性
func SetPropertySecValue(instance uintptr, propName, secPropName, value string) {
	api.DSetPropertySecValue(instance, propName, secPropName, value)
}

// LibResouces
func GetLibResouceCount() int32 {
	return api.DGetLibResouceCount()
}

func GetLibResouceItem(aIndex int32) types.TLibResouce {
	return api.DGetLibResouceItem(aIndex)
}

func GetLibResouceItems() []types.TLibResouce {
	ret := make([]types.TLibResouce, GetLibResouceCount())
	for i := 0; i < len(ret); i++ {
		ret[i] = GetLibResouceItem(int32(i))
	}
	return ret
}

func ModifyLibResouce(aPtr uintptr, aValue string) {
	api.DModifyLibResouce(aPtr, aValue)
}

// 库的信息
// 获取当前库使用的字符串编码
func LibStringEncoding() types.TStringEncoding {
	return api.DLibStringEncoding()
}

// // 共8位，2位2位的，如：$01020100 表示 1.2.1.0
func LibVersion() uint32 {
	return api.DLibVersion()
}

// Linux下liblcl使用GTK2用
func GetGDKWindowXID(handle types.HWND) types.TXID {
	return api.DGetGDKWindowXID(handle)
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
	if shift.In(types.SsCommand) {
		result += MOD_WIN
	}
	return result
}
