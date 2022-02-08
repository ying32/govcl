//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

//go:build !windows && (amd64 || arm64)
// +build !windows
// +build amd64 arm64

package types

// 消息值参见 types/messages包
type TMessage struct {
	Msg        Cardinal
	_UnusedMsg Cardinal
	WParam     WPARAM
	LParam     LPARAM
	Result     LRESULT
}
