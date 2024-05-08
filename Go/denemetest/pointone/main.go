package main

import (
	"fmt"
)

func main() {
	n := -235673526746527345
	PointOne(&n)
	fmt.Println(n)
}

func PointOne(n *int) {
	*n = 1
}
