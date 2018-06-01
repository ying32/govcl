package rtl

import (
	"reflect"

	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
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

//----------------------------Delphi/Lazarus集合操作-------------------------------------------------------
// getInterfaceUint32Val 从一个interface{}中获取一个整型并转为uint32类型
func getInterfaceUint32Val(val interface{}) uint32 {
	if val == nil {
		return 0
	}
	switch f := val.(type) {
	case int:
		return uint32(f)
	case int8:
		return uint32(f)
	case int16:
		return uint32(f)
	case int32:
		return uint32(f)
	case int64:
		return uint32(f)
	case uint:
		return uint32(f)
	case uint8:
		return uint32(f)
	case uint16:
		return uint32(f)
	case uint32:
		return uint32(f)
	case uint64:
		return uint32(f)
	case uintptr:
		return uint32(f)
	default:
		value := reflect.ValueOf(val)
		if value.IsValid() {
			switch f2 := value; value.Kind() {
			case reflect.Int, reflect.Int8, reflect.Int16, reflect.Int32, reflect.Int64:
				return uint32(f2.Int())
			case reflect.Uint, reflect.Uint8, reflect.Uint16, reflect.Uint32, reflect.Uint64, reflect.Uintptr:
				return uint32(f2.Uint())
			}
		}
	}
	return 0
}

// Include Delphi集合加法，val...中存储为位的索引，下标为0
func Include(in interface{}, val ...uint8) uint32 {
	r := getInterfaceUint32Val(in)
	for _, v := range val {
		r |= (1 << uint8(v))
	}
	return r
}

// Exclude Delphi集合减法，val...中存储为位的索引，下标为0
func Exclude(in interface{}, val ...uint8) uint32 {
	r := getInterfaceUint32Val(in)
	for _, v := range val {
		r &= ^(1 << uint8(v))
	}
	return r
}

// InSets Delphi集合类型的判断,类型，然后后面是第几位，下标为0
func InSets(in interface{}, s uint32) bool {
	r := getInterfaceUint32Val(in)
	if r&(1<<uint8(s)) != 0 {
		return true
	}
	return false
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

// InheritsFromControl 判断对象是否继承自TControl
func InheritsFromControl(obj uintptr) bool {
	return api.DInheritsFromControl(obj)
}

// InheritsFromWinControl 判断对象是否继承自TWinControl
func InheritsFromWinControl(obj uintptr) bool {
	return api.DInheritsFromWinControl(obj)
}

// InheritsFromComponent 判断对象是否继承自TComponent
func InheritsFromComponent(obj uintptr) bool {
	return api.DInheritsFromComponent(obj)
}

// LcLLoaded 是否加载的为lcl库，true表是是，false表示不是
func LcLLoaded() bool {
	return api.IsloadedLcl
}
