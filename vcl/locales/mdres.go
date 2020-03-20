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
func ModifyResources(data map[string]string) {
	for i := int32(0); i < api.DGetLibResouceCount(); i++ {
		item := api.DGetLibResouceItem(i)
		if value, ok := data[item.Name]; ok {
			api.DModifyLibResouce(item.Ptr, value)
		}
	}
}
