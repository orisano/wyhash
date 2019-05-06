package wyhash

import (
	"testing"
)

func TestDigest_Sum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0xf961f936e29c9345},
		{"a", 1, 0x6dc395f88b363baa},
		{"abc", 2, 0x3bc9d7844798ddaa},
		{"message digest", 3, 0xb31238dc2c500cd3},
		{"abcdefghijklmnopqrstuvwxyz", 4, 0xea0f542c58cddfe4},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 5, 0x1799aca591fe73b4},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 6, 0x7f0d02f53d64c1f9},
	}
	for _, test := range tests {
		h := New(test.seed)
		_, _ = h.Write([]byte(test.s))
		if got := h.Sum64(); got != test.expected {
			t.Errorf("unexpected digest. expected: %x, but got: %x", test.expected, got)
		}
	}
}

var bench = New(0)
var buf = make([]byte, 8192)

func benchmarkSize(b *testing.B, size int) {
	b.SetBytes(int64(size))
	sum := make([]byte, bench.Size())
	for i := 0; i < b.N; i++ {
		bench.Reset()
		bench.Write(buf[:size])
		bench.Sum(sum[:0])
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
