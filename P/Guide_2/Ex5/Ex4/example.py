import os
from cryptography.hazmat.primitives.ciphers import Cipher , algorithms , modes
from cryptography.hazmat.primitives import padding
from cryptography.hazmat.backends import default_backend
import argparse

def cipher(keyFile, inputFile, outputFile, algo):
    # Read key bytes from key file ( into variable key )
    with open(keyFile, "rb") as key_file:
        key = key_file.read()

    # Open input file for reading and output file for writing
    with open(inputFile, "rb") as in_file, open(outputFile, "wb") as out_file:

        if algo == "AES":
            # Get iv from file
            iv = in_file.read(16)

            cipher = Cipher( algorithms.AES( key ) , modes.CBC( iv ) , default_backend() )
            decryptor = cipher.decryptor()
        else:
            # Get iv from file
            iv = in_file.read(8)

            cipher = Cipher( algorithms.TripleDES( key ) , modes.CBC( iv ) , default_backend() )
            decryptor = cipher.decryptor()

        total_bytes = os.path.getsize(inputFile)

        if algo == "AES":
            read_bytes = 16
        else:
            read_bytes = 8

        while True : # Cicle to repeat while there is data left on the input file
            if algo == "AES":
                cgram = in_file.read(16)
            else:
                cgram = in_file.read(8)
            read_bytes += len(cgram)

            if read_bytes == total_bytes:
                text = decryptor.update(cgram)
                padding = text[-1]

                if algo == "AES":
                    text = text[0:16 - padding]
                else:
                    text = text[0:8 - padding]
                    
                out_file.write(text)
                break

            text = decryptor.update(cgram)
            out_file.write(text)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Deciphering tool')
    parser.add_argument("-k","--keyfile", help="File where the key is saved")
    parser.add_argument("-i","--infile", help="File to be deciphered")
    parser.add_argument("-o","--outfile", help="File to insert plaintext")
    parser.add_argument("-a","--algo", help="Algorithm to be used (AES or DES)")
    args = parser.parse_args()
    cipher(args.keyfile, args.infile, args.outfile, args.algo)