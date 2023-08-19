package main

import (
	"fmt"
	"strings"
)

func s_test(s string) {
	words := strings.Split(s, " ")
	fmt.Printf("word : %v, %T \n", words, words)
}

func main() {
	fmt.Printf("%v\n", wordPattern1("abba", "dog cat cat dog"))
	//s_test("dog cat cat dog")
}

func wordPattern1(pattern string, s string) bool {
	words := strings.Split(s, " ")
	sMap := map[rune]string{}
	if len(pattern) != len(words) {
		return false
	}

	for i, v := range pattern {
		if word, ok := sMap[v]; ok {
			if word != words[i] {
				return false
			}
		} else {
			for _, value := range sMap {
				if value == words[i] {
					return false
				}
			}
			sMap[v] = words[i]
		}
	}

	return true
}

func wordPattern(pattern string, s string) bool {
	words := strings.Split(s, " ")

	if len(pattern) != len(words) {
		return false
	}

	charToWord := make(map[rune]string)
	wordToChar := make(map[string]rune)

	for i, char := range pattern {
		word := words[i]

		if storedWord, exists := charToWord[char]; exists {
			if storedWord != word {
				return false
			}
		} else {
			charToWord[char] = word
		}

		if storedChar, exists := wordToChar[word]; exists {
			if storedChar != char {
				return false
			}
		} else {
			wordToChar[word] = char
		}
	}
	return true
}
