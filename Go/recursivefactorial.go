package piscine

func RecursiveFactorial(nb int) int {
	result := nb
	if nb == 0 {
		return 1
	} else if nb < 0 || nb > 25 {
		return 0
	}
	return result * RecursiveFactorial(nb-1)
}
