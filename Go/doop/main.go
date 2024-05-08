package main

import (
	"os"
)

func main() {
	if len(os.Args) < 4 {
		return
	}
	value, operator, other := TrimAtoi(os.Args[1]), os.Args[2], TrimAtoi(os.Args[3])
	if value == 0 && os.Args[1] != "0" || other == 0 && os.Args[3] != "0" || len(os.Args) < 4 {
		return
	}
	if isOverflow(value) || isOverflow(other) {
		return
	}
	switch operator {
	case "+":
		printInt(value + other)
	case "-":
		printInt(value - other)
	case "*":
		printInt(value * other)
	case "/":
		if value == 0 || other == 0 {
			os.Stdout.WriteString("No division by 0")
		} else {
			printInt(value / other)
		}
	case "%":
		if value == 0 || other == 0 {
			os.Stdout.WriteString("No modulo by 0")
		} else {
			printInt(value % other)
		}
	default:
		return
	}
	os.Stdout.WriteString("\n")
}

func isOverflow(value int) bool {
	return value >= 1<<31 || value <= -(1<<31)
}

func printInt(nbr int) {
	os.Stdout.WriteString(ConvertNbr(nbr))
}

func ConvertNbr(nbr int) string {
	str := ""
	isNegative := nbr < 0
	for i, remainder := nbr, 0; i != 0; i /= 10 {
		remainder = i % 10
		if isNegative {
			remainder = -remainder
		}
		char := rune(remainder) + rune('0')
		str += string(char)
	}
	if isNegative {
		str += "-"
	} else if nbr == 0 {
		str = "0"
	}
	return Reverse(str)
}

func Reverse(s string) string {
	runes := []rune(s)
	for i, j := 0, len(runes)-1; i < j; i, j = i+1, j-1 {
		runes[i], runes[j] = runes[j], runes[i]
	}
	return string(runes)
}

func TrimAtoi(s string) int {
	result := 0
	str := ""
	isNegative := false
	for index, char := range s {
		if index > 0 && s[index-1] == '-' && len(str) == 0 {
			isNegative = true
		}
		if isNumber(char) {
			str += string(char)
		}
	}
	str = Reverse(str)
	for index, char := range str {
		nb := int(char - '0')
		result += IterativePower(10, index) * nb
	}
	if isNegative {
		result = -result
	}
	return result
}

func isNumber(char rune) bool {
	return char >= '0' && char <= '9'
}

func IterativePower(nb int, power int) int {
	res := 1
	if power < 0 {
		return 0
	}
	for i := 1; i <= power; i++ {
		res = nb * res
	}
	return res
}
