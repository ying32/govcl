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

#include <stdint.h>
#include <stdint.h>
#include <assert.h>
#ifdef _WIN32
    #include<Windows.h>
    #define LCLAPI __stdcall
#else
    #include <dlfcn.h>

    #define LCLAPI __cdecl
#endif



#ifdef __APPLE__
    #include <Cocoa.h>
#endif

// 非Windows下的类型定义  
#ifndef _WIN32
    typedef  uintptr_t HWND;
    typedef  uintptr_t HDC;
#endif


//printf("GetFunc: %s\n", ""#name""); \
// 获取dll函数地址  
#define GET_FUNC_ADDR(name) \
if(p##name == NULL) \
   p##name = get_proc_addr(""#name""); \
assert(p##name != NULL);

// 转换参数  
#define COV_PARAM(name) \
(uintptr_t)name

// 生成eventid，实际就是函数指针  
#define MAKE_EVENT_ID(fn) \
(uintptr_t)&fn

// 集合类型，提前申明
typedef uint32_t TSet;


// 所有LCL枚举类定义  
<%enumdefs%>

 
enum TLibType {
	LtVCL,
	LtLCL
};

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

}TMessage;

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
	bool FarEast;    
	bool MiddleEast; 
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
	//Name string
	//Ptr  uintptr
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

// liblcl句柄
uintptr_t libHandle;

// 用于处理异常的模拟call  
typedef uint64_t(* LCLAPI MYSYSCALL)(void*, intptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t);

// 函数指针
MYSYSCALL pMySyscall;

// 异常处理函数实体  
inline uint64_t MySyscall(void* addr, intptr_t len = 0, uintptr_t a1 = NULL, uintptr_t a2 = NULL, uintptr_t a3 = NULL, uintptr_t a4 = NULL, 
  uintptr_t a5 = NULL, uintptr_t a6 = NULL, uintptr_t a7 = NULL, uintptr_t a8 = NULL, uintptr_t a9 = NULL, uintptr_t a10 = NULL, uintptr_t a11 = NULL, uintptr_t a12 = NULL) {
    return pMySyscall(addr, len, a1, a2, a3, a4, a5, a6, a7, a8, a9, a10, a11, a12);
}

// 所有LCL类定义  
<%typedefs%>
typedef void* TStream;

// 全局实例类定义
TApplication Application; // 应用程序
TScreen Screen;           // 屏幕
TMouse	Mouse;            // 鼠标
TClipboard	Clipboard;    // 剪切板
TPrinter Printer;         // 打印机

// 初始liblcl库
void init_lib_lcl();


// 获取过程地址
void* get_proc_addr(const char *name) {
#ifdef _WIN32
    return (void*)GetProcAddress((HMODULE)libHandle, name);
#else
    return (void*)dlsym((void*)libHandle, name); 
#endif
}

// 加载库
bool load_liblcl(const char *name) {
    if(libHandle > 0)
        return true;
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
    }
}

 
`)
}

func (c *CFile) WriteFooter() {
	c.W(`

// 模拟call  
inline uint64_t Syscall12(void* addr, uintptr_t p1 = NULL, uintptr_t p2 = NULL, uintptr_t p3 = NULL, uintptr_t p4 = NULL, uintptr_t p5 = NULL, uintptr_t p6 = NULL,
    uintptr_t p7 = NULL, uintptr_t p8 = NULL, uintptr_t p9 = NULL, uintptr_t p10 = NULL, uintptr_t p11 = NULL, uintptr_t p12 = NULL) {
    return ((uint64_t(*)(uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t, uintptr_t))addr)(
        p1, p2, p3, p4, p5, p6, p7, p8, p9, p10, p11, p12);
}

// getParam 从指定索引和地址获取事件中的参数
inline uintptr_t getParamOf(intptr_t index, void* ptr)  {
	return *((uintptr_t*)((uintptr_t)ptr + (uintptr_t)index*sizeof(uintptr_t)));
}

// 获取参数的宏  
#define GET_VAL(index) \
getParamOf(index, args)


// 事件回调
void* LCLAPI doEventCallbackProc(void* f, void* args, long argcount) {
 
    switch (argcount) {
    case 0: Syscall12(f);
        break;

    case 1: Syscall12(f, GET_VAL(0));
        break;
  
    case 2: Syscall12(f, GET_VAL(0), GET_VAL(1));
         break;

    case 3: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2));
        break;

    case 4: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(2));
        break;

    case 5: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4));
        break;

    case 6: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5));
        break;

    case 7: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6));
        break;

    case 8: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7));
        break;

    case 9: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8));
        break;

    case 10: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), GET_VAL(9));
        break;

    case 11: Syscall12(f, GET_VAL(0), GET_VAL(1), GET_VAL(2), GET_VAL(3), GET_VAL(4), GET_VAL(5), GET_VAL(6), GET_VAL(7), GET_VAL(8), GET_VAL(9), GET_VAL(10));
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
 
// 线程同步回调
void* LCLAPI doThreadSyncCallbackProc() {
    //((void(*)(void*))f)();
    return NULL;
}
 
void init_lib_lcl() {
 
    // 设置事件的回调函数 
	SetEventCallback(&doEventCallbackProc);
	// 消息回调
	SetMessageCallback(&doMessageCallbackProc);
	// 线程同步回调
	SetThreadSyncCallback(&doThreadSyncCallbackProc);

	Application = Application_Instance();
	Screen = Screen_Instance();
	Mouse = Mouse_Instance();             // 鼠标
	Clipboard = Clipboard_Instance();     // 剪切板
	Printer = Printer_Instance();         // 打印机
}

#endif // _LIBLCL_H

`)
}

func (c *CFile) Save(classDefs []string, enumDefs []byte) error {
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
		bs = bytes.Replace(bs, []byte("\n"), []byte("\r\n"), -1)
		return ioutil.WriteFile(c.fileName, bs, 0664)
	}
}
