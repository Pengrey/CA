package main

// This code was made by implementing the code from https://github.com/1r0dm480/Vigenere-Cipher-Breaker in golang
import (
	"fmt"
	"os"
	"strings"
	"time"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

// Returns the Index of Councidence for the "section" of ciphertext given
func get_index_c(ciphertext string) float64 {
	alphabet := "abcdefghijklmnopqrstuvwxyz"

	N := float64(len(ciphertext))
	frequency_sum := 0.0

	// Using Index of Coincidence formula
	for index := range alphabet {
		letter := string(alphabet[index])
		repeated := strings.Count(ciphertext, letter)
		frequency_sum += float64(repeated * (repeated - 1))
	}
	// Using Index of Coincidence formula
	ic := frequency_sum / (N * (N - 1))
	return ic
}

func get_key_length(ciphertext string) int {
	var ic_table []float64

	MAX_KEY_LENGTH_GUESS := 20

	// Splits the ciphertext into sequences based on the guessed key length from 0 until the max key length guess (20)
	// Ex. guessing a key length of 2 splits the "12345678" into "1357" and "2468"
	// This procedure of breaking ciphertext into sequences and sorting it by the Index of Coincidence
	// The guessed key length with the highest IC is the most porbable key length
	for guess_len := 0; guess_len < MAX_KEY_LENGTH_GUESS; guess_len++ {
		ic_sum := 0.0
		avg_ic := 0.0
		for i := 0; i < guess_len; i++ {
			sequence := ""
			// breaks the ciphertext into sequences
			for j := 0; j < len(ciphertext)-i; {
				sequence += string(ciphertext[i+j])

				j += guess_len
			}
			ic_sum += get_index_c(sequence)
		}
		// obviously don't want to divide by zero
		if guess_len != 0 {
			avg_ic = ic_sum / float64(guess_len)
		}
		ic_table = append(ic_table, avg_ic)
	}

	// Initialize
	var largerNumber, temp float64

	// gets highest Index of Coincidence
	for _, element := range ic_table {
		if element > temp {
			temp = element
			largerNumber = temp
		}
	}

	// returns the index of the highest Index of Coincidence (most probable key length)
	for i, val := range ic_table {
		if val == largerNumber {
			return i
		}
	}
	// Return default
	return 0
}

func main() {
	// Performance timer
	start := time.Now()

	// Get file path from arguments
	filePath := os.Args[1]

	dat, err := os.ReadFile(filePath)
	check(err)

	fmt.Printf("Estimated key size: %d", get_key_length(string(dat)))

	// calculate to exe time
	fmt.Printf("\nFile processed in %s\n", time.Since(start))
}
