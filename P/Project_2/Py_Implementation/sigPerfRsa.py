import os
import time
from cryptography.hazmat.primitives.asymmetric import padding
from cryptography.hazmat.primitives.asymmetric import utils
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives import hashes

def signAndVerify(digest, key, padding, chosen_hash, mtimes):
    # Measure the time it takes to execute the consecutive operations
    start = time.time()
    for m in range(mtimes):
        signature = key.sign(
            digest,
            padding,
            utils.Prehashed(chosen_hash)
        )
    signTime = time.time() - start

    # Verify the signature
    start = time.time()
    for m in range(mtimes):
        key.public_key().verify(
            signature,
            digest,
            padding,
            utils.Prehashed(chosen_hash)
        )
    vrfyTime = time.time() - start

    return signTime, vrfyTime

def getPerf(keys, padng):
    # Perform a loop with a certain number of iterations
    niterations = 10

    # Execute the operation a certain number of times
    mtimes = 1000

    # Hash to be used
    chosen_hash = hashes.SHA256()

    for key in keys:
        # Generate random data to sign
        data = os.urandom(1024)

        # Set hasher
        hasher = hashes.Hash(chosen_hash)

        # Generate the digest
        hasher.update(data)
        digest = hasher.finalize()

        # Keep track of the quickest time observed
        minSignTime = float("inf")
        minVrfyTime = float("inf")

        for n in range(niterations):
            signTime, vrfyTime = signAndVerify(digest, key, padng, chosen_hash, mtimes)
            
            # Update the minimum time observed
            minSignTime = min(minSignTime, signTime)
            minVrfyTime = min(minVrfyTime, vrfyTime)

        # Calculate the time each operation takes
        time_per_sign_operation = minSignTime / mtimes
        time_per_vrfy_operation = minVrfyTime / mtimes
        print("Key: %d, Padding: %s, Sign: %f, Vrfy: %f" % (key.key_size, padng.name.replace("EMSA-", "").replace("-v1_5", ""), time_per_sign_operation, time_per_vrfy_operation))


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