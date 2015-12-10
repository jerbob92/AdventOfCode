package main

import (
	"log"
	"strconv"
)

func main() {
	input := "1321131112"

	// Convert the input into a string array.
	// We use a string array as it's much faster than concatting a large string.
	inputArray := []string{}
	for _, char := range input {
		inputArray = append(inputArray, string(char))
	}

	// Part 1 has 40 passes over the input.
	for pass := 1; pass <= 40; pass++ {

		// Save the first number,  as that can never be a repeated number.
		lastChar := inputArray[0]

		// Start our pass at 1.
		currentCharCount := 1

		// Make sure we have a clean string for this pass.
		newString := []string{}

		// Loop through all characters, skip the first, as that can never be a repeated number.
		for _, char := range inputArray[1:] {

			// If the current char is the same as the last seen char, count one up
			if lastChar == char {
				currentCharCount++
			} else {

				// Else we found a new number, append our look-and-say string to the new combined string.
				newString = append(newString, strconv.Itoa(currentCharCount))
				newString = append(newString, lastChar)

				// Mark the new character as the new character.
				lastChar = char

				// Reset the character counter.
				currentCharCount = 1
			}
		}

		// Append the last found character/counter to the completed string
		newString = append(newString, strconv.Itoa(currentCharCount))
		newString = append(newString, lastChar)

		// Set our new string as the input for the next pass.
		inputArray = newString
	}

	log.Printf("Part 1: The length of the resulting string is %d", len(inputArray))
}
