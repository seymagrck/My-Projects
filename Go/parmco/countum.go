package main

import (
	"fmt"
	"os"
)

func main() {
	dizi := os.Args
	fmt.Println(len(dizi) - 1)

	sayac := 0
	for i := 1; i < len(dizi); i++ {
		sayac++
	}
	fmt.Println(sayac)
}
