package main

import (
	"bufio"
	"log"
	"os"
)

func main() {

	file, err := os.Open("src/Day5/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	niceStrings := 0

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the text of the current line.
		line := scanner.Text()

		lastChar := ""

		// Map of repeating and amount of repeating.
		repeatingCharMap := map[string]int{}
		repeatingCharMapPosition := map[string]int{}
		gotRepeatedPair := false
		for i, char := range line {

			// Create pairs, last char can't be same as current char, because then we have overlap.
			if lastChar != "" {

				// Check for overlapping patterns
				if repeatingCharMapPosition[lastChar+string(char)] == 0 || repeatingCharMapPosition[lastChar+string(char)] != (i-1) {
					repeatingCharMapPosition[lastChar+string(char)] = i
					repeatingCharMap[lastChar+string(char)]++

					if repeatingCharMap[lastChar+string(char)] > 1 {
						gotRepeatedPair = true
						break
					}
				}
			}

			// Save this character so we can use it in the next loop.
			lastChar = string(char)
		}

		// Bail out if we didn't get a repeating pair.
		if !gotRepeatedPair {
			continue
		}

		// Check if our line contains at least one letter which repeats with exactly one letter between them.
		charmap := map[int]string{}
		gotRepeatingChar := false
		for i, char := range line {
			charmap[i] = string(char)

			// Check if 2 string ago is the same as the current string.
			if i > 1 && charmap[i-2] == charmap[i] {
				gotRepeatingChar = true
				break
			}
		}

		// Bail out if we dind't get a repeating char.
		if !gotRepeatingChar {
			continue
		}

		// We got through al the requirements, this is a nice string.
		niceStrings++
	}

	log.Printf("We got %d nice strings.", niceStrings)
}
