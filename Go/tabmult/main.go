package main

import (
	"os"
	"strconv"

	"github.com/01-edu/z01"
)

func printNumber(num int) {
	if num < 10 {
		z01.PrintRune('0' + rune(num))
	} else {
		printNumber(num / 10)
		z01.PrintRune('0' + rune(num%10))
	}
}

func printMultiplicationTable(number int) {
	for i := 1; i <= 9; i++ {
		printNumber(i)
		z01.PrintRune(' ')
		z01.PrintRune('x')
		z01.PrintRune(' ')
		printNumber(number)
		z01.PrintRune(' ')
		z01.PrintRune('=')
		z01.PrintRune(' ')
		printNumber(i * number)
		z01.PrintRune('\n')
	}
}

func main() {
	// Check if the program is called with an argument
	if len(os.Args) != 2 {
		return
	}

	// Get the number from the command line argument
	numberStr := os.Args[1]

	// Convert the number argument to an integer
	number, err := strconv.Atoi(numberStr)
	if err != nil {
		return
	}

	// Display the multiplication table for the given number
	printMultiplicationTable(number)
}
