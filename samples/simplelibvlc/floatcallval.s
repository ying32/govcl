#include "textflag.h"

// 一个补丁 by：ying32
// 由于go的syscall不能返回浮点数，所以直接用汇编获取浮点寄存器xmm0的值，并返回。
// 后面要迁回dylib包

// float32 getfloat32();
TEXT ·getfloat32(SB),NOSPLIT|NOFRAME,$0
    MOVL  X0, ret(FP)
    RET

// float64 getfloat64();
TEXT ·getfloat64(SB),NOSPLIT|NOFRAME,$0
    MOVQ  X0, ret(FP)
    RET
