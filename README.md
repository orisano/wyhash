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
BenchmarkHash8Bytes-6     	 8641054	       138 ns/op	  58.14 MB/s
BenchmarkHash320Bytes-6   	 2140016	       556 ns/op	 575.52 MB/s
BenchmarkHash1K-6         	 1000000	      1132 ns/op	 904.90 MB/s
BenchmarkHash8K-6         	  158954	      7407 ns/op	1105.98 MB/s
PASS
ok  	crypto/sha1	5.502s
```

```
$ go test -bench . github.com/orisano/wyhash
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash
BenchmarkHash8Bytes-6     	165023144	         7.26 ns/op	1101.54 MB/s
BenchmarkHash320Bytes-6   	37248912	        31.8 ns/op	10047.87 MB/s
BenchmarkHash1K-6         	12369891	        96.0 ns/op	10669.18 MB/s
BenchmarkHash8K-6         	 1659277	       719 ns/op	11387.74 MB/s
PASS
ok  	github.com/orisano/wyhash	6.371s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
