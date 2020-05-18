package password

import (
	"crypto/rand"
	"errors"
	"math/big"
	"strings"
)

const (
	LowerLetters = "abcdefghijklmnopqrstuvwxyz"
	UpperLetters = "ABCDEFGHIJKLMNOPQRSTUVWXYZ"
	Digits       = "0123456789"
	Symbols      = "~!@#$%^&*()_+`-={}|[]\\:\"<>?,./"
)

var ErrorEmptyCharacterSet = errors.New("Cannot select a character from an empty set.")

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
	if len(characterSet) == 0 {
		return "", ErrorEmptyCharacterSet
	}
	randomNumber, err := rand.Int(rand.Reader, big.NewInt(int64(len(characterSet))))
	if err != nil {
		return "", err
	}
	return string(characterSet[randomNumber.Int64()]), nil
}
