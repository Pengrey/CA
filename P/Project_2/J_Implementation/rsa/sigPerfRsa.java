package rsa;

import java.security.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

public class sigPerfRsa {
    // function that signs and verifies a message using the RSA algorithm by given hash and keys, while using the PSS padding
    public static List<Long> signAndVerifyPSS(byte[] hash, KeyPair keyPair, Integer mtimes) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initSign(privateKey);
        signature.update(hash);
        byte[] sig = signature.sign();

        // Measure the time it takes to execute the consecutive operations
        long startTime = System.nanoTime();
        // For each time
        for (int k = 0; k < mtimes; k++) {
            sig = signature.sign();
        }
        long signTime = System.nanoTime() - startTime;

        // verify the signature
        signature.initVerify(publicKey);
        signature.update(hash);
        startTime = System.nanoTime();
        // For each time
        for (int k = 0; k < mtimes; k++) {
            boolean verified = signature.verify(sig);
        }
        long verifyTime = System.nanoTime() - startTime;

        List<Long> results = new ArrayList<>();
        results.add(signTime);
        results.add(verifyTime);
        return  results;
    }

    // function that signs and verifies a message using the RSA algorithm by given hash and keys, while using the PKCS1 padding
    public static List<Long> signAndVerifyPKCS1(byte[] hash, KeyPair keyPair, Integer mtimes) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initSign(privateKey);
        signature.update(hash);
        byte[] sig = signature.sign();

        // Measure the time it takes to execute the consecutive operations
        long startTime = System.nanoTime();
        // For each time
        for (int k = 0; k < mtimes; k++) {
            sig = signature.sign();
        }
        long signTime = System.nanoTime() - startTime;

        // verify the signature
        signature.initVerify(publicKey);
        signature.update(hash);
        startTime = System.nanoTime();
        // For each time
        for (int k = 0; k < mtimes; k++) {
            boolean verified = signature.verify(sig);
        }
        long verifyTime = System.nanoTime() - startTime;

        List<Long> results = new ArrayList<>();
        results.add(signTime);
        results.add(verifyTime);
        return  results;
    }

    // function that, given keys and the type of padding to use, generates random messages, hashes them and signs and verifies them with the given padding type
    public static void getPerf(KeyPair[] Keys, String padding) throws Exception {
        // Key sizes
        String[] keySizes = {"1024", "2048", "4096"};

        // Perform a loop with a certain number of iterations
        Integer niterations = 10;

        // Execute the operation a certain number of times
        Integer mtimes = 1000;

        // For each key in the array of keys
        for(int i = 0; i < Keys.length; i++) {
            // Generate random data to sign
            SecureRandom random = new SecureRandom();
            byte[] data = new byte[32];
            random.nextBytes(data);

            // Get the hash of the data
            MessageDigest digest = MessageDigest.getInstance("SHA-256");
            byte[] hash = digest.digest(data);

            // Keep track of the quickest time observed
            long minSignTime = Long.MAX_VALUE;
            long minVerifyTime = Long.MAX_VALUE;

            // For each iteration
            for (int j = 0; j < niterations; j++) {
                if (padding.equals("PSS")) {
                    // Sign and verify the hash
                    List<Long> results = signAndVerifyPSS(hash, Keys[i], mtimes);
                    long signTime = results.get(0);
                    long verifyTime = results.get(1);

                    // Update the minimum time
                    if (signTime < minSignTime) {
                        minSignTime = signTime;
                    }
                    if (verifyTime < minVerifyTime) {
                        minVerifyTime = verifyTime;
                    }
                } else if (padding.equals("PKCS1")) {
                    // Sign and verify the hash
                    List<Long> results = signAndVerifyPKCS1(hash, Keys[i], mtimes);
                    long signTime = results.get(0);
                    long verifyTime = results.get(1);

                    // Update the minimum time
                    if (signTime < minSignTime) {
                        minSignTime = signTime;
                    }
                    if (verifyTime < minVerifyTime) {
                        minVerifyTime = verifyTime;
                    }
                }
            }
            // Calculate the time each operation takes
            long signTime = minSignTime / mtimes;
            long verifyTime = minVerifyTime / mtimes;

            // Convert the time to seconds
            double signTimeSec = signTime / 1000000000.0;
            double verifyTimeSec = verifyTime / 1000000000.0;
            System.out.println("Key: " + keySizes[i] + ", Padding: " + padding + ", Sign: " + String.format(Locale.US, "%.6f", signTimeSec) + ", Verify: " + String.format(Locale.US, "%.6f", verifyTimeSec));
        }
    }

    public static void main(String[] args) throws Exception {
        // Generate key pairs
        String[] keySizes = {"1024", "2048", "4096"};
        KeyPair[] Keys = new KeyPair[keySizes.length];
        for (int i = 0; i < keySizes.length; i++) {
            // Generate a key pair
            KeyPairGenerator keyGen = KeyPairGenerator.getInstance("RSA");
            keyGen.initialize(Integer.parseInt(keySizes[i]));
            KeyPair keyPair = keyGen.generateKeyPair();

            // Store the private key
            Keys[i] = keyPair;
        }

        // Get the performance of the RSA algorithm with PKCS1 padding
        getPerf(Keys, "PKCS1");

        // Get the performance of the RSA algorithm with PSS padding
        getPerf(Keys, "PSS");
    }
}