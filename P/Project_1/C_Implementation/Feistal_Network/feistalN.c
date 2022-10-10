#include <stdio.h>

int permutation(char charToPermutate) {
    return 0;
}

int feistalIter(char firstSlice [], char secondSlice []) {
    int BLOCK_SIZE = 8;

    // Iterate over each char
    for( int i = 0 ; i < BLOCK_SIZE/2; i += 1 ){
        //printf("firstSlice: %c, secondSlice: %c\n", firstSlice[i], secondSlice[i]); // DEBUG

        char firstChar = firstSlice[i];
        char secondChar = secondSlice[i];

        // Apply permutation
        permutation(secondChar);

        // Apply XOR Li with fi and switch position between slices
        firstSlice[i] = secondChar;
        secondSlice[i] = firstChar ^ secondChar;
        //printf("Iter: %d, firstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]); // DEBUG
    }

    return 0;
}

int main() {
    int NUM_ITERATIONS = 16;

    char firstSlice[] = "IIIIIIII";
    char secondSlice[] = "OOOOOOOO";

    for( int i = 0 ; i < NUM_ITERATIONS; i += 1 ){
        //printf("Iter: %d\nFirstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]); // DEBUG
        feistalIter(firstSlice, secondSlice);
    }

    /** Check for unwind of the feistal Network

    for( int i = 0 ; i < NUM_ITERATIONS; i += 1 ){
        //printf("Iter: %d\nFirstSlice: %d, secondSlice: %d\n", i, firstSlice[0], secondSlice[0]); // DEBUG
        feistalIter(secondSlice, firstSlice);
    }

    for( int i = 0 ; i < 9; i += 1 ){
        printf("FirstSlice: %c, secondSlice: %c\n", firstSlice[0], secondSlice[0]); // DEBUG
    }
    **/

    return 0;
}