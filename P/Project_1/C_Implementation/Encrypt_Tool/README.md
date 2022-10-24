# Simple implementation of encryption tool that uses DES and E-DES to encrypt a text
The tool uses DES and E-DES to encrypt a text given in stdin

### To run the code execute the following
```bash
cc -O2 -g encrypt.c -o encrypt -lssl -lcrypto -lm && chmod +x encrypt && cat test.txt | ./encrypt && rm encrypt
```

### Example of running the code with default DES
```bash
cc -O2 -g encrypt.c -o encrypt -lssl -lcrypto -lm && chmod +x encrypt && cat test.txt | ./encrypt -d && rm encrypt
```