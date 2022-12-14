package main

import (
	"crypto/ecdsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
)

func readPrivateKey(filepath string) (*ecdsa.PrivateKey, error) {
	data, err := ioutil.ReadFile(filepath)
	if err != nil {
		return nil, err
	}

	block, _ := pem.Decode(data)
	if block == nil {
		return nil, errors.New("failed to decode PEM block containing the key")
	}

	privateKey, err := x509.ParseECPrivateKey(block.Bytes)
	if err != nil {
		return nil, err
	}

	return privateKey, nil
}

// readPrivateKeys reads the private keys from the files and returns a slice of private keys
func readPrivateKeys() []*ecdsa.PrivateKey {
	// Create a slice of private keys
	var privateKeys []*ecdsa.PrivateKey

	// Slice of file names
	fileNames := []string{"key_secp192r1.pem", "key_secp256r1.pem", "key_secp521r1.pem", "key_sect163k1.pem", "key_sect283k1.pem", "key_sect571k1.pem", "key_sect163r2.pem", "key_sect283r1.pem", "key_sect571r1.pem"}

	// Read the private keys from the files and store them in the slice
	for _, fileName := range fileNames {
		// Read the private key
		privateKey, err := readPrivateKey("./../Key_Generator/keys/" + fileName)
		if err != nil {
			log.Fatal(err)
		}

		// Append the private key to the slice
		privateKeys = append(privateKeys, privateKey)
	}

	return privateKeys
}

func main() {
	// Read the private keys
	privateKeys := readPrivateKeys()

	// Print the keys
	for _, privateKey := range privateKeys {
		fmt.Println(privateKey)
	}

}
