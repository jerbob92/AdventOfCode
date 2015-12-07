package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)


func main() {
	portMap := map[string]uint16{}
	portVisited := map[string]bool{}

	log.Printf("Part 1: The signal ultimately provided to wire %s is %d", "a", solve(portMap, portVisited, "a"))

	// Put the result from a into b.
	portMap = map[string]uint16{
		"b": portMap["a"],
	}
	portVisited = map[string]bool{
		"b": portVisited["a"],
	}

	log.Printf("Part 2: The signal ultimately provided to wire %s is %d", "a", solve(portMap, portVisited, "a"))
}

func solve(portMap map[string]uint16, portVisited map[string]bool, requestedWire string) uint16 {
	file, err := os.Open("src/Day7/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	for !portVisited[requestedWire] {
		file.Seek(0, 0)
		// We use a scanner to loop through every line.
		scanner := bufio.NewScanner(file)
		for scanner.Scan() {
			// Get the text of the current line.
			line := scanner.Text()

			// Split on operation -> destination
			parts := strings.Split(line, " -> ")
			operation := parts[0]
			destination := parts[1]

			// Bail out if we already been here.
			if portVisited[destination] {
				continue;
			}

			// Split on the operation so we know what to do.
			operations := strings.Split(operation, " ")

			// Operation with length 1 is a number or a direct port,
			if (len(operations) == 1) {
				if number, err := strconv.Atoi(operations[0]); err == nil {

					// Directly apply if the input is a number.
					portMap[destination] = uint16(number)
					portVisited[destination] = true
				} else {
					if (portVisited[operations[0]]) {
						portMap[destination] = portMap[operations[0]]
						portVisited[destination] = true
					} else {

						// Bail out if we don't have our source yet.
						continue;
					}
				}
			} else if (len(operations) == 2) {
				// Operations with length 2 are always NOT operations.
				if number, err := strconv.Atoi(operations[1]); err == nil {

					// Directly apply if the input is a number.
					portMap[destination] = ^uint16(number)
					portVisited[destination] = true
				} else {
					if (portVisited[operations[1]]) {
						portMap[destination] = ^portMap[operations[1]]
						portVisited[destination] = true
					} else {

						// Bail out if we don't have our source yet.
						continue;
					}
				}
			} else {
				haveLeft := false
				haveRight := false
				leftValue := uint16(0)
				rightValue := uint16(0)

				// Check if the left side is a number or a port source.
				if leftNumber, err := strconv.Atoi(operations[0]); err == nil {
					leftValue = uint16(leftNumber)
					haveLeft = true
				} else {
					if (portVisited[operations[0]]) {
						leftValue = portMap[operations[0]]
						haveLeft = true
					}
				}

				// Check if the right side is a number or a port source.
				if rightNumber, err := strconv.Atoi(operations[2]); err == nil {
					rightValue = uint16(rightNumber)
					haveRight = true
				} else {
					if (portVisited[operations[2]]) {
						rightValue = portMap[operations[2]]
						haveRight = true
					}
				}

				if (haveLeft && haveRight) {

					// Apply the operation that we need to do.
					switch(operations[1]) {
					case "AND":
						portMap[destination] = leftValue&rightValue
						break;
					case "OR":
						portMap[destination] = leftValue|rightValue
						break;
					case "LSHIFT":
						portMap[destination] = leftValue<<uint16(rightValue)
						break;
					case "RSHIFT":
						portMap[destination] = leftValue>>uint16(rightValue)
						break;
					}
					portVisited[destination] = true
				} else {

					// Bail out if we don't have our sources yet.
					continue
				}
			}
		}
	}

	return portMap[requestedWire]
}