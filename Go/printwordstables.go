package piscine

import "github.com/01-edu/z01"

func PrintWordsTables(table []string) {
	for _, w := range table {
		myStr := w
		for _, l := range myStr {
			z01.PrintRune(l)
		}
		z01.PrintRune('\n')
	}
}
