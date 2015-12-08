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

	numberOfOriginalCharacters := 0
	numberOfNewCharacters := 0

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line :=  scanner.Text()
		line = strings.Trim(line, " ")
		if (len(line) == 0) {
			continue;
		}

		// Count the original string.
		numberOfOriginalCharacters += len(line)

		i := 0

		// Start and end escape.
		numberOfNewCharacters += 2
		for (i < (len(line))) {

			// If we need to escape, we need an extra character. We need to escape \ and ".
			if (string(line[i]) == "\"" || string(line[i]) == "\\") {
				numberOfNewCharacters++;
			}

			// Count the acutal character.
			numberOfNewCharacters ++;

			// Move to next character.
			i++;
		}
	}

	log.Printf("Amount of characters: %d.", numberOfOriginalCharacters)
	log.Printf("Amount of new characters: %d.", numberOfNewCharacters)
	log.Printf("The total number of characters to represent the newly encoded strings minus the number of characters of code in each original string literal is %d", numberOfNewCharacters - numberOfOriginalCharacters)
}
