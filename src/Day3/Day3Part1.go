package main

import (
	"io/ioutil"
	"log"
)

func main() {

	input, err := ioutil.ReadFile("src/Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// We use a map in a map so we have xy coords with amount of presents received.
	housesGrid := map[int]map[int]int{}

	amountOfHouses := 0

	xIndex := 0
	yIndex := 0

	// Make sure we have a map here.
	// Deliver a package at the first house.
	housesGrid[xIndex] = map[int]int{}
	housesGrid[xIndex][yIndex]++
	amountOfHouses++

	for _, char := range input {
		// Char 94 == ^
		// Char 118 == v
		// Char 60 == <
		// Char 62 == >
		switch char {
		case 94:
			yIndex++
			break
		case 118:
			yIndex--
			break
		case 60:
			xIndex--
			break
		case 62:
			xIndex++
			break
		}

		// If x doesn't exist yet, create a new map.
		if housesGrid[xIndex] == nil {
			housesGrid[xIndex] = map[int]int{}
		}

		// If y doesn't exist yet, we found a new house.
		if housesGrid[xIndex][yIndex] == 0 {
			amountOfHouses++
		}

		// Deliver a package at the current house
		housesGrid[xIndex][yIndex]++
	}

	log.Printf("Part 1: amount of houses we visited: %d", amountOfHouses)
}
