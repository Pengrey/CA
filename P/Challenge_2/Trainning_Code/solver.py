#!/usr/bin/env python3

# Relation between m1 and m2:
# m1 = m + 256**62 × 133
# m2 = m + 256**62 × 147
# m1 = m2 - 256**62 × 14

# t will be the difference between the value of the random bit used on m1 and m2, this bit value is from 0 to 255
# t = 14 (in this case)

# By using the paper "Low-Exponent RSA with Related Messages" we can try to calculate the value plaintext of the message sent
# The paper says that we can calculate the value of the plaintext by using the following formula:
# m1 = alpha*m2 + beta
# m1 = m2 - 256**62 × 14
# beta = 256**62 × t
# Now we suppose that the messages are encrypted under RSA with an exponent of 17
# We have the following expression:
# ci = mi**17 mod n, where i = 1,2
#
# Then we can calculate c1, c2, alpha and beta by using the following formula:
# c1 = m**17 mod n
# c2 = (m + 256**62 × t)**17 mod n
# c2 - c1 = (m + 256**62 × t)**17 - m**17 mod n
#
# To simplify the calculation we can use the following formula instead:
# Let z denote the unknown message m. Then z satisfies the following two polynomial relations:
# z**17 - c1 = 0 mod n
# (z + 256**62 × t)**17 - c2 = 0 mod n
#
# where the ci are treated as known constants. Apply the Euclidean algorithm to find the greatest common divisor of these two univariate polynomials over the ring Z:
# gcd(z**17 - c1, (z + 256**62 × t)**17 - c2) = gcd(z**17 - c1, z**17 + 256**62 × t**17 - c2)

# Caluclate gdc of the two polynomials above and get a polynomial with sage math
from sage.all import Zmod
from helper import fast_polynomial_gcd

N = 11598576175167152956424461782394541439955617267494049729274868685493179955542949179810730474765255703849636600732122155392474369596059442252066674279501961
c1 = 4568954371398241312788046679661973370207136291806473528218813955796701824269988175094014316097935275071331825342382897830026928706806283261576848215938045
c2 = 10774777440091235821269290404723502603355688369384505562693265479722439768712101419248994746280675715013811585517386981365362526469147816190061167310265864
t = 14
e = 17

# Create the polynomial ring over Z field
x = Zmod(N)["x"].gen()

g1 = x ** e - c1
g2 = (x + 256**62 * t) ** e - c2
g = -fast_polynomial_gcd(g1, g2).monic()

print(int(g[0]))