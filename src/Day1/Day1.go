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

	// We start at 0 so we can recognize if we already have been to the basement.
	// The position is always > 0 when we have been to the basement.
	firstBasementPosition := 0

	for position, char := range input {
		// Char 40 == (
		// Char 41 == )
		// ( == Up
		// ) == Down
		if char == 40 {
			currentFloor++
		} else {
			currentFloor--
		}

		if firstBasementPosition == 0 && currentFloor == -1 {

			// We add one, we have a zero-indexed number, while we need to provide a one-indexed number.
			firstBasementPosition = (position + 1)
		}
	}

	log.Printf("Part 1: Current floor: %d", currentFloor)
	log.Printf("Part 2: Position of the character that took us to the basement floor: %d", firstBasementPosition)
}
