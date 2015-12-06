package main
import (
	"log"
	"io/ioutil"
)

func main() {

	input, err := ioutil.ReadFile("src/Day3/input.txt")
	if err != nil {
		log.Fatal(err)
	}

	// We use a map in a map so we have xy coords with amount of presents received.
	housesGrid := map[int]map[int]int{}

	amountOfHouses := 0

	xIndexSanta := 0
	yIndexSanta := 0

	xIndexRobot := 0
	yIndexRobot := 0

	// Make sure we have a map here.
	housesGrid[xIndexSanta] = map[int]int{}

	// Both robot and normal santa start at the 0,0, they get 2 presents.
	housesGrid[xIndexSanta][yIndexSanta]++;
	housesGrid[xIndexRobot][yIndexRobot]++;

	// But it's still one house.
	amountOfHouses++;

	// False = Normal Santa, True = Robot Santa.
	stepTurnRobotSanta := false

	for _, char := range input {
		// Char 94 == ^
		// Char 118 == v
		// Char 60 == <
		// Char 62 == >

		if (stepTurnRobotSanta) {
			switch(char) {
			case 94:
				yIndexRobot++
				break;
			case 118:
				yIndexRobot--
				break;
			case 60:
				xIndexRobot--
				break;
			case 62:
				xIndexRobot++
				break;
			}

			// If x doesn't exist yet, create a new map.
			if (housesGrid[xIndexRobot] == nil) {
				housesGrid[xIndexRobot] = map[int]int{}
			}

			// If y doesn't exist yet, we found a new house.
			if (housesGrid[xIndexRobot][yIndexRobot] == 0) {
				amountOfHouses++;
			}

			// Deliver a package at the current house
			housesGrid[xIndexRobot][yIndexRobot]++;

			// Let normal santa walk again.
			stepTurnRobotSanta = false
		} else {
			switch(char) {
			case 94:
				yIndexSanta++
				break;
			case 118:
				yIndexSanta--
				break;
			case 60:
				xIndexSanta--
				break;
			case 62:
				xIndexSanta++
				break;
			}

			// If x doesn't exist yet, create a new map.
			if (housesGrid[xIndexSanta] == nil) {
				housesGrid[xIndexSanta] = map[int]int{}
			}

			// If y doesn't exist yet, we found a new house.
			if (housesGrid[xIndexSanta][yIndexSanta] == 0) {
				amountOfHouses++;
			}

			// Deliver a package at the current house
			housesGrid[xIndexSanta][yIndexSanta]++;

			// Let robot santa walk again.
			stepTurnRobotSanta = true
		}
	}

	log.Printf("Part 2: amount of houses we visited: %d", amountOfHouses)
}
