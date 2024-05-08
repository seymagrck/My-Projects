package piscine

func FirstRune(s string) rune {
	if len(s) == 0 {
		return 0
	}
	runes := []rune(s)
	return runes[0]
}
