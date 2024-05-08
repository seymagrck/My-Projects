package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		z01.PrintRune('\n')
		return
	}

	input := cleanString(os.Args[1])
	if len(input) == 0 {
		z01.PrintRune('\n')
		return
	}

	printStr(input[0])
	for _, word := range input[1:] {
		z01.PrintRune(' ')
		printStr(word)
	}

	z01.PrintRune('\n')
}

func cleanString(input string) []string {
	var words []string
	var currentWord string
	inWord := false

	for _, char := range input {
		// Basit bir boşluk kontrolü
		if char == ' ' || char == '\t' || char == '\n' || char == '\r' {
			if inWord {
				words = append(words, currentWord)
				currentWord = ""
				inWord = false
			}
		} else {
			currentWord += string(char)
			inWord = true
		}
	}

	if inWord {
		words = append(words, currentWord)
	}

	return words
}

func printStr(str string) {
	for _, char := range str {
		z01.PrintRune(char)
	}
}
