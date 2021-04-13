# Metrics Optimization

Optimization exercise on an example metrics program which calculates averages and standard deviation across a mock dataset of users and their payments.

* `metrics_baseline.go` is the program as given.

* `metrics_optized.go` is an optimized version that uses an array-based vs object-oriented approach.

* `metrics_streaming.go` is a second optimized version that uses a single pass through the source files.

To run benchmarks:

```sh
➜  go test -bench=. -benchmem metrics_baseline.go metrics_test.go metrics_optimized.go metrics_streaming.go metrics_shared.go
goos: darwin
goarch: amd64
cpu: Intel(R) Core(TM) i7-9750H CPU @ 2.60GHz
BenchmarkMetrics/Metrics_Baseline-12                   1        1425283274 ns/op        544903984 B/op   5867781 allocs/op
BenchmarkMetrics/Metrics_Optimized-12                  4         311954067 ns/op        271589242 B/op   2200097 allocs/op
BenchmarkMetrics/Metrics_Streaming-12                  6         174200070 ns/op        112461762 B/op   2200009 allocs/op
PASS
ok      command-line-arguments  5.405s
```