package main

func permutation(charToPermutate byte) byte {
	return charToPermutate
}

func feistalIter(firstSlice []byte, secondSlice []byte) ([]byte, []byte) {
	BLOCK_SIZE := 8

	// Iterate over each char
	for i := 0; i < BLOCK_SIZE/2; i++ {
		//fmt.Printf("firstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]) // DEBUG

		firstChar := firstSlice[i]
		secondChar := secondSlice[i]

		// Apply permutation
		secondChar = permutation(secondChar)

		// Apply XOR Li with fi and switch position between slices
		firstSlice[i] = secondChar
		secondSlice[i] = firstChar ^ secondChar
		//fmt.Printf("Iter: %d, firstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]); // DEBUG

	}

	return firstSlice, secondSlice
}

func main() {
	NUM_ITERATIONS := 16

	firstSlice := []byte("IIIIIIII")
	secondSlice := []byte("OOOOOOOO")

	for i := 0; i < NUM_ITERATIONS; i++ {
		//fmt.Printf("Iter: %d\nFirstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]) // DEBUG
		firstSlice, secondSlice = feistalIter(firstSlice, secondSlice)
	}

	/** // Check for unwind of the feistal Network

	for i := 0; i < NUM_ITERATIONS; i++ {
		//printf("Iter: %d\nFirstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]); // DEBUG
		feistalIter(secondSlice, firstSlice)
	}

	for i := 0; i < 9; i += 1 {
		fmt.Printf("FirstSlice: %c, secondSlice: %c\n", firstSlice[0], secondSlice[0]) // DEBUG
	}
	**/
}
