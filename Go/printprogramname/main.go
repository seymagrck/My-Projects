package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[0]
	arg2 := []rune(arg)
	for _, char := range arg2[2:] {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
