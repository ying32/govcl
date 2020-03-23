//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

type TDWordFiller struct {
	Filler [4]uint8
}

//  TWMKey
type TWMKey struct {
	Msg       uint32
	MsgFiller TDWordFiller
	CharCode  [2]uint16 // 第二个元素未使用
	// CharCode: Word;
	// Unused: Word;
	CharCodeUnusedFiller TDWordFiller
	KeyData              uint32
	KeyDataFiller        TDWordFiller
	Result               uintptr
}
