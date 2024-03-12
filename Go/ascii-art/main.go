package main

import (
	"fmt"
	"os"
	"strings"
)

func main() {
	// concatenate the arguments from index[1] to one string
	concatString := strings.Join(os.Args[1:], " ")
	fmt.Println(concatString)
	// seperate the characters to a new line when \n is encountered
	sepWords := strings.Split(concatString, `\n`)
	fmt.Println(sepWords)
	// error handling
	file, err := os.ReadFile("standard.txt")
	if err != nil {
		fmt.Println("Error reading file")
		panic(err)
	}
	// count number of lines in each string length
	lines := strings.Split(string(file), "\n")
	// count number of words in each line
	for i, word := range sepWords {
		if word == "" {
			if i < len(sepWords)-1 {
				fmt.Println()
			}
			continue
		}
		// each ascii character is 9 lines tall
		for h := 1; h < 9; h++ {
			for _, l := range word {
				// calculate the range values for each ascii character
				for lineIndex, line := range lines {
					if lineIndex == (int(l)-32)*9+h {
						fmt.Print(line)
					}
				}
			}
			fmt.Println()
		}
	}
}
