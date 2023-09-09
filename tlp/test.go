package tlp

import (
	"crypto/sha256"
	"encoding/hex"
)



func solveTLP(seed string, numIterations int)(string) {

	// Initialize a variable to hold the current hash value
	currentHash := seed

	for i := 0; i < numIterations; i++ {
		hash := sha256.New()
		hash.Write([]byte(currentHash))
		hashBytes := hash.Sum(nil)
		currentHash = hex.EncodeToString(hashBytes)
	}
	return currentHash[:32]
}


