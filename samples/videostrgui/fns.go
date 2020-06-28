//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"fmt"
	"strings"

	"github.com/ying32/govcl/vcl"
	"github.com/ying32/govcl/vcl/types"
)

func showNotAllowEmpty(s *vcl.TEdit, tips string) bool {
	if strings.TrimSpace(s.Text()) == "" {
		vcl.MessageDlg(fmt.Sprintf("“%s”不能为空！", tips), types.MtInformation, types.MbOK)
		return true
	}
	return false
}
