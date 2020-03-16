//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

type NSWindowTitleVisibility uint

const (
	NSWindowTitleVisible NSWindowTitleVisibility = iota + 0
	NSWindowTitleHidden
)

const (
	// 先只提供一个吧
	NSWindowStyleMaskFullSizeContentView = 1 << 15
)
