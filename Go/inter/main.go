package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	dizi := os.Args

	if len(dizi) != 3 {
		return
	}

	metinBir := dizi[1]
	metinIki := dizi[2]
	donus := ""

	seen := make(map[rune]bool)

	for _, char1 := range metinBir {
		for _, char2 := range metinIki {
			if char1 == char2 && !seen[char1] {

				donus += string(char1)
				seen[char1] = true
				break

			}
		}
	}

	goster(donus)
}

func goster(metin string) {
	for i := 0; i < len(metin); i++ {
		z01.PrintRune(rune(metin[i]))
	}

	z01.PrintRune('\n')
}