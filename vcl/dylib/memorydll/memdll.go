// +build windows,386

/*
  此项功能稳定性还有待观查。

  作者：ying32，本为govcl项目中的一个dylib扩展包
  https://github.com/ying32/govcl

  这是一个移植自delphi BTMemoryModule.pas的，暂时只支持32bit dll， 最原始版本的c++代码
  已经支持64bit，暂时还没有想法去做移植。

*/

package memorydll

import (
	"errors"
	"strconv"
	"sync"
	"syscall"

	"unsafe"

	"fmt"

	"github.com/ying32/govcl/vcl/dylib/memorydll/btmemorymodule"
)

type LazyDLL struct {
	handle *btmemorymodule.TBTMemoryModule
	mu     sync.Mutex
	err    error
	// 加了异常处理的
	mySyscall *LazyProc
}

// NewMemoryDLL 内存DLL
func NewMemoryDLL(data []byte) *LazyDLL {
	m := new(LazyDLL)
	if len(data) > 0 {
		m.mu.Lock()
		defer m.mu.Unlock()

		m.handle = btmemorymodule.BTMemoryLoadLibary(uintptr(unsafe.Pointer(&data[0])), int64(len(data)))
		if m.handle != nil && m.handle.IsValid() {
			//runtime.SetFinalizer(m, (*LazyDLL).Close)
			// 导入调用的, 实现一个动态调用call的，主要是为了解决异常问题
			m.mySyscall = m.NewProc("MySyscall")
			if m.mySyscall.Find() != nil {
				m.mySyscall = nil
			}
		} else {
			m.err = fmt.Errorf(btmemorymodule.BTMemoryGetLastError())
		}
	}
	return m
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
	if l.handle != nil && l.handle.IsValid() {
		btmemorymodule.BTMemoryFreeLibrary(l.handle)
		l.handle = nil
	}
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
		if p.lzdll.handle != nil && p.lzdll.handle.IsValid() {
			p.p = btmemorymodule.BTMemoryGetProcAddress(p.lzdll.handle, p.Name)
		}
	}
	if p.p == 0 {
		return errors.New("proc\"" + p.Name + "\" not find.")
	}
	return nil
}

func (p *LazyProc) Call(a ...uintptr) (uintptr, uintptr, error) {
	return p.lzdll.call(p, a...)
}

func (p *LazyProc) CallOriginal(a ...uintptr) (r1, r2 uintptr, lastErr error) {
	p.Find()

	if p.p == 0 {
		return 0, 0, syscall.EINVAL
	}
	switch len(a) {
	case 0:
		return syscall.Syscall(p.p, 0, 0, 0, 0)
	case 1:
		return syscall.Syscall(p.p, 1, a[0], 0, 0)
	case 2:
		return syscall.Syscall(p.p, 2, a[0], a[1], 0)
	case 3:
		return syscall.Syscall(p.p, 3, a[0], a[1], a[2])
	case 4:
		return syscall.Syscall6(p.p, 4, a[0], a[1], a[2], a[3], 0, 0)
	case 5:
		return syscall.Syscall6(p.p, 5, a[0], a[1], a[2], a[3], a[4], 0)
	case 6:
		return syscall.Syscall6(p.p, 6, a[0], a[1], a[2], a[3], a[4], a[5])
	case 7:
		return syscall.Syscall9(p.p, 7, a[0], a[1], a[2], a[3], a[4], a[5], a[6], 0, 0)
	case 8:
		return syscall.Syscall9(p.p, 8, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], 0)
	case 9:
		return syscall.Syscall9(p.p, 9, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8])
	case 10:
		return syscall.Syscall12(p.p, 10, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], 0, 0)
	case 11:
		return syscall.Syscall12(p.p, 11, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], 0)
	case 12:
		return syscall.Syscall12(p.p, 12, a[0], a[1], a[2], a[3], a[4], a[5], a[6], a[7], a[8], a[9], a[10], a[11])

	default:
		panic("Call " + p.Name + " with too many arguments " + strconv.Itoa(len(a)) + ".")
	}
}
