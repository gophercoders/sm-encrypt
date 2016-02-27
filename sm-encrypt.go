package main

import (
	"fmt"
	"unicode/utf8"

	"github.com/gophercoders/simpleio"
)

func main() {
	// tell the user what the program does.
	fmt.Println("sm-encrypt encrypts a word according to the Caesar Cipher.")
	fmt.Println("Please enter a word to encrypt:")

	// message is the word the user wants to encrypt.
	// Your turn
	// finish the next like to create a variable called "message"
	// of the correct type. Use the cheat sheet for help
	var
	// Now you need to read in the word that the user typed.
	// Your turn
	// See if you can complete the next line to read in the word that
	// the user typed
	message =
	// Now we need to create a new varaible messageRunes - to hold
	// the individual runes or characters in the word.
	// First we create the variable
	var messageRunes []rune
	// then we convert the message to the "slice of runes" type.
	// This pattern we are using for this is in the cheat sheet
	messageRunes = []rune(message)
	// Now we want to know how many characters or runes are
	// in the word the user typed.
	// Your turn.
	// Create a new variable called "numberOfMessageRunes" of
	// the correct type.
	var
	// Your turn
	// Now you need to work out the number of characters in the
	// message and store the result in "numberOfMessageRunes"
	// Look in the cheat sheet to see how to do this.
	numberOfMessageRunes =

	// Your Turn
	// Now you need to create a new variable called "plaintext".
	// plaintext has to hold a copy of the alphabet. What type does it
	// need to be?
	var plaintext
	// plaintext holds a copy of the alphabet
	// Your Turn: What do you need to do to finish this?
	plaintext = "abcd"
	// plaintextRunes is a new variable that holds the runes or characters of the
	// plaintext
	// Your turn
	// What type does plaintextRunes need to be?
	var plaintextRunes
	// Your Turn
	// How do you convert plaintext to the type needed by paintextRunes?
	// Use the cheat sheet or look at how messageRunes does it for help
	plaintextRunes =

	// plaintextLength is a variable that holds the number of characters in the
	// plaintext
	var plaintextLength int
	// Your turn
	// How do you find out the length of the plaintext and correctly
	// set the value of plaintextLength?
	plaintextLength =

	// Your Turn
	// Now you need to create a new variable called ciphertext.
	// ciphertext has to hold a copy of the alphabet, but this time
	// shifted three letters to the right.
	// How do you do declare this variable and what type do you need?
	var
	// Your turn
	// Now you need to assign an alphabet that has been shifted three
	// letters to the right to ciphertext
	ciphertext =
	// Now we need to create another slice of runes, this time for the
	// ciphertext so that we can access the individual letters of the
	// ciphertext.
	var ciphertextRunes []rune
	ciphertextRunes = []rune(ciphertext)

	fmt.Print("The encrypted Message is: ")
	// Now we can start the encryption
	// i is the counter that records the position of the current
	// character in the message
	var i int
	i = 0
	// First we want to look at each character in the message.
	// So we need a loop that loops over the characters in the message
	// This is first loop.It is called the outer loop.
	// You want i to start at position zero and stop when it is greater
	// than or equal to the number or characters in the message
	// Your turn
	// What condition do you need for this? Remember that the loop will
	// only execute if the condition is true.
	for  {
		// j is the counter that records the position in the plaintext alphabet
		var j int
		j = 0
		// Now we need a second loop or an inner loop.
		// The inner loop loops over all of the characters in the plain text
		// alphabet.
		// Like the first loop j starts at zero and stops when j is greater
		// than or equal to the number of characters in the alphabet
		// Your Turn
		// What condition do ypu need for this so that the loop executes?
		for  {
			// Now you have to compare the current character in the message
			// with the current character in the plaintext alphabet.
			// You want to know if the characters are equal to each other.
			// Your turn
			// How do you write this condition?
			if  {
				// If the letters are the same you want to find out
				// what the ciphertext letter should be.
				// You do this by taking the position of the current letter
				// in the plaintext, j, and finding out what the letter
				// at positon j is in the cipherText.
				// Once you do this you need to convert it to a string and
				// print it out.
				// Your Turn
				// Can you remeber how to do this? The answer is in the cheat
				// sheet. You need to put your answer between the () of the
				// Print
				fmt.Print()
				// The new break keyword will stop the inner loop early.
				// So you'll jump out of this loop stright to the i = i + 1 line
				break // required!
			}
			// If the letters matched we never reach here! That's what the break does.
			// If they don't match, we need to try and match against the next
			// letter in the plaintext. To do that we have to make j larger by one
			// before the inner loop starts again
			// Your Turn
			// How do you make j bigger by one?

		}
		// Now you have encoded one letter in the message you need to start
		// the encoding process on the next letter in the message.
		// To do that you need to make i bigger by one, before the outer loop
		// starts again.
		// Your turn
		// How do you make i bigger by one?

	}
	// finally we need to print a new line on its own, so that the
	// terminal prompt appears correctly on the next line
	fmt.Println()
}
