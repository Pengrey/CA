package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"time"
)

func signAndVerify(digest []byte, key *ecdsa.PrivateKey) {
	// Sign
	r, s, _ := ecdsa.Sign(rand.Reader, key, digest)

	// Verify
	ecdsa.Verify(&key.PublicKey, digest, r, s)
}

func getPerf(keys []*ecdsa.PrivateKey) {
	// Perform a loop with a certain number of iterations
	niterations := 10

	// Execute the operation a certain number of times
	mtimes := 1000

	for k := range keys {
		// Generate random data to sign
		data := make([]byte, 1024)
		rand.Read(data)

		// Generate the digest
		digest := sha256.Sum256(data)

		// Keep track of the quickest time observed
		quickest := 1000000000.0

		// Execute the operation a certain number of times
		for n := 0; n < niterations; n++ {
			// Measure the time it takes to execute the consecutive operations
			start := time.Now()
			for m := 0; m < mtimes; m++ {
				signAndVerify(digest[:], keys[k])
			}
			elapsed := time.Since(start)

			// Update the minimum time observed
			if elapsed.Seconds() < quickest {
				quickest = elapsed.Seconds()
			}
		}

		// Calculate the time each operation takes
		elapsed := quickest / float64(mtimes)
		fmt.Printf("Curve: %s, Key size: %d, Time: %f \n", keys[k].Params().Name, keys[k].Params().BitSize, elapsed)
	}
}

func main() {
	// Load the private keys
	eccList := []elliptic.Curve{
		elliptic.P224(),
		elliptic.P256(),
		elliptic.P384(),
		elliptic.P521(),
	}

	eccKeys := make([]*ecdsa.PrivateKey, len(eccList))
	for i, curve := range eccList {
		eccKeys[i], _ = ecdsa.GenerateKey(curve, rand.Reader)
	}

	// Get performance
	getPerf(eccKeys)

}
