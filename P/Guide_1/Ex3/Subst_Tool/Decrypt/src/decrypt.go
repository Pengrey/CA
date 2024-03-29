package main

import (
	"fmt"
	"os"
	"time"
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

	dat, err := os.ReadFile(filePath)
	check(err)

	// Get alphabet size from arguments
	alphabet := os.Args[2]

	if len(alphabet) < 20 {
		panic("Alphabet two small")
	}

	// Save alphabet
	normalAlpha := []string{"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z", "0", "1", "2", "3", "4", "5", "6", "7", "8", "9"}
	alphabetMap := make(map[string]string)

	// Prepare alphabet
	for index, letter := range alphabet {
		alphabetMap[string(letter)] = normalAlpha[index]
	}

	// Decryption
	decrypted := ""

	for _, letter := range dat {
		decrypted += alphabetMap[string(letter)]
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
	f.WriteString(decrypted)

	// calculate to exe time
	fmt.Printf("\nFile Decrypted in %s\n", time.Since(start))
}
