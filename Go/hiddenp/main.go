package main

import (
	"os"

	"github.com/01-edu/z01"
)

func isHidden(s1, s2 string) bool {
	if s1 == "" {
		// An empty string is considered hidden in any string
		return true
	}

	indexS1 := 0
	for _, charS2 := range s2 {
		if charS2 == rune(s1[indexS1]) {
			indexS1++
			if indexS1 == len(s1) {
				return true
			}
		}
	}

	return false
}

func main() {
	if len(os.Args) != 3 {
		return
	}

	s1 := os.Args[1]
	s2 := os.Args[2]

	if isHidden(s1, s2) {
		z01.PrintRune('1')
		z01.PrintRune('\n')
	} else {
		z01.PrintRune('0')
		z01.PrintRune('\n')
	}
}
