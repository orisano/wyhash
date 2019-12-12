package v3

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
	if len == 0 {
		return 0
	}
	p := key
	if len < 4 {
		return wymum(wymum(wyr3(p, len)^seed^wyp0, seed^wyp1), len^wyp4)
	} else if len <= 8 {
		return wymum(wymum(wyr4(p)^seed^wyp0, wyr4(p[len-4:])^seed^wyp1), len^wyp4)
	} else if len <= 16 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[len-8:])^seed^wyp1), len^wyp4)
	} else if len <= 24 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)^wymum(wyr8(p[len-8:])^seed^wyp2, seed^wyp3), len^wyp4)
	} else if len <= 32 {
		return wymum(wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)^wymum(wyr8(p[16:])^seed^wyp2, wyr8(p[len-8:])^seed^wyp3), len^wyp4)
	}
	return sumLarge(p, len, seed)
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

func sumLarge(p []byte, len, seed uint64) uint64 {
	see1 := seed
	i := len
	if i >= 256 {
		var (
			wyp0 uint64 = wyp0
			wyp1 uint64 = wyp1
			wyp2 uint64 = wyp2
			wyp3 uint64 = wyp3
		)
		for ; i >= 256; i -= 256 {
			b := p[:256]
			seed = wymum(wyr8(b[0:])^seed^wyp0, wyr8(b[8:])^seed^wyp1) ^ wymum(wyr8(b[16:])^seed^wyp2, wyr8(b[24:])^seed^wyp3)
			see1 = wymum(wyr8(b[32:])^see1^wyp1, wyr8(b[40:])^see1^wyp2) ^ wymum(wyr8(b[48:])^see1^wyp3, wyr8(b[56:])^see1^wyp0)
			seed = wymum(wyr8(b[64:])^seed^wyp0, wyr8(b[72:])^seed^wyp1) ^ wymum(wyr8(b[80:])^seed^wyp2, wyr8(b[88:])^seed^wyp3)
			see1 = wymum(wyr8(b[96:])^see1^wyp1, wyr8(b[104:])^see1^wyp2) ^ wymum(wyr8(b[112:])^see1^wyp3, wyr8(b[120:])^see1^wyp0)
			seed = wymum(wyr8(b[128:])^seed^wyp0, wyr8(b[136:])^seed^wyp1) ^ wymum(wyr8(b[144:])^seed^wyp2, wyr8(b[152:])^seed^wyp3)
			see1 = wymum(wyr8(b[160:])^see1^wyp1, wyr8(b[168:])^see1^wyp2) ^ wymum(wyr8(b[176:])^see1^wyp3, wyr8(b[184:])^see1^wyp0)
			seed = wymum(wyr8(b[192:])^seed^wyp0, wyr8(b[200:])^seed^wyp1) ^ wymum(wyr8(b[208:])^seed^wyp2, wyr8(b[216:])^seed^wyp3)
			see1 = wymum(wyr8(b[224:])^see1^wyp1, wyr8(b[232:])^see1^wyp2) ^ wymum(wyr8(b[240:])^see1^wyp3, wyr8(b[248:])^see1^wyp0)
			p = p[256:]
		}
	}

	for ; i >= 32; i -= 32 {
		seed = wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)
		see1 = wymum(wyr8(p[16:])^see1^wyp2, wyr8(p[24:])^see1^wyp3)
		p = p[32:]
	}

	if i == 0 {
	} else if i < 4 {
		seed = wymum(wyr3(p, i)^seed^wyp0, seed^wyp1)
	} else if i <= 8 {
		seed = wymum(wyr4(p)^seed^wyp0, wyr4(p[i-4:])^seed^wyp1)
	} else if i <= 16 {
		seed = wymum(wyr8(p)^seed^wyp0, wyr8(p[i-8:])^seed^wyp1)
	} else if i <= 24 {
		seed = wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)
		see1 = wymum(wyr8(p[i-8:])^see1^wyp2, see1^wyp3)
	} else {
		seed = wymum(wyr8(p)^seed^wyp0, wyr8(p[8:])^seed^wyp1)
		see1 = wymum(wyr8(p[16:])^see1^wyp2, wyr8(p[i-8:])^see1^wyp3)
	}
	return wymum(seed^see1, len^wyp4)
}