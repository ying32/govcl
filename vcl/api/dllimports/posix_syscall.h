

static uint64_t Syscall0(uintptr_t addr) {
    return ((uint64_t(*)())addr)();
}

static uint64_t Syscall1(uintptr_t addr, uintptr_t p1) {
    return ((uint64_t(*)(uintptr_t))addr)(p1);
}

static uint64_t Syscall2(uintptr_t addr, uintptr_t p1, uintptr_t p2) {
    return ((uint64_t(*)(uintptr_t,uintptr_t))addr)(p1, p2);
}

static uint64_t Syscall3(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3);
}

static uint64_t Syscall4(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4);
}

static uint64_t Syscall5(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5);
}

static uint64_t Syscall6(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6);
}

static uint64_t Syscall7(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t, uintptr_t))addr)(p1, p2, p3, p4, p5, p6, p7);
}

static uint64_t Syscall8(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7, uintptr_t p8) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6,p7,p8);
}

static uint64_t Syscall9(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7, uintptr_t p8, uintptr_t p9) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9);
}

static uint64_t Syscall10(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7, uintptr_t p8, uintptr_t p9, uintptr_t p10) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10);
}

static uint64_t Syscall11(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7, uintptr_t p8, uintptr_t p9, uintptr_t p10, uintptr_t p11) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11);
}

static uint64_t Syscall12(uintptr_t addr, uintptr_t p1, uintptr_t p2, uintptr_t p3, uintptr_t p4, uintptr_t p5, uintptr_t p6, uintptr_t p7, uintptr_t p8, uintptr_t p9, uintptr_t p10, uintptr_t p11, uintptr_t p12) {
    return ((uint64_t(*)(uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t,uintptr_t))addr)(p1, p2, p3, p4, p5, p6,p7,p8,p9,p10,p11,p12);
}