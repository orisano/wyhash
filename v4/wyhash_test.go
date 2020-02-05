package v4

import (
	"encoding/binary"
	"testing"
)

func TestSum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0xbc98efd7661a7a1},
		{"a", 1, 0x99782e84a7cee30},
		{"abc", 2, 0x973ed17dfbe006d7},
		{"message digest", 3,0xc0189aa4012331f5},
		{"abcdefghijklmnopqrstuvwxyz", 4, 0xda133f940b62e516},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 5, 0xe062dfda99413626},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 6, 0x77092dd38803d1fa},
	}
	for _, test := range tests {
		if got := Sum64(test.seed, []byte(test.s)); got != test.expected {
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