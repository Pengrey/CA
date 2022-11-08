# Run 100 000 times the C implementations of E-DES and DES

# Init times
echo -n "1.000000" > edesEncTime.txt

# Make sure the Go implementations are compiled
go build -o edes_enc ./Metods_Code/edes_enc.go && chmod +x edes_enc

# Loop 100 000 times
for i in {1..100}
do
    # Get absolute value of a devision
    j=$(( $i / 1 ))

    # Do a progress bar
    printf "%-*s" $((j+1)) '[' | tr ' ' '#'
    printf "%*s%3d%%\r"  $((100-j))  "]" "$j"

    # Generate a random 64-bit key and save it to a file
    openssl rand -out key.bin 32 > /dev/null 2>&1

    # Run edes_enc and save the output to a file
    ./edes_enc > randomValuesEnc 2> /dev/null
done; echo

# Get the smallest time for edes_enc
minTimeedesEnc=$(cat edesEncTime.txt)

# Print the results
echo "E-DES enc: $minTimeedesEnc"

# Remove compiled if exists
rm -f edes_enc
rm -f edes_dec
rm -f des_enc
rm -f des_dec

# Remove previous results
rm -f edesEncTime.txt
rm -f edesDecTime.txt
rm -f desEncTime.txt
rm -f desDecTime.txt

# Remove the generated key
rm -f key.bin

# Remove the generated random values
rm -f randomValuesEnc