package main

import (
	"fmt"
	"math"
	"math/rand"
	"os"
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

	dat, err := os.ReadFile(filePath)
	check(err)

	// Save alphabet
	alphabet := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	normalAlpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	alphabetMap := make(map[string]string)

	// Shuffle alphabet
	rand.Seed(time.Now().UnixNano())
	rand.Shuffle(len(alphabet), func(i, j int) {
		alphabet[i], alphabet[j] = alphabet[j], alphabet[i]
	})

	// Print alphabet used
	fmt.Print("Alphabet used: ")
	for index, letter := range alphabet {
		fmt.Print(letter)
		alphabetMap[normalAlpha[index]] = letter
	}

	// Encryption
	encrypted := ""

	for _, letter := range dat {
		encrypted += alphabetMap[string(letter)]
	}

	// Save encrypted
	// create the file
	f, err := os.Create("../Output/out.txt")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer f.Close()

	// write a string
	f.WriteString(encrypted)

	// calculate to exe time
	fmt.Printf("\nFile encrypted in %s\n", time.Since(start))
}
