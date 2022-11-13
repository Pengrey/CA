# Simple implementation of E-DES
The code was made to implemente a simple implementation of E-DES with 64 bits and with 16 permutations

### To compile the code, run the following command:
```bash
cc -o eDES eDES.c -lssl -lcrypto -lnettle -lm && chmod +x eDES
```

### To run the code, run the following command:
```bash
./eDES
```