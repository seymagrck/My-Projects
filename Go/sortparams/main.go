package main

import (
	"os"

	"github.com/01-edu/z01"
)

func main() {
	args := os.Args[1:]

	sortArguments(args)

	for _, arg := range args {
		printArgument(arg)
	}
}

func sortArguments(args []string) {
	for i := 0; i < len(args); i++ {
		for j := i + 1; j < len(args); j++ {
			if args[i] > args[j] {
				args[i], args[j] = args[j], args[i]
			}
		}
	}
}

func printArgument(arg string) {
	for _, char := range arg {
		z01.PrintRune(char)
	}
	z01.PrintRune('\n')
}
