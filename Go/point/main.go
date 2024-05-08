package main

import (
	"github.com/01-edu/z01"
)

type point struct {
	x int
	y int
}

func setPoint(ptr *point) {
	ptr.x = 42
	ptr.y = 21
}

func main() {
	points := &point{}
	setPoint(points)
	sebo := "x = "
	ahmet := ", y = "
	for _, ch := range sebo {
		z01.PrintRune(ch)
	}
	PrintNbr(points.x)
	for _, xd := range ahmet {
		z01.PrintRune(xd)
	}
	PrintNbr(points.y)
	z01.PrintRune('\n')
}

func check(r int) {
	c := '0'
	for i := 1; i <= r%10; i++ {
		c++
	}
	for i := -1; i >= r%10; i-- {
		c++
	}
	if r/10 != 0 {
		check(r / 10)
	}
	z01.PrintRune(c)
}

func PrintNbr(n int) {
	check(n)
}
