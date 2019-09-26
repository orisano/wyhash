# wyhash

[![CircleCI](https://circleci.com/gh/orisano/wyhash.svg?style=svg)](https://circleci.com/gh/orisano/wyhash)
[![GoDoc](https://godoc.org/github.com/orisano/wyhash?status.svg)](https://godoc.org/github.com/orisano/wyhash)

A pure-Go wyhash implementation.

## How to use
```go
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
BenchmarkHash8Bytes-6     	 8648186	       137 ns/op	  58.61 MB/s
BenchmarkHash320Bytes-6   	 2157961	       555 ns/op	 576.73 MB/s
BenchmarkHash1K-6         	 1000000	      1129 ns/op	 907.18 MB/s
BenchmarkHash8K-6         	  161910	      7351 ns/op	1114.38 MB/s
PASS
ok  	crypto/sha1	5.505s
```

```
$ go test -bench . github.com/orisano/wyhash
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash
BenchmarkHash8Bytes-6            	213553929	         5.58 ns/op	1432.76 MB/s
BenchmarkHash320Bytes-6          	46074423	        25.5 ns/op	12543.20 MB/s
BenchmarkHash1K-6                	15385605	        77.0 ns/op	13297.57 MB/s
BenchmarkHash8K-6                	 2130975	       561 ns/op	14600.55 MB/s
BenchmarkDigest_Hash8Bytes-6     	76109239	        15.3 ns/op	 521.48 MB/s
BenchmarkDigest_Hash320Bytes-6   	35689680	        33.1 ns/op	9666.85 MB/s
BenchmarkDigest_Hash1K-6         	14600259	        81.1 ns/op	12627.28 MB/s
BenchmarkDigest_Hash8K-6         	 2078782	       575 ns/op	14254.97 MB/s
PASS
ok  	github.com/orisano/wyhash	11.450s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
