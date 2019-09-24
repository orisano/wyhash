package wyhash

import (
	"encoding/binary"
	"fmt"
	"testing"
)

func ExampleNew() {
	h := New(1)
	_, _ = h.Write([]byte("a"))
	fmt.Printf("%x\n", h.Sum64())
	// Output:
	// d81367da79aa4b2
}

func ExampleSum64() {
	d := Sum64(1, []byte("a"))
	fmt.Printf("%x\n", d)
	// Output:
	// d81367da79aa4b2
}

func TestSum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0xbc98efd7661a7a1},
		{"a", 1, 0xd81367da79aa4b2},
		{"abc", 2, 0x9ab8a05305db642a},
		{"message digest", 3, 0x37320f657213a290},
		{"abcdefghijklmnopqrstuvwxyz", 4, 0xd0b270e1d8a7019c},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 5, 0x602a1894d3bbfe7f},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 6, 0x829e9c148b75970e},
	}
	for _, test := range tests {
		if got := Sum64(test.seed, []byte(test.s)); got != test.expected {
			t.Errorf("unexpected digest. expected: %x, but got: %x", test.expected, got)
		}
	}
}

func TestDigest_Sum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0xbc98efd7661a7a1},
		{"a", 1, 0xd81367da79aa4b2},
		{"abc", 2, 0x9ab8a05305db642a},
		{"message digest", 3, 0x37320f657213a290},
		{"abcdefghijklmnopqrstuvwxyz", 4, 0xd0b270e1d8a7019c},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 5, 0x602a1894d3bbfe7f},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 6, 0x829e9c148b75970e},
	}
	for _, test := range tests {
		h := New(test.seed)
		h.Write([]byte(test.s))
		if got := h.Sum64(); got != test.expected {
			t.Errorf("unexpected digest. expected: %x, but got: %x", test.expected, got)
		}
	}
}

var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, Size)
	for i := 0; i < b.N; i++ {
		binary.LittleEndian.PutUint64(sum, Sum64(0, buf[:size]))
	}
}

func BenchmarkHash8Bytes(b *testing.B) {
	benchmarkSize(b, 8)
}

func BenchmarkHash320Bytes(b *testing.B) {
	benchmarkSize(b, 320)
}

func BenchmarkHash1K(b *testing.B) {
	benchmarkSize(b, 1024)
}

func BenchmarkHash8K(b *testing.B) {
	benchmarkSize(b, 8192)
}
