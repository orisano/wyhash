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
BenchmarkHash8Bytes-6     	 8646843	       135 ns/op	  59.19 MB/s
BenchmarkHash320Bytes-6   	 2180929	       542 ns/op	 589.94 MB/s
BenchmarkHash1K-6         	 1000000	      1105 ns/op	 926.38 MB/s
BenchmarkHash8K-6         	  160221	      7219 ns/op	1134.78 MB/s
PASS
ok  	crypto/sha1	5.416s
```

```
$ go test -bench . github.com/orisano/wyhash
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash
BenchmarkHash8Bytes-6     	170730156	         6.97 ns/op	1147.86 MB/s
BenchmarkHash320Bytes-6   	40570624	        28.6 ns/op	11176.21 MB/s
BenchmarkHash1K-6         	14049066	        84.5 ns/op	12122.42 MB/s
BenchmarkHash8K-6         	 2003247	       591 ns/op	13872.87 MB/s
PASS
ok  	github.com/orisano/wyhash	6.163s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
