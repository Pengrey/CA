package main

import (
	"crypto"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"io/ioutil"
	"log"
	"os"
	"time"
)

// readPrivateKeys reads the private keys from the files and returns a slice of private keys
func readPrivateKeys() []*rsa.PrivateKey {
	// Create a slice of private keys
	var privateKeys []*rsa.PrivateKey

	// Slice of file names
	fileNames := []string{"key_1024.pem", "key_2048.pem", "key_4096.pem"}

	// Read the private keys from the files and store them in the slice
	for _, fileName := range fileNames {
		// Open the file
		file, err := os.Open("./../Key_Generator/keys/" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		// Read the file
		fileBytes, err := ioutil.ReadAll(file)
		if err != nil {
			log.Fatal(err)
		}

		// Decode the file
		block, _ := pem.Decode(fileBytes)

		// Parse the private key
		privateKey, err := x509.ParsePKCS1PrivateKey(block.Bytes)
		if err != nil {
			log.Fatal(err)
		}

		// Append the private key to the slice
		privateKeys = append(privateKeys, privateKey)
	}

	return privateKeys
}

func signAndVerifyPSS(digest []byte, privateKey *rsa.PrivateKey) {
	// Sign the data
	signature, err := rsa.SignPSS(rand.Reader, privateKey, crypto.SHA256, digest, nil)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the signature
	err = rsa.VerifyPSS(&privateKey.PublicKey, crypto.SHA256, digest, signature, nil)
	if err != nil {
		log.Fatal(err)
	}
}

func signAndVerifyPKCS1(digest []byte, privateKey *rsa.PrivateKey) {
	// Sign the data
	signature, err := rsa.SignPKCS1v15(rand.Reader, privateKey, crypto.SHA256, digest)
	if err != nil {
		log.Fatal(err)
	}

	// Verify the signature
	err = rsa.VerifyPKCS1v15(&privateKey.PublicKey, crypto.SHA256, digest, signature)
	if err != nil {
		log.Fatal(err)
	}
}

// Function receives a private key and returns the performance of PKCS #1 v1.5 padding
func getPerfPKCS1(privateKeys []*rsa.PrivateKey) {
	// Keys are 1024, 2048, 4096 bits
	keysizes := []int{1024, 2048, 4096}

	// Perform a loop with a certain number of iterations
	niterations := 10

	// Execute the operation a certain number of times
	mtimes := 1000

	// for each private key
	for key := 0; key < len(privateKeys); key++ {
		// Generate random data to sign from urandom
		randomData := make([]byte, 100)
		_, err := rand.Read(randomData)
		if err != nil {
			log.Fatal(err)
		}

		// Set hasher
		hash := crypto.SHA256.New()

		// Generate the digest
		hash.Write(randomData)
		digest := hash.Sum(nil)

		// Keep track of the quickest time observed to a big number
		quickestTime := 9999.0

		for n := 0; n < niterations; n++ {
			// Measure the time it takes to execute the consecutive operations
			start := time.Now()
			for m := 0; m < mtimes; m++ {
				// Sign and verify data
				signAndVerifyPKCS1(digest, privateKeys[key])
			}
			elapsed := time.Since(start)

			// Update the minimum time observed
			if elapsed.Seconds() < quickestTime {
				quickestTime = elapsed.Seconds()
			}
		}

		// Calculate the time each operation takes
		operationTime := quickestTime / float64(mtimes)
		fmt.Printf("Key: %d, Padding: %s, Time: %f\n", keysizes[key], "PKCS1", operationTime)
	}
}

// Function receives a private key and returns the performance of PSS padding
func getPerfPSS(privateKeys []*rsa.PrivateKey) {
	// Keys are 1024, 2048, 4096 bits
	keysizes := []int{1024, 2048, 4096}

	// Perform a loop with a certain number of iterations
	niterations := 10

	// Execute the operation a certain number of times
	mtimes := 1000

	// for each private key
	for key := 0; key < len(privateKeys); key++ {
		// Generate random data to sign from urandom
		randomData := make([]byte, 100)
		_, err := rand.Read(randomData)
		if err != nil {
			log.Fatal(err)
		}

		// Set hasher
		hash := crypto.SHA256.New()

		// Generate the digest
		hash.Write(randomData)
		digest := hash.Sum(nil)

		// Keep track of the quickest time observed to a big number
		quickestTime := 9999.0

		for n := 0; n < niterations; n++ {
			// Measure the time it takes to execute the consecutive operations
			start := time.Now()
			for m := 0; m < mtimes; m++ {
				// Sign and verify data
				signAndVerifyPSS(digest, privateKeys[key])
			}
			elapsed := time.Since(start)

			// Update the minimum time observed
			if elapsed.Seconds() < quickestTime {
				quickestTime = elapsed.Seconds()
			}
		}

		// Calculate the time each operation takes
		operationTime := quickestTime / float64(mtimes)
		fmt.Printf("Key: %d, Padding: %s, Time: %f\n", keysizes[key], "PSS", operationTime)
	}
}

func main() {
	// Read the private keys from the files and store them in a slice
	privateKeys := readPrivateKeys()

	// Get performance of PKCS #1 v1.5 padding
	getPerfPKCS1(privateKeys)

	// Get performance of PSS padding
	getPerfPSS(privateKeys)
}
