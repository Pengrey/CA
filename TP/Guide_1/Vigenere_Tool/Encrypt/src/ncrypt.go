package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"

	"github.com/odysseus/vigenere"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	// Performance timer
	start := time.Now()

	// Get file path from arguments
	filePath := os.Args[1]

	// Get key from arguments
	key := os.Args[2]

	dat, err := os.ReadFile(filePath)
	check(err)

	// Replace non letters
	toReplace := regexp.MustCompile(`[^a-zA-Z0-9]`)

	removedNumbers := toReplace.ReplaceAllString(string(dat), "")

	// Encrypt
	encoded := vigenere.Encipher(removedNumbers, key)

	// Convert all to lowercase
	result := strings.ToLower(encoded)

	// Save encrypted
	// create the file
	f, err := os.Create("../Output/out.txt")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer f.Close()

	// write a string
	f.WriteString(result)

	// calculate to exe time
	fmt.Printf("\nFile encrypted in %s\n", time.Since(start))
}
