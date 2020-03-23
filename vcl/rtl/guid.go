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

// CN: 将TGUID转为字符形式
// EN: Convert TGUID to character form.
func GUIDToString(guid types.TGUID) string {
	return guid.ToString()
}

// CN: 将字符形式的GUID转为TGUID结构
// EN: Convert GUID in character form to TGUID structure.
func StringToGUID(str string) types.TGUID {
	return api.DStringToGUID(str)
}

// CN: 创建一个新的GUID
// EN: Create a new GUID.
func CreateGUID() types.TGUID {
	return api.DCreateGUID()
}
