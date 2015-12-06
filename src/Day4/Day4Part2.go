package main

import (
	"crypto/md5"
	"encoding/hex"
	"log"
	"strconv"
	"strings"
)

func main() {
	input := "iwrupvqb"
	hasher := md5.New()
	resultHash := ""
	lowestNumber := -1
	prefix := "000000"

	// Create a new MD5 hash until we have one that starts with 000000
	for !strings.HasPrefix(resultHash, prefix) {
		// Reset the hashing machine.
		hasher.Reset()

		// Try the next number.
		lowestNumber++

		// Generate the hash for the next number.
		hasher.Write([]byte(input + strconv.Itoa(lowestNumber)))
		resultHash = hex.EncodeToString(hasher.Sum(nil))
	}

	if lowestNumber > -1 {
		log.Printf("The lowest number that results in a hash that starts with %s with key %s + i is where i = %d", prefix, input, lowestNumber)
	}
}
