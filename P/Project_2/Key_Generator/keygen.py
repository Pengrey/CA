import os
from Crypto.PublicKey import RSA
from cryptography.hazmat.primitives.asymmetric import ec
from cryptography.hazmat.primitives import serialization

def generate_RSA_key(key_size):
    key = RSA.generate(key_size)

    # The keys are saved in the format: key_<key_size>.pem
    f = open("keys/key_" + str(key_size) + ".pem", "wb")
    f.write(key.exportKey('PEM'))
    f.close()

def generate_ECC_key(curve):
    private_key = ec.generate_private_key(curve)

    # The keys are saved in the format: key_<curve>.pem
    f = open("keys/key_" + curve.name + ".pem", "wb")
    f.write(private_key.private_bytes(
        encoding=serialization.Encoding.PEM,
        format=serialization.PrivateFormat.PKCS8,
        encryption_algorithm=serialization.NoEncryption()
    ))
    f.close()

def main():
    # The keys are saved in the folder: Key_Generator/keys
    if not os.path.exists("keys"):
        os.makedirs("keys")

    # Generate 3 different keys, with the key sizes: 1024, 2048, 4096 bits
    generate_RSA_key(1024)
    generate_RSA_key(2048)
    generate_RSA_key(4096)

    # Generate 3 keys with the curve types: NIST P, K(oblitz) and B curves
    # Allowing 3 different curve sizes: small, medium and large keys

    # NIST P curves
    generate_ECC_key(ec.SECP192R1())
    generate_ECC_key(ec.SECP256R1())
    generate_ECC_key(ec.SECP521R1())

    # NIST K curves
    generate_ECC_key(ec.SECT163K1())
    generate_ECC_key(ec.SECT283K1())
    generate_ECC_key(ec.SECT571K1())

    # B curves
    generate_ECC_key(ec.SECT163R2())
    generate_ECC_key(ec.SECT283R1())
    generate_ECC_key(ec.SECT571R1())
      

if __name__ == "__main__":
    main()