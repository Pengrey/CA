import binascii
import hashlib
import random
import pandas as pd


def iterate(binToChange):
    randN = random.randint(0, len(binToChange) - 1)

    binChanged = binToChange[randN:] + ('1' if binToChange[randN] == '0' else '0') + binToChange[:randN + 1]

    return hashlib.md5(binChanged.encode()).hexdigest()

def countDiference(referenceHash, newHash):
    referenceBin = bin(int(binascii.hexlify(bytearray(referenceHash)), 16)).zfill(8)
    newBin = bin(int(binascii.hexlify(bytearray(newHash.encode())), 16)).zfill(8)

    for i in range(10):
        sum = 0
        for i in range(len(referenceBin) - 1):
            if newBin[i] != referenceBin[i]:
                sum+=1
    return sum

if __name__ == '__main__':
    referenceBin = bin(int(binascii.hexlify(bytearray(random.sample(range(0, 128), 100))), 16)).zfill(8)

    # Take first iteration
    referenceHash = hashlib.md5(referenceBin.encode()).digest()

    # Get values from multiple iteration
    results = []
    for i in range(410):
        newHash = iterate(bin(int(binascii.hexlify(bytearray(referenceHash)), 16)).zfill(8))
        results += [countDiference(referenceHash, newHash)]
    
    # Print results
    print(results)

    # Print histogram
    #pd.Series(results).value_counts(sort=False).plot(kind='bar')

