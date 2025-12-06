package security

import (
	"crypto/hmac"
	"crypto/sha256"
	"encoding/hex"
)

func GenerateHMAC(secret, data string) string {
	mac := hmac.New(sha256.New, []byte(secret))
	mac.Write([]byte(data))
	return hex.EncodeToString(mac.Sum(nil))
}

func ValidateHMAC(secret, data, given string) bool {
	expected := GenerateHMAC(secret, data)
	return hmac.Equal([]byte(expected), []byte(given))
}
