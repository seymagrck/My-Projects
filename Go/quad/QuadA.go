package main

import "fmt"

func QuadA(x, y int) {
	if x == 1 && y == 1 {

		fmt.Printf("o")
		fmt.Printf("\n")

	} else if x == 1 && y > 1 {

		fmt.Printf("o")
		fmt.Printf("\n")

		for i := 0; i < y-2; i++ {
			fmt.Printf("|")
			fmt.Printf("\n")
		}
		fmt.Printf("o")
		fmt.Printf("\n")

	} else if x > 1 && y == 1 {
		fmt.Printf("o")
		for i := 0; i < x-2; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("o")
		fmt.Printf("\n")
	} else if x > 1 && y > 1 {

		fmt.Printf("o")
		for i := 0; i < x-2; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("o")
		fmt.Printf("\n")
		fmt.Printf("|")
		for i := 0; i < x-2; i++ {
			fmt.Printf(" ")
		}
		fmt.Printf("|")
		fmt.Printf("\n")
		fmt.Printf("o")
		for i := 0; i < x-2; i++ {
			fmt.Printf("-")
		}
		fmt.Printf("o")
		fmt.Printf("\n")

	}
}

func main() {
	QuadA(1, 5)
}
