package radix

import (
	"bufio"
	"fmt"
	"math/rand"
	"os"
	"reflect"
	"sort"
	"strings"
	"testing"
	"time"
)

func ShuffleSlice[T any](slice []T) {
	rand.New(rand.NewSource(time.Now().UnixNano())) // Seed the random number generator
	for i := range slice {
		j := rand.Intn(i + 1) // Generate a random index
		slice[i], slice[j] = slice[j], slice[i]
	}
}

func TestSortingSpeedSortSlice(t *testing.T) {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// shuffle lines
	ShuffleSlice(lines)

	sort.Strings(lines)
	result := make([]string, len(lines))
	for indx, line := range lines {
		result[indx] = line
	}

	if len(lines) != len(lines) {
		t.Fatalf("expected %d, got %d", len(lines), len(lines))
	}
}

func TestSortingSpeedUsingRadixTree(t *testing.T) {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// shuffle lines
	ShuffleSlice(lines)

	r := New()
	for _, line := range lines {
		r, _, _ = r.Insert([]byte(line), nil)
	}

	iter := r.Root().Iterator()
	result := make([]string, len(lines))
	counter := 0
	for {
		key, _, found := iter.Next()
		if !found {
			break
		}
		result[counter] = string(key)
		counter++
	}
	if counter != len(lines) {
		t.Fatalf("expected %d, got %d", len(lines), counter)
	}
}

func TestIterator_SeekPrefix(t *testing.T) {
	file, err := os.Open("words.txt")
	if err != nil {
		fmt.Println("Error opening file:", err)
		return
	}
	defer file.Close()
	scanner := bufio.NewScanner(file)

	// Read the file line by line
	var lines []string
	for scanner.Scan() {
		line := scanner.Text()
		lines = append(lines, line)
	}

	// shuffle lines
	ShuffleSlice(lines)

	r := New()
	for _, line := range lines {
		r, _, _ = r.Insert([]byte(line), nil)
	}

	prefixes := make(map[string]struct{}, 0)

	for _, line := range lines {
		prefixes[line[0:min(len(line), 2)]] = struct{}{}
	}

	for prefix, _ := range prefixes {
		iter := r.Root().Iterator()
		iter.SeekPrefix([]byte(prefix))
		results := make([]string, 0)
		counter := 0
		for {
			key, _, found := iter.Next()
			if !found {
				break
			}
			results = append(results, string(key))
			counter++
		}

		bruteResults := make([]string, 0)
		for _, line := range lines {
			if strings.HasPrefix(line, prefix) {
				bruteResults = append(bruteResults, line)
			}
		}
		if counter != len(bruteResults) {
			t.Fatalf("expected %d, got %d", len(bruteResults), counter)
		}
		sort.Strings(bruteResults)
		if !reflect.DeepEqual(results, bruteResults) {
			t.Fatalf("expected %s, got %s", results, bruteResults)
		}
	}
}
