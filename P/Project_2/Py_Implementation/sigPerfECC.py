from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import ec
from cryptography.hazmat.primitives import serialization
from cryptography.hazmat.primitives.asymmetric import utils
import os
import time

def signAndVerify(digest, chosen_hash, key):
    signature = key.sign(
        digest,
        ec.ECDSA(utils.Prehashed(chosen_hash))
    )

    # Verify the signature
    key.public_key().verify(
        signature,
        digest,
        ec.ECDSA(utils.Prehashed(chosen_hash))
    )

def getPerf(keys):
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
        min_time = float("inf")

        for n in range(niterations):
            # Measure the time it takes to execute the consecutive operations
            start = time.time()
            for m in range(mtimes):
                signAndVerify(digest, chosen_hash, key)
            elapsed = time.time() - start
            
            # Update the minimum time observed
            min_time = min(min_time, elapsed)

        # Calculate the time each operation takes
        time_per_operation = min_time / mtimes
        print("Curve: %s, Key: %d, Time: %f" % (key.curve.name, key.key_size, time_per_operation))

def main():
    # Load the private keys
    eccKeys = []
    eccList = [ ec.SECP192R1(),
                ec.SECP256R1(),
                ec.SECP521R1(),
                ec.SECT163K1(),
                ec.SECT283K1(),
                ec.SECT571K1(),
                ec.SECT163R2(),
                ec.SECT283R1(),
                ec.SECT571R1()
            ]

    for ecc in eccList:
        with open("./../Key_Generator/keys/key_" + ecc.name + ".pem", "rb") as key_file:
            private_key = serialization.load_pem_private_key(
                key_file.read(),
                password=None,
            )
            eccKeys.append(private_key)

    # Get performance
    getPerf(eccKeys)

if __name__ == "__main__":
    main()
