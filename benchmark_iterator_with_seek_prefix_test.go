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
func generateRandomPrefixFromWords(words []string, maxLen int) string {
	if len(words) == 0 {
		// Fallback to random generation if words list is empty
		length := rand.Intn(maxLen) + 1
		chars := make([]byte, length)
		for i := 0; i < length; i++ {
			chars[i] = byte(rand.Intn(26) + 'a') // 'a' through 'z'
		}
		return string(chars)
	}

	randomWord := words[rand.Intn(len(words))]
	if len(randomWord) == 0 { // Handle empty string case
		return ""
	}
	prefixLen := min(len(randomWord), maxLen)
	return randomWord[:prefixLen]
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

	prefix := generateRandomPrefixFromWords(words, 3) // Generate new random prefix for each iteration

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

const numBenchmarkPrefixes = 1000000 // Define a sufficiently large number of prefixes to pre-generate

func Benchmark_TestSeekPrefix(b *testing.B) {
	wordCount := 100000
	words := make([]string, 0)
	words = append(words, "")
	words = generateTestWords(words, wordCount) // Assumes generateTestWords is in the same package

	// Seed the global random number generator once for the benchmark.
	// In Go 1.20+, the global rand is automatically seeded. For older versions, explicit seeding is good.
	// Using rand.NewSource and rand.New to re-seed global rand's default source.
	rand.New(rand.NewSource(time.Now().UnixNano()))

	// Generate prefixes once, outside and before any b.Run blocks
	benchmarkPrefixes := make([]string, numBenchmarkPrefixes)
	for i := 0; i < numBenchmarkPrefixes; i++ {
		benchmarkPrefixes[i] = generateRandomPrefixFromWords(words, 3)
	}

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

	resultKeysImm := make([]string, 0)

	b.Run("ImmutableRadixTree_SeekPrefix", func(b *testing.B) {
		b.ResetTimer() // Reset timer before the loop, after tree setup
		for i := 0; i < b.N; i++ {
			// Use the pre-generated prefixes, cycling if b.N > numBenchmarkPrefixes
			iter := immutableTree.Root().Iterator()
			iter.SeekPrefix([]byte(benchmarkPrefixes[i % len(benchmarkPrefixes)]))
			for {
				k, _, ok := iter.Next()
				if !ok {
					break
				}
				resultKeysImm = append(resultKeysImm, string(k))
			}
		}
	})

	resultKeysOpt := make([]string, 0)

	b.Run("PrefixOptimizedRadixTree_SeekPrefix", func(b *testing.B) {
		b.ResetTimer() // Reset timer before the loop, after tree setup
		for i := 0; i < b.N; i++ {
		// Use the same pre-generated prefixes, cycling if b.N > numBenchmarkPrefixes
			iter := psoTree.Root().Iterator()
			iter.SeekPrefix([]byte(benchmarkPrefixes[i % len(benchmarkPrefixes)]))
			for {
				k, _, ok := iter.Next()
				if !ok {
					break
				}
				resultKeysOpt = append(resultKeysOpt, string(k))
			}
		}
	})
}
