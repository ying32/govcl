// +build !windows
// +build cgo

package dylib

/*
	#cgo LDFLAGS: -ldl

	//#cgo LDFLAGS: -L. -Wl,-rpath,${SRCDIR} -llcl

	#include <dlfcn.h>
	#include <limits.h>
	#include <stdlib.h>
	#include <stdint.h>
	#include <stdio.h>

	static uintptr_t libOpen(const char* path) {
	     void* h = dlopen(path, RTLD_LAZY|RTLD_GLOBAL); // RTLD_LAZY RTLD_NOW |RTLD_GLOBAL
	     if (h == NULL) {
		     printf("dlopen err: %s\n", (char*)dlerror());
	         return 0;
	     }
	     return (uintptr_t)h;
	}

	static uintptr_t libLookup(uintptr_t h, const char* name) {
	     void* r = dlsym((void*)h, name);
	     if (r == NULL) {
			 printf("dlsym err: %s\n", (char*)dlerror());
	         return 0;
	     }
	     return (uintptr_t)r;
	}

	static void libClose(uintptr_t h) {
	     if(h != 0) {
	         dlclose((void*)h);
	     }
	}


    static uint64_t Syscall0(void* addr) {
		return ((uint64_t(*)())addr)();
	}

    static uint64_t Syscall1(void* addr, void* p1) {
		return ((uint64_t(*)(void*))addr)(p1);
	}

    static uint64_t Syscall2(void* addr, void* p1, void* p2) {
		return ((uint64_t(*)(void*,void*))addr)(p1, p2);
	}

    static uint64_t Syscall3(void* addr, void* p1, void* p2, void* p3) {
		return ((uint64_t(*)(void*,void*,void*))addr)(p1, p2, p3);
	}

    static uint64_t Syscall4(void* addr, void* p1, void* p2, void* p3, void* p4) {
		return ((uint64_t(*)(void*,void*,void*,void*))addr)(p1, p2, p3, p4);
	}

    static uint64_t Syscall5(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5);
	}

    static uint64_t Syscall6(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6);
	}

    static uint64_t Syscall7(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*, void*))addr)(p1, p2, p3, p4, p5, p6, p7);
	}

    static uint64_t Syscall8(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8);
	}

    static uint64_t Syscall9(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9);
	}

    static uint64_t Syscall10(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10);
	}

    static uint64_t Syscall11(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10, void *p11) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11);
	}

    static uint64_t Syscall12(void* addr, void* p1, void* p2, void* p3, void* p4, void* p5, void* p6, void *p7, void *p8, void *p9, void *p10, void *p11, void *p12) {
		return ((uint64_t(*)(void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*,void*))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11,p12);
	}

*/
import "C"

import (
	"errors"
	"fmt"
	"os"
	"os/exec"
	"path/filepath"
	"runtime"
	"strconv"
	"sync"
	"syscall"
	"unsafe"
)

type LazyDLL struct {
	mu     sync.Mutex
	handle C.uintptr_t
	err    error
	Name   string
	// 加了异常处理的
	mySyscall *LazyProc
}

func NewLazyDLL(name string) *LazyDLL {

	m := new(LazyDLL)
	m.Name = name
	m.mu.Lock()
	defer m.mu.Unlock()

	cPath := (*C.char)(C.malloc(C.PATH_MAX + 1))
	defer C.free(unsafe.Pointer(cPath))

	cRelName := C.CString(m.libFullPath(name))
	defer C.free(unsafe.Pointer(cRelName))

	if C.realpath(cRelName, cPath) == nil {
		m.handle = C.libOpen(cRelName)
	} else {
		m.handle = C.libOpen(cPath)
	}
	if m.handle == 0 {
		m.err = fmt.Errorf("dlopen(\"%s\"), failed.", name)
		return m
	}
	// 导入调用的, 实现一个动态调用call的，主要是为了解决异常问题
	m.mySyscall = m.NewProc("MySyscall")
	if m.mySyscall.Find() != nil {
		m.mySyscall = nil
	}

	return m
}

func fileExists(path string) bool {
	_, err := os.Stat(path)
	if err == nil {
		return true
	}
	if os.IsNotExist(err) {
		return false
	}
	return false
}

func (l *LazyDLL) libFullPath(name string) string {
	if runtime.GOOS == "darwin" {
		file, _ := exec.LookPath(os.Args[0])
		libPath := filepath.Dir(file) + "/" + name
		if fileExists(libPath) {
			return libPath
		}
	}
	return name
}

func (l *LazyDLL) Load() error {
	return l.err
}

func (l *LazyDLL) NewProc(name string) *LazyProc {
	p := new(LazyProc)
	p.Name = name
	p.lzdll = l
	return p
}

func (l *LazyDLL) Close() {
	C.libClose(l.handle)
}

func (d *LazyDLL) call(proc *LazyProc, a ...uintptr) (r1, r2 uintptr, lastErr error) {
	// 没到找到我封装的那个系统函数，就使用原始的
	if d.mySyscall == nil {
		return proc.CallOriginal(a...)
	}
	addr := proc.Addr()
	if addr != 0 {
		pLen := uintptr(len(a))
		switch pLen {
		case 0:
			return d.mySyscall.CallOriginal(addr, pLen)
		case 1:
			return d.mySyscall.CallOriginal(addr, pLen, a[0])
		case 2:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1])
		case 3:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2])
		case 4:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3])
		case 5:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4])
		case 6:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5])
		case 7:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6])
		case 8:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7])
		case 9:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
		case 10:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9])
		case 11:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10])
		case 12:
			return d.mySyscall.CallOriginal(addr, pLen, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])
		default:
			panic("Call " + proc.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
		}
	}
	return 0, 0, syscall.EINVAL
}

type LazyProc struct {
	mu    sync.Mutex
	p     uintptr
	Name  string
	lzdll *LazyDLL
}

func (p *LazyProc) Addr() uintptr {
	p.Find()
	return p.p
}

func (p *LazyProc) Find() error {
	if p.p == 0 {
		p.mu.Lock()
		defer p.mu.Unlock()
		cRelName := C.CString(p.Name)
		defer C.free(unsafe.Pointer(cRelName))
		p.p = uintptr(C.libLookup(p.lzdll.handle, cRelName))
		// 这里不管了
	}
	if p.p == 0 {
		return errors.New("proc\"" + p.Name + "\" not find.")
	}
	return nil
}

func toPtr(a uintptr) unsafe.Pointer {
	return unsafe.Pointer(a)
}

func (p *LazyProc) Call(a ...uintptr) (uintptr, uintptr, error) {
	return p.lzdll.call(p, a...)
}

func (p *LazyProc) CallOriginal(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	p.Find()

	if p.p == 0 {
		return 0, 0, syscall.EINVAL
	}
	var ret C.uint64_t
	switch len(a) {
	case 0:
		ret = C.Syscall0(unsafe.Pointer(p.p))
	case 1:
		ret = C.Syscall1(toPtr(p.p), toPtr(a[0]))
	case 2:
		ret = C.Syscall2(toPtr(p.p), toPtr(a[0]), toPtr(a[1]))
	case 3:
		ret = C.Syscall3(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]))
	case 4:
		ret = C.Syscall4(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]))
	case 5:
		ret = C.Syscall5(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]))
	case 6:
		ret = C.Syscall6(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]))
	case 7:
		ret = C.Syscall7(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]))
	case 8:
		ret = C.Syscall8(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]), toPtr(a[7]))
	case 9:
		ret = C.Syscall9(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]), toPtr(a[7]), toPtr(a[8]))
	case 10:
		ret = C.Syscall10(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]), toPtr(a[7]), toPtr(a[8]), toPtr(a[9]))
	case 11:
		ret = C.Syscall11(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]), toPtr(a[7]), toPtr(a[8]), toPtr(a[9]), toPtr(a[10]))
	case 12:
		ret = C.Syscall12(toPtr(p.p), toPtr(a[0]), toPtr(a[1]), toPtr(a[2]), toPtr(a[3]), toPtr(a[4]), toPtr(a[5]), toPtr(a[6]), toPtr(a[7]), toPtr(a[8]), toPtr(a[9]), toPtr(a[10]), toPtr(a[11]))

	default:
		panic("Call " + p.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
	}
	return uintptr(ret), uintptr(ret >> 32), nil
}
