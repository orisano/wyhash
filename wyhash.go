// Package wyhash implements the wyhash hash algorithm as defined in github.com/wangyi-fudan/wyhash
package wyhash

import (
	"hash"
	"math/bits"
)

const (
	wyp0 uint64 = 0xa0761d6478bd642f
	wyp1 uint64 = 0xe7037ed1a0b428db
	wyp2 uint64 = 0x8ebc6af09c88c6e3
	wyp3 uint64 = 0x589965cc75374cc3
	wyp4 uint64 = 0x1d8e4e27c47d124f
)

// The blocksize of wyhash in bytes.
const BlockSize = 32

// The size of a wyhash checksum in bytes.
const Size = 8

// New returns a new hash.Hash64 computing the wyhash checksum.
func New(seed uint64) hash.Hash64 {
	d := &digest{seed: seed}
	d.Reset()
	return d
}

// Sum64 returns the wyhash checksum of the b
func Sum64(seed uint64, b []byte) uint64 {
	h := New(seed)
	_, _ = h.Write(b)
	return h.Sum64()
}

func mum(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}

func mix0(a, b, seed uint64) uint64 {
	return mum(a^seed^wyp0, b^seed^wyp1)
}

func mix1(a, b, seed uint64) uint64 {
	return mum(a^seed^wyp2, b^seed^wyp3)
}

type digest struct {
	seed  uint64
	state uint64
	size  int
	buf   []byte
}

func read64(b []byte) uint64 {
	return read32(b)<<32 | read32(b[4:])
}

func read32(b []byte) uint64 {
	_ = b[3]
	x := uint32(b[0]) | uint32(b[1])<<8 | uint32(b[2])<<16 | uint32(b[3])<<24
	return uint64(x)
}

func read16(b []byte) uint64 {
	_ = b[1]
	x := uint16(b[0]) | uint16(b[1])<<8
	return uint64(x)
}

func read8(b []byte) uint64 {
	x := uint8(b[0])
	return uint64(x)
}

func (d *digest) Write(p []byte) (int, error) {
	seed := d.state
	n := len(p)
	d.size += n

	buffered := len(d.buf)
	if buffered > 0 {
		rest := BlockSize - buffered
		if len(p) < rest {
			d.buf = append(d.buf, p...)
			return n, nil
		}
		d.buf = append(d.buf, p[:rest]...)
		seed = consumeBlock(seed, d.buf)
		d.buf = d.buf[:0]
		p = p[rest:]
	}
	for i, size := 0, len(p); i+BlockSize <= size; i += BlockSize {
		seed = consumeBlock(seed, p)
		p = p[BlockSize:]
	}
	if len(p) != 0 {
		d.buf = append(d.buf, p...)
	}
	d.state = seed
	return n, nil
}

func (d *digest) Sum(b []byte) []byte {
	x := d.Sum64()
	return append(b, byte((x>>56)&0xff), byte((x>>48)&0xff), byte((x>>40)&0xff), byte((x>>32)&0xff), byte((x>>24)&0xff), byte((x>>16)&0xff), byte((x>>8)&0xff), byte(x&0xff))
}

func (d *digest) Sum64() uint64 {
	seed := d.state
	seed ^= wyp0
	p := d.buf
	len1 := uint64(d.size)
	switch d.size & (BlockSize - 1) {
	case 0:
		len1 = mix0(len1, 0, seed)
	case 1:
		seed = mix0(read8(p), 0, seed)
	case 2:
		seed = mix0(read16(p), 0, seed)
	case 3:
		seed = mix0(read16(p)<<8|read8(p[2:]), 0, seed)
	case 4:
		seed = mix0(read32(p), 0, seed)
	case 5:
		seed = mix0(read32(p)<<8|read8(p[4:]), 0, seed)
	case 6:
		seed = mix0(read32(p)<<16|read16(p[4:]), 0, seed)
	case 7:
		seed = mix0(read32(p)<<24|read16(p[4:])<<8|read8(p[6:]), 0, seed)
	case 8:
		seed = mix0(read64(p), 0, seed)
	case 9:
		seed = mix0(read64(p), read8(p[8:]), seed)
	case 10:
		seed = mix0(read64(p), read16(p[8:]), seed)
	case 11:
		seed = mix0(read64(p), (read16(p[8:])<<8)|read8(p[10:]), seed)
	case 12:
		seed = mix0(read64(p), read32(p[8:]), seed)
	case 13:
		seed = mix0(read64(p), (read32(p[8:])<<8)|read8(p[12:]), seed)
	case 14:
		seed = mix0(read64(p), (read32(p[8:])<<16)|read16(p[12:]), seed)
	case 15:
		seed = mix0(read64(p), (read32(p[8:])<<24)|(read16(p[12:])<<8)|read8(p[14:]), seed)
	case 16:
		seed = mix0(read64(p), read64(p[8:]), seed)
	case 17:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read8(p[16:]), 0, seed)
	case 18:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read16(p[16:]), 0, seed)
	case 19:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1((read16(p[16:])<<8)|read8(p[18:]), 0, seed)
	case 20:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read32(p[16:]), 0, seed)
	case 21:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1((read32(p[16:])<<8)|read8(p[20:]), 0, seed)
	case 22:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1((read32(p[16:])<<16)|read16(p[20:]), 0, seed)
	case 23:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1((read32(p[16:])<<24)|(read16(p[20:])<<8)|read8(p[22:]), 0, seed)
	case 24:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), 0, seed)
	case 25:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), read8(p[24:]), seed)
	case 26:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), read16(p[24:]), seed)
	case 27:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), (read16(p[24:])<<8)|read8(p[26:]), seed)
	case 28:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), read32(p[24:]), seed)
	case 29:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), (read32(p[24:])<<8)|read8(p[28:]), seed)
	case 30:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), (read32(p[24:])<<16)|read16(p[28:]), seed)
	case 31:
		seed = mix0(read64(p), read64(p[8:]), seed) ^ mix1(read64(p[16:]), (read32(p[24:])<<24)|(read16(p[28:])<<8)|read8(p[30:]), seed)
	}
	return mum(seed^len1, wyp4)
}

func (d *digest) Reset() {
	d.state = d.seed
	d.size = 0
	if d.buf != nil {
		d.buf = d.buf[:0]
	}
}

func (d *digest) Size() int {
	return Size
}

func (d *digest) BlockSize() int {
	return BlockSize
}

func consumeBlock(seed uint64, b []byte) uint64 {
	_ = b[31]
	p1 := uint64(b[0]) | uint64(b[1])<<8 | uint64(b[2])<<16 | uint64(b[3])<<24 |
		uint64(b[4])<<32 | uint64(b[5])<<40 | uint64(b[6])<<48 | uint64(b[7])<<56
	p2 := uint64(b[8]) | uint64(b[9])<<8 | uint64(b[10])<<16 | uint64(b[11])<<24 |
		uint64(b[12])<<32 | uint64(b[13])<<40 | uint64(b[14])<<48 | uint64(b[15])<<56
	p3 := uint64(b[16]) | uint64(b[17])<<8 | uint64(b[18])<<16 | uint64(b[19])<<24 |
		uint64(b[20])<<32 | uint64(b[21])<<40 | uint64(b[22])<<48 | uint64(b[23])<<56
	p4 := uint64(b[24]) | uint64(b[25])<<8 | uint64(b[26])<<16 | uint64(b[27])<<24 |
		uint64(b[28])<<32 | uint64(b[29])<<40 | uint64(b[30])<<48 | uint64(b[31])<<56
	return mix0(p1, p2, seed) ^ mix1(p3, p4, seed)
}
