package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func main() {
	if len(os.Args) != 2 {
		return
	}

	arg := os.Args[1]
	num, err := strconv.Atoi(arg)
	if err != nil {
		printZeroes()
		return
	}

	binaryStr := strconv.FormatInt(int64(num), 2)
	printBinary(binaryStr)
}

func printBinary(binaryStr string) {
	diff := 8 - len(binaryStr)
	for i := 0; i < diff; i++ {
		z01.PrintRune('0')
	}
	for _, digit := range binaryStr {
		z01.PrintRune(digit)
	}
}

func printZeroes() {
	for i := 0; i < 8; i++ {
		z01.PrintRune('0')
	}
}
