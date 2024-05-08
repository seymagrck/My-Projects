package piscine

func StrRevOnlyDigits(s string) string {
	runes := []rune(s)
	i, j := 0, len(runes)-1

	for i < j {
		for i < j && !isDigit(runes[i]) {
			i++
		}
		for i < j && !isDigit(runes[j]) {
			j--
		}
		runes[i], runes[j] = runes[j], runes[i]
		i++
		j--
	}

	return string(runes)
}

func isDigit(r rune) bool {
	return r >= '0' && r <= '9'
}
