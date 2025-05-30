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

### Iterating on prefix with Charx -> x is changing for each test (when prefix is x)

Results for - https://github.com/asheshvidyut/prefix-search-optimized-radix/blob/main/benchmark_iterator_with_seek_prefix_test.go

#### Percentage Improvement
| Test Case  | Immutable (ns/op) | Optimized (ns/op) | Reduction (ns) | % Improvement |
| ---------- | ----------------- | ----------------- | -------------- | ------------- |
| Keys 10    | 77.26             | 40.48             | 36.78          | 47.61%        |
| Keys 100   | 221.05            | 70.36             | 150.69         | 68.17%        |
| Keys 1000  | 4296.88           | 3368.58           | 928.3          | 21.60%        |
| Keys 10000 | 4298.81           | 3393.73           | 905.08         | 21.05%        |

| Test Case  | Immutable (B/op) | Optimized (B/op) | Reduction (B) | % Improvement (Bytes) | Immutable (allocs/op) | Optimized (allocs/op) | Reduction (allocs) | % Improvement (Allocs) |
| ---------- | ---------------- | ---------------- | ------------- | --------------------- | --------------------- | --------------------- | ------------------ | ---------------------- |
| Keys 10    | 40               | 8                | 32            | 80.00%                | 2                     | 1                     | 1                  | 50.00%                 |
| Keys 100   | 40               | 8                | 32            | 80.00%                | 2                     | 1                     | 1                  | 50.00%                 |
| Keys 1000  | 88               | 8                | 80            | 90.91%                | 3                     | 1                     | 2                  | 66.67%                 |
| Keys 10000 | 88               | 8                | 80            | 90.91%                | 3                     | 1                     | 2                  | 66.67%                 |

```
BenchmarkSeekPrefixImmutableRadixTree10Chara-24             	15274972	        76.77 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charb-24             	14734117	        77.70 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charc-24             	15575943	        77.37 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chard-24             	15240331	        77.43 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chare-24             	15524208	        76.87 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charf-24             	14702588	        76.98 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charg-24             	15276706	        76.40 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charh-24             	15361718	        77.57 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chari-24             	15168244	        76.96 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charj-24             	15315992	        77.46 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chark-24             	15218184	        77.77 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charl-24             	14858880	        77.91 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charm-24             	15238627	        77.83 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charn-24             	15193464	        76.87 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charo-24             	15794480	        77.17 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charp-24             	15121131	        77.94 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charq-24             	15300223	        77.81 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charr-24             	15169582	        77.81 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chars-24             	15435277	        78.16 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chart-24             	15527197	        77.85 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charu-24             	15260644	        77.68 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charv-24             	14837864	        77.64 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charw-24             	14991537	        77.85 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charx-24             	15186418	        76.96 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chary-24             	15403572	        77.59 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charz-24             	15087324	        77.78 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chara-24       	27725242	        40.60 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charb-24       	28566375	        40.77 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charc-24       	27919021	        40.34 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chard-24       	28243780	        40.54 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chare-24       	28397722	        40.22 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charf-24       	28558468	        40.61 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charg-24       	27873542	        39.85 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charh-24       	28277292	        40.52 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chari-24       	28406402	        40.51 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charj-24       	28311682	        40.25 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chark-24       	27413666	        41.17 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charl-24       	28650004	        40.65 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charm-24       	28167554	        40.61 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charn-24       	28912549	        39.96 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charo-24       	28553086	        40.68 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charp-24       	28262220	        40.55 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charq-24       	28590102	        40.58 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charr-24       	28787506	        40.52 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chars-24       	28462762	        41.06 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chart-24       	28258160	        40.48 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charu-24       	28678009	        39.72 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charv-24       	28409912	        40.22 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charw-24       	28340013	        40.28 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charx-24       	28700918	        39.81 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chary-24       	28108147	        40.62 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charz-24       	28844434	        40.67 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chara-24            	 5444086	       223.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charb-24            	 5442054	       221.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charc-24            	 5389382	       220.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chard-24            	 5399125	       221.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chare-24            	 5387340	       221.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charf-24            	 5422624	       221.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charg-24            	 5428395	       223.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charh-24            	 5435593	       220.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chari-24            	 5357238	       220.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charj-24            	 5378626	       221.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chark-24            	 5420757	       219.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charl-24            	 5410214	       220.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charm-24            	 5428618	       224.6 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charn-24            	 5427068	       221.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charo-24            	 5383028	       222.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charp-24            	 5399582	       221.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charq-24            	 5312562	       221.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charr-24            	 5336865	       221.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chars-24            	 5424092	       220.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chart-24            	 5532978	       218.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charu-24            	 5442236	       220.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charv-24            	 5501827	       219.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charw-24            	 5213482	       221.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charx-24            	 5477402	       217.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chary-24            	 5529924	       218.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charz-24            	 5484036	       219.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chara-24      	17307858	        68.96 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charb-24      	16547337	        69.36 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charc-24      	16672669	        69.38 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chard-24      	17205880	        70.63 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chare-24      	17187978	        68.51 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charf-24      	16954696	        70.03 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charg-24      	15705817	        75.12 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charh-24      	15173550	        73.96 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chari-24      	15373519	       193.4 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charj-24      	17045876	        69.06 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chark-24      	16937253	        71.73 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charl-24      	17065579	        68.61 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charm-24      	14865972	        69.19 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charn-24      	17157252	        68.67 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charo-24      	17215304	        67.96 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charp-24      	16976742	        69.07 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charq-24      	16758081	        68.74 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charr-24      	16871250	        71.90 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chars-24      	17001807	        68.73 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chart-24      	16940792	        69.78 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charu-24      	16718720	        68.71 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charv-24      	17241154	        68.17 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charw-24      	16940616	        69.41 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charx-24      	17038724	        70.67 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chary-24      	16545594	        68.78 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charz-24      	15345169	        69.17 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chara-24           	  280803	      4251 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charb-24           	  277011	      4255 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charc-24           	  280573	      4318 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chard-24           	  282966	      4328 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chare-24           	  272220	      4270 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charf-24           	  275373	      4302 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charg-24           	  279175	      4280 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charh-24           	  267142	      4355 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chari-24           	  278161	      4323 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charj-24           	  264985	      4286 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chark-24           	  279586	      4266 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charl-24           	  283591	      4347 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charm-24           	  277065	      4391 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charn-24           	  284896	      4278 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charo-24           	  279456	      4270 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charp-24           	  274832	      4350 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charq-24           	  273030	      4277 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charr-24           	  262711	      4343 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chars-24           	  278146	      4361 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chart-24           	  260298	      4298 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charu-24           	  283363	      4272 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charv-24           	  276795	      4293 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charw-24           	  282456	      4269 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charx-24           	  249399	      4283 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chary-24           	  269348	      4295 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charz-24           	  285069	      4320 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chara-24     	  355878	      3299 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charb-24     	  352708	      3334 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charc-24     	  355800	      3375 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chard-24     	  355760	      3236 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chare-24     	  366703	      3371 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charf-24     	  354225	      3386 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charg-24     	  350968	      3368 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charh-24     	  350124	      3445 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chari-24     	  352192	      3476 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charj-24     	  358154	      3351 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chark-24     	  354216	      3342 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charl-24     	  358549	      3363 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charm-24     	  350097	      3374 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charn-24     	  356749	      3386 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charo-24     	  369632	      3365 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charp-24     	  352873	      3394 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charq-24     	  353954	      3403 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charr-24     	  349243	      3405 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chars-24     	  354596	      3414 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chart-24     	  361184	      3372 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charu-24     	  355118	      3300 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charv-24     	  357358	      3405 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charw-24     	  341574	      3398 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charx-24     	  351699	      3390 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chary-24     	  354966	      3335 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charz-24     	  346148	      3401 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chara-24          	  267594	      4288 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charb-24          	  277274	      4312 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charc-24          	  286303	      4297 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chard-24          	  281178	      4315 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chare-24          	  279016	      4272 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charf-24          	  270601	      4283 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charg-24          	  272364	      4284 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charh-24          	  276817	      4361 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chari-24          	  273630	      4318 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charj-24          	  281454	      4273 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chark-24          	  278330	      4271 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charl-24          	  282368	      4291 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charm-24          	  283530	      4281 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charn-24          	  283728	      4304 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charo-24          	  272844	      4297 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charp-24          	  281973	      4341 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charq-24          	  274827	      4294 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charr-24          	  280977	      4274 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chars-24          	  276759	      4315 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chart-24          	  278721	      4287 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charu-24          	  277722	      4291 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charv-24          	  278910	      4268 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charw-24          	  283858	      4296 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charx-24          	  280380	      4383 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chary-24          	  282062	      4314 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charz-24          	  256893	      4280 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chara-24    	  353830	      3383 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charb-24    	  355134	      3410 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charc-24    	  352788	      3409 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chard-24    	  340785	      3396 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chare-24    	  357992	      3364 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charf-24    	  355617	      3361 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charg-24    	  350512	      3439 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charh-24    	  358653	      3402 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chari-24    	  348918	      3407 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charj-24    	  345530	      3433 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chark-24    	  353206	      3418 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charl-24    	  353512	      3330 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charm-24    	  348980	      3387 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charn-24    	  358266	      3393 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charo-24    	  354746	      3371 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charp-24    	  353569	      3376 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charq-24    	  346224	      3395 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charr-24    	  355117	      3415 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chars-24    	  352862	      3375 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chart-24    	  351387	      3361 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charu-24    	  335422	      3426 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charv-24    	  341361	      3404 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charw-24    	  352033	      3372 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charx-24    	  348289	      3383 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chary-24    	  351138	      3395 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charz-24    	  350509	      3484 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	prefix-search-optimized-radix	305.336s
```