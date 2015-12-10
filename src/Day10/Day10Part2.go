package main

import (
	"log"
	"strconv"
)

func main() {
	input := "1321131112"
	inputArray := []string{}

	for _, char := range input {
		inputArray = append(inputArray, string(char))
	}

	for pass := 1; pass <= 50; pass++ {
		lastChar := ""
		currentCharCount := 1
		newString := []string{}
		for i, char := range inputArray {
			if (i == 0) {
				lastChar = char
				continue;
			}
			if (lastChar == char) {
				currentCharCount++;
			} else {
				newString = append(newString, strconv.Itoa(currentCharCount))
				newString = append(newString, lastChar)

				lastChar = char
				currentCharCount = 1
			}
		}
		newString = append(newString, strconv.Itoa(currentCharCount))
		newString = append(newString, lastChar)
		inputArray = newString
	}

	log.Printf("The length of the resulting string is %d", len(inputArray))
}
