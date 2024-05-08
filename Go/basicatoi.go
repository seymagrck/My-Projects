package piscine

func BasicAtoi(s string) int {
	var result int
	multiplier := 1
	for i := len(s) - 1; i >= 0; i-- {
		digit := int(s[i] - '0') // Karakteri rakama dönüştür
		result += digit * multiplier
		multiplier *= 10
	}
	return result
}
