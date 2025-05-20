package radix

import (
	radix "github.com/hashicorp/go-immutable-radix"
	pso_radix "prefix-search-optimized-radix/radix"
	"testing"
)

const MAX_WORDS = 1000

// generateTestWords generates a list of words with common prefixes
func generateTestWords(words []string) []string {
	if len(words) > MAX_WORDS {
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
	return generateTestWords(newWords)
}

func BenchmarkIteratorImmutableRadixTree(b *testing.B) {
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words)

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

func BenchmarkIteratorPrefixOptimizedRadixTree(b *testing.B) {
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words)

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
