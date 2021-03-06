// Code generated by command: go run gen.go -out blocks_amd64.s -stubs blocks_amd64.go. DO NOT EDIT.

// +build amd64

#include "textflag.h"

DATA wyp<>+0(SB)/8, $0xa0761d6478bd642f
DATA wyp<>+8(SB)/8, $0xe7037ed1a0b428db
DATA wyp<>+16(SB)/8, $0x8ebc6af09c88c6e3
DATA wyp<>+24(SB)/8, $0x589965cc75374cc3
GLOBL wyp<>(SB), RODATA|NOPTR, $32

// func consumeBlocks(seed uint64, b []byte) uint64
TEXT ·consumeBlocks(SB), NOSPLIT, $0-40
	MOVQ seed+0(FP), R8
	MOVQ wyp<>+0(SB), R12
	MOVQ wyp<>+8(SB), R13
	MOVQ wyp<>+16(SB), R14
	MOVQ wyp<>+24(SB), R15
	MOVQ b_base+8(FP), SI
	MOVQ b_len+16(FP), AX
	SUBQ $0x20, AX
	LEAQ (SI)(AX*1), DI
	JMP  cond

loop:
	MOVQ R8, R9
	MOVQ (SI), BX
	XORQ R12, BX
	XORQ BX, R8
	MOVQ R9, R10
	MOVQ 8(SI), BX
	XORQ R13, BX
	XORQ BX, R9
	MOVQ R8, AX
	MULQ R9
	XORQ AX, DX
	MOVQ DX, R8
	MOVQ R10, R11
	MOVQ 16(SI), BX
	XORQ R14, BX
	XORQ BX, R10
	MOVQ 24(SI), BX
	XORQ R15, BX
	XORQ BX, R11
	MOVQ R10, AX
	MULQ R11
	XORQ AX, DX
	XORQ DX, R8
	ADDQ $0x20, SI

cond:
	CMPQ SI, DI
	JLE  loop
	MOVQ R8, ret+32(FP)
	RET
