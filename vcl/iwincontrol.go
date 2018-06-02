package vcl

type IWinControl interface {
	IControl
	Parent() *TWinControl
	SetParent(IWinControl)
}
