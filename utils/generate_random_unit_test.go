package utils_test

import (
	"testing"

	"github.com/AlperRehaYAZGAN/temp-file-transfer-app/utils"
)

func TestGenerateRandomDigitString(t *testing.T) {
	// test generate random string of digits with length
	length := 10
	randomString := utils.GenerateRandomDigitString(length)
	if len(randomString) != length {
		t.Errorf("random string length is not equal to length")
	}

	// test all characters are digits
	for _, char := range randomString {
		if char < '0' || char > '9' {
			t.Errorf("random string contains non-digit character")
		}
	}
}
