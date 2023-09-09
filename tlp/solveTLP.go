package main

import (
	"crypto/sha256"
	"encoding/hex"
	"fmt"
)

func main() {
	seed := "Hello World"
	numIterations := 26665975

	// Initialize a variable to hold the current hash value
	currentHash := seed

	for i := 0; i < numIterations; i++ {
		hash := sha256.New()
		hash.Write([]byte(currentHash))
		hashBytes := hash.Sum(nil)
		currentHash = hex.EncodeToString(hashBytes)
	}

	fmt.Println("Final Hash:", currentHash)
}