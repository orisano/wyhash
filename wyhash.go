// Package wyhash implements the wyhash hash algorithm as defined in github.com/wangyi-fudan/wyhash
package wyhash

//go:generate go run ./avo/gen.go -out blocks_amd64.s -stubs blocks_amd64.go

import (
	"encoding/binary"
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
	return sum64(seed, b, uint64(len(b)))
}

func sum64(seed uint64, b []byte, len1 uint64) uint64 {
	p := b
	if len(p) >= 32 {
		seed = consumeBlocks(seed, p)
		p = p[len(p) & ^(BlockSize-1):]
	}
	switch len(p) {
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

type digest struct {
	seed  uint64
	state uint64
	size  int
	buf   []byte
}

func read64(b []byte) uint64 {
	return bits.RotateLeft64(binary.LittleEndian.Uint64(b), 32)
}

func read32(b []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(b))
}

func read16(b []byte) uint64 {
	return uint64(binary.LittleEndian.Uint16(b))
}

func read8(b []byte) uint64 {
	return uint64(b[0])
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
	if len(p) >= BlockSize {
		seed = consumeBlocks(seed, p)
		p = p[len(p) & ^(BlockSize-1):]
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
	return sum64(d.state, d.buf, uint64(d.size))
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
	p1 := binary.LittleEndian.Uint64(b[0:])
	p2 := binary.LittleEndian.Uint64(b[8:])
	p3 := binary.LittleEndian.Uint64(b[16:])
	p4 := binary.LittleEndian.Uint64(b[24:])
	return mix0(p1, p2, seed) ^ mix1(p3, p4, seed)
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
