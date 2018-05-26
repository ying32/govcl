package types

// MessageBox or MessageDlg 返回值

const (
	IdOK       = 1
	IdCancel   = 2
	IdAbort    = 3
	IdRetry    = 4
	IdIgnore   = 5
	IdYes      = 6
	IdNo       = 7
	IdClose    = 8
	IdHelp     = 9
	IdTryAgain = 10
	IdContinue = 11
	MrNone     = 0
	MrOk       = IdOK
	MrCancel   = IdCancel
	MrAbort    = IdAbort
	MrRetry    = IdRetry
	MrIgnore   = IdIgnore
	MrYes      = IdYes
	MrNo       = IdNo
	MrClose    = IdClose
	MrHelp     = IdHelp
	MrTryAgain = IdTryAgain
	MrContinue = IdContinue
	MrAll      = MrContinue + 1
	MrNoToAll  = MrAll + 1
	MrYesToAll = MrNoToAll + 1
)
