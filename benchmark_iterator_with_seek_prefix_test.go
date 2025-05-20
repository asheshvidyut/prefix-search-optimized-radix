package radix

import (
	radix "github.com/hashicorp/go-immutable-radix"
	pso_radix "prefix-search-optimized-radix/radix"
	"testing"
)

func runBenchmarkSeekPrefixImmutableRadixTreeWithWords(wordCount int, b *testing.B, char int) {
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words, wordCount)

	rtree := radix.New()
	for _, word := range words {
		rtree, _, _ = rtree.Insert([]byte(word), nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iter := rtree.Root().Iterator()
		iter.SeekPrefix([]byte(string(rune(char))))
		for {
			_, _, ok := iter.Next()
			if !ok {
				break
			}
		}
	}
}

func runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(wordCount int, b *testing.B, char int) {
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words, wordCount)

	pso_tree := pso_radix.New()
	for _, word := range words {
		pso_tree, _, _ = pso_tree.Insert([]byte(word), nil)
	}

	b.ResetTimer()
	for i := 0; i < b.N; i++ {
		iter := pso_tree.Root().Iterator()
		iter.SeekPrefix([]byte(string(rune(char))))
		for {
			_, _, ok := iter.Next()
			if !ok {
				break
			}
		}
	}
}

func BenchmarkSeekPrefixImmutableRadixTree10Chara(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charb(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charc(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chard(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chare(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charf(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charg(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charh(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chari(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charj(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chark(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charl(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charm(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charn(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charo(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charp(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charq(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charr(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chars(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chart(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charu(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charv(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charw(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charx(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chary(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charz(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10, b, int('z'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharaUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('a'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharbUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('b'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharcUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('c'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10ChardUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('d'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10ChareUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('e'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharfUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('f'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10ChargUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('g'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharhUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('h'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10ChariUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('i'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharjUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('j'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharkUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('k'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharlUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('l'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharmUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('m'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharnUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('n'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharoUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('o'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharpUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('p'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharqUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('q'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharrUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('r'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharsUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('s'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10ChartUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('t'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharuUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('u'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharvUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('v'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharwUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('w'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharxUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('x'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharyUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('y'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10CharzUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chara(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charb(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charc(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chard(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chare(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charf(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charg(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charh(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chari(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charj(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chark(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charl(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charm(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charn(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charo(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charp(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charq(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charr(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chars(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chart(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charu(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charv(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charw(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charx(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chary(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charz(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(100, b, int('z'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharaUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('a'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharbUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('b'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharcUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('c'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100ChardUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('d'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100ChareUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('e'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharfUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('f'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100ChargUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('g'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharhUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('h'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100ChariUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('i'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharjUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('j'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharkUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('k'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharlUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('l'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharmUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('m'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharnUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('n'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharoUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('o'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharpUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('p'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharqUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('q'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharrUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('r'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharsUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('s'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100ChartUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('t'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharuUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('u'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharvUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('v'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharwUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('w'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharxUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('x'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharyUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('y'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100CharzUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chara(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charb(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charc(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chard(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chare(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charf(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charg(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charh(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chari(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charj(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chark(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charl(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charm(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charn(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charo(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charp(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charq(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charr(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chars(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chart(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charu(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charv(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charw(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charx(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chary(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charz(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(1000, b, int('z'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharaUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('a'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharbUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('b'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharcUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('c'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000ChardUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('d'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000ChareUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('e'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharfUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('f'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000ChargUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('g'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharhUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('h'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000ChariUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('i'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharjUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('j'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharkUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('k'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharlUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('l'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharmUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('m'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharnUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('n'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharoUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('o'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharpUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('p'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharqUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('q'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharrUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('r'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharsUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('s'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000ChartUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('t'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharuUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('u'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharvUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('v'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharwUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('w'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharxUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('x'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharyUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('y'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000CharzUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chara(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charb(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charc(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chard(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chare(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charf(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charg(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charh(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chari(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charj(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chark(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charl(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charm(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charn(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charo(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charp(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charq(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charr(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chars(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chart(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charu(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charv(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charw(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charx(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chary(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charz(b *testing.B) {
	runBenchmarkSeekPrefixImmutableRadixTreeWithWords(10000, b, int('z'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharaUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('a'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharbUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('b'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharcUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('c'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000ChardUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('d'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000ChareUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('e'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharfUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('f'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000ChargUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('g'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharhUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('h'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000ChariUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('i'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharjUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('j'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharkUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('k'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharlUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('l'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharmUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('m'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharnUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('n'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharoUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('o'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharpUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('p'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharqUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('q'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharrUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('r'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharsUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('s'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000ChartUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('t'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharuUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('u'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharvUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('v'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharwUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('w'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharxUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('x'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharyUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('y'))
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000CharzUpper(b *testing.B) {
	runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, int('z'))
}
