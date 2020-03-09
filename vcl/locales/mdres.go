//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package locales

import "github.com/ying32/govcl/vcl/api"

// 修改资源
func ModifyResouces(data map[string]string) {
	for i := 0; i < int(api.DGetLibResouceCount()); i++ {
		item := api.DGetLibResouceItem(int32(i))
		if value, ok := data[item.Name]; ok {
			api.DModifyLibResouce(item.Ptr, value)
		}
	}
}
