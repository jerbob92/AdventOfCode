package main

import (
	"bufio"
	"log"
	"os"
	"strings"
)

func main() {

	file, err := os.Open("src/Day8/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	numberOfCharacters := 0
	numberOfValues := 0

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line := scanner.Text()

		// Trim and skip empty lines.
		line = strings.Trim(line, " ")
		if len(line) == 0 {
			continue
		}

		// Count the original string.
		numberOfCharacters += len(line)

		// Skip the begin and end quotes.
		i := 1
		for i < (len(line) - 1) {

			// \ is an escaped string.
			if string(line[i]) == "\\" {

				// If our next character is an x (hex), skip 4. Else skip 2 (normal escape, \ or ").
				if string(line[i+1]) == "x" {
					i += 4
				} else {
					i += 2
				}
			} else {
				// Normal string, just continue one.
				i += 1
			}

			// Add one to the amount of values we have.
			numberOfValues += 1
		}
	}

	log.Printf("Part 1: Amount of characters: %d.", numberOfCharacters)
	log.Printf("Part 1: Amount of values: %d.", numberOfValues)
	log.Printf("Part 1: The number of characters of code for string literals minus the number of characters in memory for the values of the strings is %d", numberOfCharacters-numberOfValues)
}
