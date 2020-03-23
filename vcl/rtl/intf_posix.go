//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

// +build !windows

package rtl

import (
	"github.com/ying32/govcl/vcl/api"
	. "github.com/ying32/govcl/vcl/types"
)

func MainInstance() uintptr {
	return 0
}

func SendMessage(hWd HWND, msg uint32, wParam, lParam uintptr) uintptr {
	return api.DSendMessage(hWd, msg, wParam, lParam)
}

func PostMessage(hWd HWND, msg uint32, wParam, lParam uintptr) bool {
	return api.DPostMessage(hWd, msg, wParam, lParam)
}

func IsIconic(hWnd HWND) bool {
	return api.DIsIconic(hWnd)
}

func IsWindow(hWnd HWND) bool {
	return api.DIsWindow(hWnd)
}

func IsZoomed(hWnd HWND) bool {
	return api.DIsZoomed(hWnd)
}

func IsWindowVisible(hWnd HWND) bool {
	return api.DIsWindowVisible(hWnd)
}

func GetDC(hWnd HWND) HDC {
	return api.DGetDC(hWnd)
}

func ReleaseDC(hWnd HWND, dc HDC) int {
	return api.DReleaseDC(hWnd, dc)
}

func SetForegroundWindow(hWnd HWND) bool {
	return api.DSetForegroundWindow(hWnd)
}

func WindowFromPoint(point TPoint) HWND {
	return api.DWindowFromPoint(point)
}

/*
// lcl接口
function IsIconic(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindow(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindowEnabled(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindowVisible(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsZoomed(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function ClientToScreen(Handle : HWND; var P : TPoint) : Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function DeleteDC(hDC: HDC): Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function DeleteObject(GDIObject: HGDIOBJ): Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function EnableWindow(hWnd: HWND; bEnable: Boolean): Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetActiveWindow : HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetBitmapBits(Bitmap: HBITMAP; Count: Longint;  Bits: Pointer): Longint; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetBkColor(DC: HDC): TColorRef; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetClientRect(handle : HWND; var Rect: TRect) : Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetDC(hWnd: HWND): HDC; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetDeviceCaps(DC: HDC; Index: Integer): Integer; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}

function GetFocus: HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetKeyState(nVirtKey: Integer): Smallint; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetForegroundWindow: HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetParent(Handle : HWND): HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function GetSystemMetrics(nIndex: Integer): Integer;  {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function InvalidateRect(aHandle : HWND; ARect : pRect; bErase : Boolean) : Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsIconic(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindow(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindowEnabled(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsWindowVisible(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function IsZoomed(handle: HWND): boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function ReleaseCapture : Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function ReleaseDC(hWnd: HWND; DC: HDC): Integer; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function SaveDC(DC: HDC): Integer; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function ScreenToClient(Handle : HWND; var P : TPoint) : Integer; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function SelectObject(DC: HDC; GDIObj: HGDIOBJ): HGDIOBJ; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function SetForegroundWindow(hWnd : HWND): Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function SetParent(hWndChild: HWND; hWndParent: HWND): HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function WindowFromPoint(Point : TPoint) : HWND; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}
function UpdateWindow(Handle: HWND): Boolean; {$IFDEF IF_BASE_MEMBER}virtual;{$ENDIF}

*/
