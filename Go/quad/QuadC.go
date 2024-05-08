package main

import "fmt"

func QuadC(x, y int) {
	if x == 1 && y == 1 {

		fmt.Printf("A")
		fmt.Printf("\n")

	} else if x == 1 && y > 1 {

		fmt.Printf("A")
		fmt.Printf("\n")
		for i := 0; i < y-2; i++ {
			fmt.Printf("B")
			fmt.Printf("\n")
		}
		fmt.Printf("C")
		fmt.Printf("\n")

	} else if x > 1 && y == 1 {
		fmt.Printf("A")
		for i := 0; i < x-2; i++ {
			fmt.Printf("B")
		}
		fmt.Printf("A")
		fmt.Printf("\n")
	} else if x > 1 && y > 1 {
		fmt.Printf("A")
		for i := 0; i < x-2; i++ {
			fmt.Printf("B")
		}
		fmt.Printf("A")
		fmt.Printf("\n")
		for i := 0; i < y-2; i++ {
			fmt.Printf("B")
			for i := 0; i < x-2; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("B")
			fmt.Printf("\n")
		}
		fmt.Printf("C")
		for i := 0; i < x-2; i++ {
			fmt.Printf("B")
		}
		fmt.Printf("C")
		fmt.Printf("\n")
	}
}
