# Simple implementation of decryption tool that uses DES and E-DES to decrypt a text
The tool uses DES and E-DES to decrypt a text given in stdin

## Compilation
To compile the tool, run the following command:
```bash
go build -o decrypt decrypt.go && chmod +x decrypt
```

## Usage

### Example of running the code with E-DES
```bash
cat test.txt | ./decrypt
```

```bash
echo 'hello world' | ./decrypt
```

### Example of running the code with default DES
```bash
cat test.txt | ./decrypt -d
```

```bash
echo 'hello world' | ./decrypt -d
```

### Example of running the code with E-DES and a key PATH
```bash
cat test.txt | ./decrypt -k PATH
```