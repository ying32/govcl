//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package win

//  MessageBox() Flags
const (
	MB_OK               = 0x00000000
	MB_OKCANCEL         = 0x00000001
	MB_ABORTRETRYIGNORE = 0x00000002
	MB_YESNOCANCEL      = 0x00000003
	MB_YESNO            = 0x00000004
	MB_RETRYCANCEL      = 0x00000005

	MB_ICONHAND        = 0x00000010
	MB_ICONQUESTION    = 0x00000020
	MB_ICONEXCLAMATION = 0x00000030
	MB_ICONASTERISK    = 0x00000040
	MB_USERICON        = 0x00000080
	MB_ICONWARNING     = MB_ICONEXCLAMATION
	MB_ICONERROR       = MB_ICONHAND
	MB_ICONINFORMATION = MB_ICONASTERISK
	MB_ICONSTOP        = MB_ICONHAND
	MB_DEFBUTTON1      = 0x00000000
	MB_DEFBUTTON2      = 0x00000100
	MB_DEFBUTTON3      = 0x00000200
	MB_DEFBUTTON4      = 0x00000300
	MB_APPLMODAL       = 0x00000000
	MB_SYSTEMMODAL     = 0x00001000
	MB_TASKMODAL       = 0x00002000
	MB_HELP            = 0x00004000 // Help Button

	MB_NOFOCUS              = 0x00008000
	MB_SETFOREGROUND        = 0x00010000
	MB_DEFAULT_DESKTOP_ONLY = 0x00020000

	MB_TOPMOST    = 0x00040000
	MB_RIGHT      = 0x00080000
	MB_RTLREADING = 0x00100000

	MB_SERVICE_NOTIFICATION      = 0x00200000
	MB_SERVICE_NOTIFICATION_NT3X = 0x00040000

	MB_TYPEMASK = 0x0000000F
	MB_ICONMASK = 0x000000F0
	MB_DEFMASK  = 0x00000F00
	MB_MODEMASK = 0x00003000
	MB_MISCMASK = 0x0000C000
)

const (
	// Registry.  Reserved Key Handles.

	HKEY_CLASSES_ROOT     = 0x80000000
	HKEY_CURRENT_USER     = 0x80000001
	HKEY_LOCAL_MACHINE    = 0x80000002
	HKEY_USERS            = 0x80000003
	HKEY_PERFORMANCE_DATA = 0x80000004
	HKEY_CURRENT_CONFIG   = 0x80000005
	HKEY_DYN_DATA         = 0x80000006
)

const (

	// The following are masks for the predefined standard access types

	DELETE                  = 0x00010000 // Renamed from DELETE
	READ_CONTROL            = 0x00020000
	WRITE_DAC               = 0x00040000
	WRITE_OWNER             = 0x00080000
	STANDARD_RIGHTS_READ    = READ_CONTROL
	STANDARD_RIGHTS_WRITE   = READ_CONTROL
	STANDARD_RIGHTS_EXECUTE = READ_CONTROL
	STANDARD_RIGHTS_ALL     = 0x001F0000
	SPECIFIC_RIGHTS_ALL     = 0x0000FFFF
	ACCESS_SYSTEM_SECURITY  = 0x01000000
	MAXIMUM_ALLOWED         = 0x02000000
	GENERIC_READ            = 0x80000000
	GENERIC_WRITE           = 0x40000000
	GENERIC_EXECUTE         = 0x20000000
	GENERIC_ALL             = 0x10000000

	// Registry Specific Access Rights.

	KEY_QUERY_VALUE        = 0x0001
	KEY_SET_VALUE          = 0x0002
	KEY_CREATE_SUB_KEY     = 0x0004
	KEY_ENUMERATE_SUB_KEYS = 0x0008
	KEY_NOTIFY             = 0x0010
	KEY_CREATE_LINK        = 0x0020
	KEY_WOW64_32KEY        = 0x0200
	KEY_WOW64_64KEY        = 0x0100
	KEY_WOW64_RES          = 0x0300

	// (STANDARD_RIGHTS_READ | KEY_QUERY_VALUE | KEY_ENUMERATE_SUB_KEYS | KEY_NOTIFY) & ^SYNCHRONIZE
	KEY_READ = 0x00020019
	// (STANDARD_RIGHTS_WRITE | KEY_SET_VALUE | KEY_CREATE_SUB_KEY) & ^SYNCHRONIZE
	KEY_WRITE = 0x00020006
	//  KEY_READ & ^SYNCHRONIZE
	KEY_EXECUTE = 0x00020019

	// (STANDARD_RIGHTS_ALL | KEY_QUERY_VALUE |
	//		KEY_SET_VALUE | KEY_CREATE_SUB_KEY | KEY_ENUMERATE_SUB_KEYS |
	//		KEY_NOTIFY | KEY_CREATE_LINK) & ^SYNCHRONIZE
	KEY_ALL_ACCESS = 0x000F003F
)

const (
	SM_SERVERR2 = 89
)

// SetWindowLongPtr GetWindowLongPtr
const (
	GWL_WNDPROC    = -4
	GWL_HINSTANCE  = -6
	GWL_HWNDPARENT = -8
	GWL_STYLE      = -16
	GWL_EXSTYLE    = -20
	GWL_USERDATA   = -21
	GWL_ID         = -12
)

// Windows Messages
const (
	WM_SYSCOMMAND = 0x0112
)

const (
	// System Menu Command Values
	SC_SIZE         = 61440
	SC_MOVE         = 61456
	SC_MINIMIZE     = 61472
	SC_MAXIMIZE     = 61488
	SC_NEXTWINDOW   = 61504
	SC_PREVWINDOW   = 61520
	SC_CLOSE        = 61536
	SC_VSCROLL      = 61552
	SC_HSCROLL      = 61568
	SC_MOUSEMENU    = 61584
	SC_KEYMENU      = 61696
	SC_ARRANGE      = 61712
	SC_RESTORE      = 61728
	SC_TASKLIST     = 61744
	SC_SCREENSAVE   = 61760
	SC_HOTKEY       = 61776
	SC_DEFAULT      = 61792
	SC_MONITORPOWER = 61808
	SC_CONTEXTHELP  = 61824
	SC_SEPARATOR    = 61455

	SCF_ISSECURE = 0x00000001

	// Obsolete names
	SC_ICON = SC_MINIMIZE
	SC_ZOOM = SC_MAXIMIZE
)

const MAX_PATH = 260

const (
	// Scroll Bar Constants
	SB_HORZ = 0
	SB_VERT = 1
	SB_CTL  = 2
	SB_BOTH = 3

	// Scroll Bar Commands
	SB_LINEUP        = 0
	SB_LINELEFT      = 0
	SB_LINEDOWN      = 1
	SB_LINERIGHT     = 1
	SB_PAGEUP        = 2
	SB_PAGELEFT      = 2
	SB_PAGEDOWN      = 3
	SB_PAGERIGHT     = 3
	SB_THUMBPOSITION = 4
	SB_THUMBTRACK    = 5
	SB_TOP           = 6
	SB_LEFT          = 6
	SB_BOTTOM        = 7
	SB_RIGHT         = 7
	SB_ENDSCROLL     = 8

	// ShowWindow() Commands
	SW_HIDE            = 0
	SW_SHOWNORMAL      = 1
	SW_NORMAL          = 1
	SW_SHOWMINIMIZED   = 2
	SW_SHOWMAXIMIZED   = 3
	SW_MAXIMIZE        = 3
	SW_SHOWNOACTIVATE  = 4
	SW_SHOW            = 5
	SW_MINIMIZE        = 6
	SW_SHOWMINNOACTIVE = 7
	SW_SHOWNA          = 8
	SW_RESTORE         = 9
	SW_SHOWDEFAULT     = 10
	SW_FORCEMINIMIZE   = 11
	SW_MAX             = 11

	// Old ShowWindow() Commands
	HIDE_WINDOW         = 0
	SHOW_OPENWINDOW     = 1
	SHOW_ICONWINDOW     = 2
	SHOW_FULLSCREEN     = 3
	SHOW_OPENNOACTIVATE = 4

	// Identifiers for the WM_SHOWWINDOW message
	SW_PARENTCLOSING = 1
	SW_OTHERZOOM     = 2
	SW_PARENTOPENING = 3
	SW_OTHERUNZOOM   = 4
)

const (
	// Code Page Default Values.
	CP_ACP        = 0     // default to ANSI code page
	CP_OEMCP      = 1     // default to OEM  code page
	CP_MACCP      = 2     // default to MAC  code page
	CP_THREAD_ACP = 3     // current thread's ANSI code page
	CP_SYMBOL     = 42    // SYMBOL translations
	CP_UTF7       = 65000 // UTF-7 translation
	CP_UTF8       = 65001 // UTF-8 translation
)

const (
	// Window Styles
	WS_OVERLAPPED   = 0
	WS_POPUP        = 0x80000000
	WS_CHILD        = 0x40000000
	WS_MINIMIZE     = 0x20000000
	WS_VISIBLE      = 0x10000000
	WS_DISABLED     = 0x8000000
	WS_CLIPSIBLINGS = 0x4000000
	WS_CLIPCHILDREN = 0x2000000
	WS_MAXIMIZE     = 0x1000000
	WS_CAPTION      = 0xC00000 // WS_BORDER or WS_DLGFRAME
	WS_BORDER       = 0x800000
	WS_DLGFRAME     = 0x400000
	WS_VSCROLL      = 0x200000
	WS_HSCROLL      = 0x100000
	WS_SYSMENU      = 0x80000
	WS_THICKFRAME   = 0x40000
	WS_GROUP        = 0x20000
	WS_TABSTOP      = 0x10000

	WS_MINIMIZEBOX = 0x20000
	WS_MAXIMIZEBOX = 0x10000

	WS_TILED   = WS_OVERLAPPED
	WS_ICONIC  = WS_MINIMIZE
	WS_SIZEBOX = WS_THICKFRAME

	// Common Window Styles
	WS_OVERLAPPEDWINDOW = WS_OVERLAPPED | WS_CAPTION | WS_SYSMENU | WS_THICKFRAME | WS_MINIMIZEBOX | WS_MAXIMIZEBOX
	WS_TILEDWINDOW      = WS_OVERLAPPEDWINDOW
	WS_POPUPWINDOW      = WS_POPUP | WS_BORDER | WS_SYSMENU
	WS_CHILDWINDOW      = WS_CHILD

	// Extended Window Styles
	WS_EX_DLGMODALFRAME  = 1
	WS_EX_NOPARENTNOTIFY = 4
	WS_EX_TOPMOST        = 8
	WS_EX_ACCEPTFILES    = 0x10
	WS_EX_TRANSPARENT    = 0x20
	WS_EX_MDICHILD       = 0x40
	WS_EX_TOOLWINDOW     = 0x80
	WS_EX_WINDOWEDGE     = 0x100
	WS_EX_CLIENTEDGE     = 0x200
	WS_EX_CONTEXTHELP    = 0x400

	WS_EX_RIGHT          = 0x1000
	WS_EX_LEFT           = 0
	WS_EX_RTLREADING     = 0x2000
	WS_EX_LTRREADING     = 0
	WS_EX_LEFTSCROLLBAR  = 0x4000
	WS_EX_RIGHTSCROLLBAR = 0

	WS_EX_CONTROLPARENT    = 0x10000
	WS_EX_STATICEDGE       = 0x20000
	WS_EX_APPWINDOW        = 0x40000
	WS_EX_OVERLAPPEDWINDOW = WS_EX_WINDOWEDGE | WS_EX_CLIENTEDGE
	WS_EX_PALETTEWINDOW    = WS_EX_WINDOWEDGE | WS_EX_TOOLWINDOW | WS_EX_TOPMOST

	WS_EX_LAYERED = 0x00080000

	WS_EX_NOINHERITLAYOUT = 0x00100000 // Disable inheritence of mirroring by children
	WS_EX_LAYOUTRTL       = 0x00400000 // Right to left mirroring
	WS_EX_COMPOSITED      = 0x02000000
	WS_EX_NOACTIVATE      = 0x08000000
)

const (
	THREAD_BASE_PRIORITY_LOWRT = 15  // value that gets a thread to LowRealtime-1
	THREAD_BASE_PRIORITY_MAX   = 2   // maximum thread base priority boost
	THREAD_BASE_PRIORITY_MIN   = -2  // minimum thread base priority boost
	THREAD_BASE_PRIORITY_IDLE  = -15 // value that gets a thread to idle

	SYNCHRONIZE              = 0x00100000
	STANDARD_RIGHTS_REQUIRED = 0x000F0000
	EVENT_MODIFY_STATE       = 0x0002
	EVENT_ALL_ACCESS         = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3
	MUTANT_QUERY_STATE       = 0x0001
	MUTANT_ALL_ACCESS        = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | MUTANT_QUERY_STATE

	SEMAPHORE_MODIFY_STATE = 0x0002
	SEMAPHORE_ALL_ACCESS   = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0x3

	PROCESS_TERMINATE         = 0x0001
	PROCESS_CREATE_THREAD     = 0x0002
	PROCESS_VM_OPERATION      = 0x0008
	PROCESS_VM_READ           = 0x0010
	PROCESS_VM_WRITE          = 0x0020
	PROCESS_DUP_HANDLE        = 0x0040
	PROCESS_CREATE_PROCESS    = 0x0080
	PROCESS_SET_QUOTA         = 0x0100
	PROCESS_SET_INFORMATION   = 0x0200
	PROCESS_QUERY_INFORMATION = 0x0400
	// if NTDDI_VERSION >= NTDDI_VISTA
	PROCESS_ALL_ACCESS = STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFFF
	// else
	//PROCESS_ALL_ACCESS        = (STANDARD_RIGHTS_REQUIRED | SYNCHRONIZE | 0xFFF);
	// endif

	PROCESSOR_INTEL_386     = 386
	PROCESSOR_INTEL_486     = 486
	PROCESSOR_INTEL_PENTIUM = 586
	PROCESSOR_INTEL_IA64    = 2200
	PROCESSOR_AMD_X8664     = 8664
	PROCESSOR_MIPS_R4000    = 4000 // incl R4101 & R3910 for Windows CE
	PROCESSOR_ALPHA_21064   = 21064
	PROCESSOR_PPC_601       = 601
	PROCESSOR_PPC_603       = 603
	PROCESSOR_PPC_604       = 604
	PROCESSOR_PPC_620       = 620
	PROCESSOR_HITACHI_SH3   = 10003  // Windows CE
	PROCESSOR_HITACHI_SH3E  = 10004  // Windows CE
	PROCESSOR_HITACHI_SH4   = 10005  // Windows CE
	PROCESSOR_MOTOROLA_821  = 821    // Windows CE
	PROCESSOR_SHx_SH3       = 103    // Windows CE
	PROCESSOR_SHx_SH4       = 104    // Windows CE
	PROCESSOR_STRONGARM     = 2577   // Windows CE - 0xA11
	PROCESSOR_ARM720        = 1824   // Windows CE - 0x720
	PROCESSOR_ARM820        = 2080   // Windows CE - 0x820
	PROCESSOR_ARM920        = 2336   // Windows CE - 0x920
	PROCESSOR_ARM_7TDMI     = 70001  // Windows CE
	PROCESSOR_OPTIL         = 0x494F // MSIL

	PROCESSOR_ARCHITECTURE_INTEL         = 0
	PROCESSOR_ARCHITECTURE_MIPS          = 1
	PROCESSOR_ARCHITECTURE_ALPHA         = 2
	PROCESSOR_ARCHITECTURE_PPC           = 3
	PROCESSOR_ARCHITECTURE_SHX           = 4
	PROCESSOR_ARCHITECTURE_ARM           = 5
	PROCESSOR_ARCHITECTURE_IA64          = 6
	PROCESSOR_ARCHITECTURE_ALPHA64       = 7
	PROCESSOR_ARCHITECTURE_MSIL          = 8
	PROCESSOR_ARCHITECTURE_AMD64         = 9
	PROCESSOR_ARCHITECTURE_IA32_ON_WIN64 = 10

	PROCESSOR_ARCHITECTURE_UNKNOWN = 0xFFFF
)

const (
	TOKEN_ASSIGN_PRIMARY    = 0x0001
	TOKEN_DUPLICATE         = 0x0002
	TOKEN_IMPERSONATE       = 0x0004
	TOKEN_QUERY             = 0x0008
	TOKEN_QUERY_SOURCE      = 0x0010
	TOKEN_ADJUST_PRIVILEGES = 0x0020
	TOKEN_ADJUST_GROUPS     = 0x0040
	TOKEN_ADJUST_DEFAULT    = 0x0080
	TOKEN_ADJUST_SESSIONID  = 0x0100

	TOKEN_ALL_ACCESS_P = STANDARD_RIGHTS_REQUIRED | TOKEN_ASSIGN_PRIMARY |
		TOKEN_DUPLICATE | TOKEN_IMPERSONATE | TOKEN_QUERY |
		TOKEN_QUERY_SOURCE | TOKEN_ADJUST_PRIVILEGES | TOKEN_ADJUST_GROUPS |
		TOKEN_ADJUST_DEFAULT

	// if _WIN32_WINNT > 0x0400 || !defined(_WIN32_WINNT)
	TOKEN_ALL_ACCESS = TOKEN_ALL_ACCESS_P | TOKEN_ADJUST_SESSIONID

	TOKEN_READ    = STANDARD_RIGHTS_READ | TOKEN_QUERY
	TOKEN_WRITE   = STANDARD_RIGHTS_WRITE | TOKEN_ADJUST_PRIVILEGES | TOKEN_ADJUST_GROUPS | TOKEN_ADJUST_DEFAULT
	TOKEN_EXECUTE = STANDARD_RIGHTS_EXECUTE
)

const (
	// constants for the biCompression field
	BI_RGB       = 0
	BI_RLE8      = 1
	BI_RLE4      = 2
	BI_BITFIELDS = 3

	// currentlly defined blend function
	AC_SRC_OVER  = 0x00
	AC_SRC_ALPHA = 0x01

	// alpha format flags
	AC_SRC_NO_PREMULT_ALPHA = 0x01
	AC_SRC_NO_ALPHA         = 0x02
	AC_DST_NO_PREMULT_ALPHA = 0x10
	AC_DST_NO_ALPHA         = 0x20

	LWA_COLORKEY = 0x00000001
	LWA_ALPHA    = 0x00000002
	ULW_COLORKEY = 0x00000001
	ULW_ALPHA    = 0x00000002
	ULW_OPAQUE   = 0x00000004
)

const (
	//
	//  File System Notification flags
	//
	SHCNRF_InterruptLevel     = 0x00000001
	SHCNRF_ShellLevel         = 0x00000002
	SHCNRF_RecursiveInterrupt = 0x00001000
	SHCNRF_NewDelivery        = 0x00008000

	SHCNE_RENAMEITEM       = 0x00000001
	SHCNE_CREATE           = 0x00000002
	SHCNE_DELETE           = 0x00000004
	SHCNE_MKDIR            = 0x00000008
	SHCNE_RMDIR            = 0x00000010
	SHCNE_MEDIAINSERTED    = 0x00000020
	SHCNE_MEDIAREMOVED     = 0x00000040
	SHCNE_DRIVEREMOVED     = 0x00000080
	SHCNE_DRIVEADD         = 0x00000100
	SHCNE_NETSHARE         = 0x00000200
	SHCNE_NETUNSHARE       = 0x00000400
	SHCNE_ATTRIBUTES       = 0x00000800
	SHCNE_UPDATEDIR        = 0x00001000
	SHCNE_UPDATEITEM       = 0x00002000
	SHCNE_SERVERDISCONNECT = 0x00004000
	SHCNE_UPDATEIMAGE      = 0x00008000
	SHCNE_DRIVEADDGUI      = 0x00010000
	SHCNE_RENAMEFOLDER     = 0x00020000
	SHCNE_FREESPACE        = 0x00040000

	// SHCNE_EXTENDED_EVENT: the extended event is identified in dwItem1,
	// packed in LPITEMIDLIST format (same as SHCNF_DWORD packing).
	// Additional information can be passed in the dwItem2 parameter
	// of SHChangeNotify (called "pidl2" below), which if present, must also
	// be in LPITEMIDLIST format.
	//
	// Unlike the standard events, the extended events are ORDINALs, so we
	// don't run out of bits.  Extended events follow the SHCNEE_* naming
	// convention.
	//
	// The dwItem2 parameter varies according to the extended event.

	SHCNE_EXTENDED_EVENT         = 0x04000000
	SHCNE_EXTENDED_EVENT_PRE_IE4 = 0x00080000

	SHCNE_ASSOCCHANGED = 0x08000000

	SHCNE_DISKEVENTS   = 0x0002381F
	SHCNE_GLOBALEVENTS = 0x0C0581E0 // Events that dont match pidls first
	SHCNE_ALLEVENTS    = 0x7FFFFFFF
	SHCNE_INTERRUPT    = 0x80000000 // The presence of this flag indicates
	// that the event was generated by an
	// interrupt.  It is stripped out before
	// the clients of SHCNNotify_ see it.

	// SHCNE_EXTENDED_EVENT extended events.  These events are ordinals.
	// This is not a bitfield.

	SHCNEE_ORDERCHANGED  = 0x00000002 // pidl2 is the changed folder
	SHCNEE_MSI_CHANGE    = 0x00000004 // pidl2 is a SHChangeProductKeyAsIDList
	SHCNEE_MSI_UNINSTALL = 0x00000005 // pidl2 is a SHChangeProductKeyAsIDList

	// Flags
	// uFlags & SHCNF_TYPE is an ID which indicates what dwItem1 and dwItem2 mean
	SHCNF_IDLIST      = 0x0000 // LPITEMIDLIST
	SHCNF_PATHA       = 0x0001 // path name
	SHCNF_PRINTERA    = 0x0002 // printer friendly name
	SHCNF_DWORD       = 0x0003 // DWORD
	SHCNF_PATHW       = 0x0005 // path name
	SHCNF_PRINTERW    = 0x0006 // printer friendly name
	SHCNF_TYPE        = 0x00FF
	SHCNF_FLUSH       = 0x1000
	SHCNF_FLUSHNOWAIT = 0x3000 // includes SHCNF_FLUSH

	SHCNF_NOTIFYRECURSIVE = 0x10000 // Notify clients registered for any child

)

const (
	// SetWindowPos Flags

	SWP_NOSIZE         = 1
	SWP_NOMOVE         = 2
	SWP_NOZORDER       = 4
	SWP_NOREDRAW       = 8
	SWP_NOACTIVATE     = 0x10
	SWP_FRAMECHANGED   = 0x20 // The frame changed: send WM_NCCALCSIZE
	SWP_SHOWWINDOW     = 0x40
	SWP_HIDEWINDOW     = 0x80
	SWP_NOCOPYBITS     = 0x100
	SWP_NOOWNERZORDER  = 0x200 // Don't do owner Z ordering
	SWP_NOSENDCHANGING = 0x400 // Don't send WM_WINDOWPOSCHANGING
	SWP_DRAWFRAME      = SWP_FRAMECHANGED
	SWP_NOREPOSITION   = SWP_NOOWNERZORDER
	SWP_DEFERERASE     = 0x2000
	SWP_ASYNCWINDOWPOS = 0x4000

	HWND_TOP       = 0
	HWND_BOTTOM    = 1
	HWND_TOPMOST   = ^(-1) // -1
	HWND_NOTOPMOST = ^(-2) // -2
)

/* Translated from WINGDI.H */

const (
	/* Binary raster ops */
	R2_BLACK       = 1    /*  0   */
	R2_NOTMERGEPEN = 2    /* DPon */
	R2_MASKNOTPEN  = 3    /* DPna */
	R2_NOTCOPYPEN  = 4    /* PN   */
	R2_MASKPENNOT  = 5    /* PDna */
	R2_NOT         = 6    /* Dn   */
	R2_XORPEN      = 7    /* DPx  */
	R2_NOTMASKPEN  = 8    /* DPan */
	R2_MASKPEN     = 9    /* DPa  */
	R2_NOTXORPEN   = 10   /* DPxn */
	R2_NOP         = 11   /* D    */
	R2_MERGENOTPEN = 12   /* DPno */
	R2_COPYPEN     = 13   /* P    */
	R2_MERGEPENNOT = 14   /* PDno */
	R2_MERGEPEN    = 15   /* DPo  */
	R2_WHITE       = 0x10 /*  1   */
	R2_LAST        = 0x10

	/* Ternary raster operations */
	SRCCOPY     = 0x00CC0020 /* dest = source                    */
	SRCPAINT    = 0x00EE0086 /* dest = source OR dest            */
	SRCAND      = 0x008800C6 /* dest = source AND dest           */
	SRCINVERT   = 0x00660046 /* dest = source XOR dest           */
	SRCERASE    = 0x00440328 /* dest = source AND (NOT dest )    */
	NOTSRCCOPY  = 0x00330008 /* dest = (NOT source)              */
	NOTSRCERASE = 0x001100A6 /* dest = (NOT src) AND (NOT dest)  */
	MERGECOPY   = 0x00C000CA /* dest = (source AND pattern)      */
	MERGEPAINT  = 0x00BB0226 /* dest = (NOT source) OR dest      */
	PATCOPY     = 0x00F00021 /* dest = pattern                   */
	PATPAINT    = 0x00FB0A09 /* dest = DPSnoo                    */
	PATINVERT   = 0x005A0049 /* dest = pattern XOR dest          */
	DSTINVERT   = 0x00550009 /* dest = (NOT dest)                */
	BLACKNESS   = 0x00000042 /* dest = BLACK                     */
	WHITENESS   = 0x00FF0062 /* dest = WHITE */

)

const (
	// SHGetSpecialFolderLocation
	//
	//  Caller should use SHGetMalloc to obtain an allocator that can free the pidl
	//
	// registry entries for special paths are kept in :
	// REGSTR_PATH_SPECIAL_FOLDERS   = REGSTR_PATH_EXPLORER + '\Shell Folders' //注释掉先

	CSIDL_DESKTOP                 = 0x0000         // <desktop>
	CSIDL_INTERNET                = 0x0001         // Internet Explorer (icon on desktop)
	CSIDL_PROGRAMS                = 0x0002         // Start Menu\Programs
	CSIDL_CONTROLS                = 0x0003         // My Computer\Control Panel
	CSIDL_PRINTERS                = 0x0004         // My Computer\Printers
	CSIDL_PERSONAL                = 0x0005         // My Documents
	CSIDL_FAVORITES               = 0x0006         // <user name>\Favorites
	CSIDL_STARTUP                 = 0x0007         // Start Menu\Programs\Startup
	CSIDL_RECENT                  = 0x0008         // <user name>\Recent
	CSIDL_SENDTO                  = 0x0009         // <user name>\SendTo
	CSIDL_BITBUCKET               = 0x000a         // <desktop>\Recycle Bin
	CSIDL_STARTMENU               = 0x000b         // <user name>\Start Menu
	CSIDL_MYDOCUMENTS             = CSIDL_PERSONAL // Personal was just a silly name for My Documents
	CSIDL_MYMUSIC                 = 0x000d         // "My Music" folder
	CSIDL_MYVIDEO                 = 0x000e         // "My Videos" folder
	CSIDL_DESKTOPDIRECTORY        = 0x0010         // <user name>\Desktop
	CSIDL_DRIVES                  = 0x0011         // My Computer
	CSIDL_NETWORK                 = 0x0012         // Network Neighborhood (My Network Places)
	CSIDL_NETHOOD                 = 0x0013         // <user name>\nethood
	CSIDL_FONTS                   = 0x0014         // windows\fonts
	CSIDL_TEMPLATES               = 0x0015
	CSIDL_COMMON_STARTMENU        = 0x0016 // All Users\Start Menu
	CSIDL_COMMON_PROGRAMS         = 0x0017 // All Users\Start Menu\Programs
	CSIDL_COMMON_STARTUP          = 0x0018 // All Users\Startup
	CSIDL_COMMON_DESKTOPDIRECTORY = 0x0019 // All Users\Desktop
	CSIDL_APPDATA                 = 0x001a // <user name>\Application Data
	CSIDL_PRINTHOOD               = 0x001b // <user name>\PrintHood
	CSIDL_LOCAL_APPDATA           = 0x001c // <user name>\Local Settings\Applicaiton Data (non roaming)
	CSIDL_ALTSTARTUP              = 0x001d // non localized startup
	CSIDL_COMMON_ALTSTARTUP       = 0x001e // non localized common startup
	CSIDL_COMMON_FAVORITES        = 0x001f
	CSIDL_INTERNET_CACHE          = 0x0020
	CSIDL_COOKIES                 = 0x0021
	CSIDL_HISTORY                 = 0x0022
	CSIDL_COMMON_APPDATA          = 0x0023 // All Users\Application Data
	CSIDL_WINDOWS                 = 0x0024 // GetWindowsDirectory()
	CSIDL_SYSTEM                  = 0x0025 // GetSystemDirectory()
	CSIDL_PROGRAM_FILES           = 0x0026 // C:\Program Files
	CSIDL_MYPICTURES              = 0x0027 // C:\Program Files\My Pictures
	CSIDL_PROFILE                 = 0x0028 // USERPROFILE
	CSIDL_SYSTEMX86               = 0x0029 // x86 system directory on RISC
	CSIDL_PROGRAM_FILESX86        = 0x002a // x86 C:\Program Files on RISC
	CSIDL_PROGRAM_FILES_COMMON    = 0x002b // C:\Program Files\Common
	CSIDL_PROGRAM_FILES_COMMONX86 = 0x002c // x86 Program Files\Common on RISC
	CSIDL_COMMON_TEMPLATES        = 0x002d // All Users\Templates
	CSIDL_COMMON_DOCUMENTS        = 0x002e // All Users\Documents
	CSIDL_COMMON_ADMINTOOLS       = 0x002f // All Users\Start Menu\Programs\Administrative Tools
	CSIDL_ADMINTOOLS              = 0x0030 // <user name>\Start Menu\Programs\Administrative Tools
	CSIDL_CONNECTIONS             = 0x0031 // Network and Dial-up Connections
	CSIDL_COMMON_MUSIC            = 0x0035 // All Users\My Music
	CSIDL_COMMON_PICTURES         = 0x0036 // All Users\My Pictures
	CSIDL_COMMON_VIDEO            = 0x0037 // All Users\My Video
	CSIDL_RESOURCES               = 0x0038 // Resource Direcotry
	CSIDL_RESOURCES_LOCALIZED     = 0x0039 // Localized Resource Direcotry
	CSIDL_COMMON_OEM_LINKS        = 0x003a // Links to All Users OEM specific apps
	CSIDL_CDBURN_AREA             = 0x003b // USERPROFILE\Local Settings\Application Data\Microsoft\CD Burning
	// unused                               0x003c
	CSIDL_COMPUTERSNEARME    = 0x003d // Computers Near Me (computered from Workgroup membership)
	CSIDL_FLAG_CREATE        = 0x8000 // combine with CSIDL_ value to force folder creation in SHGetFolderPath()
	CSIDL_FLAG_DONT_VERIFY   = 0x4000 // combine with CSIDL_ value to return an unverified folder path
	CSIDL_FLAG_DONT_UNEXPAND = 0x2000 // combine with CSIDL_ value to avoid unexpanding environment variables
	CSIDL_FLAG_NO_ALIAS      = 0x1000 // combine with CSIDL_ value to insure non-alias versions of the pidl
	CSIDL_FLAG_PER_USER_INIT = 0x0800 // combine with CSIDL_ value to indicate per-user init (eg. upgrade)
	CSIDL_FLAG_MASK          = 0xFF00 // mask for all possible flag values
)

const (
	INVALID_HANDLE_VALUE = ^(-1)
)

/* Init/Uninit */

const (
	// flags passed as the coInit parameter to CoInitializeEx.
	COINIT_MULTITHREADED     = 0 // OLE calls objects on any thread.
	COINIT_APARTMENTTHREADED = 2 // Apartment model
	COINIT_DISABLE_OLE1DDE   = 4 // Dont use DDE for Ole1 support.
	COINIT_SPEED_OVER_MEMORY = 8 // Trade memory for speed
)

const (

	/* Predefined Resource Types */
	RT_CURSOR       = 1
	RT_BITMAP       = 2
	RT_ICON         = 3
	RT_MENU         = 4
	RT_DIALOG       = 5
	RT_STRING       = 6
	RT_FONTDIR      = 7
	RT_FONT         = 8
	RT_ACCELERATOR  = 9
	RT_RCDATA       = 10
	RT_MESSAGETABLE = 11

	DIFFERENCE = 11

	RT_GROUP_CURSOR = RT_CURSOR + DIFFERENCE
	RT_GROUP_ICON   = RT_ICON + DIFFERENCE
	RT_VERSION      = 16
	RT_DLGINCLUDE   = 17
	RT_PLUGPLAY     = 19
	RT_VXD          = 20
	RT_ANICURSOR    = 21
	RT_ANIICON      = 22
	RT_HTML         = 23
	RT_MANIFEST     = 24
)

/* ====== Ranges for control message IDs ======================= */

const (
	HDM_FIRST = 0x1200 /* Header messages */
	TCM_FIRST = 0x1300 /* Tab control messages */
	PGM_FIRST = 0x1400 /* Pager control messages */
	/* For Windows >= XP */
	ECM_FIRST = 0x1500 /* Edit control messages */
	BCM_FIRST = 0x1600 /* Button control messages */
	CBM_FIRST = 0x1700 /* Combobox control messages */

	CCM_FIRST = 0x2000 /* Common control shared messages */
	CCM_LAST  = CCM_FIRST + 0x200

	CCM_SETBKCOLOR = CCM_FIRST + 1 // lParam is bkColo
)

const (
	/* Background Modes */
	TRANSPARENT = 1
	OPAQUE      = 2
	BKMODE_LAST = 2

	/* Graphics Modes */
	GM_COMPATIBLE = 1
	GM_ADVANCED   = 2
	GM_LAST       = 2
)

const (
	MOD_ALT      = 1
	MOD_CONTROL  = 2
	MOD_SHIFT    = 4
	MOD_WIN      = 8
	MOD_NOREPEAT = 0x4000
)

// Begin ShellExecuteEx and family
/* ShellExecute() and ShellExecuteEx() error codes */
const (
	/* regular WinExec() codes */
	SE_ERR_FNF          = 2 /* file not found */
	SE_ERR_PNF          = 3 /* path not found */
	SE_ERR_ACCESSDENIED = 5 /* access denied */
	SE_ERR_OOM          = 8 /* out of memory */
	SE_ERR_DLLNOTFOUND  = 32

	/* error values for ShellExecute() beyond the regular WinExec() codes */
	SE_ERR_SHARE           = 26
	SE_ERR_ASSOCINCOMPLETE = 27
	SE_ERR_DDETIMEOUT      = 28
	SE_ERR_DDEFAIL         = 29
	SE_ERR_DDEBUSY         = 30
	SE_ERR_NOASSOC         = 31

	/* Note CLASSKEY overrides CLASSNAME */
	SEE_MASK_DEFAULT   = 0x00000000
	SEE_MASK_CLASSNAME = 0x00000001
	SEE_MASK_CLASSKEY  = 0x00000003
	/* Note INVOKEIDLIST overrides IDLIST */
	SEE_MASK_IDLIST            = 0x00000004
	SEE_MASK_INVOKEIDLIST      = 0x0000000c
	SEE_MASK_ICON              = 0x00000010
	SEE_MASK_HOTKEY            = 0x00000020
	SEE_MASK_NOCLOSEPROCESS    = 0x00000040
	SEE_MASK_CONNECTNETDRV     = 0x00000080
	SEE_MASK_NOASYNC           = 0x00000100
	SEE_MASK_FLAG_DDEWAIT      = SEE_MASK_NOASYNC
	SEE_MASK_DOENVSUBST        = 0x00000200
	SEE_MASK_FLAG_NO_UI        = 0x00000400
	SEE_MASK_UNICODE           = 0x00004000
	SEE_MASK_NO_CONSOLE        = 0x00008000
	SEE_MASK_ASYNCOK           = 0x00100000
	SEE_MASK_HMONITOR          = 0x00200000 // SHELLEXECUTEINFO.hMonitor
	SEE_MASK_NOZONECHECKS      = 0x00800000
	SEE_MASK_NOQUERYCLASSSTORE = 0x01000000
	SEE_MASK_WAITFORINPUTIDLE  = 0x02000000
	SEE_MASK_FLAG_LOG_USAGE    = 0x04000000
)
