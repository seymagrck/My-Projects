package piscine

func NRune(s string, n int) rune {
	if len(s) == 0 {
		return 0
	}

	if n <= 0 || len(s) < n {
		return 0
	}
	a := n - 1
	runes := []rune(s)
	return runes[a]
}
