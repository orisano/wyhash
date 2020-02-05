package v4

import (
	"encoding/binary"
	"math/bits"
)

const (
	wyp0 = 0xa0761d6478bd642f
	wyp1 = 0xe7037ed1a0b428db
	wyp2 = 0x8ebc6af09c88c6e3
	wyp3 = 0x589965cc75374cc3
	wyp4 = 0x1d8e4e27c47d124f
)

// The size of a wyhash checksum in bytes.
const Size = 8

func Sum64(seed uint64, key []byte) uint64 {
	return sum64(key, uint64(len(key)), seed)
}

func sum64(key []byte, len, seed uint64) uint64 {
	p := key
	i := len
rest:
	if i < 4 {
		w := uint64(0)
		if i != 0 {
			w = wyr3(p, i)
		}
		return wymum(wymum(w^seed^wyp0, seed^wyp1), len^wyp4)
	} else if i <= 8 {
		return wymum(wymum(wyr4(p)^seed^wyp0, wyr4(p[i-4:])^seed^wyp1), len^wyp4)
	} else if i <= 16 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[i-8:])^seed^wyp1), len^wyp4)
	} else if i <= 32 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)^wymum(wyr8(p[i-16:])^seed^wyp2, wyr8(p[i-8:])^seed^wyp3), len^wyp4)
	} else if i <= 64 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)^wymum(wyr8(p[16:])^seed^wyp2, wyr8(p[24:])^seed^wyp3)^wymum(wyr8(p[i-32:])^seed^wyp1, wyr8(p[i-24:])^seed^wyp2)^wymum(wyr8(p[i-16:])^seed^wyp3, wyr8(p[i-8:])^seed^wyp0), len^wyp4)
	}
	see1 := seed
	see2 := seed
	see3 := seed
	for ; i >= 64; i -= 64 {
		seed = wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)
		see1 = wymum(wyr8(p[16:])^see1^wyp2, wyr8(p[24:])^see1^wyp3)
		see2 = wymum(wyr8(p[32:])^see2^wyp1, wyr8(p[40:])^see2^wyp2)
		see3 = wymum(wyr8(p[48:])^see3^wyp3, wyr8(p[56:])^see3^wyp0)
		p = p[64:]
	}
	seed ^= see1 ^ see2 ^ see3
	goto rest
}

func wyr3(p []byte, k uint64) uint64 {
	return (uint64(p[0]) << 16) | (uint64(p[k>>1]) << 8) | uint64(p[k-1])
}

func wyr4(p []byte) uint64 {
	return uint64(binary.LittleEndian.Uint32(p))
}

func wyr8(p []byte) uint64 {
	return binary.LittleEndian.Uint64(p)
}

func wymum(a, b uint64) uint64 {
	hi, lo := bits.Mul64(a, b)
	return hi ^ lo
}
