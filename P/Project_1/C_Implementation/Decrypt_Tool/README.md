# Simple implementation of decryption tool that uses DES and E-DES to decrypt a text
The tool uses DES and E-DES to decrypt a text given in stdin

### To run the code execute the following
```bash
cc -O2 -g decrypt.c -o decrypt -lssl -lcrypto -lm && chmod +x decrypt && cat test.txt | ./decrypt && rm decrypt
```

### Example of running the code with default DES
```bash
cc -O2 -g decrypt.c -o decrypt -lssl -lcrypto -lm && chmod +x decrypt && cat test.txt | ./decrypt -d && rm decrypt
```