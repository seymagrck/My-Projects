package main

import "fmt"

func QuadB(x, y int) {
	if x == 1 && y == 1 {

		fmt.Printf("/")
		fmt.Printf("\n")

	} else if x == 1 && y > 1 {

		fmt.Printf("/")
		fmt.Printf("\n")
		for i := 0; i < y-2; i++ {
			fmt.Printf("*")
			fmt.Printf("\n")
		}
		fmt.Printf("%c", 92)
		fmt.Printf("\n")

	} else if x > 1 && y == 1 {
		fmt.Printf("/")
		for i := 0; i < x-2; i++ {
			fmt.Printf("*")
		}
		fmt.Printf("%c", 92)
		fmt.Printf("\n")
	} else if x > 1 && y > 1 {
		fmt.Printf("/")
		for i := 0; i < x-2; i++ {
			fmt.Printf("*")
		}
		fmt.Printf("%c", 92)
		fmt.Printf("\n")
		for i := 0; i < y-2; i++ {
			fmt.Printf("*")
			for i := 0; i < x-2; i++ {
				fmt.Printf(" ")
			}
			fmt.Printf("*")
			fmt.Printf("\n")
		}
		fmt.Printf("%c", 92)
		for i := 0; i < x-2; i++ {
			fmt.Printf("*")
		}
		fmt.Printf("/")
		fmt.Printf("\n")
	}
}
