package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	arg := os.Args[1:] // girilen argümanları 1. ve sonrasını arg değişkenin atadım
	for _, char := range arg {
		for _, r := range char {
			z01.PrintRune(r)
		}
		z01.PrintRune('\n')
	}
}
