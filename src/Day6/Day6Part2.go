package main

import (
	"bufio"
	"log"
	"os"
	"strconv"
	"strings"
)

func main() {

	file, err := os.Open("src/Day6/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	// x,y grid and whether their brightness.
	lightGrid := map[int]map[int]int{}

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the text of the current line.
		line := scanner.Text()

		words := strings.Split(line, " ")

		// The toggle instruction.
		if words[0] == "toggle" {

			// Split the correct words, toggle x,y through x,y
			start := strings.Split(words[1], ",")
			end := strings.Split(words[3], ",")

			xStart, yStart, xEnd, yEnd := getStartAndEndCoords(start, end)

			// Loop through x positions.
			for x := xStart; x <= xEnd; x++ {

				// Ensure map.
				if lightGrid[x] == nil {
					lightGrid[x] = map[int]int{}
				}

				// Loop through y positions.
				for y := yStart; y <= yEnd; y++ {
					lightGrid[x][y] += 2
				}
			}
		} else {
			// The turn instruction.
			start := strings.Split(words[2], ",")
			end := strings.Split(words[4], ",")

			xStart, yStart, xEnd, yEnd := getStartAndEndCoords(start, end)

			// Turn brightness up or down.
			brightnessUp := false
			if words[1] == "on" {
				brightnessUp = true
			}

			// Loop through x positions.
			for x := xStart; x <= xEnd; x++ {

				// Ensure map.
				if lightGrid[x] == nil {
					lightGrid[x] = map[int]int{}
				}

				// Loop through y positions.
				for y := yStart; y <= yEnd; y++ {
					if brightnessUp {
						lightGrid[x][y]++
					} else {
						if lightGrid[x][y] > 0 {
							lightGrid[x][y]--
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