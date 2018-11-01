// +build windows linux,386 darwin,386

package types

type TMessage struct {
	Msg    Cardinal
	WParam WAPRAM
	LParam LPARAM
	Result LRESULT
}
