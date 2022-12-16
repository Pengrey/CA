package eliptic;

import java.security.*;
import org.bouncycastle.jce.provider.BouncyCastleProvider;
import org.bouncycastle.jce.spec.ECParameterSpec;
import org.bouncycastle.jce.ECNamedCurveTable;

public class sigPerfECC {
    // function that signs and verifies a message using the ECDSA algorithm by given hash and keys
    public static void signAndVerify(byte[] hash, KeyPair keyPair) throws Exception {
        // get the private key
        PrivateKey privateKey = keyPair.getPrivate();
        // get the public key
        PublicKey publicKey = keyPair.getPublic();
        // get the signature
        Signature signature = Signature.getInstance("SHA256with" + privateKey.getAlgorithm());
        signature.initSign(privateKey);
        signature.update(hash);
        byte[] sig = signature.sign();
        // verify the signature
        signature.initVerify(publicKey);
        signature.update(hash);
        boolean verified = signature.verify(sig);
    }

    // function that, given keys, generates random messages, hashes them and signs and verifies them
    public static void getPerf(KeyPair[] Keys) throws Exception {
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

            // Curves
            String[] curveCurves = {"P-192", "P-256", "P-521", "K-163", "K-283", "K-571", "B-163", "B-283", "B-571"};

            // For each iteration
            for (int j = 0; j < niterations; j++) {
                // Measure the time it takes to execute the consecutive operations
                long startTime = System.nanoTime();

                // For each time
                for (int k = 0; k < mtimes; k++) {
                    // Sign and verify the hash
                    signAndVerify(hash, Keys[i]);
                }

                // Update the minimum time observed
                long timeTaken = System.nanoTime() - startTime;
                if (timeTaken < quickestTime) {
                    quickestTime = timeTaken;
                }
            }
            // Calculate the time each operation takes
            double timePerOperation = (double) quickestTime / (double) mtimes;
            System.out.println("Curve: " + curveCurves[i] + ", Key: " + curveCurves[i].substring(2) + ", Time: " + timePerOperation);
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
