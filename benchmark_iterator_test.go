package radix

import (
	"testing"

	radix "github.com/hashicorp/go-immutable-radix"
	pso_radix "prefix-search-optimized-radix/radix"
)

// generateTestWords generates a list of words with common prefixes
func generateTestWords(words []string, maxWords int) []string {
	if len(words) > maxWords {
		return words
	}
	newWords := make([]string, len(words))
	for i, word := range words {
		newWords[i] = word
	}
	for char := 0; char < 26; char++ {
		for _, word := range words {
			newWords = append(newWords, word+string(rune(char+97)))
		}
	}
	return generateTestWords(newWords, maxWords)
}

func runBenchmarkIteratorImmutableRadixTreeWithWords(wordCount int, b *testing.B) {
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
		for {
			_, _, ok := iter.Next()
			if !ok {
				break
			}
		}
	}
}

func runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(wordCount int, b *testing.B) {
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
		for {
			_, _, ok := iter.Next()
			if !ok {
				break
			}
		}
	}
}

func BenchmarkIteratorImmutableRadixTree10(b *testing.B) {
	runBenchmarkIteratorImmutableRadixTreeWithWords(10, b)
}

func BenchmarkIteratorPrefixOptimizedRadixTree10(b *testing.B) {
	runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(10, b)
}

func BenchmarkIteratorImmutableRadixTree100(b *testing.B) {
	runBenchmarkIteratorImmutableRadixTreeWithWords(100, b)
}

func BenchmarkIteratorPrefixOptimizedRadixTree100(b *testing.B) {
	runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(100, b)
}

func BenchmarkIteratorImmutableRadixTree1000(b *testing.B) {
	runBenchmarkIteratorImmutableRadixTreeWithWords(1000, b)
}

func BenchmarkIteratorPrefixOptimizedRadixTree1000(b *testing.B) {
	runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(1000, b)
}

func BenchmarkIteratorImmutableRadixTree10000(b *testing.B) {
	runBenchmarkIteratorImmutableRadixTreeWithWords(10000, b)
}

func BenchmarkIteratorPrefixOptimizedRadixTree10000(b *testing.B) {
	runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(10000, b)
}

func BenchmarkIteratorImmutableRadixTree100000(b *testing.B) {
	runBenchmarkIteratorImmutableRadixTreeWithWords(100000, b)
}

func BenchmarkIteratorPrefixOptimizedRadixTree100000(b *testing.B) {
	runBenchmarkIteratorPrefixOptimizedRadixTreeWithWords(100000, b)
}
