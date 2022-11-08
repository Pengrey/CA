# Run 100 000 times the Go implementations of E-DES and DES

# Init times
echo -n "1.000000" > edesEncTime.txt
echo -n "1.000000" > edesDecTime.txt
echo -n "1.000000" > desEncTime.txt
echo -n "1.000000" > desDecTime.txt

# Make sure the Go implementations are compiled
go build -o edes_enc ./Metods_Code/edes_enc.go && chmod +x edes_enc
go build -o edes_dec ./Metods_Code/edes_dec.go && chmod +x edes_dec
go build -o des_enc ./Metods_Code/des_enc.go && chmod +x des_enc
go build -o des_dec ./Metods_Code/des_dec.go && chmod +x des_dec

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

    # Run edes_dec and save the output to a file
    ./edes_dec > /dev/null 2>&1

    # Run des_enc and save the output to a file
    ./des_enc > randomValuesEnc 2>&1

    # Run des_dec and save the output to a file
    ./des_dec > /dev/null 2>&1
done; echo

# Get the smallest time for edes_enc
minTimeedesEnc=$(cat edesEncTime.txt)

# Get the smallest time for edes_dec
minTimeedesDec=$(cat edesDecTime.txt)

# Get the smallest time for des_enc
minTimedesEnc=$(cat desEncTime.txt)

# Get the smallest time for des_dec
minTimedesDec=$(cat desDecTime.txt)

# Print the results
echo "E-DES enc: $minTimeedesEnc"
echo "E-DES dec: $minTimeedesDec"
echo "DES enc: $minTimedesEnc"
echo "DES dec: $minTimedesDec"

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