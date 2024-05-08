package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func isPowerOfTwo(n int) bool {
	return (n > 0) && (n&(n-1) == 0)
}

func main() {
	args := os.Args[1:]

	if len(args) == 1 {
		num, err := strconv.Atoi(args[0])
		if err != nil {
			return
		}

		result := isPowerOfTwo(num)

		if result {
			metinBas("true")
		} else {
			metinBas("false")
		}
	}
}

func metinBas(metin string) {
	for i := 0; i < len(metin); i++ {
		z01.PrintRune(rune(metin[i]))
	}
	z01.PrintRune('\n')
}
