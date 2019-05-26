# Struct padding benchmark

## Result
```
go version go1.12.5 darwin/amd64
go test -cpu 1,2,4,8 -count 5 -benchmem -bench .
Name: Intel(R) Core(TM) i7-8559U CPU @ 2.70GHz
PhysicalCores: 4
ThreadsPerCore: 2
LogicalCores: 8
Family 6 Model: 142
Features: CMOV,NX,MMX,MMXEXT,SSE,SSE2,SSE3,SSSE3,SSE4.1,SSE4.2,AVX,AVX2,FMA3,F16C,BMI1,BMI2,LZCNT,POPCNT,AESNI,CLMUL,HTT,RDRAND,RDSEED,ADX,MPX,ERMS,RDTSCP,CX16,SGX,IBPB,STIBP
Cacheline bytes: 64
L1 Data Cache: 32768 bytes
L1 Instruction Cache: 32768 bytes
L2 Cache: 262144 bytes
L3 Cache: 8388608 bytes
We have Streaming SIMD Extensions
---
goos: darwin
goarch: amd64
pkg: github.com/tai-ga/padbench
Benchmark_StructNoPad           100000000               13.0 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad           100000000               12.8 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad           100000000               12.7 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad           100000000               12.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad           100000000               12.7 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-2         30000000                49.4 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-2         30000000                49.6 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-2         30000000                45.1 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-2         30000000                49.0 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-2         30000000                49.4 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-4         30000000                55.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-4         30000000                55.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-4         30000000                55.6 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-4         30000000                55.7 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-4         30000000                55.6 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-8         30000000                48.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-8         30000000                56.1 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-8         30000000                48.0 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-8         30000000                55.4 ns/op             0 B/op          0 allocs/op
Benchmark_StructNoPad-8         30000000                55.6 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad             100000000               12.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad             100000000               12.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad             100000000               13.0 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad             100000000               12.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad             100000000               12.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-2           50000000                32.4 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-2           50000000                32.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-2           50000000                31.6 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-2           50000000                31.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-2           50000000                32.2 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-4           50000000                23.8 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-4           100000000               24.2 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-4           100000000               24.2 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-4           100000000               24.3 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-4           100000000               23.9 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-8           50000000                26.2 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-8           50000000                26.5 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-8           50000000                26.8 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-8           50000000                26.3 ns/op             0 B/op          0 allocs/op
Benchmark_StructPad-8           50000000                26.4 ns/op             0 B/op          0 allocs/op
PASS
ok      github.com/tai-ga/padbench      63.224s
```
