# wyhash

[![CircleCI](https://circleci.com/gh/orisano/wyhash.svg?style=svg)](https://circleci.com/gh/orisano/wyhash)
[![GoDoc](https://godoc.org/github.com/orisano/wyhash?status.svg)](https://godoc.org/github.com/orisano/wyhash)

A pure-Go wyhash implementation.

## How to use
```go
package main

import (
	"flag"
	"fmt"
	"io"
	"os"

	"github.com/orisano/wyhash"
)

func main() {
	seed := flag.Int("s", 0, "seed value")
	flag.Parse()

	h := wyhash.New(uint64(*seed))
	io.Copy(h, os.Stdin)
	fmt.Printf("%x\n", h.Sum64())
}
```

## Benchmark
```
$ go test -bench . crypto/sha1
goos: darwin
goarch: amd64
pkg: crypto/sha1
BenchmarkHash8Bytes-6     	 8644443	       136 ns/op	  58.67 MB/s
BenchmarkHash320Bytes-6   	 2153529	       555 ns/op	 576.73 MB/s
BenchmarkHash1K-6         	 1000000	      1131 ns/op	 905.64 MB/s
BenchmarkHash8K-6         	  159200	      7339 ns/op	1116.22 MB/s
PASS
ok  	crypto/sha1	5.483s
```

```
$ go test -bench . github.com/orisano/wyhash
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash
BenchmarkHash8Bytes-6            	259621784	         4.60 ns/op	1739.40 MB/s
BenchmarkHash320Bytes-6          	45866812	        25.6 ns/op	12488.59 MB/s
BenchmarkHash1K-6                	15543504	        76.2 ns/op	13430.11 MB/s
BenchmarkHash8K-6                	 2138480	       559 ns/op	14649.56 MB/s
BenchmarkDigest_Hash8Bytes-6     	78949834	        14.7 ns/op	 543.64 MB/s
BenchmarkDigest_Hash320Bytes-6   	37010572	        32.0 ns/op	9993.93 MB/s
BenchmarkDigest_Hash1K-6         	14721690	        79.7 ns/op	12846.16 MB/s
BenchmarkDigest_Hash8K-6         	 2085654	       573 ns/op	14302.17 MB/s
PASS
ok  	github.com/orisano/wyhash	11.331s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
