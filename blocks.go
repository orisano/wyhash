// +build !amd64

package wyhash

import (
	"encoding/binary"
)

func consumeBlocks(seed uint64, b []byte) uint64 {
	for len(b) > BlockSize {
		p := b[:BlockSize]
		p1 := binary.LittleEndian.Uint64(p[0:])
		p2 := binary.LittleEndian.Uint64(p[8:])
		p3 := binary.LittleEndian.Uint64(p[16:])
		p4 := binary.LittleEndian.Uint64(p[24:])
		seed = mix0(p1, p2, seed) ^ mix1(p3, p4, seed)
		b = b[BlockSize:]
	}
	return seed
}