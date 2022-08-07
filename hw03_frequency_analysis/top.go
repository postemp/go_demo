package hw03frequencyanalysis

import (
	"regexp"
	"strings"
)

var outWordQuantity int


func Top10(inputString string) []string {
	outWordQuantity = 10

	if len(convertStrToSlice(inputString)) == 0 {
		return nil
	}
	outputWords := SortedSlice(inputString)
	// return outputWords[:10]
	return outputWords
}

func SortedSlice(inputString string) []string {
	cleanedWords := convertStrToSlice(inputString)
	truncedWords := make([]string, len(cleanedWords))
	copy(truncedWords, cleanedWords)

	countedWords := map[string]int{}
	for _, word := range cleanedWords {
		for indx, wordFnd := range truncedWords {
			if word == wordFnd {
				countedWords[word] = countedWords[word] + 1
				truncedWords = append(truncedWords[:indx], truncedWords[indx+1:]...)
			}
		}
	}
	sortedMap := sortMap(countedWords)
	return sortedMap
}

func sortMap(inputMap map[string]int) []string  {
	outSlice := make([]string, 0)
	
	counter := 0
	for outKey, outValue := range inputMap {
		counter++
		for inKey, inValue := range inputMap {
			if (inValue > outValue) {
				outKey = inKey
				outValue = inValue
			} else if ( inValue == outValue) {
				if (inKey < outKey) {
					outKey = inKey
					outValue = inValue
				}
			}
		}
		outSlice = append(outSlice, outKey)
		delete(inputMap, outKey)
		if (outWordQuantity == counter)	{
			break
		}	
	}
	return outSlice
}

// конвертируем строку в слайс, очищаем от знаков
// пунктуации и приводим нижнему регистру.
func convertStrToSlice(inputString string) []string {
	rawWords := strings.Fields(inputString)
	cleanedWords := make([]string, 0)

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
	return cleanedWords
}