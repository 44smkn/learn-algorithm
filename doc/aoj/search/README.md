# 探索

## 線形探索

```sh
$ go test -bench  . -benchmem                             
goos: darwin
goarch: amd64
pkg: linear-serarch
BenchmarkSearchBySentinel-4     2000000000           0.01 ns/op        0 B/op          0 allocs/op
BenchmarkSearchByRange-4        1000000000           0.08 ns/op        0 B/op          0 allocs/op
PASS
ok  linear-serarch  1.100s
```
