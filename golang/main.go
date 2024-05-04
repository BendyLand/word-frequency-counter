package main

import (
	"fmt"
	"golang.org/x/text/cases"
	"golang.org/x/text/language"
	"os"
	"strings"
)

func main() {
	text := readFile("../text.txt")
	words := splitIntoWords(text)
	words = removeNonLetterChars(words)
	words = convertAllToLowercase(words)
	wordCounts := countWords(words)
	displayCountedWords(text, wordCounts)
}

func displayCountedWords(original string, wordCounts map[string]int) {
	fmt.Println("Original Text:")
	fmt.Println(original)
	fmt.Println("Each word and its respective number of occurrences in the text:")
	numWords := 0
	for k, v := range wordCounts {
		caser := cases.Title(language.English)
		k = caser.String(k) // title case
		if v == 1 {
			fmt.Printf("'%s' is found 1 time in the text.\n", k)
		} else {
			fmt.Printf("'%s' is found %d times in the text.\n", k, v)
		}
		numWords++
	}
	fmt.Printf("Number of unique words: %d\n", numWords)
}

func countWords(words []string) map[string]int {
	wordCounts := make(map[string]int)
	for _, word := range words {
		if len(word) < 1 || word == " " {
			continue
		}
		if _, ok := wordCounts[word]; ok {
			wordCounts[word] += 1
		} else {
			wordCounts[word] = 1
		}
	}
	return wordCounts
}

func convertAllToLowercase(words []string) []string {
	newWords := make([]string, len(words))
	for i, word := range words {
		newWords[i] = strings.ToLower(word)
	}
	return newWords
}

func startsWith(word string, symbols []rune) bool {
	for _, symbol := range symbols {
		if len(word) > 0 {
			if word[0] == byte(symbol) {
				return true
			}
		}
	}
	return false
}

func endsWith(word string, symbols []rune) bool {
	if len(word) > 0 {
		end := len(word) - 1
		for _, symbol := range symbols {
			if word[end] == byte(symbol) {
				return true
			}
		}
	}
	return false
}

func removeNonLetterChars(words []string) []string {
	punctuation := []rune{'.', ',', '\'', '"', ';', '(', ')', '`', '!', '{',
		'}', '[', ']'}
	var newWords []string
	for _, word := range words {
		if startsWith(word, punctuation) {
			newWord := word[1:]
			newWords = append(newWords, newWord)
		} else if endsWith(word, punctuation) {
			length := len(word)
			newWord := word[:length-1]
			newWords = append(newWords, newWord)
		} else {
			newWords = append(newWords, word)
		}
	}
	return newWords
}

func splitIntoWords(text string) []string {
	var result []string
	lines := strings.Split(text, "\n")
	for _, line := range lines {
		words := strings.Split(line, " ")
		for _, word := range words {
			result = append(result, word)
		}
	}
	return result
}

func readFile(path string) string {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Println("Error opening file.")
		os.Exit(1)
	}
	return string(data)
}
