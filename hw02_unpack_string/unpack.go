package hw02unpackstring

import (
	"errors"
	"strconv"
	"strings"
	"unicode"
)

var ErrInvalidString = errors.New("invalid string")

func Unpack(inputStr string) (string, error) {	
	var builtStr strings.Builder
    chars := []rune(inputStr)
    for i := 0; i < len(chars); i++ {
		// checking if the last symbol is digit
		if unicode.IsDigit(chars[i]) && i == len(chars) - 1 { break } 
		// checking if digit is in the beginning of string 
		if i == 0 && unicode.IsDigit(chars[i]) { return "", errors.New("некорректная строка") } 
		// checking if there are two digits
		if unicode.IsDigit(chars[i]) && unicode.IsDigit(chars[i+1]) { return "", errors.New("некорректная строка") } 
		// if chars[i] == 92 { return "\\", errors.New("символ 92") }
		if i+1 < len(chars) && unicode.IsDigit(chars[i+1]) { // checking if a next symbol is a digit
			repeatTimes, err := strconv.Atoi(string(chars[i+1]))
			if err != nil { return "", errors.New("ошибка при конвертации символа в цифру") }
				builtStr.WriteString(strings.Repeat(string(chars[i]),repeatTimes))
			continue
		}
		if unicode.IsDigit(chars[i]) {continue}
		builtStr.WriteString(string(chars[i]))
    }

	return builtStr.String(), nil
}