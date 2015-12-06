package main
import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func main() {

	file, err := os.Open("src/Day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// x,y grid and whether their brightness.
	lightGrid := map[int]map[int]int{}

	// Loop through x positions.
	for i := 0; i <= 999; i++ {

		if lightGrid[i] == nil {
			lightGrid[i] = map[int]int{}
		}

		// Loop through y positions.
		for i2 := 0; i2 <= 999; i2++ {
			lightGrid[i][i2] = 0
		}
	}

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the text of the current line.
		line := scanner.Text()

		words := strings.Split(line, " ")

		// The toggle instruction.
		if (words[0] == "toggle") {

			// Split the correct words, toggle x,y through x,y
			start := strings.Split(words[1], ",")
			end := strings.Split(words[3], ",")

			xStart, yStart, xEnd, yEnd := getStartAndEndCoords(start, end)

			instructionLightGrid := getLightsInGrid(xStart, yStart, xEnd, yEnd)

			for x, ymap := range instructionLightGrid {
				// Add  2 to the brightness for toggles.
				for y, _ := range ymap {
					lightGrid[x][y] += 2;
				}
			}
		} else {
			// The turn instruction.
			start := strings.Split(words[2], ",")
			end := strings.Split(words[4], ",")

			xStart, yStart, xEnd, yEnd := getStartAndEndCoords(start, end)

			instructionLightGrid := getLightsInGrid(xStart, yStart, xEnd, yEnd)

			// Turn brightness up or down.
			brightnessUp := false
			if words[1] == "on" {
				brightnessUp = true
			}

			for x, ymap := range instructionLightGrid {
				for y, _ := range ymap {
					if brightnessUp {
						lightGrid[x][y]++
					} else {
						lightGrid[x][y]--
						if (lightGrid[x][y] < 0) {
							lightGrid[x][y] = 0
						}
					}
				}
			}
		}
	}

	totalLigthBrightness := 0

	// Count brightness of all lights.
	for _, ymap := range lightGrid {
		for _, lightBrightness := range ymap {
			totalLigthBrightness += lightBrightness
		}
	}

	log.Printf("Total brightness of all lights combined %d", totalLigthBrightness)
}


func getStartAndEndCoords(start []string, end []string) (int, int, int, int) {
	xStart, _ := strconv.Atoi(start[0])
	yStart, _ := strconv.Atoi(start[1])

	xEnd, _ := strconv.Atoi(end[0])
	yEnd, _ := strconv.Atoi(end[1])
	return xStart, yStart, xEnd, yEnd
}

func getLightsInGrid(xStart int, yStart int, xEnd int, yEnd int) map[int]map[int]bool {

	lightGrid := map[int]map[int]bool{}

	// Loop through x positions.
	for i := xStart; i <= xEnd; i++ {

		if lightGrid[i] == nil {
			lightGrid[i] = map[int]bool{}
		}

		// Loop through y positions.
		for i2 := yStart; i2 <= yEnd; i2++ {
			lightGrid[i][i2] = true
		}
	}

	return lightGrid
}
