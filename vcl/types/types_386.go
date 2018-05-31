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
	KeyData              uint32 // 第二个元素未使用
	KeyDataFiller        TDWordFiller
	Result               uintptr
}
