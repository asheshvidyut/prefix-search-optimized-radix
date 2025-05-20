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

func BenchmarkSeekPrefixImmutableRadixTree10Chara(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charb(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charc(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chard(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chare(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charf(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charg(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charh(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chari(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charj(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chark(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charl(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charm(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charn(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charo(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charp(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charq(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charr(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chars(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chart(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charu(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charv(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charw(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charx(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Chary(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10Charz(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharAUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('A'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharBUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('B'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharCUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('C'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharDUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('D'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharEUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('E'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharFUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('F'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharGUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('G'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharHUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('H'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharIUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('I'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharJUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('J'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharKUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('K'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharLUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('L'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharMUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('M'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharNUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('N'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharOUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('O'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharPUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('P'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharQUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('Q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharRUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('R'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharSUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('S'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharTUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('T'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharUUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('U'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharVUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('V'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharWUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('W'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharXUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('X'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharYUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('Y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10CharZUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10, b, int('Z'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chara(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charb(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charc(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chard(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chare(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charf(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charg(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charh(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chari(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charj(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chark(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charl(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charm(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charn(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charo(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charp(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charq(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charr(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chars(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chart(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charu(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charv(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charw(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charx(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Chary(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree100Charz(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharAUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('A'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharBUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('B'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharCUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('C'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharDUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('D'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharEUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('E'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharFUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('F'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharGUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('G'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharHUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('H'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharIUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('I'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharJUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('J'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharKUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('K'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharLUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('L'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharMUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('M'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharNUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('N'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharOUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('O'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharPUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('P'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharQUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('Q'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharRUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('R'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharSUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('S'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharTUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('T'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharUUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('U'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharVUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('V'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharWUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('W'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharXUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('X'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharYUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('Y'))
}

func BenchmarkSeekPrefixImmutableRadixTree100CharZUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(100, b, int('Z'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chara(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charb(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charc(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chard(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chare(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charf(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charg(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charh(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chari(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charj(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chark(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charl(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charm(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charn(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charo(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charp(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charq(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charr(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chars(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chart(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charu(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charv(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charw(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charx(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Chary(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000Charz(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharAUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('A'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharBUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('B'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharCUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('C'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharDUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('D'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharEUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('E'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharFUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('F'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharGUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('G'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharHUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('H'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharIUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('I'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharJUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('J'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharKUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('K'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharLUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('L'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharMUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('M'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharNUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('N'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharOUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('O'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharPUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('P'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharQUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('Q'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharRUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('R'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharSUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('S'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharTUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('T'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharUUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('U'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharVUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('V'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharWUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('W'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharXUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('X'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharYUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('Y'))
}

func BenchmarkSeekPrefixImmutableRadixTree1000CharZUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(1000, b, int('Z'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chara(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('a'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charb(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('b'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charc(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('c'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chard(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('d'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chare(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('e'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charf(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('f'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charg(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('g'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charh(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('h'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chari(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('i'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charj(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('j'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chark(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('k'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charl(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('l'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charm(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('m'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charn(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('n'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charo(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('o'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charp(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('p'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charq(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charr(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('r'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chars(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('s'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chart(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('t'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charu(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('u'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charv(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('v'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charw(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('w'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charx(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('x'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Chary(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000Charz(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('z'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharAUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('A'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharBUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('B'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharCUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('C'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharDUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('D'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharEUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('E'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharFUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('F'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharGUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('G'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharHUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('H'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharIUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('I'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharJUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('J'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharKUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('K'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharLUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('L'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharMUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('M'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharNUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('N'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharOUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('O'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharPUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('P'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharQUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('Q'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharRUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('R'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharSUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('S'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharTUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('T'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharUUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('U'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharVUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('V'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharWUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('W'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharXUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('X'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharYUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('Y'))
}

func BenchmarkSeekPrefixImmutableRadixTree10000CharZUpper(b *testing.B) {
	runBehcmkarSeekPrefixImmutableRadixTreeWithWords(10000, b, int('Z'))
}
