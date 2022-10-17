#!/usr/bin/env python3
import hashlib, os
from alive_progress import alive_bar

HASH_LIMIT = 3
PERMUTATIONS = 100

def get_collision():
    firstHash = hashlib.sha3_256(os.urandom(4)).digest()[:HASH_LIMIT]

    attempts = 1

    while True:
        if firstHash == hashlib.sha3_256(os.urandom(4)).digest()[:HASH_LIMIT]:
            return attempts
        else: 
            attempts += 1

def get_average():
    with alive_bar(PERMUTATIONS) as bar:
        avg = 0
        for i in range(PERMUTATIONS):
            attempts = get_collision()
            avg +=  (attempts / PERMUTATIONS)
            print(f"Collided in {attempts} attempts.")

            bar()
    
    return avg

if __name__ == '__main__':
    print(f"Hash size used: {HASH_LIMIT}\nTheorical attempts needed: 2^({HASH_LIMIT * 4}/2) == {pow(2,HASH_LIMIT * 2)}\nAverage attempts obtained: {get_average()}")