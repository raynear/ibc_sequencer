package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
	"time"
)

func createPuzzle(seed string, timeInSeconds int) (string, int) {
	key := seed
	startTime := time.Now()
	endTime := startTime.Add(time.Duration(timeInSeconds) * time.Second)
	iters := 0

	for time.Now().Before(endTime) {
		// Repetitively hash and count iterations
		// The goal is to get number of iterations during the specified time
		// Given the seed and the number of iterations, the solution to the puzzle, the encryption key could only be reached if the seed have been repetitievly hashed the {iterations} times
		hash := sha256.New()
		hash.Write([]byte(key))
		hashBytes := hash.Sum(nil)
		key = hex.EncodeToString(hashBytes)
		iters++
	}

	return key, iters
}

func main() {
	seed := "Hello World"
	timeInSeconds := 10

	key, iters := createPuzzle(seed, timeInSeconds)
	fmt.Printf("Final Key: %s\n", key)
	fmt.Printf("Number of Iterations: %d\n", iters)
}
