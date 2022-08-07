package hw03frequencyanalysis

import (
	"regexp"
	"strings"
)

var outWordQuantity int

func Top10(inputString string) []string {
	outWordQuantity = 10

	rawWords := strings.Fields(inputString)
	cleanedWords := make([]string, 0)

	// конвертируем строку в слайс, очищаем от знаков
	// пунктуации и приводим нижнему регистру.
	rePunctuationMarks := regexp.MustCompile(`[!/@#$%^&*(),..:]`)
	for _, word := range rawWords {
		if word == "-" {
			continue
		}
		word = rePunctuationMarks.ReplaceAllString(word, "")
		if word == "" {
			continue
		}
		word = strings.ToLower(word)
		cleanedWords = append(cleanedWords, word)
	}

	if len(cleanedWords) == 0 {
		return nil
	}

	outputWords := SortedSlice(cleanedWords)
	return outputWords
}

func SortedSlice(inputSlice []string) []string {
	truncedWords := make([]string, len(inputSlice))
	copy(truncedWords, inputSlice)
	countedWords := map[string]int{}
	for _, word := range inputSlice {
		for indx, wordFnd := range truncedWords {
			if word == wordFnd {
				countedWords[word] = countedWords[word] + 1 //nolint:gocritic
				truncedWords = append(truncedWords[:indx], truncedWords[indx+1:]...)
			}
		}
	}
	sortedMap := sortMap(countedWords)
	return sortedMap
}

func sortMap(inputMap map[string]int) []string {
	outSlice := make([]string, 0)

	counter := 0
	for outKey, outValue := range inputMap {
		counter++
		for inKey, inValue := range inputMap {
			if inValue > outValue {
				outKey = inKey
				outValue = inValue
			} else if inValue == outValue {
				if inKey < outKey {
					outKey = inKey
					outValue = inValue
				}
			}
		}
		outSlice = append(outSlice, outKey)
		delete(inputMap, outKey)
		if outWordQuantity == counter {
			break
		}
	}
	return outSlice
}
