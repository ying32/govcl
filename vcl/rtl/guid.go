
//----------------------------------------
// 
// Copyright © ying32. All Rights Reserved.
// 
// Licensed under Apache License 2.0
//
//----------------------------------------


package rtl

import (
	"github.com/ying32/govcl/vcl/api"
	"github.com/ying32/govcl/vcl/types"
)

// GUIDToString 将GUID转为字符形式
func GUIDToString(guid types.TGUID) string {
	return guid.ToString()
}

// StringToGUID 将字符形式的GUID转为TGUID结构
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// CreateGUID 创建一个新的GUID
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}
