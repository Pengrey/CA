# Simple implementation of encryption tool that uses DES and E-DES to encrypt a text
The tool uses DES and E-DES to encrypt a text given in stdin

## Compilation
To compile the tool, run the following command:
```bash
cc -O2 -g encrypt.c -o encrypt -lssl -lcrypto -lnettle -lm && chmod +x encrypt
```

## Usage

### Example of running the code with E-DES
```bash
cat test.txt | ./encrypt
```

```bash
echo 'hello world' | ./encrypt
```

### Example of running the code with default DES
```bash
cat test.txt | ./encrypt -d
```

```bash
echo 'hello world' | ./encrypt -d
```

### Example of running the code with E-DES and a key PATH
```bash
cat test.txt | ./encrypt -k PATH
```