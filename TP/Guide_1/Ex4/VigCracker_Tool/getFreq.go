package main

// This code was made by implementing the code from https://github.com/1r0dm480/Vigenere-Cipher-Breaker in golang
import (
	"fmt"
	"os"
	"sort"
	"strconv"
	"strings"
	"sync"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// FreqMap records the frequency of each rune in a given text.
type FreqMap map[rune]int

// Frequency counts the frequency of each rune in a given text and returns this
// data as a FreqMap.
func Frequency(s string) FreqMap {
	m := FreqMap{}
	for _, r := range s {
		m[r]++
	}
	return m
}
func ConcurrentFrequency(texts []string) FreqMap {
	res := make(FreqMap)
	ch := make(chan FreqMap, 10)
	var wg sync.WaitGroup
	wg.Add(len(texts))
	for _, text := range texts {
		go func(t string) {
			ch <- Frequency(t)
			wg.Done()
		}(text)
	}
	go func() {
		wg.Wait()
		close(ch)
	}()
	for freqmap := range ch {
		for letter, freq := range freqmap {
			res[letter] += freq
		}
	}
	return res
}

func sum(array []float32) float32 {
	result := float32(0)
	for _, v := range array {
		result += v
	}
	return result
}

// Get frequency of letters in text
func get_freq(cyphertext string, keyLength int) {

	// Prepare slice
	slice := ""
	for j := 0; j < len(cyphertext); {
		slice += string(cyphertext[j])

		j += keyLength
	}

	// Calculate letter frequency in slice
	result := ConcurrentFrequency(strings.Split(slice, ""))

	// Sort results
	keys := make([]string, 0, len(result))

	// Total number of letters
	totalCount := 0

	// Copy map for sort and count number of letters present
	for key := range result {
		keys = append(keys, string(key))
		totalCount += result[key]
	}

	// Sort by descending order
	sort.SliceStable(keys, func(i, j int) bool {
		return result[[]rune(keys[i])[0]] > result[[]rune(keys[j])[0]]
	})

	// Print results
	fmt.Println("+--------+-------+--------+------------------------------------------------------------------------------------------------------+")
	fmt.Println("| Letter | Count | Percen | Graph                                                                                                |")
	fmt.Println("+--------+-------+--------+------------------------------------------------------------------------------------------------------+")
	for _, k := range keys {
		percentage := float32(result[[]rune(k)[0]]) / float32(totalCount) * 100
		graph := ""
		for i := 0; int(percentage) > i; i++ {
			graph += "="
		}

		if int(percentage/10) < 1 {
			fmt.Printf("| %6v | %5d | 0%.2f %%| %-100s |\n", k, result[[]rune(k)[0]], percentage, graph)
		} else {
			fmt.Printf("| %6v | %5d | %.2f %%| %-100s |\n", k, result[[]rune(k)[0]], percentage, graph)
		}
	}
	fmt.Println("+--------+-------+--------+------------------------------------------------------------------------------------------------------+")
}

func main() {

	// Performance timer
	start := time.Now()

	// Get file path from arguments
	filePath := os.Args[1]

	// Get key length
	keyLength, _ := strconv.Atoi(os.Args[2])

	dat, err := os.ReadFile(filePath)
	check(err)

	// Get frequency from file text
	get_freq(string(dat), keyLength)

	// calculate to exe time
	fmt.Printf("\nFile processed in %s\n", time.Since(start))
}
