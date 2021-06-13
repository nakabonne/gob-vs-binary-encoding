
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

```
$ go run main.go
encoding/binary encoded binary length: 16
encoding/binary encoded binary: [154 153 153 153 153 153 185 63 1 0 0 0 0 0 0 0]
encoding/binary decoded point: &main.DataPoint{Value:0.1, Timestamp:1}

encoding/gob encoded binary length: 64
encoding/gob encoded binary: [47 255 129 3 1 1 9 68 97 116 97 80 111 105 110 116 1 255 130 0 1 2 1 5 86 97 108 117 101 1 8 0 1 9 84 105 109 101 115 116 97 109 112 1 4 0 0 0 15 255 130 1 248 154 153 153 153 153 153 185 63 1 2 0]
encoding/gob decoded point: &main.DataPoint{Value:0.1, Timestamp:1}
```
