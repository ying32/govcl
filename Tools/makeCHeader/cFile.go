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
	LtVCL,
	LtLCL
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
