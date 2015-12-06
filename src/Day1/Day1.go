package main
import (
	"log"
	"io/ioutil"
)

func main() {

	input, err := ioutil.ReadFile("src/Day1/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// We start at floor 0
	currentFloor := 0

	// Loop through every character, we don't need the index.
	for _, char := range input {
		// Char 40 == (
		// Char 41 == )
		// ( == Up
		// ) == Down
		if char == 40 {
			currentFloor++
		} else {
			currentFloor--
		}
	}

	log.Printf("Current floor: %d", currentFloor)
}
