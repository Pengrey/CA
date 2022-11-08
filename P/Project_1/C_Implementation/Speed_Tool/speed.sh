# Run 100 000 times the C implementations of E-DES and DES

# Init times
echo "1" > edesEncTime.txt
echo "1" > edesDecTime.txt

# Make sure the C implementations are compiled
cc -O3 -o edes_enc ./Metods_Code/edes_enc.c -lssl -lcrypto && chmod +x edes_enc
cc -O3 -o edes_dec ./Metods_Code/edes_dec.c -lssl -lcrypto && chmod +x edes_dec

# Loop 100 000 times
for i in {1..1000}
do
    # Get absolute value of a devision
    j=$(( $i / 10 ))

    # Do a progress bar
    printf "%-*s" $((j+1)) '[' | tr ' ' '#'
    printf "%*s%3d%%\r"  $((100-j))  "]" "$j"

    # Generate a random 64-bit key and save it to a file
    openssl rand -out key.bin 32 > /dev/null 2>&1

    # Run edes_enc and save the output to a file
    ./edes_enc > randomValuesEnc 2> /dev/null

    # Run edes_dec and save the output to a file
    ./edes_dec > /dev/null 2>&1
done; echo

# Get the smallest time for edes_enc
minTimeEnc=$(cat edesEncTime.txt)

# Get the smallest time for edes_dec
minTimeDec=$(cat edesDecTime.txt)

# Print the results
echo "edes_enc: $minTimeEnc"
echo "edes_dec: $minTimeDec"

# Remove compiled if exists
rm -f edes_enc
rm -f edes_dec

# Remove previous results
rm -f edesEncTime.txt
rm -f edesDecTime.txt

# Remove the generated key
rm -f key.bin

# Remove the generated random values
rm -f randomValuesEnc