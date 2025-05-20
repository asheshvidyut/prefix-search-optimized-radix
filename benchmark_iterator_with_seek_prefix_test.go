package radix

import (
	radix "github.com/hashicorp/go-immutable-radix"
	pso_radix "prefix-search-optimized-radix/radix"
	"testing"
)

func runBehcmkarSeekPrefixImmutableRadixTreeWithWords(wordCount int, b *testing.B, char int) {
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

func BenchmarkSeekPrefixImmutableRadixTree10(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, char+65)
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, char+97)
	}
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, char+65)
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10, b, char+97)
	}
}

func BenchmarkSeekPrefixImmutableRadixTree100(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, char+65)
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, char+97)
	}
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree100(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, char+65)
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(100, b, char+97)
	}
}

func BenchmarkSeekPrefixImmutableRadixTree1000(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, char+65)
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, char+97)
	}
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree1000(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, char+65)
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(1000, b, char+97)
	}
}

func BenchmarkSeekPrefixImmutableRadixTree10000(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, char+65)
		runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, char+97)
	}
}

func BenchmarkSeekPrefixPrefixOptimizedRadixTree10000(b *testing.B) {
	for char := 0; char < 26; char++ {
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, char+65)
		runBenchmarkSeekPrefixPrefixOptimizedRadixTreeWithWords(10000, b, char+97)
	}
}
