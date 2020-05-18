package password

import (
	"crypto/rand"
	"math/big"
	"strings"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

func Generate(length int, includeLowerLetters, includeUpperLetters, includeDigits, includeSymbols bool) (string, error) {
	var password string
	characterSet := buildCharacterSet(includeLowerLetters, includeUpperLetters, includeDigits, includeSymbols)
	for i := 0; i < length; i++ {
		character, err := getRandomElement(characterSet)
		if err != nil {
			return "", err
		}
		password += character
	}
	return password, nil
}

func buildCharacterSet(includeLowerLetters, includeUpperLetters, includeDigits, includeSymbols bool) string {
	var characterSetBuilder strings.Builder
	if includeLowerLetters {
		characterSetBuilder.WriteString(LowerLetters)
	}
	if includeUpperLetters {
		characterSetBuilder.WriteString(UpperLetters)
	}
	if includeDigits {
		characterSetBuilder.WriteString(Digits)
	}
	if includeSymbols {
		characterSetBuilder.WriteString(Symbols)
	}
	return characterSetBuilder.String()
}

func getRandomElement(characterSet string) (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
	if err != nil {
		return "", nil
	}
	return string(characterSet[randomNumber.Int64()]), nil
}
