# Simple edes_enc test for the C implementation of the edes_enc tool

## Test E-DES encryption with the C implementation of the edes_enc tool
```bash
cc -O3 -o edes_enc edes_enc.c -lssl -lcrypto && chmod +x edes_enc && ./edes_enc > randomValuesEnc && rm edes_enc
```

## Test E-DES decryption with the C implementation of the edes_dec tool
```bash
cc -O3 -o edes_dec edes_dec.c -lssl -lcrypto && chmod +x edes_dec && ./edes_dec > /dev/null && rm edes_dec
```