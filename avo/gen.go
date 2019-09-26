package main

import (
	. "github.com/mmcloughlin/avo/build"
	. "github.com/mmcloughlin/avo/operand"
	. "github.com/mmcloughlin/avo/reg"
)

func main() {
	wypData := GLOBL("wyp", NOPTR|RODATA)
	DATA(0, U64(0xa0761d6478bd642f))
	DATA(8, U64(0xe7037ed1a0b428db))
	DATA(16, U64(0x8ebc6af09c88c6e3))
	DATA(24, U64(0x589965cc75374cc3))

	TEXT("sum64_amd64", NOSPLIT, "func(seed uint64, b []byte) uint64")
	s1 := R8
	s2 := R9
	s3 := R10
	s4 := R11

	wyp0 := R12
	wyp1 := R13
	wyp2 := R14
	wyp3 := R15

	Load(Param("seed"), s1)
	MOVQ(wypData.Offset(0), wyp0)
	MOVQ(wypData.Offset(8), wyp1)
	MOVQ(wypData.Offset(16), wyp2)
	MOVQ(wypData.Offset(24), wyp3)

	b := Load(Param("b").Base(), RSI)
	len0 := Load(Param("b").Len(), GP64())
	SUBQ(Imm(32), len0)
	end := RDI
	LEAQ(Mem{Base: b, Index: len0, Scale: 1}, end)

	Label("loop")
	CMPQ(b, end)
	JGT(LabelRef("done"))

	MOVQ(s1, s2)
	round(Mem{Base: b, Disp: 0}, wyp0, s1)
	MOVQ(s2, s3)
	round(Mem{Base: b, Disp: 8}, wyp1, s2)
	MOVQ(s3, s4)
	round(Mem{Base: b, Disp: 16}, wyp2, s3)
	round(Mem{Base: b, Disp: 24}, wyp3, s4)

	MOVQ(s1, RAX)
	MULQ(s2)
	XORQ(RAX, RDX)
	MOVQ(RDX, s1)

	MOVQ(s3, RAX)
	MULQ(s4)
	XORQ(RAX, RDX)
	XORQ(RDX, s1)
	ADDQ(Imm(32), b)
	JMP(LabelRef("loop"))

	Label("done")
	Store(s1, ReturnIndex(0))
	RET()
	Generate()
}

func round(m Mem, p, state Register) {
	v := GP64()
	MOVQ(m, v)
	XORQ(p, v)
	XORQ(v, state)
}