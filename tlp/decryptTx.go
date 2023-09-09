package tlp

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"
)

func decrypt(key []byte, ciphertextHex string) ([]byte, error) {
	// Convert the hexadecimal string back to a byte slice
	ciphertext, err := hex.DecodeString(ciphertextHex)
	if err != nil {
		return nil, err
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	// Extract the IV from the ciphertext
	iv := ciphertext[:aes.BlockSize]
	ciphertext = ciphertext[aes.BlockSize:]

	// Create a new AES cipher block mode
	stream := cipher.NewCTR(block, iv)

	// Decrypt the ciphertext
	plaintext := make([]byte, len(ciphertext))
	stream.XORKeyStream(plaintext, ciphertext)

	return plaintext, nil
}


/* 
func main() {
	// Define a 256-bit (32-byte) secret key
	key := []byte("3bad4a7985551159ab3431fe63246c5e")
	encryptedData := "af23ce80c77246afb857bec7668df760a7c2ecb03448db28e45c45"

	// Decrypt the data
	decryptedData, err := decrypt(key, encryptedData)
	if err != nil {
		fmt.Println("Decryption error:", err)
		return
	}

	fmt.Println("Decrypted Data:", string(decryptedData))
}

 */
