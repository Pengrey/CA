import os
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives import hashes

# Read the private keys from the files
keys = []
for sizeB in [1024, 2048, 4096]:
    with open(f"./keys/key_{sizeB}.pem", "rb") as key_file:
        keys.append(serialization.load_pem_private_key(
        key_file.read(),
        password=None,
        ))
    
# PKCS #1 padding
# Generate random data to sign
data = os.urandom(1024)

# Sign the data with each key using PKCS #1 padding
for key in keys:
    signature = key.sign(
        data,
        padding.PKCS1v15(),
        hashes.SHA256()
    )

    # Verify the signature
    key.public_key().verify(
        signature,
        data,
        padding.PKCS1v15(),
        hashes.SHA256()
    )
    print("PKCS #1 signature is valid")

# PSS padding
# Generate random data to sign
data = os.urandom(1024)

# Use the same configuration for PSS padding
mgf = padding.MGF1(hashes.SHA256())
salt_length = padding.PSS.MAX_LENGTH

# Sign the data with each key using PSS padding
for key in keys:
    signature = key.sign(
        data,
        padding.PSS(
            mgf=mgf,
            salt_length=salt_length
        ),
        hashes.SHA256()
    )

    # Verify the signature
    key.public_key().verify(
        signature,
        data,
        padding.PSS(
            mgf=mgf,
            salt_length=salt_length
        ),
        hashes.SHA256()
    )
    print("PSS signature is valid")