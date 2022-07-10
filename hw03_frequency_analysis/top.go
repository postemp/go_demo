package hw03frequencyanalysis

import (
	"regexp"
	"sort"
	"strings"
)

// конвертируем строку в слайс, очищаем от знаков
// пунктуации и приводим нижнему регистру.
func convertStrToSlice(inputString string) []string {
	rawWords := strings.Fields(inputString)
	cleanedWords := make([]string, 0)

	rePunctuationMarks := regexp.MustCompile(`[,..:]`)
	for _, word := range rawWords {
		if word == "-" {
			continue
		}
		word = rePunctuationMarks.ReplaceAllString(word, "")
		word = strings.ToLower(word)
		cleanedWords = append(cleanedWords, word)
	}
	return cleanedWords
}

func SortedStruct(inputString string) []string {
	cleanedWords := convertStrToSlice(inputString)
	truncedWords := make([]string, len(cleanedWords))
	copy(truncedWords, cleanedWords)

	var countedWords []struct {
		Word    string
		Counter int
	}
	var cntWord struct {
		Word    string
		Counter int
	}

	// превращаем слайс слов в структуру, где у каждого слова есть количество вхождений
	for _, word := range cleanedWords {
		ifFounded := 0
		for indx, wordFnd := range truncedWords {
			if word == wordFnd {
				ifFounded++
				cntWord.Word = word
				cntWord.Counter = 0
				if ifFounded == 1 {
					countedWords = append(countedWords, cntWord)
				}
				countedWords[len(countedWords)-1].Counter++
				truncedWords = append(truncedWords[:indx], truncedWords[indx+1:]...) // delete by index
				// fmt.Println(word)
				// fmt.Println(countedWords[len(countedWords)-1].Counter)
			}
		}
	}
	sort.Slice(countedWords, func(i, j int) bool { return countedWords[i].Counter > countedWords[j].Counter })
	// fmt.Println(countedWords[:11])

	var tmpCountedWords []struct {
		Word    string
		Counter int
	}

	var outCountedWords []struct {
		Word    string
		Counter int
	}

	// сортируем слова с одинаковым количеством вхождений
	for indx, item := range countedWords {
		// fmt.Println("indx = ", indx)
		if item.Counter == countedWords[indx+1].Counter {
			tmpCountedWords = append(tmpCountedWords, item)
		} else {
			tmpCountedWords = append(tmpCountedWords, item)
			sort.Slice(tmpCountedWords, func(i, j int) bool { return tmpCountedWords[i].Word < tmpCountedWords[j].Word })
			outCountedWords = append(outCountedWords, tmpCountedWords...)
			tmpCountedWords = nil
		}
		if indx == len(countedWords)-2 {
			break
		}
	}
	// fmt.Println(outCountedWords[:10]).

	// формируем конечный слайс
	outputWords := make([]string, 0)

	for _, item := range outCountedWords {
		outputWords = append(outputWords, item.Word)
	}
	return outputWords
}

func Top10(inputString string) []string {
	if len(inputString) == 0 {
		return nil
	}
	outputWords := SortedStruct(inputString)
	return outputWords[:10]
}
