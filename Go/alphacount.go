package piscine

func AlphaCount(s string) int {
	var count int
	for i := len(s) - 1; i >= 0; i-- {
		if s[i] >= 'A' && s[i] <= 'Z' || s[i] >= 'a' && s[i] <= 'z' {
			count++
		}
	}
	return count
}
