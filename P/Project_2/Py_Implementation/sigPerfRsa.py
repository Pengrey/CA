import os
import time
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives.asymmetric import rsa
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives import hashes

def signAndVerify(data, key, padding, hashAlg):
    signature = key.sign(
        data,
        padding,
        hashAlg
    )

    # Verify the signature
    key.public_key().verify(
        signature,
        data,
        padding,
        hashAlg
    )

def getPerf(keys, padng):
    # Perform a loop with a certain number of iterations
    niterations = 10

    # Execute the operation a certain number of times
    mtimes = 1000

    for key in keys:
        # Generate random data to sign
        data = os.urandom(1024)

        # Keep track of the quickest time observed
        min_time = float("inf")

        for i in range(niterations):
            # Measure the time it takes to execute the consecutive operations
            start = time.time()
            for j in range(mtimes):
                signAndVerify(data, key, padng, hashes.SHA256())
            elapsed = time.time() - start
            
            # Update the minimum time observed
            min_time = min(min_time, elapsed)

        # Calculate the time each operation takes
        time_per_operation = min_time / mtimes
        print("Key: %d, Padding: %s, Time: %f" % (key.key_size, padng.name.replace("EMSA-", "").replace("-v1_5", ""), time_per_operation))


def main():
    # Read the private keys from the files
    keys = []
    for sizeB in [1024, 2048, 4096]:
        with open(f"./../Key_Generator/keys/key_{sizeB}.pem", "rb") as key_file:
            keys.append(serialization.load_pem_private_key(
            key_file.read(),
            password=None,
            ))
        
    # Get performance of PKCS #1 padding
    getPerf(keys, padding.PKCS1v15())

    # Get performance of PSS padding
    # Use the same configuration for PSS padding
    mgf = padding.MGF1(hashes.SHA256())
    salt_length = padding.PSS.MAX_LENGTH
    getPerf(keys, padding.PSS(
        mgf=mgf,
        salt_length=salt_length
    ))

if __name__ == "__main__":
    main()