package tlp

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"encoding/hex"
	"io"
)

func encrypt(key []byte, payload []byte) (string, error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}

	// Create a random IV (Initialization Vector)
	iv := make([]byte, aes.BlockSize)
	if _, err := io.ReadFull(rand.Reader, iv); err != nil {
		return "", err
	}

	// Create a new AES cipher block mode
	stream := cipher.NewCTR(block, iv)

	// Encrypt the payload
	ciphertext := make([]byte, len(payload))
	stream.XORKeyStream(ciphertext, payload)

	// Prepend the IV to the ciphertext
	ciphertext = append(iv, ciphertext...)

	// Convert the byte slice to a hexadecimal string
	ciphertextHex := hex.EncodeToString(ciphertext)

	return ciphertextHex, nil
}


/* 
func main() {
	// Define a 256-bit (32-byte) secret key
	key := []byte("3bad4a7985551159ab3431fe63246c5e")

	// Data to be encrypted
	payload := []byte("Hello World")

	// Encrypt the data
	encryptedData, err := encrypt(key, payload)
	if err != nil {
		fmt.Println("Encryption error:", err)
		return
	}

	fmt.Println("Encrypted Data (Hex):", encryptedData)
}
 */