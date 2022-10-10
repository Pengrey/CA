package main

import (
	"fmt"
	"math"
	"os"
	"sort"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func ngrams(words []string, size int) (count map[string]uint32) {

	count = make(map[string]uint32, 0)
	offset := int(math.Floor(float64(size / 2)))

	max := len(words)
	for i, _ := range words {
		if i < offset || i+size-offset > max {
			continue
		}
		gram := strings.Join(words[i-offset:i+size-offset], " ")
		count[gram]++
	}

	return count
}

func main() {
	// Performance timer
	start := time.Now()

	// Get file path from arguments
	filePath := os.Args[1]

	// Get ngram size from arguments
	ngramSize, _ := strconv.Atoi(os.Args[2])

	dat, err := os.ReadFile(filePath)
	check(err)

	ngrams := ngrams(strings.Split(string(dat), ""), ngramSize)

	// Sort results
	keys := make([]string, 0, len(ngrams))

	// Copy map for sort and count number of letters present
	for key := range ngrams {
		keys = append(keys, string(key))
	}

	// Sort by descending order
	sort.SliceStable(keys, func(i, j int) bool {
		return ngrams[keys[i]] > ngrams[keys[j]]
	})

	fmt.Printf("Ngrams of size %d\n", ngramSize)
	for _, k := range keys {
		fmt.Printf("Ngram: %s : Count: %d\n", k, ngrams[k])
	}

	/**
	fmt.Println("3 grams")
	for value, ngram := range ngrams(strings.Split(string(dat), ""), 3) {
		fmt.Printf("Ngram: %s : Count: %d\n", value, ngram)
	}
	**/

	// calculate to exe time
	fmt.Printf("\nFile analysed in %s\n", time.Since(start))
}
