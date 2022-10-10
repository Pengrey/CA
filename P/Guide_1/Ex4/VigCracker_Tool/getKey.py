#!/usr/bin/python2
# Testing code from https://github.com/1r0dm480/Vigenere-Cipher-Breaker
MAX_KEY_LENGTH_GUESS = 20
alphabet = 'abcdefghijklmnopqrstuvwxyz'

# Returns the Index of Councidence for the "section" of ciphertext given
def get_index_c(ciphertext):
	
	N = float(len(ciphertext))
	frequency_sum = 0.0

	# Using Index of Coincidence formula
	for letter in alphabet:
		frequency_sum+= ciphertext.count(letter) * (ciphertext.count(letter)-1)

	# Using Index of Coincidence formula
	ic = frequency_sum/(N*(N-1))
	return ic

# Returns the key length with the highest average Index of Coincidence
def get_key_length(ciphertext):
	ic_table=[]

	# Splits the ciphertext into sequences based on the guessed key length from 0 until the max key length guess (20)
	# Ex. guessing a key length of 2 splits the "12345678" into "1357" and "2468"
	# This procedure of breaking ciphertext into sequences and sorting it by the Index of Coincidence
	# The guessed key length with the highest IC is the most porbable key length
	for guess_len in range(MAX_KEY_LENGTH_GUESS):
		ic_sum=0.0
		avg_ic=0.0
		for i in range(guess_len):
			sequence=""
			# breaks the ciphertext into sequences
			for j in range(0, len(ciphertext[i:]), guess_len):
				sequence += ciphertext[i+j]

			ic_sum+=get_index_c(sequence)
			
		# obviously don't want to divide by zero
		if not guess_len==0:
			avg_ic=ic_sum/guess_len
		ic_table.append(avg_ic)

	# returns the index of the highest Index of Coincidence (most probable key length)
	best_guess = ic_table.index(sorted(ic_table, reverse = True)[0])
	second_best_guess = ic_table.index(sorted(ic_table, reverse = True)[1])
	print(sorted(ic_table, reverse = True)[0])
	print(second_best_guess)
	# Since this program can sometimes think that a key is literally twice itself, or three times itself, 
	# it's best to return the smaller amount.
	# Ex. the actual key is "dog", but the program thinks the key is "dogdog" or "dogdogdog"
	# (The reason for this error is that the frequency distribution for the key "dog" vs "dogdog" would be nearly identical)
	if best_guess % second_best_guess == 0:
		return second_best_guess
	else:
		return best_guess

def main():
    f = open("./Input/out.txt","r")
    lines = f.readlines()
    ciphertext = ''.join(x.lower() for x in lines if x.isalpha())
    print("Estimated key size: " + str(get_key_length(ciphertext)))

if __name__ == '__main__':
	main()
