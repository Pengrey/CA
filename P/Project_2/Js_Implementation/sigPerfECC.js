// Import the required libraries
const crypto = require('crypto');
const fs = require('fs');
const os = require('os');

function signAndVerify(digest, key, padding, chosenHash) {
  const signature = key.sign(
    digest,
    padding,
    utils.Prehashed(chosenHash)
  );

  // Verify the signature
  key.public_key().verify(
    signature,
    digest,
    padding,
    utils.Prehashed(chosenHash)
  );
}

function getPerf(keys, padng) {
  // Perform a loop with a certain number of iterations
  const niterations = 10;

  // Execute the operation a certain number of times
  const mtimes = 1000;

  // Hash to be used
  const chosenHash = hashes.SHA256();

  for (const key of keys) {
    // Generate random data to sign
    const data = os.urandom(1024);

    // Set hasher
    const hasher = hashes.Hash(chosenHash);

    // Generate the digest
    hasher.update(data);
    const digest = hasher.finalize();

    // Keep track of the quickest time observed
    let minTime = Number.POSITIVE_INFINITY;

    for (let n = 0; n < niterations; n += 1) {
      // Measure the time it takes to execute the consecutive operations
      const start = performance.now();
      for (let m = 0; m < mtimes; m += 1) {
        signAndVerify(digest, key, padng, chosenHash);
      }
      const elapsed = performance.now() - start;

      // Update the minimum time observed
      minTime = Math.min(minTime, elapsed);
    }

    // Print the results
    console.log(`Time taken for ${key.key_size()} bits: ${minTime}`);
    }
}

function main() {
  // Read the private keys from the files
    const keys = [1024, 2048, 4096].map((bits) => {
      const pem = fs.readFileSync(`./../Key_Generator/keys/key_${bits}.pem`, 'utf-8');
        // Import RSA private key
        return crypto.createPrivateKey({
            key: pem,
            format: 'pem',
            type: 'pkcs1'
        });
    });

    // Get performance of PKCS #1 padding
    console.log('PKCS #1 padding');
    getPerf(keys, padding.PKCS1v15());

    // Get performance of PSS padding
    console.log('PSS padding');
    getPerf(keys, padding.PSS(
        mgf=padding.MGF1(hashes.SHA256()),
        salt_length=padding.PSS.MAX_LENGTH
    ));
}

main();