package main

import (
	"crypto/des"
	"crypto/sha256"
	"fmt"
	"io/ioutil"
	"log"
	"strconv"
	"strings"
	"time"
)

func getKey(keyFile string) string {
	// Read file
	file, err := ioutil.ReadFile(keyFile)
	if err != nil {
		log.Fatal(err)
	}

	// Convert file to SHA256 hash
	hash := sha256.Sum256(file)

	// Return hash as string
	return fmt.Sprintf("%x", hash)
}

func getDESBlock(key string, data [8]uint8) {
	// Convert key to byte array
	keyBytes := []byte(key)

	// Create cipher
	c, err := des.NewCipher(keyBytes)
	if err != nil {
		log.Fatal(err)
	}

	// Convert [8]uint8 data to byte array
	dataBytes := []byte{data[0], data[1], data[2], data[3], data[4], data[5], data[6], data[7]}

	// Create byte array to store decrypted data
	decrypted := make([]byte, 8)

	// Decrypt data
	c.Decrypt(decrypted, dataBytes)

	// Print encrypted data
	fmt.Printf("%s", decrypted)
}

func getDESCipher(key string, ciphertext []byte) {

	//fmt.Println("Running DES")

	data := [8]uint8{0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00}

	charIndex := 0

	// Read data char by char from 4KB of data
	for _, c := range ciphertext {

		// Save input in data
		data[charIndex] = uint8(c)

		charIndex++

		// When we have full block we cipher them and we reset the clock
		if charIndex == 8 {
			getDESBlock(key, data)
			charIndex = 0
		}
	}
}

func main() {
	// Read key of 256 bit from file and generate the seed
	seed := strings.ToUpper(getKey("key.bin"))

	// Read 4KB data from file
	data, err := ioutil.ReadFile("randomValuesEnc")
	if err != nil {
		log.Fatal(err)
	}

	// Convert seed to string with size 8
	seed = seed[:8]

	// Start timer to measure time taken to encrypt
	start := time.Now()

	getDESCipher(seed, data)

	// Stop timer
	elapsed := time.Since(start)

	// Read the time taken to encrypt from the file edesEncTime.txt
	file, err := ioutil.ReadFile("desDecTime.txt")
	if err != nil {
		log.Fatal(err)
	}

	// Convert the time taken to encrypt from string to float64
	encTime, err := strconv.ParseFloat(string(file), 64)

	// Save the time taken to encrypt if it is less than the time taken to encrypt in the previous run
	if float64(elapsed.Seconds()) < encTime {
		ioutil.WriteFile("desDecTime.txt", []byte(strconv.FormatFloat(float64(elapsed.Seconds()), 'f', 6, 64)), 0644)
	}
}
