# Prefix Iteration Optimized Mutable Radix Tree

This repo introduces some modification in Mutable Radix Trees to make prefix iteration faster and memory efficient.

## Benchmarking Results

Results for - https://github.com/asheshvidyut/prefix-search-optimized-radix/blob/main/benchmark_iterator_test.go

### Iterating on all keys in Radix Tree (when prefix is "")

```
goos: linux
goarch: amd64
pkg: prefix-search-optimized-radix
cpu: AMD EPYC 7B12
BenchmarkIteratorImmutableRadixTree10-24                    	 6561556	       184.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree10-24              	39115611	        29.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree100-24                   	  310989	      3748 ns/op	      88 B/op	       3 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree100-24             	  723337	      1629 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree1000-24                  	    6063	    182730 ns/op	     184 B/op	       4 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree1000-24            	   10000	    117636 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree10000-24                 	    6105	    183487 ns/op	     184 B/op	       4 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree10000-24           	   10000	    116628 ns/op	       0 B/op	       0 allocs/op
```

#### Percentage Improvement (Speed)
| Test Case   | Immutable (ns/op) | Optimized (ns/op) | Reduction (ns) | % Improvement |
| ----------- | ----------------- | ----------------- | -------------- | ------------- |
| Keys 10     | 188.4             | 29.3              | 159.1          | 84.45%        |
| Keys 100    | 3746              | 1613              | 2133           | 56.94%        |
| Keys 1000   | 182540            | 120358            | 62182          | 34.06%        |
| Keys 10000  | 183450            | 122619            | 60831          | 33.16%        |
| Keys 100000 | 34890604          | 24042748          | 10847856       | 31.09%        |

#### Percentage Improvement (Memory)
| Test Case   | Immutable (B/op) | Optimized (B/op) | Reduction (B) | % Improvement (Bytes) | Immutable (allocs/op) | Optimized (allocs/op) | Reduction (allocs) | % Improvement (Allocs) |
| ----------- | ---------------- | ---------------- | ------------- | --------------------- | --------------------- | --------------------- | ------------------ | ---------------------- |
| Keys 10     | 40               | 0                | 40            | 100.00%               | 2                     | 0                     | 2                  | 100.00%                |
| Keys 100    | 88               | 0                | 88            | 100.00%               | 3                     | 0                     | 3                  | 100.00%                |
| Keys 1000   | 184              | 0                | 184           | 100.00%               | 4                     | 0                     | 4                  | 100.00%                |
| Keys 10000  | 184              | 0                | 184           | 100.00%               | 4                     | 0                     | 4                  | 100.00%                |
| Keys 100000 | 184              | 0                | 184           | 100.00%               | 4                     | 0                     | 4                  | 100.00%                |

### Iterating on keys with common prefix

```
Benchmark_TestSeekPrefix/ImmutableRadixTree_SeekPrefix-24         	  105685	     15320 ns/op	    3854 B/op	      39 allocs/op
Benchmark_TestSeekPrefix/PrefixOptimizedRadixTree_SeekPrefix-24   	  149748	     10279 ns/op	    2857 B/op	      37 allocs/op
```

#### Percentage Improvement

| Metric    | Immutable Radix Tree | Prefix Optimized Radix Tree | Difference | % Improvement (Reduction) |
| --------- | -------------------- | --------------------------- | ---------- | ------------------------- |
| ns/op     | 15320                | 10279                       | 5041       | 32.90%                    |
| B/op      | 3854                 | 2857                        | 997        | 25.87%                    |
| allocs/op | 39                   | 37                          | 2          | 5.13%                     |