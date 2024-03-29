#!/usr/bin/env python3

# Relation between m1 and m2:
# m1 = m + 256**254 × r1
# m2 = m + 256**254 × r2
# m1 = m2 - 256**254 × (r2 - r1)

# t will be the difference between the value of the random bit used on m1 and m2, this bit value is from 0 to 255

# By using the paper "Low-Exponent RSA with Related Messages" we can try to calculate the value plaintext of the message sent
# The paper says that we can calculate the value of the plaintext by using the following formula:
# m1 = alpha*m2 + beta
# m1 = m2 - 256**254 × (r2 - r1)
# beta = 256**254 × (r2 - r1)
# Now we suppose that the messages are encrypted under RSA with an exponent of 65537
# We have the following expression:
# ci = mi**65537 mod n, where i = 1,2
#
# Then we can calculate c1, c2, alpha and beta by using the following formula:
# c1 = m**65537 mod n
# c2 = (m + 256**254 × (r2 - r1))**65537 mod n
# c2 - c1 = (m + 256**254 × (r2 - r1))**65537 - m**65537 mod n
#
# To simplify the calculation we can use the following formula instead:
# Let z denote the unknown message m. Then z satisfies the following two polynomial relations:
# z**65537 - c1 = 0 mod n
# t = (r2 - r1)
# (z + 256**254 × t)**65537 - c2 = 0 mod n
#
# where the ci are treated as known constants. Apply the Euclidean algorithm to find the greatest common divisor of these two univariate polynomials over the ring Z:
# gcd(z**65537 - c1, (z + 256**254 × t)**65537 - c2)

# The following code is based on the code from the paper "Low-Exponent RSA with Related Messages" and the code from the github repository "https://github.com/jvdsn/crypto-attacks/tree/f9bd04b8311aaed12ef807155efdcbd0230e669d/rsa-related-messages"
from sage.all import Zmod
from helper import fast_polynomial_gcd
N = 32256584853881668819643226749483462544517280786196195255582678074125522841390956540782962014408431565909877544245718140900922831414113933094382685804181028715628789173848541552090856381959547800109346250918023077411316486779059209905502142149463496915675425171779940802141236802539100446445864948469009929745424996161237068507360821386446982187061271005035396213978536661615000916707597139419106311209764388516023878827425893886409701653277417646979000696619615284135939446467741679314349581234059292706721475682874495779225854759884984965933278249977898357852295765822796204816960068087631764580573733332129094366877
e = 65537
c1 = 13999706255490144554213084007362100309456625583629535725983824440742450114510021923628499385800655768704697513873892730862062179064503017442228577850631048180035574442293934041730503877415326236005866550461714630750446977657222269945701441873316626527820354240870874188564326954720906083178551774216174147120037614811028371139001418608902352823024836164449629163263287594329753060294903583026297409928003541808413010042765306020034705666435003576321720840945288688623014953034534515521905658806072559250268784415441059276528955550985148160043527400785300898790453936675517505792768625755579501439957445247311207034994
c2 = 7330646540408462499487732988734435881882600143917250125234725645295265436844855956345867015460257851867152226082992604583380722353979123823769153663750127378381181947124303124717022628317971445602980523815915703319718627892152151763224999403696019840134053621522434775044673990084838715747559195164385146719276035186448119274230106962051610865462775732255984526858086851588534412481383870183265812676914984402719899236047964431291778087348829328052439726843671068196575220476292945021334692065439186080314715245574910218285041885631078655049765029101223651172336086726816604586051074113029899282669743291791900574376
# t = ?
# t is the difference between the value of the random bit used on m1 and m2, this bit value is from 0 to 255
# The value of t is unknown, so we we need to bruteforce it to find the value of the plaintext

# We try every possible value of t from -255 to 255
for t in range(-255, 256):
    print("Trying t = {}".format(t), end="\r")
    # We calculate the value of the plaintext by using the formula from the paper
    # We use the function fast_polynomial_gcd to calculate the gcd of the two polynomials
    # The function fast_polynomial_gcd is based on the code from the github repository "https://github.com/jvdsn/crypto-attacks"

    # Create the polynomial ring Z[x]
    x = Zmod(N)["x"].gen()

    # Create the two polynomials
    p1 = x**e - c1
    p2 = (x + 256**254 * t)**e - c2
    # Calculate the gcd of the two polynomials
    gcd = -fast_polynomial_gcd(p1, p2).monic()

    # Try to find the value of the plaintext
    try:
        m = int(gcd[0])

        # Remove the padding
        m = m.to_bytes((m.bit_length() + 7) // 8, 'little')

        # Remove random bit at the end of the message
        m = m[:-1]

        # Convert to ascii
        m = m.decode('ascii')

        print("Flag:", m)
        break
    except:
        pass