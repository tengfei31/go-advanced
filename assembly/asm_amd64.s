#include "textflag.h"

// func AsmCallCAdd(cfun uintptr, a, b int64) int64
TEXT ·AsmCallCAdd(SB), NOSPLIT, $0
    MOVQ cfun+0(FP), AX //cfun
    MOVQ a+8(FP), DI // a
    MOVQ b+8*2(FP), SI // b
    CALL AX
    MOVQ AX, ret+8*3(FP)
    RET

