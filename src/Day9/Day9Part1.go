package main

import (
	"bufio"
	"log"
	"os"
	"strings"
	"strconv"
)

type Location struct  {
	Destinations map[string]int
}

func main() {

	file, err := os.Open("src/Day9/input.txt")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	locationMap := map[string]map[string]int{}

	// We use a scanner to loop through every line.
	// We save every destination with it's destinations distances.
	scanner := bufio.NewScanner(file)
	for scanner.Scan() {
		line :=  scanner.Text()
		parts := strings.Split(line, " ")

		distance, _ := strconv.Atoi(parts[4])

		// Make sure we have the maps.
		if (locationMap[parts[0]] == nil) {
			locationMap[parts[0]] = map[string]int{}
		}
		if (locationMap[parts[2]] == nil) {
			locationMap[parts[2]] = map[string]int{}
		}

		// Add both start -> destination and destination -> start to the map.
		locationMap[parts[0]][parts[2]] = distance
		locationMap[parts[2]][parts[0]] = distance
	}

	minRouteDistance := 0;
	minRoute := ""

	// Loop through all the locations to find the best route.
	for key, _ := range locationMap {
		currentRoute := []string{key}

		// Get best route for this starting location.
		route, newMinDistance := findLocationEnd(locationMap, currentRoute, key, 1, 0)

		// Is this the first or shorter than any found?
		if (minRouteDistance == 0 || newMinDistance < minRouteDistance) {
			minRouteDistance = newMinDistance
			minRoute = route
		}
	}

	log.Printf("The shortest route was %d using route %s.", minRouteDistance, minRoute);
}

func findLocationEnd(locationMap map[string]map[string]int, currentRoute []string, currentLocation string, currentLength int, currentDistance int) (string, int) {
	if (currentLength == len(locationMap)) {

		// Bail out when we went to every city.
		return strings.Join(currentRoute, " -> "), currentDistance;
	} else {

		// Loop through all possible connections
		maxDistance := 0;
		maxRoute := "";
		for key, distanceBetween := range locationMap[currentLocation] {

			// Check if the current destination isn't already visited in this loop, we only need to visit every place once.
			if (!inCurrentRoute(currentRoute, key)) {

				// Combine the current destination into the route.
				newCurrentRoute := append(currentRoute, key)

				// Remember we are one destination further.
				newLength := currentLength + 1;

				// Add the current destination to the total destional.
				newCurrentDistance := currentDistance + distanceBetween
				newMaxRoute, newCurrentDistance := findLocationEnd(locationMap, newCurrentRoute, key, newLength, newCurrentDistance)

				// Check if this is the first combination or this one is faster than the known routes.
				if (maxDistance == 0 || newCurrentDistance < maxDistance) {
					maxDistance = newCurrentDistance;
					maxRoute = newMaxRoute
				}
			}
		}
		return maxRoute, maxDistance;
	}
}

// Helper to see if the matches destination is in our route.
func inCurrentRoute(currentRoute []string, currentLocation string) bool {
	for _, location := range currentRoute {
		if (location == currentLocation) {
			return true
		}
	}
	return false
}