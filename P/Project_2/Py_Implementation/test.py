import ecdsa

# Set the curve type for each curve.
# NIST P-256 curve:
curve_p256 = ecdsa.NIST256p

# Koblitz curve secp256k1:
curve_secp256k1 = ecdsa.SECP256k1

# B-233 curve:
curve_b233 = ecdsa.BRAINPOOLP256r1

# Set the message to be signed.
message = b"This is the message to be signed"

# Generate private and public keys for each curve.
# NIST P-256 curve:
priv_key_p256 = ecdsa.SigningKey.generate(curve=curve_p256)
pub_key_p256 = priv_key_p256.get_verifying_key()

# Koblitz curve secp256k1:
priv_key_secp256k1 = ecdsa.SigningKey.generate(curve=curve_secp256k1)
pub_key_secp256k1 = priv_key_secp256k1.get_verifying_key()

# B-233 curve:
priv_key_b233 = ecdsa.SigningKey.generate(curve=curve_b233)
pub_key_b233 = priv_key_b233.get_verifying_key()

# Sign the message using each private key.
# NIST P-256 curve:
signature_p256 = priv_key_p256.sign(message)

# Koblitz curve secp256k1:
signature_secp256k1 = priv_key_secp256k1.sign(message)

# B-233 curve:
signature_b233 = priv_key_b233.sign(message)

# Verify the signature using each public key.
# NIST P-256 curve:
pub_key_p256.verify(signature_p256, message)

# Koblitz curve secp256k1:
pub_key_secp256k1.verify(signature_secp256k1, message)

# B-233 curve:
pub_key_b233.verify(signature_b233, message)

# Print the results.
print("NIST P-256 curve: ", pub_key_p256.verify(signature_p256, message))
print("Koblitz curve secp256k1: ", pub_key_secp256k1.verify(signature_secp256k1, message))
print("B-233 curve: ", pub_key_b233.verify(signature_b233, message))

