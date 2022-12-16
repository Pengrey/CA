package rsa;

import java.security.*;

public class sigPerfRsa {
    // function that signs and verifies a message using the RSA algorithm by given hash and keys, while using the PSS padding
    public static void signAndVerifyPSS(byte[] hash, KeyPair keyPair) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initSign(privateKey);
        signature.update(hash);
        byte[] sig = signature.sign();
        // verify the signature
        signature.initVerify(publicKey);
        signature.update(hash);
        boolean verified = signature.verify(sig);
    }

    // function that signs and verifies a message using the RSA algorithm by given hash and keys, while using the PKCS1 padding
    public static void signAndVerifyPKCS1(byte[] hash, KeyPair keyPair) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256withRSA");
        signature.initSign(privateKey);
        signature.update(hash);
        byte[] sig = signature.sign();
        // verify the signature
        signature.initVerify(publicKey);
        signature.update(hash);
        boolean verified = signature.verify(sig);
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
            long quickestTime = Long.MAX_VALUE;

            // For each iteration
            for (int j = 0; j < niterations; j++) {
                // Measure the time it takes to execute the consecutive operations
                long startTime = System.nanoTime();

                // For each time
                for (int k = 0; k < mtimes; k++) {
                    // Sign and verify the hash
                    if (padding.equals("PSS")) {
                        signAndVerifyPSS(hash, Keys[i]);
                    } else if (padding.equals("PKCS1")) {
                        signAndVerifyPKCS1(hash, Keys[i]);
                    }
                }

                // Measure the time it takes to execute the consecutive operations
                long endTime = System.nanoTime();

                // Calculate the time it took to execute the operations
                long duration = (endTime - startTime);

                // If the time it took to execute the operations is quicker than the quickest time observed
                if (duration < quickestTime) {
                    // Update the quickest time observed
                    quickestTime = duration;
                }
            }

            // Calculate the time each operation takes
            double timePerOperation = (double) quickestTime / (double) mtimes;
            // Print the time each operation takes
            System.out.println("Key: " + keySizes[i] + ", Padding: " + padding + ", Time: " + timePerOperation);
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