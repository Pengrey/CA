package main

import (
	"fmt"
	"os"
	"regexp"
	"strings"
	"time"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func removeAccents(s string) string {
	t := transform.Chain(norm.NFD, runes.Remove(runes.In(unicode.Mn)), norm.NFC)
	output, _, e := transform.String(t, s)
	if e != nil {
		panic(e)
	}
	return output
}

func main() {
	// Performance timer
	start := time.Now()

	// Get file path from arguments
	filePath := os.Args[1]

	dat, err := os.ReadFile(filePath)
	check(err)

	// Remove accents
	nonAccentText := removeAccents(string(dat))

	// Replace non letters
	toReplace := regexp.MustCompile(`[^a-zA-Z0-9]`)

	onlyLetterText := toReplace.ReplaceAllString(nonAccentText, "")

	// Convert all to lowercase
	result := strings.ToLower(onlyLetterText)

	// Print result
	// fmt.Print(result)

	// Save result to folder

	// create the file
	f, err := os.Create("./Prepared_Files/out.txt")
	if err != nil {
		fmt.Println(err)
	}
	// close the file with defer
	defer f.Close()

	// write a string
	f.WriteString(result)

	// calculate to exe time
	fmt.Printf("File prepared in %s\n", time.Since(start))
}
