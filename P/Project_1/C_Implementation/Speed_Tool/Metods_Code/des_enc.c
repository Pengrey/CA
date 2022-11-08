#include <stdio.h>
#include <stdlib.h>
#include <string.h>
#include <time.h>
#include <openssl/sha.h>
#include <nettle/des.h>

typedef unsigned char uint8_t;

void readKey(char *key, int keySize, char *keyFile) {
    FILE *fp = fopen(keyFile, "r");
    if (fp == NULL) {
        printf("Error opening key file\n");
        exit(1);
    }

    // Check error
    if (fread(key, sizeof(char), keySize, fp) != keySize) {
        printf("Error reading key file\n");
        exit(1);
    }
    fclose(fp);
}

void readData(char *data, int dataSize, char *dataFile) {
    FILE *fp = fopen(dataFile, "r");
    if (fp == NULL) {
        printf("Error opening data file\n");
        exit(1);
    }

    // Check error
    if (fread(data, sizeof(char), dataSize, fp) != dataSize) {
        printf("Error reading data file\n");
        exit(1);
    }
    fclose(fp);
}

void getDESBlock(uint8_t key[8], uint8_t data[8]) {
    struct des_ctx ctx;
    uint8_t out[8];

    des_set_key(&ctx, key);
    
    des_encrypt(&ctx, 8, out, data);
    
    for (int i = 0; i < 8; i++) {
        printf("%c", out[i]);
    }
}

void getDESCipher(uint8_t key[8], char *ciphertext)
{
    //printf("[!] Using DES\n");

    uint8_t data[8] = {0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00, 0x00};

    char c;
    int charIndex = 0;

    // Read data char by char from 4KB of data
    for (int i = 0 ; i < 4096 ; i++){
        c = ciphertext[i];

        // Save input in data
        data[charIndex] = (uint8_t) c;

        charIndex++;

        // When we have full block we cipher them and we reset the clock
        if (charIndex == 8) {
            getDESBlock(key, data);
            charIndex = 0;
        }
    }
}

int main(int argc, char *argv[])
{
    // Read key of 256 bit from file
    unsigned char seed[65];
    readKey(seed, 32, "key.bin");

    uint8_t key[8];

    // Convert seed to 8 bytes
    for (int i = 0 ; i < 8 ; i++){
        key[i] = (uint8_t) seed[i];
    }
    
    // Read 4KB data from file
    char *data = malloc(4096);
    readData(data, 4096, "randomValues");

    // Start timer to measure time taken to encrypt
    clock_t start = clock();

    // Encrypt the data
    getDESCipher(key, data);

            // Stop timer
    clock_t end = clock();

    // Save time taken to encrypt in a file
    double time_taken = ((double) (end - start)) / CLOCKS_PER_SEC;

    // Read the time taken to encrypt from the file
    FILE *fp = fopen("desEncTime.txt", "r");
    // Handle error if file is not found
    if (fp == NULL) {
        printf("Error opening file!\n");
        exit(1);
    }

    double time_taken2;

    // Check fscanf() return value
    if (fscanf(fp, "%lf", &time_taken2) != 1) {
        printf("Error reading file!\n");
        exit(1);
    }


    fclose(fp);

    // Print the time taken to encrypt if it is less than the time taken to encrypt in the previous run
    if (time_taken < time_taken2) {
        fp = fopen("desEncTime.txt", "w");
        fprintf(fp, "%lf", time_taken);
        fclose(fp);
    }
}