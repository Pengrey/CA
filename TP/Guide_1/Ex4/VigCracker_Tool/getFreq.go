package main

// This code was made by implementing the code from https://github.com/1r0dm480/Vigenere-Cipher-Breaker in golang
import (
	"fmt"
	"os"
	"strconv"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
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
	var stringSlices []string
	totalLettersAdded := 0
	var lettersCounted = map[string]int{
		"a": 0,
		"b": 0,
		"c": 0,
		"d": 0,
		"e": 0,
		"f": 0,
		"g": 0,
		"h": 0,
		"i": 0,
		"j": 0,
		"k": 0,
		"l": 0,
		"m": 0,
		"n": 0,
		"o": 0,
		"p": 0,
		"q": 0,
		"r": 0,
		"s": 0,
		"t": 0,
		"u": 0,
		"v": 0,
		"w": 0,
		"x": 0,
		"y": 0,
		"z": 0,
	}

	// Prepare segments with key length
	for i := 0; i+keyLength < len(cyphertext); {
		stringSlices = append(stringSlices, cyphertext[i:i+keyLength])
		i += keyLength
	}

	// Check letter frequency in each slice
	for _, word := range stringSlices {
		// Prevent repeated analyses on the same slice
		analysedLetters := ""
		mostFrequent := ""

		for _, letter := range word {
			char := string(letter)
			// Skip repeated letters
			if strings.Contains(analysedLetters, char) {
				continue
			}

			if mostFrequent == "" || float32(strings.Count(word, char)) > float32(strings.Count(word, mostFrequent)) {
				mostFrequent = char
			}

			analysedLetters += char
		}

		// Add frequency to map
		lettersCounted[mostFrequent] += 1
		totalLettersAdded += 1
	}

	// Get average letter frequency for each letter
	for key, val := range lettersCounted {
		fmt.Println(key, (float32(val)/float32(totalLettersAdded))*100, "%")
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
