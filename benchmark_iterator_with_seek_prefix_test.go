package radix

import (
	"math/rand"
	"reflect"
	"testing"
	"time"

	radix "github.com/hashicorp/go-immutable-radix"
	pso_radix "prefix-search-optimized-radix/radix"
)

// generateRandomPrefix generates a random string of lowercase letters
// with a length between 1 and maxLen (inclusive).
func generateRandomPrefix(maxLen int) string {
	if maxLen <= 0 {
		maxLen = 1 // Ensure at least length 1 if invalid maxLen is given
	}
	// rand.Intn(n) returns a value in [0, n).
	// So, rand.Intn(maxLen) gives [0, maxLen-1].
	// Adding 1 makes it [1, maxLen].
	length := rand.Intn(maxLen) + 1
	chars := make([]byte, length)
	for i := 0; i < length; i++ {
		chars[i] = byte(rand.Intn(26) + 'a') // 'a' through 'z'
	}
	return string(chars)
}

func TestSeekPrefix(t *testing.T) {
	wordCount := 100000
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words, wordCount) // Assumes generateTestWords is in the same package

	// Seed the global random number generator once for the benchmark.
	// In Go 1.20+, the global rand is automatically seeded. For older versions, explicit seeding is good.
	// Using rand.NewSource and rand.New to re-seed global rand's default source.
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Setup for immutable radix tree (hashicorp/go-immutable-radix)
	immutableTree := radix.New()
	for _, word := range words {
		immutableTree, _, _ = immutableTree.Insert([]byte(word), nil)
	}

	// Setup for prefix-optimized radix tree (prefix-search-optimized-radix/radix)
	psoTree := pso_radix.New()
	for _, word := range words {
		psoTree, _, _ = psoTree.Insert([]byte(word), nil)
	}

	prefix := generateRandomPrefix(3) // Generate new random prefix for each iteration

	resultsImmutable := make([]string, 0)
	resultsOptimized := make([]string, 0)

	iter := immutableTree.Root().Iterator()
	iter.SeekPrefix([]byte(prefix))
	for {
		k, _, ok := iter.Next()
		resultsImmutable = append(resultsImmutable, string(k))
		if !ok {
			break
		}
	}

	iterPso := psoTree.Root().Iterator()
	iterPso.SeekPrefix([]byte(prefix))
	for {
		k, _, ok := iterPso.Next()
		resultsOptimized = append(resultsOptimized, string(k))
		if !ok {
			break
		}
	}

	if len(resultsImmutable) != len(resultsOptimized) {
		t.Fatalf("expected %d, got %d", len(resultsImmutable), len(resultsOptimized))
	}
	// assert both slices are equal
	if !reflect.DeepEqual(resultsImmutable, resultsOptimized) {
		t.Fatalf("expected %s, got %s", resultsImmutable, resultsOptimized)
	}
}

func Benchmark_TestSeekPrefix(b *testing.B) {
	wordCount := 100000
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words, wordCount) // Assumes generateTestWords is in the same package

	// Seed the global random number generator once for the benchmark.
	// In Go 1.20+, the global rand is automatically seeded. For older versions, explicit seeding is good.
	// Using rand.NewSource and rand.New to re-seed global rand's default source.
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Setup for immutable radix tree (hashicorp/go-immutable-radix)
	immutableTree := radix.New()
	for _, word := range words {
		immutableTree, _, _ = immutableTree.Insert([]byte(word), nil)
	}

	// Setup for prefix-optimized radix tree (prefix-search-optimized-radix/radix)
	psoTree := pso_radix.New()
	for _, word := range words {
		psoTree, _, _ = psoTree.Insert([]byte(word), nil)
	}

	b.Run("ImmutableRadixTree_SeekPrefix", func(b *testing.B) {
		b.ResetTimer()                    // Reset timer before the loop, after tree setup
		prefix := generateRandomPrefix(3) // Generate new random prefix for each iteration
		for i := 0; i < b.N; i++ {
			iter := immutableTree.Root().Iterator()
			iter.SeekPrefix([]byte(prefix))
			for {
				_, _, ok := iter.Next()
				if !ok {
					break
				}
			}
		}
	})

	b.Run("PrefixOptimizedRadixTree_SeekPrefix", func(b *testing.B) {
		b.ResetTimer()                    // Reset timer before the loop, after tree setup
		prefix := generateRandomPrefix(3) // Generate new random prefix for each iteration
		for i := 0; i < b.N; i++ {
			iter := psoTree.Root().Iterator()
			iter.SeekPrefix([]byte(prefix))
			for {
				_, _, ok := iter.Next()
				if !ok {
					break
				}
			}
		}
	})
}
