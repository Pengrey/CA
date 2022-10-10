import argparse
from cryptography . hazmat . primitives import hashes
from cryptography . hazmat . primitives . kdf . pbkdf2 import PBKDF2HMAC
from cryptography . hazmat . backends import default_backend

def generate(pwd, outFile):
    # The PBKDF2 generator of Python receives as input the number of bytes to generate , instead of bits
    salt = b'\ x00 '
    kdf = PBKDF2HMAC( hashes.SHA1() , 16 , salt , 1000 , default_backend() )
    key = kdf.derive( bytes( pwd , 'UTF -8 ' ) )

    # Write key in a file
    with open(outFile, "wb") as key_file: key_file.write(key)

if __name__ == '__main__':
    parser = argparse.ArgumentParser(description='Ciphering tool')
    parser.add_argument("-p","--pwd", help="Password to be used")
    parser.add_argument("-o","--outfile", help="File to insert key")
    args = parser.parse_args()
    generate(args.pwd, args.outfile)