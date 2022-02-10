#include "textflag.h"

// Id int
GLOBL ·Id(SB),NOPTR,$8
DATA ·Id+0(SB)/8,$0x00002537
// DATA ·Id+0(SB)/1,$0x37
// DATA ·Id+1(SB)/1,$0x25
// DATA ·Id+2(SB)/1,$0x00
// DATA ·Id+3(SB)/1,$0x00
// DATA ·Id+4(SB)/1,$0x00
// DATA ·Id+5(SB)/1,$0x00
// DATA ·Id+6(SB)/1,$0x00
// DATA ·Id+7(SB)/1,$0x00

// Name string
GLOBL ·NameData(SB),NOPTR,$8
DATA ·NameData(SB)/8,$"gopher"

GLOBL ·Name(SB),NOPTR,$16
DATA ·Name+0(SB)/8,$·NameData(SB)
DATA ·Name+8(SB)/8,$6

// PrintHelloWorld
TEXT ·PrintHelloWorld(SB), $16-0
    MOVQ ·helloWorld+0(SB), AX; MOVQ AX, 0(SP)
    MOVQ ·helloWorld+8(SB), BX; MOVQ BX, 8(SP)
    CALL runtime·printstring(SB)
    CALL runtime·printnl(SB)
    RET

// Num int数组
GLOBL ·Num(SB),NOPTR,$16
DATA ·Num+0(SB)/8,$0
DATA ·Num+8(SB)/8,$1

// bool
GLOBL ·boolValue(SB),NOPTR,$1

GLOBL ·trueValue(SB),NOPTR,$1
DATA ·trueValue(SB)/1,$1

GLOBL ·falseValue(SB),NOPTR,$1
DATA ·falseValue(SB)/1,$0

// int
GLOBL ·int32Value(SB),NOPTR,$4
DATA ·int32Value(SB)/4,$0x01020304//$0x04030201
//DATA ·int32Value+0(SB)/1,$0x01
//DATA ·int32Value+1(SB)/1,$0x02
//DATA ·int32Value+2(SB)/2,$0x03

GLOBL ·uint32Value(SB),NOPTR,$4
DATA ·uint32Value(SB)/4,$0x01020304

// float
GLOBL ·float32Value(SB),NOPTR,$4
DATA ·float32Value(SB)/4,$1.5

GLOBL ·float64Value(SB),NOPTR,$8
DATA ·float64Value(SB)/8,$1.51

// string
GLOBL text<>(SB),NOPTR,$16
DATA text<>+0(SB)/8,$"Hello Wo"
DATA text<>+8(SB)/8,$"rld!"

GLOBL ·helloworld(SB),NOPTR,$16
DATA ·helloworld+0(SB)/8,$text<>(SB)
DATA ·helloworld+8(SB)/8,$12

// slice
GLOBL ·sliceValue(SB),NOPTR,$24
DATA ·sliceValue+0(SB)/8,$text<>(SB)
DATA ·sliceValue+8(SB)/8,$12
DATA ·sliceValue+16(SB)/8,$16

// map、chan
GLOBL ·mapValue(SB),NOPTR,$8
DATA ·mapValue+0(SB)/8,$0

GLOBL ·chanValue(SB),NOPTR,$8
DATA ·chanValue+0(SB)/8,$0

// func Swap(int, int) (int, int)
TEXT ·Swap(SB),NOSPLIT,$0-32
    MOVQ a+0(FP), AX
    MOVQ b+8(FP), BX
    MOVQ BX, ret0+16(FP)
    MOVQ AX, ret1+24(FP)
    RET

// func Foo(bool, int16) []byte
TEXT ·Foo(SB),$0-32
    MOVQ a+0(FP), AX
    MOVQ b+2(FP), BX
    MOVQ c_data+8*1(FP), CX
    MOVQ c_len+8*2(FP), DX
    MOVQ c_cap+8*3(FP), DI
    RET

// func Foo1()
TEXT ·Foo1(SB),$0
    MOVQ a-32(SP), AX
    MOVQ b-30(SP), BX
    MOVQ c_data-24(SP), CX
    MOVQ c_len-16(SP), DX
    MOVQ c_cap-8(SP), DI
    RET

// func If(bool, int, int) int
TEXT ·If(SB),NOSPLIT,$0-32
    MOVQ ok+8*0(FP), CX
    MOVQ a+8*1(FP), AX
    MOVQ b+8*2(FP), BX

    CMPQ CX, $0
    JZ   L
    MOVQ AX, ret+8*3(FP)
    RET
L:
    MOVQ BX, ret+8*3(FP)
    RET

// func LoopAdd(cnt, v0, step int) int
TEXT ·LoopAdd(SB),NOSPLIT,$0-32
    MOVQ cnt+0(FP), AX // cnt
    MOVQ v0+8(FP), BX // v0/result
    MOVQ step+8*2(FP), CX // step

LOOP_BEGIN:
    MOVQ $0, DX // i

LOOP_IF:
    CMPQ DX, AX // compare i, cnt
    JL LOOP_BODY // if i < cnt: goto LOOP_BODY
    JMP LOOP_END

LOOP_BODY:
    ADDQ $1, DX // i++
    ADDQ CX, BX // result += step
    JMP LOOP_IF

LOOP_END:
    MOVQ BX, ret+8*3(FP) // return result
    RET

// func CopySlice_AVX2(dst, src []byte, len int)
TEXT ·CopySlice_AVX2(SB), NOSPLIT, $0
    MOVQ dst_data+0(FP), DI
    MOVQ src_data+24(FP), SI
    MOVQ len+32(FP), BX
    MOVQ $0, AX

LOOP:
    VMOVDQU 0(SI)(AX*1), Y0
    VMOVDQU Y0, 0(DI)(AX*1)
    ADDQ $32, AX
    CMPQ AX, BX
    JL LOOP
    RET


// func getG() unsafe.Pointer
TEXT ·getG(SB), NOSPLIT, $0-8
    MOVQ (TLS), AX // get runtime.g
    MOVQ $type·runtime·g(SB), BX // get runtime.g type
    MOVQ AX, ret+0(FP)
    RET

