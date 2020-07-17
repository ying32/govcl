//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package types

// 光标定义
const (
	CrHigh = TCursor(0)

	CrDefault   = TCursor(0)
	CrNone      = TCursor(-1)
	CrArrow     = TCursor(-2)
	CrCross     = TCursor(-3)
	CrIBeam     = TCursor(-4)
	CrSize      = TCursor(-22)
	CrSizeNESW  = TCursor(-6) // diagonal north east - south west
	CrSizeNS    = TCursor(-7)
	CrSizeNWSE  = TCursor(-8)
	CrSizeWE    = TCursor(-9)
	CrSizeNW    = TCursor(-23)
	CrSizeN     = TCursor(-24)
	CrSizeNE    = TCursor(-25)
	CrSizeW     = TCursor(-26)
	CrSizeE     = TCursor(-27)
	CrSizeSW    = TCursor(-28)
	CrSizeS     = TCursor(-29)
	CrSizeSE    = TCursor(-30)
	CrUpArrow   = TCursor(-10)
	CrHourGlass = TCursor(-11)
	CrDrag      = TCursor(-12)
	CrNoDrop    = TCursor(-13)
	CrHSplit    = TCursor(-14)
	CrVSplit    = TCursor(-15)
	CrMultiDrag = TCursor(-16)
	CrSQLWait   = TCursor(-17)
	CrNo        = TCursor(-18)
	CrAppStart  = TCursor(-19)
	CrHelp      = TCursor(-20)
	CrHandPoint = TCursor(-21)
	CrSizeAll   = TCursor(-22)
	CrLow       = TCursor(-30)
)
