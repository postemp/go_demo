package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) { //nolint:gocognit
	var builtStr strings.Builder
	chars := []rune(inputStr)
	for i := 0; i < len(chars); i++ {
		// checking if the last symbol is digit
		if unicode.IsDigit(chars[i]) && i == len(chars)-1 {
			break
		}
		// checking if digit is in the beginning of string
		if i == 0 && unicode.IsDigit(chars[i]) {
			return "", errors.New("некорректная строка")
		}
		// checking if there are two digits
		if unicode.IsDigit(chars[i]) && unicode.IsDigit(chars[i+1]) {
			return "", errors.New("некорректная строка")
		}
		// checking if a symbol is '\' - 92
		if chars[i] == 92 && (unicode.IsDigit(chars[i+1]) || chars[i+1] == 92) {
			if i+2 < len(chars) && unicode.IsDigit(chars[i+2]) {
				repeatTimes, err := strconv.Atoi(string(chars[i+2]))
				if err != nil {
					return "", errors.New("ошибка при конвертации символа в цифру")
				}
				builtStr.WriteString(strings.Repeat(string(chars[i+1]), repeatTimes))
				i++
				continue
			}
			builtStr.WriteString(string(chars[i+1]))
			i++
			continue
		}
		if i+1 < len(chars) && unicode.IsDigit(chars[i+1]) { // checking if the next symbol is a digit
			repeatTimes, err := strconv.Atoi(string(chars[i+1]))
			if err != nil {
				return "", errors.New("ошибка при конвертации символа в цифру")
			}
			builtStr.WriteString(strings.Repeat(string(chars[i]), repeatTimes))
			continue
		}
		if unicode.IsDigit(chars[i]) {
			continue
		}
		builtStr.WriteString(string(chars[i]))
	}
	return builtStr.String(), nil
}
