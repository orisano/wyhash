package wyhash

import (
	"hash"
)

const (
	wyp0 uint64 = 0xa0761d6478bd642f
	wyp1 uint64 = 0xe7037ed1a0b428db
	wyp2 uint64 = 0x8ebc6af09c88c6e3
	wyp3 uint64 = 0x589965cc75374cc3
	wyp4 uint64 = 0x1d8e4e27c47d124f
	wyp5 uint64 = 0xeb44accab455d165
)

const (
	blockSize = 32
)

func New() hash.Hash {
	d := new(digest)
	d.Reset()
	return d
}

func mum(a, b uint64) uint64 {
	ha, la := a>>32, a&0xffffffff
	hb, lb := b>>32, b&0xffffffff
	rh := ha * hb
	rm0 := ha * lb
	rm1 := hb * la
	rl := la * lb
	t := rl + (rm0 << 32)
	c := uint64(0)
	if t < rl {
		c++
	}
	lo := t + (rm1 << 32)
	if lo < t {
		c++
	}
	hi := rh + (rm0 >> 32) + c
	return hi ^ lo
}

type digest struct {
	seed uint64
	size int
	buf  []byte
}

func mread64(b []byte) uint64 {
	return read32(b)<<32 | read32(b[4:])
}

func read64(b []byte) uint64 {
	_ = b[7]
	x := uint64(b[7]) | uint64(b[6])<<8 | uint64(b[5])<<16 | uint64(b[4])<<24 |
		uint64(b[3])<<32 | uint64(b[2])<<40 | uint64(b[1])<<48 | uint64(b[0])<<56
	return x
}

func read32(b []byte) uint64 {
	_ = b[3]
	x := uint32(b[3]) | uint32(b[2])<<8 | uint32(b[1])<<16 | uint32(b[0])<<24
	return uint64(x)
}

func read16(b []byte) uint64 {
	_ = b[1]
	x := uint16(b[1]) | uint16(b[0])<<8
	return uint64(x)
}

func read8(b []byte) uint64 {
	x := uint8(b[0])
	return uint64(x)
}

func (d *digest) Write(p []byte) (int, error) {
	seed := d.seed
	n := len(p)
	buffered := len(d.buf)
	if buffered > 0 {
		rest := blockSize - buffered
		if len(p) < rest {
			d.buf = append(d.buf, p...)
			return n, nil
		}
		d.buf = append(d.buf, p[:rest]...)
		seed = consumeBlock(seed, d.buf)
		d.buf = d.buf[:0]
		p = p[rest:]
	}
	for i, size := 0, len(p); i+32 <= size; i += 32 {
		seed = consumeBlock(seed, p)
		p = p[32:]
	}
	if len(p) != 0 {
		d.buf = append(d.buf, p...)
	}
	d.seed = seed
	d.size += n
	return n, nil
}

func (d *digest) Sum(b []byte) []byte {
	x := d.Sum64()
	return append(b, byte((x>>56)&0xff), byte((x>>48)&0xff), byte((x>>40)&0xff), byte((x>>32)&0xff), byte((x>>24)&0xff), byte((x>>16)&0xff), byte((x>>8)&0xff), byte(x&0xff))
}

func (d *digest) Sum64() uint64 {
	seed := d.seed
	seed ^= wyp0
	p := d.buf
	switch d.size & 31 {
	case 1:
		seed = mum(seed, read8(p)^wyp1)
	case 2:
		seed = mum(seed, read16(p)^wyp1)
	case 3:
		seed = mum(seed, (read16(p)<<8|read8(p[2:]))^wyp1)
	case 4:
		seed = mum(seed, read32(p)^wyp1)
	case 5:
		seed = mum(seed, (read32(p)<<8|read8(p[4:]))^wyp1)
	case 6:
		seed = mum(seed, (read32(p)<<16|read16(p[4:]))^wyp1)
	case 7:
		seed = mum(seed, (read32(p)<<24|read16(p[4:])<<8|read8(p[6:]))^wyp1)
	case 8:
		seed = mum(seed, mread64(p)^wyp1)
	case 9:
		seed = mum(mread64(p)^seed, read8(p[8:])^wyp2)
	case 10:
		seed = mum(mread64(p)^seed, read16(p[8:])^wyp2)
	case 11:
		seed = mum(mread64(p)^seed, ((read16(p[8:])<<8)|read8(p[10:]))^wyp2)
	case 12:
		seed = mum(mread64(p)^seed, read32(p[8:])^wyp2)
	case 13:
		seed = mum(mread64(p)^seed, ((read32(p[8:])<<8)|read8(p[12:]))^wyp2)
	case 14:
		seed = mum(mread64(p)^seed, ((read32(p[8:])<<16)|read16(p[12:]))^wyp2)
	case 15:
		seed = mum(mread64(p)^seed, ((read32(p[8:])<<24)|(read16(p[12:])<<8)|read8(p[14:]))^wyp2)
	case 16:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2)
	case 17:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, read8(p[16:])^wyp3)
	case 18:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, read16(p[16:])^wyp3)
	case 19:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, (read16(p[16:])<<8)|read8(p[18:])^wyp3)
	case 20:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, read32(p[16:])^wyp3)
	case 21:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, (read32(p[16:])<<8)|read8(p[20:])^wyp3)
	case 22:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, (read32(p[16:])<<16)|read16(p[20:])^wyp3)
	case 23:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, (read32(p[16:])<<24)|(read16(p[20:])<<8)|read8(p[22:])^wyp3)
	case 24:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(seed, mread64(p[16:])^wyp3)
	case 25:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, read8(p[24:])^wyp4)
	case 26:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, read16(p[24:])^wyp4)
	case 27:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, ((read16(p[24:])<<8)|read8(p[26:]))^wyp4)
	case 28:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, read32(p[24:])^wyp4)
	case 29:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, ((read32(p[24:])<<8)|read8(p[28:]))^wyp4)
	case 30:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, ((read32(p[24:])<<16)|read16(p[28:]))^wyp4)
	case 31:
		seed = mum(mread64(p)^seed, mread64(p[8:])^wyp2) ^ mum(mread64(p[16:])^seed, ((read32(p[24:])<<24)|(read16(p[28:])<<8)|read8(p[30:]))^wyp4)
	}
	return mum(seed, uint64(d.size)^wyp5)
}

func (d *digest) Reset() {
	d.seed = 0
	d.size = 0
	if d.buf != nil {
		d.buf = d.buf[:0]
	}
}

func (d *digest) Size() int {
	return 8
}

func (d *digest) BlockSize() int {
	return blockSize
}

func consumeBlock(seed uint64, p []byte) uint64 {
	return mum(seed^wyp0, mum(read64(p)^wyp1, read64(p[8:])^wyp2)^mum(read64(p[16:])^wyp3, read64(p[24:])^wyp4))
}
