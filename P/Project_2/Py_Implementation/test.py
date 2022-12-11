# Sign random data with a ECDSA key and verify it

from cryptography.hazmat.primitives import hashes
from cryptography.hazmat.primitives.asymmetric import ec
from cryptography.hazmat.primitives import serialization

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

    # Sign the data
    data = b"Hello World"

    for ecc in eccKeys:
        signature = ecc.sign(
            data,
            ec.ECDSA(hashes.SHA256())
        )

        # Verify the signature
        public_key = ecc.public_key()
        public_key.verify(
            signature,
            data,
            ec.ECDSA(hashes.SHA256())
        )

        print("Signature verified")

if __name__ == "__main__":
    main()


