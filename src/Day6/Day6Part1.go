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

	// x,y grid and whether they are turned on or not.
	lightGrid := map[int]map[int]bool{}

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
					lightGrid[x] = map[int]bool{}
				}

				// Loop through y positions.
				for y := yStart; y <= yEnd; y++ {
					if lightGrid[x][y] {
						lightGrid[x][y] = false
					} else {
						lightGrid[x][y] = true
					}
				}
			}

		} else {
			// The turn instruction.
			start := strings.Split(words[2], ",")
			end := strings.Split(words[4], ",")

			xStart, yStart, xEnd, yEnd := getStartAndEndCoords(start, end)

			// Turn on or off?
			newStatus := false
			if words[1] == "on" {
				newStatus = true
			}

			// Loop through x positions.
			for x := xStart; x <= xEnd; x++ {

				// Ensure map.
				if lightGrid[x] == nil {
					lightGrid[x] = map[int]bool{}
				}

				// Loop through y positions.
				for y := yStart; y <= yEnd; y++ {
					lightGrid[x][y] = newStatus
				}
			}
		}
	}

	litLights := 0
	// Count amount of lights that are turned on.
	for _, ymap := range lightGrid {
		for _, lightStatus := range ymap {
			if lightStatus {
				litLights++
			}
		}
	}

	log.Printf("Part 1: There are %d lights lit", litLights)
}

func getStartAndEndCoords(start []string, end []string) (int, int, int, int) {
	xStart, _ := strconv.Atoi(start[0])
	yStart, _ := strconv.Atoi(start[1])

	xEnd, _ := strconv.Atoi(end[0])
	yEnd, _ := strconv.Atoi(end[1])
	return xStart, yStart, xEnd, yEnd
}
