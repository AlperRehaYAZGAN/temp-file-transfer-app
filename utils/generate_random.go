package utils

import (
	"math/rand"
	"time"
)

func GenerateRandomDigitString(length int) string {
	// generate random string of digits with length
	domain := "0123456789"
	var bytes = make([]byte, length)

	// increase randomness
	now := time.Now().UnixNano()
	rand.Seed(now)

	for i := 0; i < length; i++ {
		// get random index
		index := rand.Intn(len(domain))
		// get random character
		bytes[i] = domain[index]
		// append character to bytes
	}

	return string(bytes)
}
