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

	"github.com/orisano/wyhash/v2"
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
$ go test -bench . -count=5
goos: linux
goarch: amd64
pkg: github.com/orisano/wyhash/v2
BenchmarkHash8Bytes-36             	239112354	         5.02 ns/op	1593.40 MB/s
BenchmarkHash8Bytes-36             	238774428	         5.02 ns/op	1593.65 MB/s
BenchmarkHash8Bytes-36             	239230683	         5.02 ns/op	1593.35 MB/s
BenchmarkHash8Bytes-36             	233548692	         5.03 ns/op	1589.73 MB/s
BenchmarkHash8Bytes-36             	240766962	         5.01 ns/op	1596.49 MB/s
BenchmarkHash320Bytes-36           	39866637	        30.2 ns/op	10602.81 MB/s
BenchmarkHash320Bytes-36           	38263518	        30.1 ns/op	10626.23 MB/s
BenchmarkHash320Bytes-36           	39905398	        30.2 ns/op	10607.83 MB/s
BenchmarkHash320Bytes-36           	39833031	        30.1 ns/op	10632.60 MB/s
BenchmarkHash320Bytes-36           	39827767	        30.1 ns/op	10648.74 MB/s
BenchmarkHash1K-36                 	13151994	        91.5 ns/op	11191.01 MB/s
BenchmarkHash1K-36                 	13109916	        91.5 ns/op	11187.74 MB/s
BenchmarkHash1K-36                 	13137547	        91.6 ns/op	11183.96 MB/s
BenchmarkHash1K-36                 	13094876	        91.6 ns/op	11173.48 MB/s
BenchmarkHash1K-36                 	13119465	        91.7 ns/op	11170.75 MB/s
BenchmarkHash8K-36                 	 1818872	       662 ns/op	12378.54 MB/s
BenchmarkHash8K-36                 	 1815982	       657 ns/op	12475.84 MB/s
BenchmarkHash8K-36                 	 1810904	       675 ns/op	12143.65 MB/s
BenchmarkHash8K-36                 	 1790918	       670 ns/op	12232.80 MB/s
BenchmarkHash8K-36                 	 1815848	       660 ns/op	12407.03 MB/s
BenchmarkDigest_Hash8Bytes-36      	75824784	        15.8 ns/op	 504.92 MB/s
BenchmarkDigest_Hash8Bytes-36      	75778485	        15.8 ns/op	 505.87 MB/s
BenchmarkDigest_Hash8Bytes-36      	75971301	        16.1 ns/op	 496.39 MB/s
BenchmarkDigest_Hash8Bytes-36      	73296864	        15.8 ns/op	 506.77 MB/s
BenchmarkDigest_Hash8Bytes-36      	75797804	        15.8 ns/op	 507.25 MB/s
BenchmarkDigest_Hash320Bytes-36    	32235518	        37.4 ns/op	8567.48 MB/s
BenchmarkDigest_Hash320Bytes-36    	31975928	        38.2 ns/op	8384.57 MB/s
BenchmarkDigest_Hash320Bytes-36    	32122788	        37.5 ns/op	8544.50 MB/s
BenchmarkDigest_Hash320Bytes-36    	31351216	        37.5 ns/op	8542.51 MB/s
BenchmarkDigest_Hash320Bytes-36    	32137383	        37.5 ns/op	8522.40 MB/s
BenchmarkDigest_Hash1K-36          	12877767	        92.9 ns/op	11020.89 MB/s
BenchmarkDigest_Hash1K-36          	12920068	        92.8 ns/op	11028.94 MB/s
BenchmarkDigest_Hash1K-36          	12926998	        92.9 ns/op	11024.32 MB/s
BenchmarkDigest_Hash1K-36          	12930212	        93.1 ns/op	10996.36 MB/s
BenchmarkDigest_Hash1K-36          	12938821	        93.0 ns/op	11012.81 MB/s
BenchmarkDigest_Hash8K-36          	 1798011	       675 ns/op	12134.26 MB/s
BenchmarkDigest_Hash8K-36          	 1794015	       668 ns/op	12258.51 MB/s
BenchmarkDigest_Hash8K-36          	 1796629	       668 ns/op	12261.52 MB/s
BenchmarkDigest_Hash8K-36          	 1795533	       669 ns/op	12251.80 MB/s
BenchmarkDigest_Hash8K-36          	 1793821	       664 ns/op	12329.72 MB/s
PASS
ok  	github.com/orisano/wyhash/v2	58.638s
```

## Author
Nao YONASHIRO (@orisano)

## License
MIT

## Reference
* [wangyi-fudan/wyhash](https://github.com/wangyi-fudan/wyhash)
