# Metrics Optimization

Optimization exercise on an example metrics program which calculates averages and standard deviation across a mock dataset of users and their payments.

* `metrics_baseline.go` is the program as given.

* `metrics_array.go` is an optimized version that uses an array-based vs object-oriented approach.

* `metrics_streaming.go` is a second optimized version that uses a single pass through the source files.

To run benchmarks:

```sh
➜  go test -bench=. -benchmem metrics_baseline.go metrics_test.go metrics_optimized.go metrics_streaming.go metrics_shared.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Metrics_Baseline-12                   1        1450591279 ns/op  544892904 B/op   5867714 allocs/op
BenchmarkMetrics/Metrics_Array-12                      4         313136749 ns/op  271589240 B/op   2200097 allocs/op
BenchmarkMetrics/Metrics_Streaming-12                  6         172643944 ns/op  112461893 B/op   2200010 allocs/op
PASS
ok      command-line-arguments  6.349s
```