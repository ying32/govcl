
#include "textflag.h"

TEXT ·SetFPMask(SB),NOSPLIT,$8
	STMXCSR	0(SP)
	MOVL	0(SP), AX
	//ANDL	$~0x3F, AX
	//ORL	$(0x3F<<7), AX
	MOVL $0x137F, AX
	MOVL	AX, 0(SP)
	LDMXCSR	0(SP)
	RET
