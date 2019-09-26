// +build !amd64

package wyhash

func consumeBlocks(seed uint64, b []byte) uint64 {
	for i, size := 0, len(b); i + BlockSize <= size; i += BlockSize {
		seed = consumeBlock(seed, b)
		b = b[BlockSize:]
	}
	return seed
}