package checkfilevalidity

import (
	"crypto/sha256"
	"encoding/hex"
)

// Function CheckFileValidity takes a slice of bytes and returns a hexadecimal string. The function uses a hashing algorithm to determine the integrity and correctness of a banner file.
func CheckFileValidity(bannerFileData []byte) string {
	hasher := sha256.New()
	hasher.Write(bannerFileData)
	hashInBytes := hasher.Sum(nil)
	fileHash := hex.EncodeToString(hashInBytes)
	return fileHash
}
