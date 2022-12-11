import os
from Crypto.PublicKey import RSA

def generate_key(key_size):
    key = RSA.generate(key_size)

    # The keys are saved in the format: key_<key_size>.pem
    f = open("keys/key_" + str(key_size) + ".pem", "wb")
    f.write(key.exportKey('PEM'))
    f.close()

def main():
    # The keys are saved in the folder: Key_Generator/keys
    if not os.path.exists("keys"):
        os.makedirs("keys")

    # Generate 3 different keys, with the key sizes: 1024, 2048, 4096 bits
    generate_key(1024)
    generate_key(2048)
    generate_key(4096)

if __name__ == "__main__":
    main()