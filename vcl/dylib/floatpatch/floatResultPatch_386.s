#include "textflag.h"

// 一个补丁 by：ying32
// 由于go的syscall不能返回浮点数，所以直接用汇编获取浮点寄存器F0的值，并返回。

// float32 Getfloat32();
TEXT ·Getfloat32(SB),NOSPLIT|NOFRAME,$0
    FMOVFP F0, ret(FP)
    RET

// float64 Getfloat64();
TEXT ·Getfloat64(SB),NOSPLIT|NOFRAME,$0
	FMOVDP F0, ret(FP)
    RET
