# Prefix Iteration Optimized Mutable Radix Tree

This repo introduces some modification in Mutable Radix Trees to make prefix iteration faster and memory efficient.

## Benchmarking Results

Results for - https://github.com/asheshvidyut/prefix-search-optimized-radix/blob/main/benchmark_iterator_test.go

### Iterating on all keys in Radix Tree

```
goos: linux
goarch: amd64
pkg: prefix-search-optimized-radix
cpu: AMD EPYC 7B12
BenchmarkIteratorImmutableRadixTree10-24                    	 6680989	       183.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree10-24              	37727648	        30.47 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree100-24                   	  329053	      3607 ns/op	      88 B/op	       3 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree100-24             	  709500	      1661 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree1000-24                  	    6451	    180995 ns/op	     184 B/op	       4 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree1000-24            	    9609	    120672 ns/op	       0 B/op	       0 allocs/op
BenchmarkIteratorImmutableRadixTree10000-24                 	    5900	    181655 ns/op	     184 B/op	       4 allocs/op
BenchmarkIteratorPrefixOptimizedRadixTree10000-24           	    9662	    117438 ns/op	       0 B/op	       0 allocs/op

```

#### Percentage Improvement (Speed)
| Test Case      | Immutable (ns/op) | Optimized (ns/op) | % Improvement |
| -------------- | ----------------- | ----------------- | ------------- |
| Iterator 10    | 183.1             | 30.47             | **83.36%**    |
| Iterator 100   | 3607              | 1661              | **53.96%**    |
| Iterator 1000  | 180,995           | 120,672           | **33.37%**    |
| Iterator 10000 | 181,655           | 117,438           | **35.34%**    |

#### Percentage Improvement (Memory)
| Test Case      | Immutable (B/op) | Optimized (B/op) | Reduction | % Improvement |
| -------------- | ---------------- | ---------------- | --------- | ------------- |
| Iterator 10    | 40               | 0                | 40 B      | **100%**      |
| Iterator 100   | 88               | 0                | 88 B      | **100%**      |
| Iterator 1000  | 184              | 0                | 184 B     | **100%**      |
| Iterator 10000 | 184              | 0                | 184 B     | **100%**      |

### Iterating on prefix with Charx -> x is changing for each test

Results for - https://github.com/asheshvidyut/prefix-search-optimized-radix/blob/main/benchmark_iterator_with_seek_prefix_test.go

#### Percentage Improvement
| Test Case       | Average Time (ns/op) Improvement | Average Allocations (allocs/op) Improvement |
|-----------------|----------------------------------|---------------------------------------------|
| Iterator 10     | **43.53%**                       | **50%**                                     |                                
| Iterator 100    | **68.86%**                       | **50%**                                     |
| Iterator 1000   | **21.56%**	                      | **66.67%**                                  | 
| Iterator 10000  | **21.78%**	                      | **66.67%**                                  |


```
BenchmarkSeekPrefixImmutableRadixTree10Chara-24             	15237057	        77.50 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charb-24             	15061152	        76.44 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charc-24             	15482919	        76.28 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chard-24             	15479586	        76.44 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chare-24             	15407764	        76.52 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charf-24             	15407218	        75.46 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charg-24             	15706352	        75.31 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charh-24             	15882207	        76.89 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chari-24             	15351138	        76.32 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charj-24             	15411552	        76.64 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chark-24             	15420194	        76.33 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charl-24             	15265862	        77.18 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charm-24             	15218582	        76.60 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charn-24             	15352039	        74.96 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charo-24             	15583879	        75.13 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charp-24             	15596484	        75.61 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charq-24             	15568441	        75.56 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charr-24             	15454342	        76.10 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chars-24             	15560924	        76.31 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chart-24             	15561338	        75.43 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charu-24             	15516292	        75.40 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charv-24             	15242829	        75.63 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charw-24             	15358824	        76.43 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charx-24             	15646808	        75.56 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Chary-24             	15594718	        77.70 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10Charz-24             	15288328	        77.35 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chara-24       	27246858	        42.85 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charb-24       	27221868	        42.94 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charc-24       	26861552	        43.01 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chard-24       	26965446	        43.01 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chare-24       	26223480	        42.78 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charf-24       	25498790	        42.83 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charg-24       	27703897	        42.16 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charh-24       	27149704	        43.06 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chari-24       	26734849	        43.36 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charj-24       	27010269	        42.81 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chark-24       	26851586	        43.09 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charl-24       	26986380	        45.76 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charm-24       	25764355	        43.60 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charn-24       	26848168	        42.42 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charo-24       	27129658	        43.08 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charp-24       	27066294	        43.36 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charq-24       	26250321	        43.15 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charr-24       	26101332	        43.74 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chars-24       	26532721	        43.16 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chart-24       	26900650	        43.15 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charu-24       	27417564	        42.67 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charv-24       	27442863	        43.05 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charw-24       	27317718	        42.70 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charx-24       	26976412	        42.31 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Chary-24       	26679416	        43.25 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10Charz-24       	26539058	        42.88 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chara-24            	 5445049	       221.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charb-24            	 5405463	       219.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charc-24            	 5385795	       222.6 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chard-24            	 5448735	       220.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chare-24            	 5422670	       219.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charf-24            	 5468924	       220.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charg-24            	 5372131	       223.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charh-24            	 5236040	       219.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chari-24            	 5281958	       223.6 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charj-24            	 5399079	       222.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chark-24            	 5525913	       218.0 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charl-24            	 5280544	       223.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charm-24            	 5348899	       223.3 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charn-24            	 5523660	       217.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charo-24            	 5346825	       223.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charp-24            	 5343248	       222.6 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charq-24            	 5438364	       221.1 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charr-24            	 5368182	       223.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chars-24            	 5450229	       218.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chart-24            	 5524689	       218.4 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charu-24            	 5416837	       220.5 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charv-24            	 5426809	       216.2 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charw-24            	 5484987	       216.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charx-24            	 5527123	       216.8 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Chary-24            	 5528174	       216.9 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixImmutableRadixTree100Charz-24            	 5488987	       217.7 ns/op	      40 B/op	       2 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chara-24      	17391624	        68.47 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charb-24      	17130208	        68.90 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charc-24      	17258212	        68.15 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chard-24      	17089500	        68.75 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chare-24      	17057587	        67.99 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charf-24      	17206858	        68.54 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charg-24      	17369353	        70.25 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charh-24      	16820812	        68.00 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chari-24      	16576806	        68.83 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charj-24      	16531254	        68.28 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chark-24      	16982082	        68.88 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charl-24      	17117295	        68.47 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charm-24      	16966894	        68.78 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charn-24      	17248047	        67.99 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charo-24      	16701649	        67.82 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charp-24      	17193670	        72.04 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charq-24      	17258937	        68.44 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charr-24      	16448228	        68.68 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chars-24      	16766734	        68.94 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chart-24      	16851560	        68.97 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charu-24      	16854603	        68.48 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charv-24      	17376705	        68.00 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charw-24      	16852324	        68.82 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charx-24      	17232915	        68.28 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Chary-24      	17159155	        68.08 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree100Charz-24      	17040301	        68.90 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chara-24           	  283544	      4247 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charb-24           	  286922	      4245 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charc-24           	  268425	      4236 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chard-24           	  265020	      4241 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chare-24           	  284378	      4279 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charf-24           	  282463	      4269 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charg-24           	  285868	      4215 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charh-24           	  280488	      4249 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chari-24           	  284598	      4228 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charj-24           	  281661	      4212 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chark-24           	  281950	      4232 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charl-24           	  283839	      4271 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charm-24           	  279786	      4242 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charn-24           	  279920	      4224 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charo-24           	  282759	      4222 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charp-24           	  281307	      4250 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charq-24           	  280767	      4227 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charr-24           	  285573	      4274 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chars-24           	  279924	      4251 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chart-24           	  286081	      4229 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charu-24           	  279108	      4260 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charv-24           	  285782	      4313 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charw-24           	  273714	      4296 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charx-24           	  259119	      4290 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Chary-24           	  277081	      4309 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree1000Charz-24           	  270703	      4216 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chara-24     	  358516	      3345 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charb-24     	  366708	      3391 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charc-24     	  355910	      3291 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chard-24     	  356186	      3351 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chare-24     	  359236	      3324 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charf-24     	  361432	      3277 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charg-24     	  362738	      3361 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charh-24     	  365374	      3359 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chari-24     	  360614	      3373 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charj-24     	  358464	      3372 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chark-24     	  359397	      3342 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charl-24     	  340284	      3389 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charm-24     	  354487	      3268 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charn-24     	  357226	      3293 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charo-24     	  358980	      3288 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charp-24     	  341439	      3326 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charq-24     	  358168	      3355 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charr-24     	  360522	      3404 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chars-24     	  353095	      3325 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chart-24     	  354814	      3352 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charu-24     	  361990	      3321 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charv-24     	  359072	      3312 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charw-24     	  360568	      3259 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charx-24     	  361780	      3291 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Chary-24     	  360934	      3319 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree1000Charz-24     	  366404	      3359 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chara-24          	  282409	      4228 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charb-24          	  283872	      4270 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charc-24          	  284821	      4352 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chard-24          	  266284	      4235 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chare-24          	  285926	      4255 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charf-24          	  264063	      4247 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charg-24          	  284806	      4215 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charh-24          	  286014	      4261 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chari-24          	  278306	      4262 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charj-24          	  275646	      4221 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chark-24          	  284934	      4303 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charl-24          	  265808	      4272 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charm-24          	  284284	      4216 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charn-24          	  282660	      4244 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charo-24          	  280000	      4259 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charp-24          	  272690	      4277 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charq-24          	  275565	      4232 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charr-24          	  283431	      4244 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chars-24          	  278584	      4251 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chart-24          	  275676	      4278 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charu-24          	  284701	      4254 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charv-24          	  286598	      4249 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charw-24          	  284222	      4257 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charx-24          	  277394	      4190 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Chary-24          	  273170	      4244 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixImmutableRadixTree10000Charz-24          	  278019	      4301 ns/op	      88 B/op	       3 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chara-24    	  353516	      3398 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charb-24    	  347608	      3378 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charc-24    	  359282	      3334 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chard-24    	  364342	      3283 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chare-24    	  346372	      3297 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charf-24    	  368419	      3293 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charg-24    	  367656	      3294 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charh-24    	  351650	      3330 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chari-24    	  353445	      3382 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charj-24    	  350163	      3362 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chark-24    	  357517	      3356 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charl-24    	  363393	      3281 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charm-24    	  357282	      3313 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charn-24    	  356734	      3304 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charo-24    	  363284	      3274 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charp-24    	  363535	      3315 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charq-24    	  363540	      3326 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charr-24    	  349428	      3358 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chars-24    	  361802	      3348 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chart-24    	  357376	      3344 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charu-24    	  354115	      3345 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charv-24    	  358956	      3332 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charw-24    	  350799	      3254 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charx-24    	  363972	      3314 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Chary-24    	  341826	      3286 ns/op	       8 B/op	       1 allocs/op
BenchmarkSeekPrefixPrefixOptimizedRadixTree10000Charz-24    	  355557	      3358 ns/op	       8 B/op	       1 allocs/op
PASS
ok  	prefix-search-optimized-radix	291.558s
```
