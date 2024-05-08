package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	argument := os.Args
	if len(argument) == 4 {
		first := []rune(argument[1])
		second := []rune(argument[2])
		third := []rune(argument[3])

		for i := 0; i < len(first); i++ {
			if first[i] == second[0] {
				first[i] = third[0]
			}
		}
		for i := 0; i < len(first); i++ {
			z01.PrintRune(first[i])
		}
		z01.PrintRune('\n')
	}
}
