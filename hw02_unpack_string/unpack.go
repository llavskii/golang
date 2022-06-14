package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	notValidString := checkPackedString(source)
	if notValidString != nil {
		return source, notValidString
	}
	var result = ""
	runes := []rune(source)
	for i, char := range runes {
		if unicode.IsDigit(char) {
			continue
		}
		var repeat int
		if i+1 == len(runes) || !unicode.IsDigit(runes[i+1]) { // Is it last symbol or next symbol is not digit?
			repeat = 1
		} else {
			repeat = int(runes[i+1] - '0')
		}
		result += strings.Repeat(string(char), repeat)
	}

	return result, nil
}

func checkPackedString(source string) error {
	if len(source) == 0 { // Is empty string?
		return nil
	}
	for i, char := range source { // Check that first char is not number
		if i == 0 && unicode.IsDigit(char) {
			return ErrInvalidString
		}
	}
	var isPreviousNumber bool
	for _, char := range source { // Search sequence of numbers
		if unicode.IsDigit(char) {
			if isPreviousNumber {
				return ErrInvalidString
			}
			isPreviousNumber = true
		} else {
			isPreviousNumber = false
		}
	}
	return nil
}
