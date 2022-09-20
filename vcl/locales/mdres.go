//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package locales

import "github.com/ying32/govcl/vcl/api"

// ModifyResources
//  修改资源
//  Modify resources.
func ModifyResources(data map[string]string) {
	for i := int32(0); i < api.DGetLibResourceCount(); i++ {
		item := api.DGetLibResourceItem(i)
		if value, ok := data[item.Name]; ok {
			api.DModifyLibResource(item.Ptr, value)
		}
	}
}
