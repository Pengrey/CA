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
	alphabetLetters := []rune{'a', 'b', 'c', 'd', 'e', 'f', 'g', 'h', 'i', 'j', 'k', 'l', 'm', 'n', 'o', 'p', 'q', 'r', 's', 't', 'u', 'v', 'w', 'x', 'y', 'z'}
	//commonLetters := []rune{'a', 'e', 'o', 's', 'r', 'i', 'n', 'd', 'm', 'u', 't', 'c', 'l', 'p', 'v', 'g', 'h', 'q', 'b', 'f', 'z', 'j', 'x', 'k', 'w', 'y'} // PT
	commonLetters := []rune{'e', 't', 'a', 'o', 'i', 'n', 's', 'h', 'r', 'd', 'l', 'c', 'u', 'm', 'w', 'f', 'g', 'y', 'p', 'b', 'v', 'k', 'j', 'x', 'q', 'z'} // ENG
	//commonLetters := []rune{'a', 'i', 't', 'n', 'e', 's', 'l', 'o', 'k', 'u', 'm', 'h', 'v', 'r', 'j', 'p', 'y', 'd', 'g', 'c', 'b', 'f', 'w', 'z', 'x', 'q'} // FIN

	// Number of iterations equal to the length of the supposed key
	for interation := 0; interation < keyLength; interation++ {

		// Prepare slice
		slice := ""
		for j := interation; j < len(cyphertext); {
			slice += string(cyphertext[j])

			j += keyLength
		}

		// Calculate letter frequency in slice
		result := ConcurrentFrequency(strings.Split(slice, ""))

		// Sort results
		keys := make([]string, 0, len(result))

		// Copy map for sort and count number of letters present
		for key := range result {
			keys = append(keys, string(key))
		}

		// Sort by descending order
		sort.SliceStable(keys, func(i, j int) bool {
			return result[[]rune(keys[i])[0]] > result[[]rune(keys[j])[0]]
		})

		cypheredLetters := []string{}
		var totalShift int32

		for _, k := range keys {
			cypheredLetters = append(cypheredLetters, k)
		}

		// Get average shift
		equalPeaks := 3
		for index, letter := range cypheredLetters[:equalPeaks] {
			totalShift += []rune(letter)[0] - commonLetters[index]
		}

		avgShift := int(totalShift / int32(len(cypheredLetters[:equalPeaks])))

		if avgShift <= 0 {
			fmt.Print(string(alphabetLetters[len(alphabetLetters)+avgShift-1]))
		} else {
			fmt.Print(string(alphabetLetters[avgShift]))
		}
	}
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
