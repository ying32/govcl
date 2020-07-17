//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build arm 386

package types

type TDWordFiller struct {
}

//  TWMKey
type TWMKey struct {
	Msg       uint32
	MsgFiller TDWordFiller
	CharCode  [2]uint16
	// CharCode: Word;
	// Unused: Word;
	CharCodeUnusedFiller TDWordFiller
	KeyData              uint32
	KeyDataFiller        TDWordFiller
	Result               uintptr
}
