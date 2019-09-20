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

	TEXT("sum64_avx", NOSPLIT, "func(seed uint64, b []byte) uint64")
	seed := Load(Param("seed"), RCX)
	b := Load(Param("b").Base(), RBX)
	len0 := Load(Param("b").Len(), RDI)
	SUBQ(Imm(32), len0)

	wyp := YMM()
	VMOVDQU(wypData, wyp)

	i := RSI
	XORQ(i, i)

	Label("loop")
	CMPQ(i, len0)
	JGT(LabelRef("done"))

	xs := XMM()
	MOVQ(seed, xs)
	ys := YMM()
	VPBROADCASTQ(xs, ys)

	block := YMM()
	VMOVDQU(Mem{Base: b}, block)

	VPXOR(wyp, ys, ys)
	VPXOR(block, ys, ys)

	buf := AllocLocal(32)
	VMOVDQU(ys, buf)

	m := GP64()
	MOVQ(buf.Offset(0), RAX)
	MOVQ(buf.Offset(8), m)
	MULQ(m)
	XORQ(RDX, RAX)
	MOVQ(RAX, seed)

	MOVQ(buf.Offset(16), RAX)
	MOVQ(buf.Offset(24), m)
	MULQ(m)
	XORQ(RDX, RAX)
	XORQ(RAX, seed)
	ADDQ(Imm(32), b)
	JMP(LabelRef("loop"))

	Label("done")
	Store(seed, ReturnIndex(0))
	RET()
	Generate()
}
