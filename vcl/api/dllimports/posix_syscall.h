

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