# E-DES

## Abstract

This work consists of an implementation of a variant of the asymmetric key algorithm DES (Data Encryption Standard), named E-DES, with a key size of 256 bits and Feistel networks that are faster than the original algorithm.

This project aims to study a variant of DES, called E-DES. In this variant, Feistel Networks and S-Boxes are the only operations used to implement a cipher similar to DES. As with many other ciphers, DES relies on S-Boxes as a building block. It employs static S-Boxes, which raises concerns about hidden cryptoanalysis trapdoors. We will use variable, key-dependent S-Boxes in E-DES and longer, 256-bit keys compared to DES.


## Project Structure

The project is divided into two directories:

* C_Implementation: Contains the C implementation of the tools.
* Go_Implementation: Contains the Go implementation of the tools.

In each implementation, the tools are divided into four directories:

* Decrypt_Tool: Contains the tool to decrypt a file.
* Encrypt_Tool: Contains the tool to encrypt a file.
* E-DES: Contains the implementation of the E-DES cipher, in this directory, the implementation of the cipher is divided into two files:
    * Feistel_Network: Contains the implementation of the Feistel Network and their respective tests.
    * S_Box: Contains the implementation of the S-Boxes and their respective tests.
* Speed_Tool: Contains the tool to measure the speed of the cipher, in this directory there is present a file named randomValues, which contains a set of random values to be used in the speed test and a directory named Metods_Code, which contains the modules used to measure the speed of the cipher.

A report of the project is also present in the root directory of the project and is named CA.pdf.