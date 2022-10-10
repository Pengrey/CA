import os
from cryptography.hazmat.primitives.ciphers import Cipher , algorithms , modes
from cryptography.hazmat.primitives import padding
from cryptography.hazmat.backends import default_backend
import argparse

def cipher(keyFile, inputFile, outputFile):
    # Read key bytes from key file ( into variable key )
    with open(keyFile, "rb") as key_file:
        key = key_file.read()


    # Setup cipher : AES in CBC mode , w/ a random IV and PKCS #7 padding ( similar to PKCS #5)
    iv = os.urandom( algorithms.AES.block_size // 8 );
    cipher = Cipher( algorithms.AES( key ) , modes.CBC( iv ) , default_backend() )
    encryptor = cipher.encryptor()
    padder = padding.PKCS7( algorithms.AES.block_size ).padder()

    # Open input file for reading and output file for writing
    with open(inputFile, "rb") as in_file, open(outputFile, "wb") as out_file:

        # Write the contents of iv in the output file
        out_file.write(iv)

        while True : # Cicle to repeat while there is data left on the input file
            # Read a chunk of the input file to the plaintext variable
            plaintext = in_file.read(16)

            if not plaintext :
                ciphertext = encryptor.update( padder.finalize() )
                # Write the contents of ciphertext in the output file
                out_file.write(ciphertext)
                break
            else :
                ciphertext = encryptor.update( padder.update( plaintext ) )
                
                # Write the ciphertext in the output file
                out_file.write(ciphertext)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Ciphering tool')
    parser.add_argument("-k","--keyfile", help="File where the key is saved")
    parser.add_argument("-i","--infile", help="File to be ciphered")
    parser.add_argument("-o","--outfile", help="File to insert ciphertext")
    args = parser.parse_args()
    cipher(args.keyfile, args.infile, args.outfile)