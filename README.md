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
BenchmarkHash8Bytes-4     	 6575694	       178 ns/op	  44.87 MB/s
BenchmarkHash320Bytes-4   	 1710038	       700 ns/op	 457.02 MB/s
BenchmarkHash1K-4         	  760449	      1369 ns/op	 748.13 MB/s
BenchmarkHash8K-4         	  134654	      8690 ns/op	 942.69 MB/s
PASS
ok  	crypto/sha1	5.616s
```

```
$ go test -bench . github.com/orisano/wyhash
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash
BenchmarkHash8Bytes-4     	48438282	        24.6 ns/op	 324.95 MB/s
BenchmarkHash320Bytes-4   	15014091	        77.3 ns/op	4140.85 MB/s
BenchmarkHash1K-4         	 5443318	       220 ns/op	4664.48 MB/s
BenchmarkHash8K-4         	  648306	      1590 ns/op	5151.00 MB/s
PASS
ok  	github.com/orisano/wyhash	4.949s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
