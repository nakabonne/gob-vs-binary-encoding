
```
$ go test -benchmem -bench=. .
goos: darwin
goarch: amd64
pkg: github.com/nakabonne/hoge
cpu: Intel(R) Core(TM) i5-7267U CPU @ 3.10GHz
BenchmarkEncodeDecodeWithBinary-4        3364623               358.8 ns/op           160 B/op          5 allocs/op
BenchmarkEncodeDecodeWithGob-4             59617             20889 ns/op            7856 B/op        197 allocs/op
PASS
ok      github.com/nakabonne/hoge       3.376s
```
