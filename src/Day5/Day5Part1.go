package main
import (
	"os"
	"bufio"
	"log"
	"strings"
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


		// First filter out the bad strings.
		if (strings.Contains(line, "ab") || strings.Contains(line, "cd") || strings.Contains(line, "pq") || strings.Contains(line, "xy")) {
			continue
		}

		// Now loop through the line to see if we got a repeating character.
		hasRepeatingCharacter := false

		lastChar := ""
		vowelsFound := 0;
		for _, char := range line {

			// Save this character so we can match it in the next loop.
			if string(char) == lastChar {
				hasRepeatingCharacter = true
			}

			lastChar = string(char)

			// Check for vowels.
			if lastChar == "a" || lastChar == "e" || lastChar == "i" || lastChar == "o" || lastChar == "u" {
				vowelsFound++
			}
		}

		// If we didn't have a repeating character, skip this line.
		if !hasRepeatingCharacter {
			continue
		}

		// If we didn't have enough vowels, skip this line.
		if vowelsFound < 3 {
			continue;
		}

		// We got through al the requirements, this is a nice string.
		niceStrings++
	}

	log.Printf("We got %d nice strings.", niceStrings)
}
