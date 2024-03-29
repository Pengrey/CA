package eliptic;

import java.security.*;
import java.util.ArrayList;
import java.util.List;
import java.util.Locale;

import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.jce.spec.ECParameterSpec;
import org.bouncycastle.jce.ECNamedCurveTable;

public class sigPerfECC {
    // function that signs and verifies a message using the ECDSA algorithm by given hash and keys
    public static List<Long> signAndVerify(byte[] hash, KeyPair keyPair, Integer mtimes) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256with" + privateKey.getAlgorithm());
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

    // function that, given keys, generates random messages, hashes them and signs and verifies them
    public static void getPerf(KeyPair[] Keys) throws Exception {
        // Perform a loop with a certain number of iterations
        Integer niterations = 100;

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

            // Curves
            String[] curveCurves = {"P-192", "P-256", "P-521", "K-163", "K-283", "K-571", "B-163", "B-283", "B-571"};

            // For each iteration
            for (int j = 0; j < niterations; j++) {
                // Sign and verify the hash
                List<Long> results = signAndVerify(hash, Keys[i], mtimes);
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
            // Calculate the time each operation takes
            long signTime = minSignTime / mtimes;
            long verifyTime = minVerifyTime / mtimes;

            // Convert the time to seconds
            double signTimeSec = signTime / 1000000000.0;
            double verifyTimeSec = verifyTime / 1000000000.0;
            System.out.println("Curve: " + curveCurves[i] + ", Key: " + curveCurves[i].substring(2) + ", Sign: " + String.format(Locale.US, "%.6f", signTimeSec) + ", Verify: " + String.format(Locale.US, "%.6f", verifyTimeSec));
        }
    }

    // This is the main method
    public static void main(String[] args) throws Exception {
        // Load private keys from .pem files located in "./../Key_Generator/keys/"
        String[] curveCurves = {"P-192", "P-256", "P-521", "K-163", "K-283", "K-571", "B-163", "B-283", "B-571"};

        // Arrays to store the private keys
        KeyPair[] Keys = new KeyPair[curveCurves.length];

        // For each curve name, generate a key pair and store the private key
        Security.addProvider(new BouncyCastleProvider());
        for (int i = 0; i < curveCurves.length; i++) {
            // Generate a key pair
            ECParameterSpec ecSpec = ECNamedCurveTable.getParameterSpec(curveCurves[i]);
            KeyPairGenerator g = KeyPairGenerator.getInstance("ECDSA", "BC");
            g.initialize(ecSpec, new SecureRandom());

            // Store the private key
            Keys[i] = g.generateKeyPair();
        }

        // Get performance
        getPerf(Keys);
    }
}
