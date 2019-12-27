# wyhash/v3

## Benchmark
```
$ go test -bench . -count=5
goos: darwin
goarch: amd64
pkg: github.com/orisano/wyhash/v3
BenchmarkHash8Bytes-4     	204881978	         5.75 ns/op	1391.04 MB/s
BenchmarkHash8Bytes-4     	208127184	         5.81 ns/op	1375.82 MB/s
BenchmarkHash8Bytes-4     	205134412	         5.84 ns/op	1370.35 MB/s
BenchmarkHash8Bytes-4     	208115768	         5.78 ns/op	1385.04 MB/s
BenchmarkHash8Bytes-4     	208113566	         5.75 ns/op	1392.24 MB/s
BenchmarkHash320Bytes-4   	28030105	        40.5 ns/op	7903.32 MB/s
BenchmarkHash320Bytes-4   	27906342	        40.5 ns/op	7905.20 MB/s
BenchmarkHash320Bytes-4   	27775752	        40.5 ns/op	7902.56 MB/s
BenchmarkHash320Bytes-4   	27639873	        40.5 ns/op	7899.02 MB/s
BenchmarkHash320Bytes-4   	28170300	        40.8 ns/op	7839.74 MB/s
BenchmarkHash1K-4         	12459498	        93.4 ns/op	10963.34 MB/s
BenchmarkHash1K-4         	12679180	        92.8 ns/op	11034.17 MB/s
BenchmarkHash1K-4         	12455362	        92.8 ns/op	11035.81 MB/s
BenchmarkHash1K-4         	12675535	        92.9 ns/op	11025.75 MB/s
BenchmarkHash1K-4         	12679039	        92.8 ns/op	11039.17 MB/s
BenchmarkHash8K-4         	 1714108	       697 ns/op	11758.66 MB/s
BenchmarkHash8K-4         	 1716477	       697 ns/op	11757.37 MB/s
BenchmarkHash8K-4         	 1712962	       741 ns/op	11058.60 MB/s
BenchmarkHash8K-4         	 1713458	       703 ns/op	11659.84 MB/s
BenchmarkHash8K-4         	 1714802	       702 ns/op	11665.51 MB/s
PASS
ok  	github.com/orisano/wyhash/v3	30.795s
```
