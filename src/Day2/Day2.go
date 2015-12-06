package main
import (
	"os"
	"bufio"
	"log"
	"strings"
	"strconv"
)

func main() {

	file, err := os.Open("src/Day2/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	neededWrappingPaper := 0

	// We use a scanner to loop through every line.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		// Get the text of the current line.
		packageDimension := scanner.Text()

		// Split on x, x marks a part of the dimension.
		// packageDimensions[0] = l
		// packageDimensions[1] = w
		// packageDimensions[2] = h
		packageDimensionsStrings := strings.Split(packageDimension, "x")

		// Convert the dimensions to ints
		packageDimensionsInts := map[int]int{}
		packageDimensionsInts[0], _ = strconv.Atoi(packageDimensionsStrings[0])
		packageDimensionsInts[1], _ = strconv.Atoi(packageDimensionsStrings[1])
		packageDimensionsInts[2], _ = strconv.Atoi(packageDimensionsStrings[2])

		// Calculate l*w
		lwsheet := (packageDimensionsInts[0] * packageDimensionsInts[1])
		slackPaper := lwsheet

		// Calculate w*h
		whsheet := (packageDimensionsInts[1] * packageDimensionsInts[2])
		if whsheet < slackPaper {
			slackPaper = whsheet
		}

		// Calculate h*l
		hlsheet := (packageDimensionsInts[2] * packageDimensionsInts[0])
		if hlsheet < slackPaper {
			slackPaper = hlsheet
		}

		// Add the sheets to the needed wrapping paper.
		// We need 2 of every sheet.
		neededWrappingPaper += lwsheet*2
		neededWrappingPaper += whsheet*2
		neededWrappingPaper += hlsheet*2

		// Add the slack paper.
		neededWrappingPaper += slackPaper
	}

	log.Printf("We need to order %d quare feet of wrapping paper.", neededWrappingPaper)
}
