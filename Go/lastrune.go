package piscine

func LastRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	a := len(s) - 1
	runes := []rune(s)
	return runes[a]
}
