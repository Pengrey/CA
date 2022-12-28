package main

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"fmt"
	"math/big"
	"time"
)

func signAndVerify(digest []byte, key *ecdsa.PrivateKey, mtimes int) (float64, float64) {
	var r, s *big.Int
	// Sign the data
	start := time.Now()
	for m := 0; m < mtimes; m++ {
		r, s, _ = ecdsa.Sign(rand.Reader, key, digest)
	}
	end := time.Now()
	signTime := end.Sub(start).Seconds()

	// Verify
	start = time.Now()
	for m := 0; m < mtimes; m++ {
		ecdsa.Verify(&key.PublicKey, digest, r, s)
	}
	end = time.Now()
	vrfyTime := end.Sub(start).Seconds()

	return signTime, vrfyTime
}

func getPerf(keys []*ecdsa.PrivateKey) {
	// Perform a loop with a certain number of iterations
	niterations := 100

	// Execute the operation a certain number of times
	mtimes := 1000

	for k := range keys {
		// Generate random data to sign
		data := make([]byte, 1024)
		rand.Read(data)

		// Generate the digest
		digest := sha256.Sum256(data)

		// Keep track of the quickest time observed to a big number
		minSignTime := 9999.0
		minVrfyTime := 9999.0

		// Execute the operation a certain number of times
		for n := 0; n < niterations; n++ {
			// Sign and verify data
			signTime, vrfyTime := signAndVerify(digest[:], keys[k], mtimes)

			// Update the minimum time observed
			if signTime < minSignTime {
				minSignTime = signTime
			}
			if vrfyTime < minVrfyTime {
				minVrfyTime = vrfyTime
			}
		}

		// Calculate the time each operation takes
		signTime := minSignTime / float64(mtimes)
		vrfyTime := minVrfyTime / float64(mtimes)

		fmt.Printf("Curve: %s, Key size: %d, Sign: %f, Verify: %f\n", keys[k].Params().Name, keys[k].Params().BitSize, signTime, vrfyTime)
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
