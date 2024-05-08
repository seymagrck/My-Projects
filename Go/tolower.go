package piscine

func ToLower(s string) string {
	runes := []rune(s)
	for i := 0; i < len(runes); i++ {
		if runes[i] >= 'A' && runes[i] <= 'Z' {
			runes[i] = runes[i] - 'A' + 'a'
		}
	}
	return string(runes)
}
