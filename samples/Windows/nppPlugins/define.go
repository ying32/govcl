package main

/*
   #include <windows.h>
*/
import "C"

type HWND = uintptr

type LRESULT = uintptr

type UINT = uint32

type WPARAM = uintptr

type LPARAM = uintptr

type PTCHAR = uintptr

type NppData struct {
	NppHandle             HWND
	ScintillaMainHandle   HWND
	ScintillaSecondHandle HWND
}

// PFuncSetInfo
//type PFUNCSETINFO   func() procedure(NppData: TNppData);cdecl;
type PFUNCSETINFO = uintptr

// PFuncPluginCmd
//PFUNCPLUGINCMD = procedure;cdecl;
type PFUNCPLUGINCMD = uintptr

// PBeNotified
//PBENOTIFIED    = procedure(var SCNotification: TSCNotification);cdecl;
type PBENOTIFIED = uintptr

// PMessageProc
//PMESSAGEPROC   = function(iMessage: UINT; wParam: WPARAM; lParam: LPARAM):LRESULT;cdecl;
type PMESSAGEPROC = uintptr

type ShortcutKey struct {
	IsCtrl  bool
	IsAlt   bool
	IsShift bool
	Key     byte
}

const nbChar = 64

type FuncItem struct {
	ItemName   [nbChar]uint16
	PFunc      PFUNCPLUGINCMD
	CmdID      int32
	Init2Check bool
	PShKey     *ShortcutKey
}
