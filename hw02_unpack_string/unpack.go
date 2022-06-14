package hw02unpackstring

import (
	"errors"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(source string) (string, error) {
	runes := []rune(source)
	for i, char := range runes { // Check that first char is not number
		if i == 0 && unicode.IsDigit(char) {
			return source, ErrInvalidString
		}
	}
	var isPreviousNumber bool
	for _, char := range runes { // Search sequence of numbers
		if unicode.IsDigit(char) {
			if isPreviousNumber {
				return source, ErrInvalidString
			}
			isPreviousNumber = true
		} else {
			isPreviousNumber = false
		}
	}
	var result = ""
	for i, char := range runes { // Unpacking validated string
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
