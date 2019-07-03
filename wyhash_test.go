package wyhash

import (
	"fmt"
	"testing"
)

func ExampleNew() {
	h := New(1)
	_, _ = h.Write([]byte("a"))
	fmt.Printf("%x\n", h.Sum64())
	// Output:
	// c71ba35f06089cd6
}

func ExampleSum64() {
	d := Sum64(1, []byte("a"))
	fmt.Printf("%x\n", d)
	// Output:
	// c71ba35f06089cd6
}

func TestDigest_Sum64(t *testing.T) {
	tests := []struct {
		s        string
		seed     uint64
		expected uint64
	}{
		{"", 0, 0x5f03f00e3f460a7},
		{"a", 1, 0xc71ba35f06089cd6},
		{"abc", 2, 0xcedc5099a34d885c},
		{"message digest", 3, 0x1caa6019f2274307},
		{"abcdefghijklmnopqrstuvwxyz", 4, 0xe3089173e34144d3},
		{"ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789", 5, 0xed4731baa4233e09},
		{"12345678901234567890123456789012345678901234567890123456789012345678901234567890", 6, 0xb6ec251785c0d299},
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
