//----------------------------------------
//
// Copyright © ying32. All Rights Reserved.
//
// Licensed under Apache License 2.0
//
//----------------------------------------

package main

import (
	"bytes"
	"fmt"
	"io/ioutil"
)

type CFile struct {
	buff     *bytes.Buffer
	fileName string
}

func NewCFile(fileName string) *CFile {
	f := new(CFile)
	f.buff = bytes.NewBuffer(nil)
	f.fileName = fileName
	return f
}

func (c *CFile) W(s string) {
	c.buff.WriteString(s)
}

func (c *CFile) WLn() {
	c.buff.WriteString("\n")
}

func (c *CFile) String() string {
	return c.buff.String()
}

func (c *CFile) WComment(s string) {
	c.WLn()
	c.W("// " + s)
	c.WLn()
}

func (c *CFile) WriteHeader() {
	c.WLn()
	c.W(`

// 代码由makeCHeader工具自动生成。  
// 作者：ying32  
// https://github.com/ying32  

#ifndef _LIBLCL_H
#define _LIBLCL_H

#ifdef __cplusplus
//extern "C" {
#define CChar const 
#else
#define CChar
#endif

#ifdef __GNUC__
   // #pragma GCC diagnostic ignored "-Wint-to-pointer-cast"
#endif

#include <stdint.h>
#include <stdio.h>
#include <assert.h>

#ifdef __GNUC__
#include <pthread.h>
#endif

#ifdef _WIN32
    #include<Windows.h>
    #define LCLAPI __stdcall
#else
    #include <dlfcn.h>
    #include <stddef.h>
    #include <stdbool.h>
    // __cdecl 默认
    #define LCLAPI 
    #define TRUE 1
    #define FALSE 0
#endif

#ifndef NULL
    #define NULL 0
#endif

//#ifdef __APPLE__
//    #include <Cocoa/Cocoa.h>
//#endif

// 非Windows下的类型定义  
#ifndef _WIN32
    typedef uintptr_t HWND;
    typedef uintptr_t HDC;
    typedef int32_t BOOL; 
    typedef uintptr_t HPEN;
    typedef uintptr_t HMENU;
    typedef uintptr_t HPALETTE;
    typedef uintptr_t HICON;
    typedef uintptr_t HKEY;
    typedef uintptr_t HMONITOR;
    typedef uintptr_t HBRUSH;
    typedef uintptr_t HBITMAP;
    typedef uintptr_t HFONT;
#endif


#ifdef __linux__

typedef uintptr_t PGdkWindow;
typedef uintptr_t TXId;
typedef uintptr_t PGtkFixed;

#endif

#ifdef __APPLE__
   //#include <Cocoa/Cocoa.h>
typedef *void MyNSWindow;
#endif

// printf("GetFunc: %s=%p\n", ""#name"", p##name); 

// 获取dll函数地址  
#define GET_FUNC_ADDR(name) \
if(!p##name) \
   p##name = get_proc_addr(""#name""); \
assert(p##name != NULL); 

// 定义dll函数指针
#define DEFINE_FUNC_PTR(name) \
static void* p##name; 

// 转换参数  
#define COV_PARAM(name) \
(uintptr_t)name

// 生成eventid，实际就是函数指针  
//#define MAKE_EVENT_ID(fn) \
//(uintptr_t)(void*)&fn

// 集合类型，提前申明
typedef uint32_t TSet;


// 常量
// keycode
// Virtual Keys, Standard Set
#define vkLButton     0x01 //   1
#define vkRButton     0x02 //   2
#define vkCancel      0x03 //   3
#define vkMButton     0x04 //   4
#define vkXButton1    0x05 //   5
#define vkXButton2    0x06 //   6
#define vkBack        0x08 //   8
#define vkTab         0x09 //   9
#define vkLineFeed    0x0A //  10
#define vkClear       0x0C //  12
#define vkReturn      0x0D //  13
#define vkShift       0x10 //  16
#define vkControl     0x11 //  17
#define vkMenu        0x12 //  18
#define vkPause       0x13 //  19
#define vkCapital     0x14 //  20
#define vkKana        0x15 //  21
#define vkHangul      0x15 //  21
#define vkJunja       0x17 //  23
#define vkFinal       0x18 //  24
#define vkHanja       0x19 //  25
#define vkKanji       0x19 //  25
#define vkConvert     0x1C //  28
#define vkNonConvert  0x1D //  29
#define vkAccept      0x1E //  30
#define vkModeChange  0x1F //  31
#define vkEscape      0x1B //  27
#define vkSpace       0x20 //  32
#define vkPrior       0x21 //  33
#define vkNext        0x22 //  34
#define vkEnd         0x23 //  35
#define vkHome        0x24 //  36
#define vkLeft        0x25 //  37
#define vkUp          0x26 //  38
#define vkRight       0x27 //  39
#define vkDown        0x28 //  40
#define vkSelect      0x29 //  41
#define vkPrint       0x2A //  42
#define vkExecute     0x2B //  43
#define vkSnapshot    0x2C //  44
#define vkInsert      0x2D //  45
#define vkDelete      0x2E //  46
#define vkHelp        0x2F //  47
// vk0 thru vk9 are the same as ASCII '0' thru '9' (0x30 - 0x39)
#define vk0  0x30 //  48
#define vk1  0x31 //  49
#define vk2  0x32 //  50
#define vk3  0x33 //  51
#define vk4  0x34 //  52
#define vk5  0x35 //  53
#define vk6  0x36 //  54
#define vk7  0x37 //  55
#define vk8  0x38 //  56
#define vk9  0x39 //  57
// vkA thru vkZ are the same as ASCII 'A' thru 'Z' (0x41 - 0x5A)
#define vkA          0x41 //  65
#define vkB          0x42 //  66
#define vkC          0x43 //  67
#define vkD          0x44 //  68
#define vkE          0x45 //  69
#define vkF          0x46 //  70
#define vkG          0x47 //  71
#define vkH          0x48 //  72
#define vkI          0x49 //  73
#define vkJ          0x4A //  74
#define vkK          0x4B //  75
#define vkL          0x4C //  76
#define vkM          0x4D //  77
#define vkN          0x4E //  78
#define vkO          0x4F //  79
#define vkP          0x50 //  80
#define vkQ          0x51 //  81
#define vkR          0x52 //  82
#define vkS          0x53 //  83
#define vkT          0x54 //  84
#define vkU          0x55 //  85
#define vkV          0x56 //  86
#define vkW          0x57 //  87
#define vkX          0x58 //  88
#define vkY          0x59 //  89
#define vkZ          0x5A //  90
#define vkLWin       0x5B //  91
#define vkRWin       0x5C //  92
#define vkApps       0x5D //  93
#define vkSleep      0x5F //  95
#define vkNumpad0    0x60 //  96
#define vkNumpad1    0x61 //  97
#define vkNumpad2    0x62 //  98
#define vkNumpad3    0x63 //  99
#define vkNumpad4    0x64 // 100
#define vkNumpad5    0x65 // 101
#define vkNumpad6    0x66 // 102
#define vkNumpad7    0x67 // 103
#define vkNumpad8    0x68 // 104
#define vkNumpad9    0x69 // 105
#define vkMultiply   0x6A // 106
#define vkAdd        0x6B // 107
#define vkSeparator  0x6C // 108
#define vkSubtract   0x6D // 109
#define vkDecimal    0x6E // 110
#define vkDivide     0x6F // 111
#define vkF1         0x70 // 112
#define vkF2         0x71 // 113
#define vkF3         0x72 // 114
#define vkF4         0x73 // 115
#define vkF5         0x74 // 116
#define vkF6         0x75 // 117
#define vkF7         0x76 // 118
#define vkF8         0x77 // 119
#define vkF9         0x78 // 120
#define vkF10        0x79 // 121
#define vkF11        0x7A // 122
#define vkF12        0x7B // 123
#define vkF13        0x7C // 124
#define vkF14        0x7D // 125
#define vkF15        0x7E // 126
#define vkF16        0x7F // 127
#define vkF17        0x80 // 128
#define vkF18        0x81 // 129
#define vkF19        0x82 // 130
#define vkF20        0x83 // 131
#define vkF21        0x84 // 132
#define vkF22        0x85 // 133
#define vkF23        0x86 // 134
#define vkF24        0x87 // 135

#define vkCamera        0x88 // 136
#define vkHardwareBack  0x89 // 137

#define vkNumLock   0x90 // 144
#define vkScroll    0x91 // 145
#define vkLShift    0xA0 // 160
#define vkRShift    0xA1 // 161
#define vkLControl  0xA2 // 162
#define vkRControl  0xA3 // 163
#define vkLMenu     0xA4 // 164
#define vkRMenu     0xA5 // 165

#define vkBrowserBack        0xA6 // 166
#define vkBrowserForward     0xA7 // 167
#define vkBrowserRefresh     0xA8 // 168
#define vkBrowserStop        0xA9 // 169
#define vkBrowserSearch      0xAA // 170
#define vkBrowserFavorites   0xAB // 171
#define vkBrowserHome        0xAC // 172
#define vkVolumeMute         0xAD // 173
#define vkVolumeDown         0xAE // 174
#define vkVolumeUp           0xAF // 175
#define vkMediaNextTrack     0xB0 // 176
#define vkMediaPrevTrack     0xB1 // 177
#define vkMediaStop          0xB2 // 178
#define vkMediaPlayPause     0xB3 // 179
#define vkLaunchMail         0xB4 // 180
#define vkLaunchMediaSelect  0xB5 // 181
#define vkLaunchApp1         0xB6 // 182
#define vkLaunchApp2         0xB7 // 183

#define vkSemicolon     0xBA // 186
#define vkEqual         0xBB // 187
#define vkComma         0xBC // 188
#define vkMinus         0xBD // 189
#define vkPeriod        0xBE // 190
#define vkSlash         0xBF // 191
#define vkTilde         0xC0 // 192
#define vkLeftBracket   0xDB // 219
#define vkBackslash     0xDC // 220
#define vkRightBracket  0xDD // 221
#define vkQuote         0xDE // 222
#define vkPara          0xDF // 223

#define vkOem102      0xE2 // 226
#define vkIcoHelp     0xE3 // 227
#define vkIco00       0xE4 // 228
#define vkProcessKey  0xE5 // 229
#define vkIcoClear    0xE6 // 230
#define vkPacket      0xE7 // 231
#define vkAttn        0xF6 // 246
#define vkCrsel       0xF7 // 247
#define vkExsel       0xF8 // 248
#define vkErEof       0xF9 // 249
#define vkPlay        0xFA // 250
#define vkZoom        0xFB // 251
#define vkNoname      0xFC // 252
#define vkPA1         0xFD // 253
#define vkOemClear    0xFE // 254
#define vkNone        0xFF // 255

// 消息框的
#define IdOK         1
#define IdCancel     2
#define IdAbort      3
#define IdRetry      4
#define IdIgnore     5
#define IdYes        6
#define IdNo         7
#define IdClose      8
#define IdHelp       9
#define IdTryAgain   10
#define IdContinue   11
#define mrNone       0
#define mrOk         IdOK
#define mrCancel     IdCancel
#define mrAbort      IdAbort
#define mrRetry      IdRetry
#define mrIgnore     IdIgnore
#define mrYes        IdYes
#define mrNo         IdNo
#define mrClose      IdClose
#define mrHelp       IdHelp
#define mrTryAgain   IdTryAgain
#define mrContinue   IdContinue
#define mrAll        MrContinue + 1
#define mrNoToAll    MrAll + 1
#define mrYesToAll   MrNoToAll + 1


// 所有LCL枚举类定义  
//  TCursor = -32768..32767;
typedef int16_t TCursor;
 
#define	crHigh  (TCursor)0;
#define	crDefault    (TCursor)0
#define	crNone       (TCursor)-1
#define	crArrow      (TCursor)-2
#define	crCross      (TCursor)-3
#define	crIBeam      (TCursor)-4
#define	crSize       (TCursor)-22
#define	crSizeNESW   (TCursor)-6 // diagonal north east - south west
#define	crSizeNS     (TCursor)-7
#define	crSizeNWSE   (TCursor)-8
#define	crSizeWE     (TCursor)-9
#define	crSizeNW     (TCursor)-23
#define	crSizeN      (TCursor)-24
#define	crSizeNE     (TCursor)-25
#define	crSizeW      (TCursor)-26
#define	crSizeE      (TCursor)-27
#define	crSizeSW     (TCursor)-28
#define	crSizeS      (TCursor)-29
#define	crSizeSE     (TCursor)-30
#define	crUpArrow    (TCursor)-10
#define	crHourGlass  (TCursor)-11
#define	crDrag       (TCursor)-12
#define	crNoDrop     (TCursor)-13
#define	crHSplit     (TCursor)-14
#define	crVSplit     (TCursor)-15
#define	crMultiDrag  (TCursor)-16
#define	crSQLWait    (TCursor)-17
#define	crNo         (TCursor)-18
#define	crAppStart   (TCursor)-19
#define	crHelp       (TCursor)-20
#define	crHandPoint  (TCursor)-21
#define	crSizeAll    (TCursor)-22
#define	crLow  (TCursor)-30
 

// 补充的一部分

// TShiftState
typedef enum {
    ssShift = 0,
    ssAlt,
    ssCtrl,
    ssLeft,
    ssRight,
    ssMiddle,
    ssDouble,
    // Extra additions
    ssMeta,
    ssSuper,
    ssHyper,
    ssAltGr,
    ssCaps,
    ssNum,
    ssScroll,
    ssTriple,
    ssQuad,
    ssExtra1,
    ssExtra2,
}TShiftStates;

// TControlState
typedef enum {
	csLButtonDown = 0,
	csClicked,
	csPalette,
	csReadingState,
	csFocusing,
	csCreating, // not used, exists for Delphi compatibility
	csPaintCopy,
	csCustomPaint,
	csDestroyingHandle,
	csDocking,
	csVisibleSetInLoading,
}TControlStates;


// TGridDrawState
typedef enum {
	gdSelected = 0,
	gdFocused,
	gdFixed,
	gdHot,
	gdPushed,
	gdRowHighlight,
}TGridDrawStates;

// TCustomDrawState
typedef enum {
	cdsSelected = 0,
	cdsGrayed,
	cdsDisabled,
	cdsChecked,
	cdsFocused,
	cdsDefault,
	cdsHot,
	cdsMarked,
	cdsIndeterminate,
}TCustomDrawStates;

<%enumdefs%>

typedef int32_t TLeftRight; 
typedef TBorderStyle TFormBorderStyle;
typedef int32_t TColorBoxStyle;
typedef TAlignment TLinkAlignment;
typedef TMenuItemAutoFlag TMenuAutoFlag; 
typedef int32_t TNumGlyphs;
typedef uint16_t TShortCut; 
typedef int16_t TScrollBarInc;
 
typedef enum  {
	ltVCL,
	ltLCL
}TLibType;


// 重定义
typedef int32_t TModalResult;
typedef uint32_t TColor;
typedef void* THelpEventData;
typedef uint16_t TTabOrder;
typedef void* PFNLVCOMPARE;
typedef void* PFNTVCOMPARE;
typedef uint8_t TFontCharset;
typedef int32_t TSpacingSize;
typedef uint16_t Char;
typedef uintptr_t TClass;

typedef uintptr_t HINST;
typedef uintptr_t LResult;
typedef uintptr_t WParam;
typedef uintptr_t LParam;
typedef uintptr_t THandle;


typedef struct TMessage {
    uint32_t Message;
#if defined(_WIN64) || defined(__x86_64__)
    uint32_t _UnusedMsg;
#endif
    WParam WParam;
    LParam LParam;
    LResult LResult;
}TMessage;

typedef struct TDWordFiller {
#if defined(_WIN64) || defined(__x86_64__)
    uint8_t Filler[4];
#endif
} TDWordFiller;

typedef struct TWMKey {
	uint32_t Msg;       
	TDWordFiller MsgFiller; 
	uint16_t CharCode[2]; // 第二个元素未使用
	TDWordFiller CharCodeUnusedFiller; 
	uint32_t KeyData;               
	TDWordFiller KeyDataFiller;        
	LResult Result;               
} TWMKey;

typedef struct TGridCoord {
   int32_t x, y;
} TGridCoord;

typedef void* TCustomData;


typedef struct TRect {
    int32_t Left, Top, Right, Bottom;
}TRect;

typedef TRect TGridRect;

typedef void* IObjectArray;

typedef struct TSysLocale {
	//Delphi compat fields
	int32_t DefaultLCID; 
	int32_t PriLangID;
	int32_t SubLangID;   

	// win32 names
	BOOL FarEast;    
	BOOL MiddleEast; 
}TSysLocale;

typedef struct TSmallPoint {
	int16_t X; 
	int16_t Y;
}TSmallPoint;

typedef struct TGUID {
	uint32_t D1; 
	uint16_t D2; 
	uint16_t D3; 
	uint8_t  D4[8];
}TGUID;

// LibResouces
typedef struct TLibResouce {
	char* Name;
	char* Value;   
}TLibResouce;

// TConstraintSize = 0..MaxInt;
typedef int32_t TConstraintSize; 

typedef struct TAlignInfo {
	void *AlignList; //: TList;
	int32_t ControlIndex;
	TAlign  Align;        
	int32_t Scratch;      
}TAlignInfo;

typedef struct TPoint{
    int32_t X, Y;
}TPoint;

typedef struct TSize {
	int32_t Cx, Cy;
}TSize;

typedef struct TResItem {
   char* Name;      
   char* ValuePtr;  
}TResItem;

// 所有LCL类定义  
<%typedefs%>
typedef void* TStream;


// 所有事件定义
<%eventdefs%>


// 一些其它函数

// Lazarus集合加法，val...中存储为位的索引，下标为0
TSet Include(TSet s, uint8_t val) {
    return (TSet)(s | (1 << val));
}
//TSet Include(TSet s, ...) {
//    uint32_t r = (uint32_t)s;
//    va_list varlist;     
//    va_start(varlist, s);
//    uint8_t val;
//    while ((val = va_arg(varlist, int)) != -1) {
//        r |= (1 << (uint8_t)val);
//    }
//    va_end(varlist); 
//    return (TSet)r;
//}

// Lazarus集合类型的判断,类型，然后后面是第几位，下标为0
TSet Exclude(TSet s, uint8_t val) {
    return (TSet)(s & (~(1 << val)));
}
//TSet Exclude(TSet s, ...) {
//    uint32_t r = (uint32_t)s;
//    va_list varlist;
//    va_start(varlist, s);
//    uint8_t val;
//    while ((val = va_arg(varlist, int)) != -1) {
//        r &= ~(1 << (uint8_t)val);
//    }
//    va_end(varlist);
//    return (TSet)r;
//}

// Lazarus集合类型的判断,类型，然后后面是第几位，下标为0
BOOL InSet(uint32_t s, uint8_t val) {
    if ((s&(1 << val)) != 0) {
        return TRUE;
    }
    return FALSE;
}



// liblcl句柄
static uintptr_t libHandle;

// 用于处理异常的模拟call  
typedef uint64_t(LCLAPI *MYSYSCALL)(void*, intptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);

// 函数指针
static MYSYSCALL pMySyscall;

// 异常处理函数实体  
#define MySyscall(addr, len, a1, a2 , a3, a4, a5, a6, a7, a8, a9, a10, a11, a12) \
    pMySyscall((void*)addr, (intptr_t)len, COV_PARAM(a1), COV_PARAM(a2), COV_PARAM(a3), COV_PARAM(a4), COV_PARAM(a5), COV_PARAM(a6), COV_PARAM(a7), COV_PARAM(a8), COV_PARAM(a9), COV_PARAM(a10), COV_PARAM(a11), COV_PARAM(a12))

// 全局实例类定义
static TApplication Application; // 应用程序
static TScreen Screen;           // 屏幕
static TMouse	Mouse;            // 鼠标
static TClipboard	Clipboard;    // 剪切板
static TPrinter Printer;         // 打印机

// 全局互斥锁
#ifdef __GNUC__
static pthread_mutex_t threadSyncMutex;
#else
static RTL_CRITICAL_SECTION threadSyncMutex;
#endif

// 初始liblcl库
void init_lib_lcl();
// 反向初始liblcl库
void un_init_lib_lcl();


// 获取过程地址
void* get_proc_addr(const char *name) {
#ifdef _WIN32
    return (void*)GetProcAddress((HMODULE)libHandle, name);
#else
    return (void*)dlsym((void*)libHandle, name); 
#endif
}

// 加载库
BOOL load_liblcl(const char *name) {
    if(libHandle > 0)
        return TRUE;
#ifdef _WIN32
    libHandle = (uintptr_t)LoadLibraryA(name);
#else
    libHandle = (uintptr_t)dlopen(name, RTLD_LAZY|RTLD_GLOBAL); 
#endif
    if(libHandle > 0) {
         pMySyscall = (MYSYSCALL)get_proc_addr("MySyscall");
         // 初始库
         init_lib_lcl();
    }
    return libHandle > 0;
}

// 关闭库
void close_liblcl() {
    if(libHandle > 0) {
	#ifdef _WIN32
	    FreeLibrary((HMODULE)libHandle);
	#else
	    dlclose((void*)libHandle);
	#endif
        libHandle = 0;
        un_init_lib_lcl();
    }
}

 
`)
}

func (c *CFile) WriteFooter() {
	c.W(`


// 模拟call  
#define Syscall12(addr, p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12) \
    ((uint64_t(*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))addr)(p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12)

// getParam 从指定索引和地址获取事件中的参数
#define getParamOf(index, ptr) \
 (*((uintptr_t*)((uintptr_t)ptr + (uintptr_t)index*sizeof(uintptr_t))))


// 获取参数的宏  
#define GET_VAL(index) \
getParamOf(index, args)


// 事件回调
void* LCLAPI doEventCallbackProc(void* f, void* args, long argcount) {
 
    switch (argcount) {
    case 0: Syscall12(f, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);
       break;
	
    case 1: Syscall12(f, GET_VAL(0), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);
       break;
	
    case 2: Syscall12(f, GET_VAL(0), GET_VAL(1), 0, 0, 0, 0, 0, 0, 0, 0, 0, 0);
        break;
	
    case 3: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), 0, 0, 0, 0, 0, 0, 0, 0, 0);
       break;
	
    case 4: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(2), 0, 0, 0, 0, 0, 0, 0, 0);
       break;
	
    case 5: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), 0, 0, 0, 0, 0, 0, 0);
       break;
	
    case 6: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), 0, 0, 0, 0, 0, 0);
       break;
	
    case 7: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), 0, 0, 0, 0, 0);
       break;
	
    case 8: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), 0, 0, 0, 0);
       break;
	
    case 9: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), 0, 0, 0);
       break;
	
    case 10: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), GET_VAL(9), 0, 0);
       break;
	
    case 11: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), GET_VAL(9), GET_VAL(10), 0);
       break;
	
    case 12: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), GET_VAL(9), GET_VAL(10), GET_VAL(11));
       break;
    }
    return NULL;
}
 
 
// 消息回调
void* LCLAPI doMessageCallbackProc(void* f, void* msg) {
   ((void(*)(void*))f)(msg);
    return NULL;
}
 
// 线程同步过程
TThreadProc threadSyncProc;
// 线程同步回调
void* LCLAPI doThreadSyncCallbackProc() {
    if (threadSyncProc) {
        ((TThreadProc)threadSyncProc)();
        threadSyncProc = NULL;
    }
    return NULL;
}

// 线程同步方法
// 无参数，无返回值的一个函数
void ThreadSync(TThreadProc fn) {
    // 加锁
#ifdef __GNUC__
    pthread_mutex_lock(&threadSyncMutex);
#else
    EnterCriticalSection(&threadSyncMutex);
#endif
    threadSyncProc = fn;
    Synchronize(FALSE);
    threadSyncProc = NULL;
#ifdef __GNUC__
    pthread_mutex_unlock(&threadSyncMutex);
#else
    LeaveCriticalSection(&threadSyncMutex);
#endif
   
}
 
#define GET_CALLBACK(name) \
  (void*)&name
 
void init_lib_lcl() {
#ifdef __GNUC__
    pthread_mutex_init(&threadSyncMutex, NULL);
#else
    InitializeCriticalSection(&threadSyncMutex);
#endif

    // 设置事件的回调函数 
	SetEventCallback(GET_CALLBACK(doEventCallbackProc));
	// 消息回调
	SetMessageCallback(GET_CALLBACK(doMessageCallbackProc));
	// 线程同步回调
	SetThreadSyncCallback(GET_CALLBACK(doThreadSyncCallbackProc));

	Application = Application_Instance();
	Screen = Screen_Instance();
	Mouse = Mouse_Instance();             // 鼠标
	Clipboard = Clipboard_Instance();     // 剪切板
	Printer = Printer_Instance();         // 打印机

#ifdef _WIN32
    // 尝试加载exe中名为MAINICON的图标为应用程序图标
    if(Application) {
        TIcon icon = Application_GetIcon(Application);
        if(icon) {
            Icon_SetHandle(icon, LoadIconA(GetModuleHandleA(NULL), "MAINICON"));
        } 
    }
#endif
}

void un_init_lib_lcl() {
#ifdef __GNUC__
    pthread_mutex_destroy(&threadSyncMutex);
#else
    DeleteCriticalSection(&threadSyncMutex);
#endif
}

//#ifdef __cplusplus
//}
//#endif

#endif // _LIBLCL_H

`)
}

func (c *CFile) Save(classDefs []string, enumDefs []byte, eventDefs []byte) error {
	if len(classDefs) == 0 {
		return ioutil.WriteFile(c.fileName, c.buff.Bytes(), 0664)
	} else {
		getClassDefs := func() []byte {
			buf := bytes.NewBuffer(nil)
			for _, class := range classDefs {
				buf.WriteString(fmt.Sprintf("typedef void* %s;\n", class))
			}
			return buf.Bytes()
		}

		bs := bytes.Replace(c.buff.Bytes(), []byte("<%typedefs%>"), getClassDefs(), 1)
		bs = bytes.Replace(bs, []byte("<%enumdefs%>"), enumDefs, 1)
		bs = bytes.Replace(bs, []byte("<%eventdefs%>"), eventDefs, 1)
		bs = bytes.Replace(bs, []byte("\n"), []byte("\r\n"), -1)
		return ioutil.WriteFile(c.fileName, bs, 0664)
	}
}
