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
	p1 := R8
	p2 := R9
	p3 := R10
	p4 := R11

	wyp0 := R12
	wyp1 := R13
	wyp2 := R14
	wyp3 := R15

	Load(Param("seed"), p1)
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

	MOVQ(p1, p2)
	MOVQ(p1, p3)
	MOVQ(p1, p4)

	round(p1, b, wyp0)
	round(p2, b, wyp1)
	round(p3, b, wyp2)
	round(p4, b, wyp3)

	MOVQ(p1, RAX)
	MULQ(p2)
	XORQ(RAX, RDX)
	MOVQ(RDX, p1)

	MOVQ(p3, RAX)
	MULQ(p4)
	XORQ(RAX, RDX)
	XORQ(RDX, p1)
	JMP(LabelRef("loop"))

	Label("done")
	Store(p1, ReturnIndex(0))
	RET()
	Generate()
}

func round(state, base, p Register) {
	v := GP64()
	MOVQ(Mem{Base: base}, v)
	ADDQ(Imm(8), base)
	XORQ(p, v)
	XORQ(v, state)
}