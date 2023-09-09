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
	fmt.Printf("Final Encryption-Decryption Key: %s\n", key)
	fmt.Printf("Number of Iterations: %d\n", iters)
}

/* 

Final Encryption-Decryption Key: 36bcbe93fdfcb8b1838d61a186270a7bdf9578243c82f6ad680c3c7c20899365
Number of Iterations: 52497602

Final Encryption-Decryption Key: 1e93ac7f1eb5938e6ffde179bee107b28b7aa2ee55c62a3d5069f3a4884a84d7
Number of Iterations: 53339407

Final Encryption-Decryption Key: d7914d745741c1b5a42a6b8431e2e7cb6ea92e71a8b7e9714ef19b5ccc2efcb9
Number of Iterations: 53154103

Final Encryption-Decryption Key: d287756aeaabd73717c112f7430bac83fd0d47386dc6ed68d8ac4d46496011da
Number of Iterations: 53309261

Final Encryption-Decryption Key: 5989e5c786f5158645e5960d1f04b75f08a45a365a1c5f122a5649b80a5d19d8
Number of Iterations: 52848296

Final Encryption-Decryption Key: 2e3617be04786b8c8457a775ebfa0fd0677cdd34a01f379e55608e9bb476b3ca
Number of Iterations: 53403837

Final Encryption-Decryption Key: 33afabfb0c8e16873aadd38162dc819416a6045e0ecc1f2fc4ad6dc99b6614e2
Number of Iterations: 53075165

Final Encryption-Decryption Key: 752650d1ec082a17d41056d32fa6118f0831c78ce2d66a0079122fcef63cbbd2
Number of Iterations: 53051990

Final Encryption-Decryption Key: 4c6297ae041014c2133ccba2546c8a61a39121795da7f25dbfa8dc12f7398491
Number of Iterations: 52915720

Final Encryption-Decryption Key: a34f8e293e176b43dc70643205cbd37bc187858f75c9661aeb129a34a78b117e
Number of Iterations: 52805174

*/