package utils

import (
	"crypto/sha512"
	"encoding/base64"
)

func HashPassword(text, salt string) string {
	if text == "" {
		return ""
	}

	clearTextBytes := []byte(text)
	saltBytes := []byte(salt)
	sha512Hasher := sha512.New()

	clearTextBytes = append(clearTextBytes, saltBytes...)
	sha512Hasher.Write(clearTextBytes)

	hashedPasswordBytes := sha512Hasher.Sum(nil)

	return base64.URLEncoding.EncodeToString(hashedPasswordBytes)
}

func CompareHash(hashClient, expectedHash string) bool {
	return hashClient == expectedHash
}
