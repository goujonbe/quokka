package password

import (
	"crypto/rand"
	"math/big"
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

func buildCharacterSet(includeLowerLetters, includeUpperLetters, includeDigits, includeSymbols bool) (characterSet string) {
	if includeLowerLetters {
		characterSet += LowerLetters
	}
	if includeUpperLetters {
		characterSet += UpperLetters
	}
	if includeDigits {
		characterSet += Digits
	}
	if includeSymbols {
		characterSet += Symbols
	}
	return
}

func getRandomElement(characterSet string) (string, error) {
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
	if err != nil {
		return "", nil
	}
	return string(characterSet[randomNumber.Int64()]), nil
}
